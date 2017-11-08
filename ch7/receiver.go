// Receiver demonstrates value receiver versus pointer receiver
package main

import (
	"fmt"
)

type SetGeter interface {
	Get() int
	Set(int) bool
}

type Age struct {
	age int
}

func (a Age) Get() int {
	return a.age
}

func (a Age) Set(s int) bool {
	a.age = s
	return true
}

type Height struct {
	height int
}

func (h Height) Get() int {
	return h.height
}

func (h *Height) Set(s int) bool {
	h.height = s
	return true
}

type Weight struct {
	weight int
}

func (w *Weight) Get() int {
	return w.weight
}

func (w *Weight) Set(s int) bool {
	w.weight = s
	return true
}

func main() {
	var sg SetGeter
	fmt.Println("--- both value receiver -----")
	sg = Age{age: 5}
	sg.Set(10)
	fmt.Println(sg.Get())
	sg = new(Age)
	sg.Set(9)
	fmt.Println(sg.Get())

	fmt.Println("--- one value receiver, one pointer receiver -----")
	//sg = Height{height: 5}
	sg = &Height{height: 5}
	sg.Set(10)
	fmt.Println(sg.Get())
	sg = new(Height)
	sg.Set(9)
	fmt.Println(sg.Get())

	fmt.Println("--- both pointer receiver -----")
	//sg = Weight{weight: 5}
	sg = &Weight{weight: 5}
	sg.Set(10)
	fmt.Println(sg.Get())
	sg = new(Weight)
	sg.Set(9)
	fmt.Println(sg.Get())
}
