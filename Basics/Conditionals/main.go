package main

import "fmt"

func main() {
	x1 := 1
	x2 := "Teste"

	if x1 == 1 {
		fmt.Println(x1)
	}

	// switch x2 {
	// case "Testee":
	// 	fmt.Println("true")
	// case "":
	// 	fmt.Println("false")
	// default:
	// 	fmt.Println("default")
	// }

	switch x2 {
	case "Teste":
		fmt.Println(x2)
		fallthrough
	case "Bunda":
		fmt.Println("false 2")
	default:
		fmt.Print("default")
	}
}
