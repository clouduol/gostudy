// TestAssign tests assign for different structures
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Circle struct {
	x, y, r int
}

func main() {
	p := Point{1, 2}
	c := Circle{3, 4, 1}
	fmt.Printf("%#v\n", p)
	fmt.Printf("%#v\n", c)
	a := p
	fmt.Printf("%#v\n", a)
	// a = c
	a = Point{7, 8}
	fmt.Printf("%#v\n", a)

	fmt.Printf("test nil\n")
	//var arr [3]int
	sli := []int{}
	mp := map[string]int{}
	//var st Point
	//fmt.Printf("%t\n", arr == nil)
	fmt.Printf("%t\n", sli == nil)
	fmt.Printf("%t\n", mp == nil)
	//fmt.Printf("%t\n", st == nil)
}
