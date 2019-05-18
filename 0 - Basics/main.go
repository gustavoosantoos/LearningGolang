package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	printFoo()

	for i := 0; i < 100; i++ {
		fmt.Println(i)

		if i % 5 == 0 {
			fmt.Println("Is 5 multiple")
		}
	}
}

func printFoo() {
	fmt.Println("I'am " +
		"a foo text")
}
