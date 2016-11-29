package basic

import (
	"testing"
	"golang.org/x/crypto/sha3"
	"fmt"
)

func TestCrypto0(t *testing.T) {
	h := sha3.New256()
	h.Write([]byte{0,1,2})
	data := h.Sum(nil)
	fmt.Println(data)
}

