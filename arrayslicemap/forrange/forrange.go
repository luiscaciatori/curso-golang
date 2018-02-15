package main

import (
	"fmt"
)

func main() {
	var numeros = [...]int{1, 2, 3, 4, 5, 6, 7}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Println(numero)
	}
}
