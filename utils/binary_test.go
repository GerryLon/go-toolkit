package utils

import (
	"testing"
)

func TestIsLittleEndian(t *testing.T) {
	// 操作系统一般是小端存储， 网络传输一般是大端存储
	if !IsLittleEndian() {
		t.Error("Your OS not use little endian")
	}
}
