package main

import "fmt"

type Employee struct {
	name string
	id int
	salary float64
}

func updateSalary(e *Employee, ns float64){
	e.salary = ns
}

func main(){
	emp1 := Employee{
		name: "Shankar",
		id : 1,
		salary: 100000,
	}
	fmt.Println("Before update:", emp1)
	updateSalary(&emp1,230000)
	fmt.Println("After update:",emp1)
}