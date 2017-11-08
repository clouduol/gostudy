// Valueaddr demonstrates if value addr exists
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

func retArrEm() [3]int {
	arr := [3]int{4, 5, 6}
	return arr
}

func retArr(arr [3]int) [3]int {
	arr[0] = 0
	return arr
}

func retStrEm() Point {
	str := Point{9, 0}
	return str
}

func retStr(str Point) *Point {
	str.y = 10
	return &str
}

func main() {
	fmt.Println("--- initialization ----")
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr[0])
	fmt.Println((&arr)[1])

	var one = [3]int{1, 2, 3}[0]
	fmt.Println(one)

	var x = Point{2, 4}.x
	fmt.Println(x)

	fmt.Println("--- function ----")
	third := retArrEm()[0]
	fmt.Println(third)
	y := retStrEm().y
	fmt.Println(y)
	zero := retArr(arr)[0]
	fmt.Println(zero)
	fmt.Println(arr)
	y = retStr(Point{1, 2}).y
	fmt.Println(y)

	fmt.Println("--- initial & function return value assignment----")
	//[3]int{1, 2, 3}[0] = 10
	//&[3]int{1, 2, 3}[0] = 10
	var p = &[3]int{1, 2, 3}
	p[0] = 10
	// return pointer value, assignment is right, but value is wrong---- cannt get address
	retStr(Point{1, 2}).y = 100
}
