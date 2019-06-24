package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please supply a path to search")
		os.Exit(1)
	}
	target := os.Args[1]
	imports, err := Imports{path: target}.scan()

	if err != nil {
		log.Fatal("Something went wrong", err)
		os.Exit(2)
	}


	fmt.Println("Imports:", imports)
}

type Imports struct {
	path string
}

func(i Imports) scan() ([]string, error) {
	return []string{}, nil
}
ad