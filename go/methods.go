package main

import "fmt"
import "math"

type Point struct {
	X, Y float64
}

func (p *Point) Scale(s float64) {
	p.X *= s
	p.Y *= s
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	x := &Point{3, 4}
	fmt.Println(x.Abs())

	x.Scale(5)
	fmt.Println(x.Abs())
}
