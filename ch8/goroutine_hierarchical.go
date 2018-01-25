package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("leave func main")
	fmt.Println("enter func main")
	go level1()
	time.Sleep(1 * time.Second)
	return
}

func level1() {
	defer fmt.Println("leave func level1")
	fmt.Println("enter func level1")
	go level2()
	time.Sleep(2 * time.Second)
	return
}

func level2() {
	defer fmt.Println("leave func level2")
	fmt.Println("enter func level2")
	time.Sleep(3 * time.Second)
}
