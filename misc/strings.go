package main

import (
	"fmt"
	//"strings"
)

func main() {
	str := "This is a test"
	tail := "test"
	subStr := str[:len(str)-len(tail)]
	fmt.Printf("%T\n", subStr)
	fmt.Printf("%s\n", subStr)
}
