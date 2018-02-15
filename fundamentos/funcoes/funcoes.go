package main

import (
	"fmt"
)

func soma(a int, b int) int {
	return a + b
}

func imprimir(a int) {
	fmt.Println(a)
}

func main() {
	resultado := soma(17, 2)
	imprimir(resultado)
}
