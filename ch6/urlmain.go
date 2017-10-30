// Urlmain test package url
package main

import (
	"fmt"
	"url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("q"))
	fmt.Println(m.Get("item"))
	fmt.Println(m["item"])

	var mp = make(url.Values)
	fmt.Println(mp["test"])
	mp["test"] = []string{"1", "2"}
	fmt.Println(mp["test"])

	m = nil
	if m == nil {
		fmt.Println("nil")
	}
	fmt.Println(m)
	fmt.Println(m.Get("item"))
	//	m.Add("item", "3")

	var sl []int
	if sl == nil {
		fmt.Println("nil")
	}
	fmt.Println(sl)
	sl = append(sl, 3)
	fmt.Println(sl)
}
