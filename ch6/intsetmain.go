// Intsetmain test package bit
package main

import (
	"bit"
	"fmt"
)

func main() {
	var x, y bit.IntSet
	x.Add(1)
	x.Add(8)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(1)
	y.Add(4)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(10))
}
