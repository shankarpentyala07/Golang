package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Single input
	var name string
	fmt.Printf("Enter Your Name:")
	fmt.Scanf("%s", &name)
	fmt.Println("Hey there", name)

	// Multiple Inputs
	var fname string
	var hispanicStr string
	fmt.Println("Enter your first name & say true for hispanic or else false")
	fmt.Scanf("%s %s", &fname, &hispanicStr)
	hispanic, err := strconv.ParseBool(hispanicStr)
	if err != nil {
		fmt.Println("Error parsing hispanic value:", err)
		return
	}
	fmt.Println("The person's first name is", fname, "and is hispanic:", hispanic)

	// Scanf Return values
	var lname string
	var age int
	fmt.Println("Enter your lname and age")
	count, err := fmt.Scanf("%s %d", &lname, &age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("lname:", lname)
	fmt.Println("age:", age)
	fmt.Println("Count:", count)
}
