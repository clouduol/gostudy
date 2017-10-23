// Noempty is an example of an in-place slice algorithm
package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty(data))
	fmt.Printf("%q\n", data)

	ptr := &data
	fmt.Printf("%q\t%T %q\n", data, data, data[0])
	//fmt.Printf("%q\t%T %q\n", ptr, ptr, ptr[0])
	fmt.Printf("%q\t%T\n", ptr, ptr)

	fmt.Println("test param")
	testParam(data)
	fmt.Printf("%q\n", data)
	testParam2(ptr)
	fmt.Printf("%q\n", (*ptr))

	fmt.Println("test empty slice")
	var x []int
	fmt.Printf("%t\t%d\t%d\n", x == nil, len(x), cap(x))
	x = append(x, 3)
	fmt.Printf("%d\n", x)
	y := []int{}
	fmt.Printf("%t\t%d\t%d\n", y == nil, len(y), cap(y))
	y = append(y, 5)
	fmt.Printf("%d\n", y)

}

func noempty(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func testParam(strings []string) {
	strings[0] = "changed"
}

func testParam2(ptr *[]string) {
	(*ptr)[1] = "changed"
}
