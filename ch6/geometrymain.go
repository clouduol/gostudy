// Geometrymain tests package geometry
package main

import (
	"fmt"
	"geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance())

	p.ScaleBy(2.0)
	(&q).ScaleBy(2.0)
	fmt.Println(p)
	fmt.Println(q)
}
