package main

import "fmt"

func main() {
	m := map[string]int{
		"Gustavo":  24,
		"Fernanda": 24,
		"Rosana":   42,
	}

	fmt.Println(m)
	fmt.Println(m["Gustavo"])
	fmt.Println(m["Adriano"])

	if v, ok := m["Adriano"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Não existe")
	}
}
