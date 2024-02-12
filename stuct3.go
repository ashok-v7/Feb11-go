package main

import (
	"fmt"
)

// METHOD RECEVIER

// Circle represents a circle with a position (x,y), a radius, and an area.
type Circle struct {
	x      int
	y      int
	radius float64
	area   float64
}

// calcArea calculates the area of the circle and updates the area field.
func (c *Circle) calcArea() {
	const PI float64 = 3.14
	// Directly update the area of the circle.
	c.area = PI * c.radius * c.radius
}

func main() {
	// Initialize a Circle instance.
	c := Circle{x: 5, y: 5, radius: 5}
	fmt.Printf("Before calculating area: %+v \n", c)
	// Calculate the area of the circle.
	c.calcArea()
	fmt.Printf("After calculating area: %+v \n", c)
}
