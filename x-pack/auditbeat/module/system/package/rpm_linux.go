// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build linux,cgo

package pkg

import (
	"errors"
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"github.com/coreos/pkg/dlopen"
	"github.com/joeshaw/multierror"
)

/*
#include <stdio.h>
#include <stdlib.h>

#include <rpm/rpmlib.h>
#include <rpm/header.h>
#include <rpm/rpmts.h>
#include <rpm/rpmdb.h>
#include <rpm/rpmsq.h>

rpmts
my_rpmtsCreate(void *f) {
  rpmts (*rpmtsCreate)();
  rpmtsCreate = (rpmts (*)())f;

  return rpmtsCreate();
}

int
my_rpmReadConfigFiles(void *f) {
  int (*rpmReadConfigFiles)(const char*, const char*);
  rpmReadConfigFiles = (int (*)(const char*, const char*))f;
  return rpmReadConfigFiles(NULL, NULL);
}

rpmdbMatchIterator
my_rpmtsInitIterator(void *f, rpmts ts) {
  rpmdbMatchIterator (*rpmtsInitIterator)(const rpmts, rpmTag, const void*, size_t);
  rpmtsInitIterator = (rpmdbMatchIterator (*)(const rpmts, rpmTag, const void*, size_t))f;

  return rpmtsInitIterator(ts, RPMDBI_PACKAGES, NULL, 0);
}

Header
my_rpmdbNextIterator(void *f, rpmdbMatchIterator mi) {
  Header (*rpmdbNextIterator)(rpmdbMatchIterator);
  rpmdbNextIterator = (Header (*)(rpmdbMatchIterator))f;

  return rpmdbNextIterator(mi);
}

Header
my_headerLink(void *f, Header h) {
  Header (*headerLink)(Header);
  headerLink = (Header (*)(Header))f;

  return headerLink(h);
}

// Note: Using int32_t instead of rpmTag/rpmTagVal in definitions
// to make it work on CentOS 6.x, 7.x, and Fedora 29.
const char *
my_headerGetString(void *f, Header h, int32_t tag) {
  const char * (*headerGetString)(Header, int32_t);
  headerGetString = (const char * (*)(Header, int32_t))f;

  return headerGetString(h, tag);
}

// Note: Using int32_t instead of rpmTag/rpmTagVal in definitions
// to make it work on CentOS 6.x, 7.x, and Fedora 29.
uint64_t
my_headerGetNumber(void *f, Header h, int32_t tag) {
  uint64_t (*headerGetNumber)(Header, int32_t);
  headerGetNumber = (uint64_t (*)(Header, int32_t))f;

  return headerGetNumber(h, tag);
}

void
my_headerFree(void *f, Header h) {
  Header (*headerFree)(Header);
  headerFree = (Header (*)(Header))f;

  headerFree(h);
}

void
my_rpmdbFreeIterator(void *f, rpmdbMatchIterator mi) {
  rpmdbMatchIterator (*rpmdbFreeIterator)(rpmdbMatchIterator);
  rpmdbFreeIterator = (rpmdbMatchIterator (*)(rpmdbMatchIterator))f;

  rpmdbFreeIterator(mi);
}

void
my_rpmtsFree(void *f, rpmts ts) {
  rpmts (*rpmtsFree)(rpmts);
  rpmtsFree = (rpmts (*)(rpmts))f;

  rpmtsFree(ts);
}

// By default, librpm is going to trap various UNIX signals including SIGINT and SIGTERM
// which will prevent Beats from shutting down correctly.
//
// This disables that behavior. We should be very dilligent in
// cleaning up in our use of librpm.
//
// More recent versions of librpm have a new function rpmsqSetInterruptSafety()
// to do this, see below.
//
// See also:
// - librpm traps signals and calls exit(1) to terminate the whole process incl. our Go code: https://github.com/rpm-software-management/rpm/blob/rpm-4.11.3-release/lib/rpmdb.c#L640
// - has caused problems for gdb before, calling rpmsqEnable(_, NULL) is the workaround they also use: https://bugzilla.redhat.com/show_bug.cgi?id=643031
// - the new rpmsqSetInterruptSafety(), unfortunately only available in librpm>=4.14.0 (CentOS 7 has 4.11.3): https://github.com/rpm-software-management/rpm/commit/56f49d7f5af7c1c8a3eb478431356195adbfdd25
void
my_disableLibrpmSignalTraps(void *f) {
	int (*rpmsqEnable)(int, rpmsqAction_t);
	rpmsqEnable = (int (*)(int, rpmsqAction_t))f;

	// Disable all traps
	rpmsqEnable(-SIGHUP, NULL);
	rpmsqEnable(-SIGINT, NULL);
	rpmsqEnable(-SIGTERM, NULL);
	rpmsqEnable(-SIGQUIT, NULL);
	rpmsqEnable(-SIGPIPE, NULL);
}

void
my_rpmsqSetInterruptSafety(void *f, int on) {
	void (*rpmsqSetInterruptSafety)(int);
	rpmsqSetInterruptSafety = (void (*)(int))f;

	rpmsqSetInterruptSafety(on);
}
*/
import "C"

