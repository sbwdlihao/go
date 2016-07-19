package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
	v.Scale(2)
	fmt.Println(v)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
