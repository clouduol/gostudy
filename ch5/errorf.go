// Errorf demonstrates function variable length parameters
package main

import (
	"fmt"
	"os"
)

func main() {
	linenum, name := 76, "ssssh"
	errorf(linenum, "program: %s", name)
	fmt.Fprintf(os.Stderr, "%s\n", "bracket?")
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args)
	fmt.Fprintln(os.Stderr)
}
