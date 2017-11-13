// TypeAssert demonstrate interface type assertion
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int64

func (b *ByteCounter) Write(p []byte) (n int, err error) { return 0, nil }

func main() {
	var w io.Writer
	fmt.Printf("w: %T\n", w)
	w = os.Stdout
	f := w.(*os.File)
	fmt.Printf("f: %T\n", f)

	rw := w.(io.ReadWriter)
	fmt.Printf("rw: %T\n", rw)

	w = new(ByteCounter)
	rw, ok := w.(io.ReadWriter)
	if ok {
		fmt.Printf("rw: %T\n", rw)
	} else {
		fmt.Printf("rw == nil\n")
	}

	c, ok := w.(*bytes.Buffer)
	//fmt.Printf("c: %T\n", w.(*bytes.Buffer))
	if ok {
		fmt.Printf("c: %v\n", c)
	} else {
		fmt.Printf("c == nil\n")
	}

	fmt.Println("main end")
}
