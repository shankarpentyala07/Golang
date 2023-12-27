package main

import "fmt"

func calcSquares(num int, squareop chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num = num / 10
	}
	squareop <- sum

}

func calcCubes(num int, cubeop chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num = num / 10
	}
	cubeop <- sum
}

func main() {
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(123, sqrch)
	go calcCubes(123, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
