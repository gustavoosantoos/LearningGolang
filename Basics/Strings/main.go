package main

import "fmt"

func main() {
	s := "Hello, world!"
	s1 := `Hello,
		   World!`

	fmt.Println(s)
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	// Strings, como no C# e Java, são um array de byes (que na vdd é um alias para uint8)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	if len(s) > 3 || len(s) < 3 {
		fmt.Println(len(s))
	}

	x1 := 1

	// Faz o papel de um loop infinito
	// for {
	// 	fmt.Println("Loop infinito")
	// }

	// Faz o papel de While
	for x1 < 3 {
		fmt.Println(x1)
		x1++
	}

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U", s1[i])
	}

	for i, v := range s {
		fmt.Println(i, v)
	}
}
