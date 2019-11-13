package main

import "fmt"

var a = 42

func main() {
	fmt.Println(a)

	// Imprime o tipo da variável
	fmt.Printf("%T\n", a)

	fmt.Printf("%b\n", a)
	fmt.Printf("%x\n", a)
	fmt.Printf("%#x\n", a)
}
