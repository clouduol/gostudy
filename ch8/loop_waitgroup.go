package main

import (
	"fmt"
	"log"
	"sync"
	//"time"
)

func main() {
	filenames := []string{"pig", "dog", "chick"}
	chfilename := make(chan string)
	done := make(chan int)

	var results []string
	go func() {
		results = modfilenames(chfilename, done)
	}()

	for _, filename := range filenames {
		chfilename <- filename
	}
	// MUST DO IT to end range loop
	close(chfilename)

	//time.Sleep(3 * time.Second)
	<-done
	for _, result := range results {
		fmt.Println(result)
	}

	//time.Sleep(3 * time.Second)
}

func modfilenames(filenames <-chan string, done chan<- int) []string {
	defer fmt.Println("leave func modfilenames")
	fmt.Println("enter func modfilenames")
	chmodfilename := make(chan string)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		//worker
		go func(f string) {
			defer wg.Done()
			name, err := modfilename(f)
			if err != nil {
				log.Println(err)
				return
			}
			chmodfilename <- name
		}(f)
	}

	//closer
	go func() {
		defer fmt.Println("leave func closer")
		fmt.Println("enter func closer")
		wg.Wait()
		close(chmodfilename)
	}()

	var results []string
	for modfn := range chmodfilename {
		fmt.Printf("in func modfilenames: modfn %s\n", modfn)
		results = append(results, modfn)
	}

	done <- 0
	return results
}

func modfilename(filename string) (string, error) {
	defer fmt.Println("leave func modfilename")
	fmt.Println("enter func modfilename")
	return filename + ".png", nil
}
