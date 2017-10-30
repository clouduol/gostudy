// Cache demonstrates anonymous struct with method
package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func main() {
	cache.mapping["first"] = "1"
	cache.mapping["second"] = "2"

	fmt.Println(cache)
	fmt.Println(Lookup("first"))
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
