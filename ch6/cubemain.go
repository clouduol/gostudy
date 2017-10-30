// Heightpointmain test struct Cube
package main

import (
	"fmt"
	"geometry"
)

func main() {
	var c geometry.Cube
	c.X = 1
	c.Y = 2
	c.Z = 3
	fmt.Println(c)
	c.Multiple(2.0)
	fmt.Println(c)
	//c.ScaleBy(2.0)
	c.Point.ScaleBy(2.0)
	fmt.Println(c)
}
