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

	// Sprint faz a mesma coisa que o printf mas retorna o format ao inves de plotar no stdout
	a8 := fmt.Sprintf("%b", a)
	fmt.Println(a8)
}