// Constants in sync with /usr/include/rpm/rpmtag.h
const (
	RPMTAG_NAME        = 1000
	RPMTAG_VERSION     = 1001
	RPMTAG_RELEASE     = 1002
	RPMTAG_SUMMARY     = 1004
	RPMTAG_LICENSE     = 1014
	RPMTAG_URL         = 1020
	RPMTAG_ARCH        = 1022
	RPMTAG_SIZE        = 1009
	RPMTAG_INSTALLTIME = 1008
)

type cFunctions struct {
	rpmtsCreate             unsafe.Pointer
	rpmReadConfigFiles      unsafe.Pointer
	rpmtsInitIterator       unsafe.Pointer
	rpmdbNextIterator       unsafe.Pointer
	headerLink              unsafe.Pointer
	headerGetString         unsafe.Pointer
	headerGetNumber         unsafe.Pointer
	headerFree              unsafe.Pointer
	rpmdbFreeIterator       unsafe.Pointer
	rpmtsFree               unsafe.Pointer
	rpmsqEnable             unsafe.Pointer
	rpmsqSetInterruptSafety unsafe.Pointer
}

var cFun *cFunctions

func dlopenCFunctions() (*cFunctions, error) {
	var librpmNames = []string{
		"/usr/lib64/librpm.so",
	}
	var cFun cFunctions

	librpm, err := dlopen.GetHandle(librpmNames)
	if err != nil {
		return nil, err
	}

	cFun.rpmtsCreate, err = librpm.GetSymbolPointer("rpmtsCreate")
	if err != nil {
		return nil, err
	}

	cFun.rpmReadConfigFiles, err = librpm.GetSymbolPointer("rpmReadConfigFiles")
	if err != nil {
		return nil, err
	}

	cFun.rpmtsInitIterator, err = librpm.GetSymbolPointer("rpmtsInitIterator")
	if err != nil {
		return nil, err
	}

	cFun.rpmdbNextIterator, err = librpm.GetSymbolPointer("rpmdbNextIterator")
	if err != nil {
		return nil, err
	}

	cFun.headerLink, err = librpm.GetSymbolPointer("headerLink")
	if err != nil {
		return nil, err
	}

	cFun.headerGetString, err = librpm.GetSymbolPointer("headerGetString")
	if err != nil {
		return nil, err
	}

	cFun.headerGetNumber, err = librpm.GetSymbolPointer("headerGetNumber")
	if err != nil {
		return nil, err
	}

	cFun.headerFree, err = librpm.GetSymbolPointer("headerFree")
	if err != nil {
		return nil, err
	}

	cFun.rpmdbFreeIterator, err = librpm.GetSymbolPointer("rpmdbFreeIterator")
	if err != nil {
		return nil, err
	}

	cFun.rpmtsFree, err = librpm.GetSymbolPointer("rpmtsFree")
	if err != nil {
		return nil, err
	}

	// Only available in librpm>=4.13.0
	cFun.rpmsqSetInterruptSafety, err = librpm.GetSymbolPointer("rpmsqSetInterruptSafety")
	if err != nil {
		var err2 error
		// Only available in librpm<4.14.0
		cFun.rpmsqEnable, err2 = librpm.GetSymbolPointer("rpmsqEnable")
		if err2 != nil {
			var errs multierror.Errors
			errs = append(errs, err, err2)
			return nil, errs.Err()
		}
	}

	return &cFun, nil
}

