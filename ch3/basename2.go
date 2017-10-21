// Basename2 removes directory components and a .suffix
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("%s\n", basename(arg))
	}
}

func basename(s string) string {
	// Discard last "/" and everything before
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	// Preserve everything before last '.'
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
