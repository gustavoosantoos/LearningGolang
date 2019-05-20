package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	printFoo()
	printCatchError()
	printIgnoringError()

	for i := 0; i < 100; i++ {
		fmt.Println(i)

		if i%5 == 0 {
			fmt.Println("Is 5 multiple")
		}
	}
}

func printFoo() {
	fmt.Println("I'am " +
		"a foo text")
}

func printCatchError() {
	n, e := fmt.Println("Printing message catch possible errors")
	fmt.Println(n)
	fmt.Println(e)
}

func printIgnoringError() {
	n, _ := fmt.Println("Printing message ignoring possible errors")
	fmt.Println(n)
}
