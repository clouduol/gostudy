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
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//Printer(in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}

}
