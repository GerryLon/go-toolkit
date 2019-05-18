package utils

import (
	"github.com/GerryLon/go-toolkit/consts"
	"math/rand"
	"time"
)

func RandInt(min, max int) (result int) {
	if min >= max {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RandStrCustom(source string, minLen, maxLen int) (result string) {
	sourceLen := len(source)
	targetLen := RandInt(minLen, maxLen)
	if sourceLen == 0 || sourceLen < targetLen {
		return source
	}
	return source[0:targetLen]
}

func RandStr(minLen, maxLen int) (result string) {
	return RandStrCustom(consts.AllLettersAndNumbers, minLen, maxLen)
}
