package main

import (
	"fmt"
)

func main() {
	a1 := 1
	a2 := "Texto"

	// Imprime a variável e abaixo, o tipo dela com o %T
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	fmt.Println(a2)
	fmt.Printf("%T\n", a2)
}
