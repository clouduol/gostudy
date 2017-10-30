// SameName demonstrates what happens when tow functions have the same name
package main

import "fmt"

type add int

func add(a, b int) int {
	return a + b
}

//func add(a int) int {
//	return a + 6
//}
//
//func add(a, b int) {
//	fmt.Printf("%d\n", a+b)
//}

func main() {
	a := 5
	b := 9
	c := add(a, b)
	d := add(a)
	fmt.Printf("%d\n", c)
	fmt.Printf("%d\n", d)
	add(c, d)
}
