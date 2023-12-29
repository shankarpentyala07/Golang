/*
Type Aliases:
Custom names for existing types
Useful for readibility or while transitioning code.For example, type bytedt int64 creates a bytedt as an alias for type int64

*/

package main

import "fmt"

func main(){
	type bytedt int64
	var x bytedt = 10
	fmt.Printf("Type of x is %T \n",x)
	fmt.Printf("value of x is %v",x)

}