/*
channels are used for communication b/w go routines(concurrent threads).They can be synchronous or asynchronous
c := make(chan int) //unbuffered channel of integers
*/
package main

import (
	"fmt"
	"time"
)

type Order struct{
	id int
	status string
}

func processOrder(id int, orderCh chan Order){
	time.Sleep(200* time.Millisecond)
	orderCh <- Order{
		id: id,
		status: "Completed",
	}
}

func main(){
	orderCh := make(chan Order)

	for i:=0;i<=5;i++{
		go processOrder(i,orderCh)
	}

	for i:=0;i<=5;i++{
		order := <- orderCh
		fmt.Printf("order:  id = %d , status = %s \n", order.id, order.status)
	}
}