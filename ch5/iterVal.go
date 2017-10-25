// IterVal demostrates for statement iteral variables and anonymous function warnings
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//right()
	wrong()
}

func wrong() {
	var rmdirs []func()
	for _, dir := range tempDirs() {
		fmt.Printf("---- &dir: %x\n", &dir)
		dir := dir
		fmt.Printf("---- &dir: %x\n", &dir)
		fmt.Printf("create dir: %s\n", dir)
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		fmt.Printf("remove dir\n")
		rmdir()
	}
}

func right() {
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d
		fmt.Printf("---- &dir: %x\t &d: %x\n", &dir, &d)
		fmt.Printf("create dir: %s\n", dir)
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	for _, rmdir := range rmdirs {
		fmt.Printf("remove dir\n")
		rmdir()
	}
}

func tempDirs() []string {
	var dirs []string
	for i := 0; i < 5; i++ {
		dir := "/tmp/test/" + strconv.Itoa(i)
		dirs = append(dirs, dir)
	}
	return dirs
}
