package basic

import (
	"fmt"
	"testing"
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

