// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package perfmon

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

var (
	modpdh = windows.NewLazySystemDLL("pdh.dll")

	procPdhOpenQueryW               = modpdh.NewProc("PdhOpenQueryW")
	procPdhAddEnglishCounterW       = modpdh.NewProc("PdhAddEnglishCounterW")
	procPdhCollectQueryData         = modpdh.NewProc("PdhCollectQueryData")
	procPdhGetFormattedCounterValue = modpdh.NewProc("PdhGetFormattedCounterValue")
	procPdhCloseQuery               = modpdh.NewProc("PdhCloseQuery")
)

func _PdhOpenQuery(dataSource *uint16, userData uintptr, query *PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhOpenQueryW.Addr(), 3, uintptr(unsafe.Pointer(dataSource)), uintptr(userData), uintptr(unsafe.Pointer(query)))
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhAddCounter(query PdhQueryHandle, counterPath string, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	var _p0 *uint16
	_p0, errcode = syscall.UTF16PtrFromString(counterPath)
	if errcode != nil {
		return
	}
	return __PdhAddCounter(query, _p0, userData, counter)
}

func __PdhAddCounter(query PdhQueryHandle, counterPath *uint16, userData uintptr, counter *PdhCounterHandle) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhAddEnglishCounterW.Addr(), 4, uintptr(query), uintptr(unsafe.Pointer(counterPath)), uintptr(userData), uintptr(unsafe.Pointer(counter)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhCollectQueryData(query PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhCollectQueryData.Addr(), 1, uintptr(query), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhGetFormattedCounterValue(counter PdhCounterHandle, format PdhCounterFormat, counterType *uint32, value *PdhCounterValue) (errcode error) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterValue.Addr(), 4, uintptr(counter), uintptr(format), uintptr(unsafe.Pointer(counterType)), uintptr(unsafe.Pointer(value)), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}

func _PdhCloseQuery(query PdhQueryHandle) (errcode error) {
	r0, _, _ := syscall.Syscall(procPdhCloseQuery.Addr(), 1, uintptr(query), 0, 0)
	if r0 != 0 {
		errcode = syscall.Errno(r0)
	}
	return
}
