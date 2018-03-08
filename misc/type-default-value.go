package main

import (
	"fmt"
)

type UID string

func main() {
	var uid UID
	if uid == "" {
		fmt.Println("empty string")
	}

	/*
		if uid == nil {
			fmt.Println("nil")
		}
	*/

	fmt.Println(uid)

	uid = "123456"
	fmt.Println(uid)

	var s string = "this is a string"
	/*
		s = uid
	*/
	fmt.Println(s)

	s = "assign another string"
	fmt.Println(s)
	/*
		s[0] = '2'
	*/
	fmt.Printf("%c\n", rune(s[0]))
	fmt.Printf("%c\n", s[0])

	/*
		uid = s
	*/
	uid = "use string literal value"
	fmt.Println(uid)

	uid = UID(s)
	fmt.Println(uid)
}
