// PopcountTest test package popcount
package main

import (
	"fmt"
	"popcount"
)

var testData [5]uint64
var testData2 [5]uint64

func init() {
	for i, _ := range testData {
		testData[i] = uint64(10*i + 5)
	}
}

func init() {
	for i, _ := range testData2 {
		testData2[i] = uint64(2*i + 1)
	}
}

func main() {
	for _, data := range testData {
		fmt.Printf("%d\t%d\n", data, popcount.PopCount(uint64(data)))
	}
	for _, data := range testData2 {
		fmt.Printf("%d\t%d\n", data, popcount.PopCount(uint64(data)))
	}
}
