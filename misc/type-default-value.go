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
}
