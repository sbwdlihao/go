package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)
	s[0] = 4
	fmt.Println(primes)

	s1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(s1)

	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2[:0])
	printSlice(s2[:4])
	printSlice(s2[2:])

	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil")
	}

	b := make([]int, 5)
	printSlice(b)

	c := make([]int, 0, 5)
	printSlice(c)

	d := c[:2]
	printSlice(d)

	e := d[2:5]
	printSlice(e)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}
