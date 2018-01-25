package main

import (
	"fmt"
)

func main() {
	filenames := []string{"pig", "dog", "chick"}

	type item struct {
		modfile string
		err     error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.modfile, it.err = modfilename(f)
			ch <- it
		}(f)
	}

	var modfilenames []string
	for range filenames {
		it := <-ch
		if it.err != nil {
			return
		}
		modfilenames = append(modfilenames, it.modfile)
	}

	for _, x := range modfilenames {
		fmt.Println(x)
	}
}

func modfilename(f string) (string, error) {
	return f + ".png", nil
}
