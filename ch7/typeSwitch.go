// TypeSwitch demonstrates interface type switch
package main

import (
	"fmt"
)

type Point struct{ X, Y, z int }

func (p *Point) get(a string) int {
	switch a {
	case "x":
		return p.X
	case "y":
		return p.Y
	case "z":
		return p.z
	default:
		return 0
	}
}

type Getter interface {
	get(string) int
}

func main() {
	var ei Getter
	ei = &Point{1, 2, 3}
	fmt.Printf("ei: %T\n", ei)
	fmt.Printf("X: %d,\tY: %d\tz: %d\n", ei.get("x"), ei.get("y"), ei.get("z"))
	//fmt.Printf("X: %d,\tY: %d\tz: %d\n", ei.X, ei.get("y"), ei.get("z"))

	p := ei.(*Point)
	fmt.Printf("p: %T\n", p)
	fmt.Printf("X: %d,\tY: %d\tz: %d\n", p.X, p.Y, p.z)

	fmt.Println("--- type switch --- ")
	interOrConcre(nil)
	interOrConcre(89)
	interOrConcre("hello world")
	interOrConcre(Point{1, 2, 3})
}

func interOrConcre(v interface{}) {
	switch x := v.(type) {
	case nil:
		fmt.Printf("nil\n")
	case int:
		fmt.Printf("%d\n", x)
	case string:
		fmt.Printf("%q\n", x)
	case Point:
		fmt.Printf("x.get(\"x\"): %d\t", x.get("x"))
		// x.X is effective, so x is the result of type assertion
		fmt.Printf("x.X: %d\n", x.X)
	default:
		fmt.Printf("others\n")
	}
}
