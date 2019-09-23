// Copyright 2019 Tobias Klauser
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

package numcpus_test

import (
	"runtime"
	"testing"

	"github.com/tklauser/numcpus"
)

func TestNumcpus(t *testing.T) {
	n, err := numcpus.GetOffline()
	if err != nil && runtime.GOOS == "linux" {
		t.Fatalf("GetOffline: %v", err)
	}
	t.Logf("KernelMax = %v", n)

	n, err = numcpus.GetOffline()
	if err != nil && runtime.GOOS == "linux" {
		t.Fatalf("GetOffline: %v", err)
	}
	t.Logf("Offline = %v", n)

	n, err = numcpus.GetOnline()
	if err != nil {
		t.Fatalf("GetOnline: %v", err)
	}
	t.Logf("Online = %v", n)

	n, err = numcpus.GetPossible()
	if err != nil {
		t.Fatalf("GetPossible: %v", err)
	}
	t.Logf("Possible = %v", n)

	n, err = numcpus.GetPresent()
	if err != nil {
		t.Fatalf("GetPresent: %v", err)
	}
	t.Logf("Present = %v", n)
}
