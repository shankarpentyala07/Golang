package main
import (
	"fmt"
	"strconv"

)

func  main (){
	var i int = 22
	var f float64 = float64(i)
	fmt.Printf("The variable f is  %.2f \n", f)

	var f1 float64 = 42.32
	var i1 int = int(f1)
	fmt.Printf("The variable i1 is %v \n", i1)

	//using String Conversion package

	var i3 int = 79
	fmt.Printf("The String is %q \n", strconv.Itoa(i3))

	var fname string = "42"
	i4,err := strconv.Atoi(fname)
	fmt.Printf("The integer is %v  and the conversion error is %v \n",i4, err )


}