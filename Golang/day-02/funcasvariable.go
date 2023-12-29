/*
Functions can be assigned to variables. For instance var f func(int) int declares
f as a variable that can hold a function taking an int returning an int
*/

package main

import "fmt"

func main(){
var f func(int) int
f = func(x int)int { return x*x}
fmt.Println(f(5))
}