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
}
