package numcpus

import "testing"

func TestParsCPURange(t *testing.T) {
	testCases := []struct {
		cpus string
		n    int
	}{
		{"", 0},
		{"0", 1},
		{"0-1", 2},
		{"0-7", 8},
		{"1-7", 7},
		{"1-15", 15},
		{"0,2-4,7", 5},
		{"0,2-4,7-15", 13},
	}

	for _, tc := range testCases {
		n, err := parseCPURange(tc.cpus)
		if err != nil {
			t.Errorf("failed to parse CPU range: %v", err)
		}

		if n != tc.n {
			t.Errorf("parseCPURange(%s) = %d, expected %d", tc.cpus, n, tc.n)
		}
	}
}
