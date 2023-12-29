package main

import (
	"day2/shapes"
	"fmt"
)

func main() {
	c := shapes.Circle{4}
	fmt.Println("Area of Circle: ", c.AreaOfCircle())

	s := shapes.Square{5}
	fmt.Println("Area of Square:", s.AreaOfSquare())
}
