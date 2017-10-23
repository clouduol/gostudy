// Graph depicts map as map value
package main

import (
	"fmt"
)

var graph = make(map[string]map[string]bool)

func main() {
	fmt.Println("add edge (a,b)")
	fmt.Println("add edge (a,c)")
	addEdge("a", "b")
	addEdge("a", "c")
	fmt.Printf("has (a,c): %t\n", hasEdge("a", "c"))
	fmt.Printf("has (b,c): %t\n", hasEdge("b", "c"))

	fmt.Println("test ptr")
	a := make(map[string]int)
	a["first"] = 1
	a["second"] = 2
	b := &a
	fmt.Printf("%T\t%d\t%v\n", a, a["first"], a)
	//fmt.Printf("%T\t%d\t%v\n", b, b["first"], b)
	fmt.Printf("%T\t%d\t%v\n", b, (*b)["first"], b)

	fmt.Println("test map param")
	testParam(a)
	fmt.Printf("%v\n", a)
	testParam2(b)
	fmt.Printf("%v\n", *b)

}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func testParam(m map[string]int) {
	m["first"] = 100
}

func testParam2(m *map[string]int) {
	(*m)["second"] = 200
}
