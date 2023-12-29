package shapes

type Square struct {
	Length float64
}

//exported function
func(s Square) AreaOfSquare() float64 {
	return s.Length* s.Length
}

//unexported function
func(s Square) permieterOfSqaure() float64 {
	return 4 *s.Length
}