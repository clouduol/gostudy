// Conversion demostrates some conversions between numbers and strings
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 234
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println(strconv.FormatInt(int64(x), 2))
	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)
	a, err := strconv.Atoi("789")
	if err != nil {
		fmt.Println(a + 1)
	}
	b, err := strconv.ParseInt("789", 10, 64)
	if err != nil {
		fmt.Println(b + 2)
	}
	c := 3.1415
	d := int64(c)
	fmt.Printf("%g\t%d\n", c, d)
	e := false
	f := 1
	fmt.Println(btoi(e))
	fmt.Println(itob(f))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}
