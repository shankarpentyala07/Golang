package main
import "fmt"

func main(){
var s string  = "Hello World"
var i int  = 10
var f float64  = 4.69
var b bool = true
var fname = "shankar"
var lname = "pentyala"
fmt.Print(fname,"\n")
fmt.Print(lname, "\n")
fmt.Println(s , i)
fmt.Println(i)
fmt.Println(f)
fmt.Println(b)

// printf - Function for printing formatted i/p & o/package
fmt.Printf("My first name is %v" , fname)
fmt.Println()
fmt.Printf("The value of i is %d", i)
fmt.Println()
fmt.Printf("The Jersey number for %v is %d",fname,i)

//Short Hand Declaration of same data types

var s1,t1 string = "siva" , "kumar"
fmt.Println(s1)
fmt.Println(t1)

// Short Hand Declaration for different data types

var (
	s2 string = "sam"
	i3 int = 100
)
fmt.Println(s2)
fmt.Println(i3)


//Short Variable Declaration
 s4 := "Hello World"
 fmt.Println(s4)

 s5 := "Cristiano"
 s5 = "Ronaldo"
 fmt.Println(s5)


}
