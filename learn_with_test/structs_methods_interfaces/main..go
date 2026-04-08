package main
// import "math"

const pi = 3.142
type Rectangle struct{
	Height float64
	Width float64
}

type Circle struct{
	radius float64
}
func Perimeter(r Rectangle)float64{
	return 2*(r.Height+r.Width)
}

func Area(r Rectangle)float64{
	return r.Height*r.Width
}

func CirlePeri(r Circle)float64{
	return 2*pi*r.radius
}
func CirleArea(r Circle)float64{
	return  2*pi*r.radius*r.radius
}