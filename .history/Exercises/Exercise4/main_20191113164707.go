package main

import "fmt"

type newint int

var x newint

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
