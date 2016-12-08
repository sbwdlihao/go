package basic

import (
	"testing"
	"regexp"
	"fmt"
	"encoding/hex"
)

func TestRegex0(t *testing.T) {
	reg := regexp.MustCompile(`\x45{2}`)
	b := []byte{0x45,0x45,0x46}
	b1 := []byte{0x44, 0x45,0x45,0x46}
	b2 := []byte{0x44, 0x45,0x45,0x46, 0x45, 0x45}
	fmt.Println(reg.Find(b)) // [69 69]
	fmt.Println(reg.FindIndex(b)) // [0 2]
	fmt.Println(reg.FindIndex(b1)) // [1 3]
	fmt.Println(reg.FindIndex(b2)) // [1 3]
	fmt.Println(reg.FindAllIndex(b2, -1)) // [[1 3] [4 6]]
	fmt.Println(reg.FindAllIndex(b2, 0)) // []
	fmt.Println(reg.FindAllIndex(b2, 1)) // [[1 3]]
	fmt.Println(reg.FindAllIndex(b2, 2)) // [[1 3] [4 6]]

	b3, _ := hex.DecodeString("444545464545")
	fmt.Println(reg.FindAllIndex(b3, -1))

	// 貌似只能对ascii可显示字符进行匹配
	reg1 := regexp.MustCompile(`\x80{2}`)
	b4, _ := hex.DecodeString("8080")
	fmt.Println(reg1.FindAllIndex(b4, -1)) // []
}
