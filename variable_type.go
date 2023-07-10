package main
import (
	"fmt"
	"reflect"
)

func main() {

	var no int = 42
	var message string = "Good Day"
	var isCorrect bool = true
	var bill float32 = 420.69

	fmt.Printf("Variable no =  %v is of type %T \n", no ,no)
	fmt.Printf("Variable message =  %v is of type %T \n", message ,message)
	fmt.Printf("Variable isCorrect =  %v is of type %T \n", isCorrect ,isCorrect)
	fmt.Printf("Variable bill =  %v is of type %T \n", bill ,bill)
    
	fmt.Println("Using reflect function")

	fmt.Printf("Variable no =  %v is of type %v \n", no ,reflect.TypeOf(no))
	fmt.Printf("Variable message =  %v is of type %v \n", message ,reflect.TypeOf(message))
	fmt.Printf("Variable isCorrect =  %v is of type %v \n", isCorrect ,reflect.TypeOf(isCorrect))
	fmt.Printf("Variable bill =  %v is of type %v \n", bill ,reflect.TypeOf(bill))
}