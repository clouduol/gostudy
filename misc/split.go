// Split demonstrates strings.Split functions
package main

import (
	"fmt"
	"strings"
)

func main() {
	test1 := "nginx:latest"
	test2 := "nginx:"
	test3 := ":nginx:latest"
	fmt.Printf("%q\n%d\n%v\n", test1, len(strings.Split(test1, ":")), strings.Split(test1, ":"))
	fmt.Printf("%q\n%d\n%v\n", test2, len(strings.Split(test2, ":")), strings.Split(test2, ":"))
	fmt.Printf("%q\n%d\n%v\n", test3, len(strings.Split(test3, ":")), strings.Split(test3, ":"))
}
