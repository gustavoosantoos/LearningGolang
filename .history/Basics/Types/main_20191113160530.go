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
	multiplas linhas e com "aspas"
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

	// Valores default de variáveis (ZERO Values em Go, default em C#)
	// numericos = 0
	// string = string.Empty
	// outros = nil
	var a5 string
	var a6 int
	var a7 *int

	fmt.Println(a5)
	fmt.Printf("%T\n", a5)

	fmt.Println(a6)
	fmt.Printf("%T\n", a6)

	fmt.Println(a7)
	fmt.Printf("%T\n", a7)
}
