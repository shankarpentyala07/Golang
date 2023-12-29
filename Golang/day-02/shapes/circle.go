package shapes

type Circle struct {
	Radius float64
}

const pi = 3.14

//exported function
func(s Circle) AreaOfCircle() float64 {
	return pi * s.Radius* s.Radius
}

//unexported function
func(s Circle) permieterOfSqaure() float64 {
	return 4 *s.Radius
}