package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")

	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <valores> <unidade>")
		os.Exit(0)
	}
}
