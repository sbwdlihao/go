package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	V1 = Vertex{1, 2}
	V2 = Vertex{X: 1}
	V3 = Vertex{}
	V4 = &Vertex{3, 4}
)

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(v)

	p := &v
	p.Y = 10
	fmt.Println(*p)

	fmt.Println(V1, V2, V3, V4)
}
