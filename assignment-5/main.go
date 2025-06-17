package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	width, height float64
}

// making method Area for Circle
func (c Circle) Area() float64 {
	ans := math.Pi * c.radius * c.radius
	//fmt.Printf("Area of Circle is : %f", ans)
	return ans
}

// method Area for Rectangle
func (r Rectangle) Area() float64 {
	ans := r.width * r.height
	return ans
}
func main() {
	fmt.Println("Simple Interface Implementation")

	var radius float64
	var width, height float64
	fmt.Println("Enter radius for circle: ")
	fmt.Scan(&radius)
	fmt.Println("Enter height for rectangle: ")
	fmt.Scan(&height)
	fmt.Println("Enter Width for rectangle: ")
	fmt.Scan(&width)

	var s Shape
	c := Circle{radius}
	s = c
	PrintArea(s)

	r := Rectangle{width, height}
	s = r
	PrintArea(s)

}
func PrintArea(s Shape) {
	fmt.Println("Area of Circle: ", s.Area())
}
