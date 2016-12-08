package basic

import (
	"fmt"
	"testing"
	"encoding/hex"
	"bytes"
)

// how to reset a slice
func TestSlice0(t *testing.T) {
	bytes := make([]byte, 8)
	fmt.Println(len(bytes)) // 8
	fmt.Println(cap(bytes)) // 8
	bytes = bytes[:0]
	fmt.Println(len(bytes)) // 0
	fmt.Println(cap(bytes)) // 8
}

func TestSlice1(t *testing.T) {
	// 没办法像python那样一次性初始化所有的元素为同一个值，除了0以外
	a := [3]byte{0xEE}
	fmt.Println(a) // 238 0 0
}

func TestSlice2(t *testing.T) {
	a, _ := hex.DecodeString("010203eeee0405eeee")
	fmt.Println(bytes.Index(a, []byte{0xee, 0xee})) // 3
	fmt.Println(bytes.LastIndex(a, []byte{0xee, 0xee})) // 7
}

