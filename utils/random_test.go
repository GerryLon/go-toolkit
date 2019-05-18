package utils

import (
	"fmt"
	"github.com/GerryLon/go-toolkit/consts"
	"strings"
	"testing"
	"time"
)

func TestRandInt(t *testing.T) {
	if RandInt(0, 0) != 0 {
		t.Error("RandInt(0, 0) != 0")
	}

	for i := 0; i < 100; i++ {
		time.Sleep(time.Nanosecond)
		if r := RandInt(0, 10); r < 0 || r > 10 {
			t.Error("RandInt(0, 10) not in [0, 10], r=", r)
		} else {
			fmt.Println("r=", r)
		}
	}
}

func TestRandStrCustom(t *testing.T) {
	min := 1
	max := 7
	randStr := RandStrCustom(consts.AllLettersAndNumbers, min, max)
	length := len(randStr)
	if length < min || length > max {
		t.Errorf("RandStrCustom(%s, %d, %d) length error\n", consts.AllLettersAndNumbers, min, max)
	}
	for i := 0; i < length; i++ {
		if strings.Index(consts.AllLettersAndNumbers, randStr[i:i+1]) < 0 {
			t.Error("unexpected char")
		}
	}
}
