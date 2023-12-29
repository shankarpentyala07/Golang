package main

import "fmt"

func main(){
	 i := 10
	 var p *int = &i
	 fmt.Println(p)
	 fmt.Println(*p)
}

/* output
0xc0000120b0
10
*/