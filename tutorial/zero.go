package main

import (
	"fmt"
)

func main() {
	var i int
	var b bool
	var f float64
	var s string
	fmt.Printf("%v %v %v %v\n", i, b, f, s)

	i0 := 3
	var f0 float64 = float64(i0)
	fmt.Println(f0)
}

