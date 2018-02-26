package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown, Press return to abort")
	tick := time.Tick(1 * time.Second)
	for countdown := 1; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:

		case <-abort:
			fmt.Println("Launch Aborted!")
			return
		}
	}

	fmt.Println("Launch!!!")

	// random select if multiple eligible
	case1 := make(chan int, 1)
	case2 := make(chan int, 1)
	case1count := 0
	case2count := 0
	defaultcount := 0

	for count := 10; count > 0; count-- {
		select {
		case case1 <- 1:
			fmt.Println("case1")
			<-case1
			case1count++
		case case2 <- 2:
			fmt.Println("case2")
			<-case2
			case2count++
		default:
			fmt.Println("default")
			defaultcount++
		}
	}
	fmt.Printf("case1 count: %d\n", case1count)
	fmt.Printf("case2 count: %d\n", case2count)
	fmt.Printf("default count: %d\n", defaultcount)
}
