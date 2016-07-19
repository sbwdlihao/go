package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum - 1
	y = sum - 2
	return
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))
}