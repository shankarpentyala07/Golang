package main
import "fmt"

func main(){
	const age = 12
	const h_name,h_age = "Shankar", 26

	fmt.Printf("%v : %T \n", age,age)
	fmt.Printf("%v : %T \n", h_age,h_age)
	fmt.Printf("%v : %T \n", h_name,h_name)



	const name string = "sam kuran"
	const age1 int = 22
	fmt.Printf("%v : %T \n", name,name)
	fmt.Printf("%v : %T \n", age1,age1)

}