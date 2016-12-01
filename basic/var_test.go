package basic

import (
	"testing"
	"fmt"
)

func TestVar0(t *testing.T) {
	fmt.Println(byte(0)) // 0
	fmt.Println(byte(255)) // 255
	//fmt.Println(byte(256)) constant 256 overflows byte
	//fmt.Println(byte(257)) constant 257 overflows byte
	//fmt.Println(byte(-1)) constant -1 overflows byte
	var a uint64
	a = 256
	fmt.Println(byte(a)) // 0
	a = 257
	fmt.Println(byte(a)) // 1
	var b int64
	b = -1
	fmt.Println(byte(b)) // 255

	// 围绕1个256个刻度时钟，起点是0，终点是255
}

func TestVar1(t *testing.T) {
	one := [32]byte{1} // [1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(one)
}