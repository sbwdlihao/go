package basic

import (
	"testing"
	"encoding/binary"
	"fmt"
)

func TestEncoding0(t *testing.T)  {
	var l [10]byte
	binary.PutUvarint(l[:], 0) // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(l)
	var l1 [10]byte
	binary.PutUvarint(l1[:], 1234) // [210 9 0 0 0 0 0 0 0 0]
	fmt.Println(l1)
	binary.PutUvarint(l1[:], 0xff) // [255 1 0 0 0 0 0 0 0 0]
	fmt.Println(l1)
	binary.PutUvarint(l1[:], 256) // [128 2 0 0 0 0 0 0 0 0]
	fmt.Println(l1)
	binary.PutUvarint(l1[:], 0x7fffffffffffffff) // [255 255 255 255 255 255 255 255 127 0]
	fmt.Println(l1)
	binary.PutUvarint(l1[:], 0xffffffffffffffff) // [255 255 255 255 255 255 255 255 255 1]
	fmt.Println(l1)
}

