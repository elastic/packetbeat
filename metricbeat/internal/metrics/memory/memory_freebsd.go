// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package memory

import (
	"unsafe"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/metricbeat/internal/metrics"
)

/*
#include <sys/param.h>
#include <sys/mount.h>
#include <sys/ucred.h>
#include <sys/types.h>
#include <sys/sysctl.h>
#include <stdlib.h>
#include <stdint.h>
#include <unistd.h>
#include <time.h>
*/
import "C"

func get(_ string) (Memory, error) {
	val := C.uint32_t(0)
	sc := C.size_t(4)

	memData := Memory{}

	name := C.CString("vm.stats.vm.v_page_count")
	_, err := C.sysctlbyname(name, unsafe.Pointer(&val), &sc, nil, 0)
	C.free(unsafe.Pointer(name))
	if err != nil {
		return memData, errors.Errorf("error in vm.stats.vm.v_page_count")
	}
	pagecount := uint64(val)

	name = C.CString("vm.stats.vm.v_page_size")
	_, err = C.sysctlbyname(name, unsafe.Pointer(&val), &sc, nil, 0)
	C.free(unsafe.Pointer(name))
	if err != nil {
		return memData, errors.Errorf("error in vm.stats.vm.v_page_size")
	}
	pagesize := uint64(val)

	name = C.CString("vm.stats.vm.v_free_count")
	_, err = C.sysctlbyname(name, unsafe.Pointer(&val), &sc, nil, 0)
	C.free(unsafe.Pointer(name))
	if err != nil {
		return memData, errors.Errorf("error in vm.stats.vm.v_free_count")
	}
	memData.Free = metrics.NewUintValue(uint64(val) * pagesize)

	name = C.CString("vm.stats.vm.v_inactive_count")
	_, err = C.sysctlbyname(name, unsafe.Pointer(&val), &sc, nil, 0)
	C.free(unsafe.Pointer(name))
	if err != nil {
		return memData, errors.Errorf("error in vm.stats.vm.v_inactive_count")
	}
	kern := uint64(val)

	memData.Total = metrics.NewUintValue(uint64(pagecount * pagesize))

	memData.Used.Bytes = metrics.NewUintValue(memData.Total.ValueOrZero() - memData.Free.ValueOrZero())
	memData.Actual.Free = metrics.NewUintValue(memData.Free.ValueOrZero() + (kern * pagesize))
	memData.Actual.Used.Bytes = metrics.NewUintValue(memData.Used.Bytes.ValueOrZero() - (kern * pagesize))

	return memData, nil
}
