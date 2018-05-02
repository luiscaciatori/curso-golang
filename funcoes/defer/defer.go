package main

import (
	"fmt"
)

func imprimeValor(valor int) int {
	defer fmt.Println("Acabou!")

	if valor > 5000 {
		fmt.Println("Valor alto")
		return valor
	}

	fmt.Println("Valor baixo")
	return valor
}

func main() {
	fmt.Println(imprimeValor(4000))
	fmt.Println(imprimeValor(10000))
}
