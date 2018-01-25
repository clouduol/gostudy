package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	/*
		if naturals == nil {
			fmt.Println("channel naturals is nil")
		} else {
			fmt.Println("channel naturals is not nil")
		}

		m := make(map[string]int)
		if m == nil {
			fmt.Println("map m is nil")
		} else {
			fmt.Println("map m is not nil")
		}
	*/

	squares := make(chan int)

	//Counter
	go counter(naturals)

	//Squarer
	go squarer(naturals, squares)

	//Printer(in main goroutine)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
