// Structini demonstrates what happened when using declare a struct
package main

import (
	"fmt"
)

type Point struct{ x, y int }

var DefaultPoint = &defaultPoint
var defaultPoint Point

func main() {
	DefaultPoint.x = 10
	DefaultPoint.y = 20
	fmt.Println(DefaultPoint)

	var point = Point{20, 30}
	fmt.Println(point)

	var m map[string]int
	m["first"] = 1
	fmt.Println(m)
	fmt.Printf("%t\n", m == nil)
	var p = make(map[string]int)
	p["second"] = 2
	fmt.Println(p)
	fmt.Printf("%t\n", p == nil)
	var a = new(map[string]int)
	//a["third"] = 3
	(*a)["third"] = 3
	fmt.Println(a)
	fmt.Printf("%t\n", a == nil)

	//var s []int
	//fmt.Println(s[0])
}
