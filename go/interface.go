package main

import "fmt"
import "math"

type Point struct {
	X, Y float64
}

type intt int

// define method on intt similar to:
// class intt
//   def negative
//     ...
//   end
// end
func (self intt) negative() bool {
	var result bool = false

	if self < 0 {
		result = true
	}

	return result
}

// func name return_type { body }
// methods have receiver in front, similar to python's
// def self.__init__(self)
// func (receiver) name return_type { body }
func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// interface only applies to methods defined
func (p *Point) NotAbs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

type Magnitude interface {
	Abs() float64
}

func main() {
	x := &Point{3, 4}

	var m Magnitude

	// can do this since both m & x implement Magnitude
	// interface
	m = x

	fmt.Println(m)

	var integer1 intt = 0
	fmt.Println(integer1.negative())

	var integer2 intt = -1
	fmt.Println(integer2.negative())
}
