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

import "testing"

func TestCountCPURange(t *testing.T) {
	testCases := []struct {
		str     string
		n       int
		wantErr bool
	}{
		{str: "", n: 0},
		{str: "0", n: 1},
		{str: "0-1", n: 2},
		{str: "1-1", n: 1},
		{str: "0-7", n: 8},
		{str: "1-7", n: 7},
		{str: "1-15", n: 15},
		{str: "0-3,7", n: 5},
		{str: "0,2-4", n: 4},
		{str: "0,2-4,7", n: 5},
		{str: "0,2-4,7-15", n: 13},
		{str: "0,2-4,6,8-10", n: 8},
		{str: ",", wantErr: true},
		{str: "invalid", n: 0, wantErr: true},
		{str: "-", wantErr: true},
		{str: ",", wantErr: true},
		{str: ",1", wantErr: true},
		{str: "0,", wantErr: true},
		{str: "0-", wantErr: true},
		{str: "0,2-", wantErr: true},
		{str: "0-,1", wantErr: true},
		{str: "0,-3,5", wantErr: true},
		{str: "42-0", wantErr: true},
		{str: "0,5-3", wantErr: true},
	}

	for _, tc := range testCases {
		n, err := countCPURange(tc.str)
		if !tc.wantErr && err != nil {
			t.Errorf("countCPURange(%q) = %v, expected no error", tc.str, err)
		} else if tc.wantErr && err == nil {
			t.Errorf("countCPURange(%q) expected error", tc.str)
		}

		if n != tc.n {
			t.Errorf("countCPURange(%q) = %d, expected %d", tc.str, n, tc.n)
		}
	}
}

func TestGetFromCPUAffinity(t *testing.T) {
	nAffinity, err := getFromCPUAffinity()
	if err != nil {
		t.Fatalf("getFromCPUAffinity: %v", err)
	}

	cpus := "online"
	nSysfs, err := readCPURangeWith(cpus, countCPURange)
	if err != nil {
		t.Fatalf("counting CPU ranges from %q failed: %v", cpus, err)
	}

	if nAffinity != nSysfs {
		t.Errorf("getFromCPUAffinity() = %d, readCPURange(%q) = %d, want the same return value", nAffinity, cpus, nSysfs)
	}
}
