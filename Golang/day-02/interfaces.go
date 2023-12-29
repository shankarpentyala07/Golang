/*
Interfaces define set of methods.
A type implements an interface by implementing its methods.
Interfaces are central to Go's type system and polymorphism
*/

package main

import "fmt"

type Geometry interface{
	Area() float64
	Perimeter() float64
}

type Rectangle struct{
	length,width float64
}

type Square struct{
	length float64
}

func(s Square) Area() float64{
	return s.length * s.length
}

func (s Square) Perimeter() float64{
	return 4*s.length
}

func (r Rectangle) Area() float64{
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.length + r.width)
}

func main(){
	var w Geometry= Rectangle{10,20}
	area := w.Area()
	perimeter := w.Perimeter()

	fmt.Println("Area:",area)
	fmt.Println("permieter:",perimeter)
	var s Geometry = Square{4}
	area = s.Area()
	perimeter = s.Perimeter()

	fmt.Println("Area:",area)
	fmt.Println("permieter:",perimeter)


}