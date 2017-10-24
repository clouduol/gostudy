// TestRet test if return value is a copy
package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

var arr = [...]int{4: 3}
var sli = []int{3: 7}
var mp = map[string]int{"second": 2}
var pt = Point{3, 4}

func testArr() [5]int {
	return arr
}

func testSli() []int {
	return sli
}

func testMp() map[string]int {
	return mp
}

func testPt() Point {
	return pt
}

func main() {
	fmt.Println("initial value")
	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", sli)
	fmt.Printf("%v\n", mp)
	fmt.Printf("%v\n", pt)

	arr1 := testArr()
	sli1 := testSli()
	mp1 := testMp()
	pt1 := testPt()
	arr1[0] = 100
	sli1[0] = 200
	mp1["second"] = 200
	pt1.x = 300

	fmt.Println("changed return value")
	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", sli1)
	fmt.Printf("%v\n", mp1)
	fmt.Printf("%v\n", pt1)

	fmt.Println("original value")
	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", sli)
	fmt.Printf("%v\n", mp)
	fmt.Printf("%v\n", pt)
}
