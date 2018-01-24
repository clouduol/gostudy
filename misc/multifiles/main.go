package main

import "fmt"

const foobar string = "foobar"

func main() {
	fmt.Println("--- test multiple files in one package ---")
	bar("success")
	fmt.Println(foo)
}
