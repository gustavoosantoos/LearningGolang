package main

import "fmt"

func main() {
	x := 10
	o, _ := fmt.Println(x)
	x = 99
	fmt.Println(x)

	var y int = 10
	fmt.Println(y)
}
