package main

import (
	"log"
)

func init() {
	log.SetPrefix("[LOG] ")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	log.Printf("printf %s\n", "test")
	log.Printf("printf %s\n", "test2")
}
