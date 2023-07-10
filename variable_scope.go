package main
import "fmt"
var country string = "usa"
func main (){
	city := "London"
	{
		street := "awesome"
		fmt.Println("Name of", city , " street is" , street)
	}
	fmt.Println("Outer Block City is ", city)

	fmt.Println("The country name is ", country)
}