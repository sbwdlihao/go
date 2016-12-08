package basic

import (
	"testing"
	"encoding/binary"
	"fmt"
	"bytes"
	"encoding/hex"
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

func TestEncoding1(t *testing.T) {
	var pi uint32
	b := []byte{0x01, 0x02, 0x03, 0x04}
	buf := bytes.NewReader(b)
	binary.Read(buf, binary.LittleEndian, &pi) // 0x04030201 = 67305985
	fmt.Println(pi)
	var pi1 uint32
	b1 := []byte{0x01, 0x02, 0x03, 0x04} // 0x01020304 = 16909060
	buf1 := bytes.NewReader(b1)
	binary.Read(buf1, binary.BigEndian, &pi1)
	fmt.Println(pi1)

	b2, _ := hex.DecodeString("000000004f2ea239532b2e77bb46c03b86643caac3fe92959a31fd2d03979c34")
	fmt.Println(b2)
	HashSize := 32
	for j := 0; j < HashSize/8; j++ {
		i := j * 4
		b2[i], b2[i+1], b2[i+2], b2[i+3], b2[HashSize-4-i], b2[HashSize-3-i], b2[HashSize-2-i], b2[HashSize-1-i] = b2[HashSize-4-i], b2[HashSize-3-i], b2[HashSize-2-i], b2[HashSize-1-i], b2[i], b2[i+1], b2[i+2], b2[i+3]
	}
	fmt.Println(hex.EncodeToString(b2[:])) // 03979c349a31fd2dc3fe929586643caabb46c03b532b2e774f2ea23900000000
}

