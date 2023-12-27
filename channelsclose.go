package main

import "fmt"

func digits(number int, dchn1 chan int) {
	for number != 0 {
		dchn1 <- number % 10
		number = number / 10
	}
	close(dchn1)
}

func calcSquares(num int, squareop chan int) {
	sum := 0
	dchn1 := make(chan int)
	go digits(num, dchn1)
	for digit := range dchn1 {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(num int, cubeop chan int) {
	sum := 0
	dchn1 := make(chan int)
	go digits(num, dchn1)
	for digit := range dchn1 {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 123
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
