// Using interfaces to calculate the area of different shapes

package main

import (
	"fmt"
	"math"
)

// Shape interface defines a method to calculate area
type Shape interface {
	Area() float64
}

// Circle struct represents a circle with a radius
type Circle struct {
	Radius float64
}

// Area method calculates the area of the circle
func (c Circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}	

// Rectangle struct represents a rectangle with width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method calculates the area of the rectangle
func (r Rectangle) Area() float64{
	return r.Width * r.Height
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Calculate and print the area of the circle
	fmt.Printf("Area of Circle: %.2f\n", calculateArea(circle))

	// Calculate and print the area of the rectangle
	fmt.Printf("Area of Rectangle: %.2f\n", calculateArea(rectangle))
}
