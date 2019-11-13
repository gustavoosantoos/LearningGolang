package main

import (
	"fmt"
)

func main() {
	a1 := 1
	a2 := "Texto"
	a3 := `Texto com "aspas", é uma string literal`

	a4 := `Podemos ter também
	string em
	multiplas linhas
	`

	// Imprime a variável e abaixo, o tipo dela com o %T
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	fmt.Println(a2)
	fmt.Printf("%T\n", a2)

	fmt.Println(a3)
	fmt.Printf("%T\n", a3)

	fmt.Println(a4)
	fmt.Printf("%T\n", a4)
}
