// Copyright 2018 Tobias Klauser
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

package numcpus

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

const sysfsCPUBasePath = "/sys/devices/system/cpu"

func getFromCPUAffinity() (int, error) {
	var cpuSet unix.CPUSet
	if err := unix.SchedGetaffinity(0, &cpuSet); err != nil {
		return 0, err
	}
	return cpuSet.Count(), nil
}

func readCPURangeWith[T any](file string, f func(cpus string) (T, error)) (T, error) {
	var zero T
	buf, err := os.ReadFile(filepath.Join(sysfsCPUBasePath, file))
	if err != nil {
		return zero, err
	}
	return f(strings.Trim(string(buf), "\n "))
}

func countCPURange(cpus string) (int, error) {
	n := int(0)
	for _, cpuRange := range strings.Split(cpus, ",") {
		if len(cpuRange) == 0 {
			continue
		}
		from, to, found := strings.Cut(cpuRange, "-")
		first, err := strconv.ParseUint(from, 10, 32)
		if err != nil {
			return 0, err
		}
		if !found {
			n++
			continue
		}
		last, err := strconv.ParseUint(to, 10, 32)
		if err != nil {
			return 0, err
		}
		if last < first {
			return 0, fmt.Errorf("last CPU in range (%d) less than first (%d)", last, first)
		}
		n += int(last - first + 1)
	}
	return n, nil
}

func getConfigured() (int, error) {
	d, err := os.Open(sysfsCPUBasePath)
	if err != nil {
		return 0, err
	}
	defer d.Close()
	fis, err := d.Readdir(-1)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, fi := range fis {
		if name := fi.Name(); fi.IsDir() && strings.HasPrefix(name, "cpu") {
			_, err := strconv.ParseInt(name[3:], 10, 64)
			if err == nil {
				count++
			}
		}
	}
	return count, nil
}

func getKernelMax() (int, error) {
	buf, err := os.ReadFile(filepath.Join(sysfsCPUBasePath, "kernel_max"))
	if err != nil {
		return 0, err
	}
	n, err := strconv.ParseInt(strings.Trim(string(buf), "\n "), 10, 32)
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func getOffline() (int, error) {
	return readCPURangeWith("offline", countCPURange)
}

func getOnline() (int, error) {
	if n, err := getFromCPUAffinity(); err == nil {
		return n, nil
	}
	return readCPURangeWith("online", countCPURange)
}

func getPossible() (int, error) {
	return readCPURangeWith("possible", countCPURange)
}

func getPresent() (int, error) {
	return readCPURangeWith("present", countCPURange)
}
