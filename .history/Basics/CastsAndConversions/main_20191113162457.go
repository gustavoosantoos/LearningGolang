package main

import "fmt"

type hotdog int

func main() {
	a1 := 1
	var a2 hotdog = 3

	fmt.Println(a1)
	fmt.Println(a2)

	// Em GO, não há cast e sim conversions, somente semântica, pois é a mesma coisa
	a1 = int(a2)

	fmt.Println(a1)
}
