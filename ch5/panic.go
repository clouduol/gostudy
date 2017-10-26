// Panic demonstrates panic exceptions
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var m map[string]int
	fmt.Printf("%T\t%v\n", m, m)
	//m["panic"] = 1
	m = map[string]int{}
	m["panic"] = 1

	fmt.Println("---------")
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
