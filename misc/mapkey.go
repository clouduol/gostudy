// Mapkey demonstrates map with key "" and nil, etc.
package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["first"] = 1
	m["second"] = 2
	m[""] = 0
	fmt.Println(m)

	m2 := make(map[*int]int)
	m2[nil] = 0
	first := 1
	m2[&first] = 1
	fmt.Println(m2)
}
