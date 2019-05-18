package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	printFoo()
}

func printFoo() {
	fmt.Println("I'am " +
		"a foo text")
}