func listRPMPackages() ([]*Package, error) {
	// In newer versions, librpm is using the thread-local variable
	// `disableInterruptSafety` in rpmio/rpmsq.c to disable signal
	// traps. To make sure our settings remain in effect throughout
	// our function calls we have to lock the OS thread here, since
	// Golang can otherwise use any thread it likes for each C.* call.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	if cFun == nil {
		var err error
		cFun, err = dlopenCFunctions()
		if err != nil {
			return nil, err
		}
	}

	if cFun.rpmsqSetInterruptSafety != nil {
		C.my_rpmsqSetInterruptSafety(cFun.rpmsqSetInterruptSafety, 0)
	}

	rpmts := C.my_rpmtsCreate(cFun.rpmtsCreate)
	if rpmts == nil {
		return nil, fmt.Errorf("Failed to get rpmts")
	}
	defer C.my_rpmtsFree(cFun.rpmtsFree, rpmts)
	res := C.my_rpmReadConfigFiles(cFun.rpmReadConfigFiles)
	if int(res) != 0 {
		return nil, fmt.Errorf("Error: %d", int(res))
	}

	mi := C.my_rpmtsInitIterator(cFun.rpmtsInitIterator, rpmts)
	if mi == nil {
		return nil, fmt.Errorf("Failed to get match iterator")
	}
	if cFun.rpmsqEnable != nil {
		C.my_disableLibrpmSignalTraps(cFun.rpmsqEnable)
	}
	defer C.my_rpmdbFreeIterator(cFun.rpmdbFreeIterator, mi)

	var packages []*Package
	for header := C.my_rpmdbNextIterator(cFun.rpmdbNextIterator, mi); header != nil; header = C.my_rpmdbNextIterator(cFun.rpmdbNextIterator, mi) {

		pkg, err := packageFromHeader(header, cFun)
		if err != nil {
			return nil, err
		}

		packages = append(packages, pkg)
	}

	return packages, nil
}

func packageFromHeader(header C.Header, cFun *cFunctions) (*Package, error) {

	header = C.my_headerLink(cFun.headerLink, header)
	if header == nil {
		return nil, fmt.Errorf("Error calling headerLink")
	}
	defer C.my_headerFree(cFun.headerFree, header)

	pkg := Package{}

	name := C.my_headerGetString(cFun.headerGetString, header, RPMTAG_NAME)
	if name != nil {
		pkg.Name = C.GoString(name)
	} else {
		pkg.Error = errors.New("Failed to get package name")
	}

	version := C.my_headerGetString(cFun.headerGetString, header, RPMTAG_VERSION)
	if version != nil {
		pkg.Version = C.GoString(version)
	} else {
		pkg.Error = errors.New("Failed to get package version")
	}

	release := C.my_headerGetString(cFun.headerGetString, header, RPMTAG_RELEASE)
	if release != nil {
		pkg.Release = C.GoString(release)
	}

	license := C.my_headerGetString(cFun.headerGetString, header, RPMTAG_LICENSE)
	if license != nil {
		pkg.License = C.GoString(license)
	}

	arch := C.my_headerGetString(cFun.headerGetString, header, RPMTAG_ARCH)
	if arch != nil {
		pkg.Arch = C.GoString(arch)
	}

	url := C.my_headerGetString(cFun.headerGetString, header, RPMTAG_URL)
	if url != nil {
		pkg.URL = C.GoString(url)
	}

	summary := C.my_headerGetString(cFun.headerGetString, header, RPMTAG_SUMMARY)
	if summary != nil {
		pkg.Summary = C.GoString(summary)
	}

	size := C.my_headerGetNumber(cFun.headerGetNumber, header, RPMTAG_SIZE)
	if size != 0 {
		pkg.Size = uint64(size)
	}

	installTime := C.my_headerGetNumber(cFun.headerGetNumber, header, RPMTAG_INSTALLTIME)
	if installTime != 0 {
		pkg.InstallTime = time.Unix(int64(installTime), 0)
	}

	return &pkg, nil
}
