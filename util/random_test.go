package util

import "testing"

func TestRandInt(t *testing.T) {
	if RandInt(0, 0) != 0 {
		t.Error("RandInt(0, 0) != 0")
	}

	for i := 0; i < 100000; i++ {
		if r := RandInt(0, 10); r < 0 || r > 10 {
			t.Error("RandInt(0, 10) not in [0, 10], r=", r)
		}
	}
}
