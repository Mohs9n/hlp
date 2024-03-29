package main

import (
	"fmt"
	"os"

	"github.com/Mohs9n/hlp"
)

func main() {
	h, err := os.Open("/home/mohsen/src/go/gophers/hlp/test/ex4.html")
	if err != nil {
		panic(err)
	}

	links, err := hlp.Parse(h)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)
}
