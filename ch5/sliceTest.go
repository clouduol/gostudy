// SliceTest test slice as a parameter
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []string{"first", "second", "third"}
	fmt.Printf("%v %d\n", s, len(s))
	sliceParam(s)
	fmt.Printf("%v %d\n", s, len(s))
	//fmt.Printf("%s\n", s[4])

	m := map[string]int{"first": 1, "second": 2}
	fmt.Printf("%v %d\n", m, len(m))
	mapParam(m)
	fmt.Printf("%d\n", len(m))

	a := []int{1, 2, 3}
	fmt.Printf("%v %d\n", a, len(a))
	b := a
	b = append(b, 4)
	fmt.Printf("%v %d\n", b, len(b))
	fmt.Printf("%v %d\n", a, len(a))
}

func sliceParam(str []string) {
	str = append(str, "tail")
	str[0] = "0"
	fmt.Printf("%v %d\n", str, len(str))
	//return str
}

func mapParam(mp map[string]int) {
	for i := 10; i < 10000; i++ {
		key := "key" + strconv.Itoa(i)
		mp[key] = i
	}
	fmt.Printf("%d\n", len(mp))
}
