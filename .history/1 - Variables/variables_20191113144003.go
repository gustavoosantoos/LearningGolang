package main

import "fmt"

var y = 99

func main() {
	// Short declaration, declara variavel e atribui um valor
	x := 10
	str := "Hello, World!"

	// Atribuição
	x = 99

	var a1 = 30
	var a2 int = 32

	// Short declartion with err handling
	o, _ := fmt.Println(x)

	fmt.Println(str)
	fmt.Println(o)
	fmt.Println(y)
	fmt.Println(x)
}

func foo() {
	fmt.Print("Again:", y)
}