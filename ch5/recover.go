// Recover demonstrates recover features
package main

import (
	"fmt"
)

func main() {
	err := recoverFunc()
	fmt.Printf("%d\n", err)
}

func recoverFunc() (err int) {
	defer func() {
		if p := recover(); p != nil {
			err = 1
		} else {
			err = 0
		}
	}()
	x := 0
	y := 1
	z := y / x
	fmt.Printf("%f\n", z)
	panic("division by zero")
}
