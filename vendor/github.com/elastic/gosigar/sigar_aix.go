// +build aix

package gosigar

/*
#cgo LDFLAGS: -L/usr/lib -lperfstat

#include <libperfstat.h>
#include <procinfo.h>
#include <unistd.h>
#include <utmp.h>
#include <sys/mntctl.h>
#include <sys/proc.h>
#include <sys/types.h>
#include <sys/vmount.h>

*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

var system struct {
	ticks    uint64
	btime    uint64
	pagesize uint64
}

func init() {
	// sysconf(_SC_CLK_TCK) returns the number of ticks by second.
	system.ticks = uint64(C.sysconf(C._SC_CLK_TCK))
	system.pagesize = uint64(os.Getpagesize())
}

// utmp can't be used by "encoding/binary" if generated by cgo,
// some pads will be explicitly missing.
type utmp struct {
	User        [256]uint8
	ID          [14]uint8
	Line        [64]uint8
	XPad1       int16
	Pid         int32
	Type        int16
	XPad2       int16
	Time        int64
	Termination int16
	Exit        int16
	Host        [256]uint8
	XdblWordPad int32
	XreservedA  [2]int32
	XreservedV  [6]int32
}

func bootTime() (uint64, error) {
	if system.btime != 0 {
		return system.btime, nil
	}

	// Get boot time from /etc/utmp
	file, err := os.Open("/etc/utmp")
	if err != nil {
		return 0, fmt.Errorf("error while opening /etc/utmp: %s", err)
	}

	defer file.Close()

	for {
		var utmp utmp
		if err := binary.Read(file, binary.BigEndian, &utmp); err != nil {
			break
		}

		if utmp.Type == C.BOOT_TIME {
			system.btime = uint64(utmp.Time)
			break
		}
	}
	return system.btime, nil
}

func tick2msec(val uint64) uint64 {
	return val * 1000 / system.ticks
}

// Get returns the list of file systems
func (self *FileSystemList) Get() error {
	var size C.int
	_, err := C.mntctl(C.MCTL_QUERY, C.sizeof_int, (*C.char)(unsafe.Pointer(&size)))
	if err != nil {
		return fmt.Errorf("error while retrieving file system number: %s", err)
	}

	buf := make([]byte, size)
	num, err := C.mntctl(C.MCTL_QUERY, C.ulong(size), (*C.char)(&buf[0]))
	if err != nil {
		return fmt.Errorf("error while retrieving file system list: %s", err)
	}

	// Vmount structure has a fixed size area for common data (type,
	// offsets, etc) and another area with variable length data (devname,
	// options, etc). These data can be accessed based on the offsets
	// stored in an array inside the fixed part. They can be retrieve
	// using index given by C define.
	vmt2data := func(buf []byte, ent *C.struct_vmount, idx int, baseOff int) []byte {
		off := int(ent.vmt_data[idx].vmt_off)
		size := int(ent.vmt_data[idx].vmt_size)
		return buf[baseOff+off : baseOff+off+size]
	}

	entOff := 0

	fslist := make([]FileSystem, num)
	for i := 0; i < int(num); i++ {
		ent := (*C.struct_vmount)(unsafe.Pointer(&buf[entOff]))
		fs := &fslist[i]

		// Correspondances taken for /etc/vfs
		switch ent.vmt_gfstype {
		case C.MNT_AIX:
			fs.SysTypeName = "jfs2"
		case C.MNT_NAMEFS:
			fs.SysTypeName = "namefs"
		case C.MNT_NFS:
			fs.SysTypeName = "nfs"
		case C.MNT_JFS:
			fs.SysTypeName = "jfs"
		case C.MNT_CDROM:
			fs.SysTypeName = "cdrom"
		case C.MNT_PROCFS:
			fs.SysTypeName = "proc"
		case C.MNT_SFS:
			fs.SysTypeName = "sfs"
		case C.MNT_CACHEFS:
			fs.SysTypeName = "cachefs"
		case C.MNT_NFS3:
			fs.SysTypeName = "nfs3"
		case C.MNT_AUTOFS:
			fs.SysTypeName = "autofs"
		case C.MNT_POOLFS:
			fs.SysTypeName = "poolfs"
		case C.MNT_UDF:
			fs.SysTypeName = "udfs"
		case C.MNT_NFS4:
			fs.SysTypeName = "nfs4"
		case C.MNT_CIFS:
			fs.SysTypeName = "cifs"
		case C.MNT_PMEMFS:
			fs.SysTypeName = "pmemfs"
		case C.MNT_AHAFS:
			fs.SysTypeName = "ahafs"
		case C.MNT_STNFS:
			fs.SysTypeName = "stnfs"
		default:
			if ent.vmt_flags&C.MNT_REMOTE != 0 {
				fs.SysTypeName = "network"
			} else {
				fs.SysTypeName = "none"
			}
		}

		fs.DirName = convertBytesToString(vmt2data(buf, ent, C.VMT_STUB, entOff))
		fs.Options = convertBytesToString(vmt2data(buf, ent, C.VMT_ARGS, entOff))
		devname := convertBytesToString(vmt2data(buf, ent, C.VMT_OBJECT, entOff))
		if ent.vmt_flags&C.MNT_REMOTE != 0 {
			hostname := convertBytesToString(vmt2data(buf, ent, C.VMT_OBJECT, entOff))
			fs.DevName = hostname + ":" + devname
		} else {
			fs.DevName = devname
		}

		entOff += int(ent.vmt_length)
	}

	self.List = fslist

	return nil
}

// Get returns the CPU load average
func (self *LoadAverage) Get() error {
	cpudata := C.perfstat_cpu_total_t{}

	if _, err := C.perfstat_cpu_total(nil, &cpudata, C.sizeof_perfstat_cpu_total_t, 1); err != nil {
		return fmt.Errorf("perfstat_cpu_total: %s", err)
	}

	// from libperfstat.h:
	// "To calculate the load average, divide the numbers by (1<<SBITS).
	//  SBITS is defined in <sys/proc.h>."
	fixedToFloat64 := func(x uint64) float64 {
		return float64(x) / (1 << C.SBITS)
	}
	self.One = fixedToFloat64(uint64(cpudata.loadavg[0]))
	self.Five = fixedToFloat64(uint64(cpudata.loadavg[1]))
	self.Fifteen = fixedToFloat64(uint64(cpudata.loadavg[2]))

	return nil
}

// Get returns the system uptime
func (self *Uptime) Get() error {
	btime, err := bootTime()
	if err != nil {
		return err
	}
	uptime := time.Now().Sub(time.Unix(int64(btime), 0))
	self.Length = uptime.Seconds()
	return nil
}

// Get returns the current system memory
func (self *Mem) Get() error {
	meminfo := C.perfstat_memory_total_t{}
	_, err := C.perfstat_memory_total(nil, &meminfo, C.sizeof_perfstat_memory_total_t, 1)
	if err != nil {
		return fmt.Errorf("perfstat_memory_total: %s", err)
	}

	self.Total = uint64(meminfo.real_total) * system.pagesize
	self.Free = uint64(meminfo.real_free) * system.pagesize

	kern := uint64(meminfo.numperm) * system.pagesize // number of pages in file cache

	self.Used = self.Total - self.Free
	self.ActualFree = self.Free + kern
	self.ActualUsed = self.Used - kern

	return nil
}

// Get returns the current system swap memory
func (self *Swap) Get() error {
	ps := C.perfstat_pagingspace_t{}
	id := C.perfstat_id_t{}

	id.name[0] = 0

	for {
		// errno can be set during perfstat_pagingspace's call even
		// if it succeeds. Thus, only check it when the result is -1.
		if r, err := C.perfstat_pagingspace(&id, &ps, C.sizeof_perfstat_pagingspace_t, 1); r == -1 && err != nil {
			return fmt.Errorf("perfstat_memory_total: %s", err)
		}

		if ps.active != 1 {
			continue
		}

		// convert MB sizes to bytes
		self.Total += uint64(ps.mb_size) * 1024 * 1024
		self.Used += uint64(ps.mb_used) * 1024 * 1024

		if id.name[0] == 0 {
			break
		}
	}

	self.Free = self.Total - self.Used

	return nil
}

// Get returns information about a CPU
func (self *Cpu) Get() error {
	cpudata := C.perfstat_cpu_total_t{}

	if _, err := C.perfstat_cpu_total(nil, &cpudata, C.sizeof_perfstat_cpu_total_t, 1); err != nil {
		return fmt.Errorf("perfstat_cpu_total: %s", err)
	}

	self.User = tick2msec(uint64(cpudata.user))
	self.Sys = tick2msec(uint64(cpudata.sys))
	self.Idle = tick2msec(uint64(cpudata.idle))
	self.Wait = tick2msec(uint64(cpudata.wait))

	return nil
}

// Get returns the list of CPU used by the system
func (self *CpuList) Get() error {
	cpudata := C.perfstat_cpu_t{}
	id := C.perfstat_id_t{}
	id.name[0] = 0

	// Retrieve the number of cpu using perfstat_cpu
	capacity, err := C.perfstat_cpu(nil, nil, C.sizeof_perfstat_cpu_t, 0)
	if err != nil {
		return fmt.Errorf("error while retrieving CPU number: %s", err)
	}
	list := make([]Cpu, 0, capacity)

	for {
		if _, err := C.perfstat_cpu(&id, &cpudata, C.sizeof_perfstat_cpu_t, 1); err != nil {
			return fmt.Errorf("perfstat_cpu: %s", err)
		}

		cpu := Cpu{}
		cpu.User = tick2msec(uint64(cpudata.user))
		cpu.Sys = tick2msec(uint64(cpudata.sys))
		cpu.Idle = tick2msec(uint64(cpudata.idle))
		cpu.Wait = tick2msec(uint64(cpudata.wait))

		list = append(list, cpu)

		if id.name[0] == 0 {
			break
		}
	}

	self.List = list

	return nil
}

// Get returns the list of all active processes
func (self *ProcList) Get() error {
	info := C.struct_procsinfo64{}
	pid := C.pid_t(0)

	var list []int

	for {
		// getprocs first argument is a void*
		num, err := C.getprocs(unsafe.Pointer(&info), C.sizeof_struct_procsinfo64, nil, 0, &pid, 1)
		if err != nil {
			return err
		}

		list = append(list, int(info.pi_pid))

		if num == 0 {
			break
		}
	}

	self.List = list

	return nil
}

// Get returns information about a process
func (self *ProcState) Get(pid int) error {
	info := C.struct_procsinfo64{}
	cpid := C.pid_t(pid)

	num, err := C.getprocs(unsafe.Pointer(&info), C.sizeof_struct_procsinfo64, nil, 0, &cpid, 1)
	if err != nil {
		return err
	}
	if num != 1 {
		return syscall.ESRCH
	}

	self.Name = C.GoString(&info.pi_comm[0])
	self.Ppid = int(info.pi_ppid)
	self.Pgid = int(info.pi_pgrp)
	self.Nice = int(info.pi_nice)
	self.Tty = int(info.pi_ttyd)
	self.Priority = int(info.pi_pri)

	switch info.pi_state {
	case C.SACTIVE:
		self.State = RunStateRun
	case C.SIDL:
		self.State = RunStateIdle
	case C.SSTOP:
		self.State = RunStateStop
	case C.SZOMB:
		self.State = RunStateZombie
	case C.SSWAP:
		self.State = RunStateSleep
	default:
		self.State = RunStateUnknown
	}

	// Get process username. Fallback to UID if username is not available.
	uid := strconv.Itoa(int(info.pi_uid))
	userID, err := user.LookupId(uid)
	if err == nil && userID.Username != "" {
		self.Username = userID.Username
	} else {
		self.Username = uid
	}

	thrinfo := C.struct_thrdsinfo64{}
	tid := C.tid_t(0)

	if _, err := C.getthrds(cpid, unsafe.Pointer(&thrinfo), C.sizeof_struct_thrdsinfo64, &tid, 1); err != nil {
		self.Processor = int(thrinfo.ti_affinity)
	}

	return nil
}

//Get returns the current memory usage of a process
func (self *ProcMem) Get(pid int) error {
	info := C.struct_procsinfo64{}
	cpid := C.pid_t(pid)

	num, err := C.getprocs(unsafe.Pointer(&info), C.sizeof_struct_procsinfo64, nil, 0, &cpid, 1)
	if err != nil {
		return err
	}
	if num != 1 {
		return syscall.ESRCH
	}

	self.Size = uint64(info.pi_size) * system.pagesize
	self.Share = uint64(info.pi_sdsize) * system.pagesize
	self.Resident = uint64(info.pi_drss+info.pi_trss) * system.pagesize

	self.MinorFaults = uint64(info.pi_minflt)
	self.MajorFaults = uint64(info.pi_majflt)
	self.PageFaults = self.MinorFaults + self.MajorFaults

	return nil
}

// Get returns a process uptime
func (self *ProcTime) Get(pid int) error {
	info := C.struct_procsinfo64{}
	cpid := C.pid_t(pid)

	num, err := C.getprocs(unsafe.Pointer(&info), C.sizeof_struct_procsinfo64, nil, 0, &cpid, 1)
	if err != nil {
		return err
	}
	if num != 1 {
		return syscall.ESRCH
	}

	self.StartTime = uint64(info.pi_start) * 1000
	self.User = uint64(info.pi_utime) * 1000
	self.Sys = uint64(info.pi_stime) * 1000
	self.Total = self.User + self.Sys

	return nil
}

// Get returns arguments of a process
func (self *ProcArgs) Get(pid int) error {
	/* If buffer is not large enough, args are truncated */
	buf := make([]byte, 8192)
	info := C.struct_procsinfo64{}
	info.pi_pid = C.pid_t(pid)

	if _, err := C.getargs(unsafe.Pointer(&info), C.sizeof_struct_procsinfo64, (*C.char)(&buf[0]), 8192); err != nil {
		return err
	}

	bbuf := bytes.NewBuffer(buf)

	var args []string

	for {
		arg, err := bbuf.ReadBytes(0)
		if err == io.EOF || arg[0] == 0 {
			break
		}
		if err != nil {
			return err
		}

		args = append(args, string(chop(arg)))
	}

	self.List = args
	return nil
}

