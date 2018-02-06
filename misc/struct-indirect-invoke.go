package main

import (
	"fmt"
)

type LowerStruct struct {
	secondname string
}

type UpperStruct struct {
	firstname string
	LowerStruct
}

func (l LowerStruct) PrintName() {
	fmt.Println("l.secondname")
}

func main() {
	low := LowerStruct{"second name"}
	up := UpperStruct{"first name", low}

	fmt.Printf("low: %v\n", low)
	fmt.Printf("up: %v\n", up)

	low.PrintName()
	up.PrintName()

	up.secondname = "last name"
	fmt.Printf("low: %v\n", low)
	fmt.Printf("up: %v\n", up)
	low.PrintName()
	up.PrintName()
}
