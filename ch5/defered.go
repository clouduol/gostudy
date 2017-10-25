// Defered demonstrates defer statement features
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("---- 1 ----")
	bigSlowOperation()
	fmt.Println("---- 2 ----")
	bigSlowOperation2()
	fmt.Println("---- 3 ----")
	fmt.Printf("%d\n", triple(10))
}

func bigSlowOperation() {
	fmt.Println("before defer")
	defer trace("bigSlowOperation")()
	fmt.Println("after defer")
	time.Sleep(5 * time.Second)
}

func bigSlowOperation2() {
	fmt.Println("before defer")
	defer trace("bigSlowOperation")
	fmt.Println("after defer")
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s\n", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
}
