package main

import "fmt"

func main() {
	x := 10
	x = 99
	var y int = 10
	str := "Hello, World!"
	o, _ := fmt.Println(x)

	fmt.Println(str)
	fmt.Println(o)
	fmt.Println(y)
	fmt.Println(x)
}
