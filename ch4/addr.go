// Addr tests address of anonymous variable
package main

import "fmt"

func addr() string {
	s := "hello world!"
	return s
}

func main() {
	s := addr()
	p := &addr()
	fmt.Println(p)
}
