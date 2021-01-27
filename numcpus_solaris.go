// Copyright 2021 Tobias Klauser
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build solaris

package numcpus

import (
	"syscall"
	"unsafe"
)

// TODO: remove sysconf wrapper once https://golang.org/cl/286593 is merged and x/sys updated.

//go:cgo_import_dynamic libc_sysconf sysconf "libc.so"
//go:linkname procSysconf libc_sysconf

var procSysconf uintptr

//go:linkname sysvicall6 syscall.sysvicall6
func sysvicall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

func sysconf(name int) (int64, error) {
	n, _, errno := sysvicall6(uintptr(unsafe.Pointer(&procSysconf)), 1, uintptr(name), 0, 0, 0, 0, 0)
	if errno != 0 {
		return -1, errno
	}
	return int64(n), nil
}

// taken from /usr/include/sys/unistd.h
const (
	_SC_NPROCESSORS_CONF = 14
	_SC_NPROCESSORS_ONLN = 15
	_SC_NPROCESSORS_MAX  = 516
)

func getKernelMax() (int, error) {
	n, err := sysconf(_SC_NPROCESSORS_MAX)
	return int(n), err
}

func getOffline() (int, error) {
	return 0, ErrNotSupported
}

func getOnline() (int, error) {
	n, err := sysconf(_SC_NPROCESSORS_ONLN)
	return int(n), err
}

func getPossible() (int, error) {
	n, err := sysconf(_SC_NPROCESSORS_CONF)
	return int(n), err
}

func getPresent() (int, error) {
	n, err := sysconf(_SC_NPROCESSORS_CONF)
	return int(n), err
}
