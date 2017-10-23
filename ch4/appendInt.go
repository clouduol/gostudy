// AppendInt shows slice append operation
package main

import (
	"fmt"
)

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	x = appendInt(x, y...)
	fmt.Printf("%d\n", x)

	a := []int{2: 3, 5: 4}
	fmt.Printf("%d\n", a)
	a = append(a, 8)
	fmt.Printf("%d\n", a)
	a = append(a, 2, 3, 4)
	fmt.Printf("%d\n", a)
	a = append(a, a...)
	fmt.Printf("%d\n", a)
}

func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}