// Get returns the environment of a process
func (self *ProcEnv) Get(pid int) error {
	if self.Vars == nil {
		self.Vars = map[string]string{}
	}

	/* If buffer is not large enough, args are truncated */
	buf := make([]byte, 8192)
	info := C.struct_procsinfo64{}
	info.pi_pid = C.pid_t(pid)

	if _, err := C.getevars(unsafe.Pointer(&info), C.sizeof_struct_procsinfo64, (*C.char)(&buf[0]), 8192); err != nil {
		return err
	}

	bbuf := bytes.NewBuffer(buf)

	delim := []byte{61} // "="

	for {
		line, err := bbuf.ReadBytes(0)
		if err == io.EOF || line[0] == 0 {
			break
		}
		if err != nil {
			return err
		}

		pair := bytes.SplitN(chop(line), delim, 2)
		if len(pair) != 2 {
			return fmt.Errorf("Error reading process environment for PID: %d", pid)
		}
		self.Vars[string(pair[0])] = string(pair[1])
	}

	return nil
}

// Get returns the path of the process executable
func (self *ProcExe) Get(pid int) error {
	/* If buffer is not large enough, args are truncated */
	buf := make([]byte, 8192)
	info := C.struct_procsinfo64{}
	info.pi_pid = C.pid_t(pid)

	if _, err := C.getargs(unsafe.Pointer(&info), C.sizeof_struct_procsinfo64, (*C.char)(&buf[0]), 8192); err != nil {
		return err
	}

	bbuf := bytes.NewBuffer(buf)

	// retrieve the first argument
	cmd, err := bbuf.ReadBytes(0)
	if err != nil {
		return err
	}
	self.Name = string(chop(cmd))

	cwd, err := os.Readlink("/proc/" + strconv.Itoa(pid) + "/cwd")
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	self.Cwd = cwd

	return nil
}

// Get returns process filesystem usage. Not implimented on AIX.
func (*ProcFDUsage) Get(_ int) error {
	return ErrNotImplemented{runtime.GOOS}
}

// Get returns filesytem usage. Not implimented on AIX.
func (*FDUsage) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}

// Get returns huge pages info. Not implimented on AIX.
func (*HugeTLBPages) Get() error {
	return ErrNotImplemented{runtime.GOOS}
}
