package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 3

	fmt.Println(x)
	fmt.Printf("%T", x)
}
