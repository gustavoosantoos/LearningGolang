package main

import "fmt"

type newint int

var x newint
var y int

// Rune é o mesmo que char em C#
var z rune

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

	y := int(x)
	fmt.Println(y)

	z = 'C'
	fmt.Println(z)
}
