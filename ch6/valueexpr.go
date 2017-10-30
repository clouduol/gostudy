// Valueexpr demonstrates method value and method expression
package main

import (
	"fmt"
	"geometry"
)

func main() {
	fmt.Println("method expression")
	var path geometry.Path
	path = append(path, geometry.Point{1, 2})
	path = append(path, geometry.Point{3, 4})
	path = append(path, geometry.Point{9, 7})
	fmt.Println(path)
	TranslateBy(path, geometry.Point{1, 1}, true)
	fmt.Println(path)
	TranslateBy(path, geometry.Point{1, 1}, false)
	fmt.Println(path)

	fmt.Println("method value")
	p := geometry.Point{1, 2}
	q := geometry.Point{9, 10}
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q))
	var origin geometry.Point
	fmt.Println(distanceFromP(origin))
}

func TranslateBy(path geometry.Path, offset geometry.Point, add bool) {
	var op func(geometry.Point, geometry.Point) geometry.Point
	if add {
		op = geometry.Point.Add
	} else {
		op = geometry.Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}
