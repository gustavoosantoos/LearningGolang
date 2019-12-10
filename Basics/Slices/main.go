package main

import "fmt"

func main() {
	x1 := []int{1, 5, 2, 19, 30}
	fmt.Println(x1)
	fmt.Printf("%T\n", x1)

	x1 = x1[:len(x1)-1]

	for index, value := range x1 {
		fmt.Println(index, value)
	}

	x2 := x1[1:4]
	fmt.Println(x2)
}
