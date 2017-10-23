// Treesort sort a array by tree
package main

import (
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

type point struct {
	x, y, z int
}

func main() {
	var values = []int{4, 7, 3, 1, 8, 6, 2, 5}
	fmt.Printf("%v\n", values)
	Sort(values)
	fmt.Printf("%v\n", values)

	fmt.Println("test ptr ret value")
	point1 := point{1, 2, 3}
	point2 := point{x: 10, y: 11, z: 12}
	fmt.Printf("%v\n", point1)
	fmt.Printf("%v\n", point2)
	changeX(&point1, 100).z = 300
	point3 := changeX(&point1, 100)
	changeX(&point2, 100).z = 300
	point4 := changeY(&point2, 200)
	fmt.Printf("%v\n", point3)
	fmt.Printf("%v\n", point4)
	fmt.Printf("%d\n", point3.x)
	fmt.Printf("%d\n", point4.y)
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func changeX(ptr *point, x int) *point {
	ptr.x = x
	return ptr
}

func changeY(ptr *point, y int) point {
	ptr.y = y
	return *ptr
}
