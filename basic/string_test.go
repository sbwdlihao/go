package basic

import (
	"testing"
	"fmt"
)

func TestString0(t *testing.T)  {
	a := "a"
	a1 := "a"
	b := "b"
	fmt.Println(a == a1) // true
	fmt.Println(a == b) // false
}
