package main

import (
	"fmt"
)

func main() {
	// Todo map deve ser inicializado
	aprovados := make(map[int]string)
	aprovados[1] = "Erica"
	aprovados[2] = "Luiz Gustavo"
	aprovados[3] = "Roberto"

	fmt.Println(aprovados)

	for id, nome := range aprovados {
		fmt.Printf("%s ID: %d\n", nome, id)
	}

	fmt.Println(aprovados[1])

	delete(aprovados, 3)

	fmt.Println(aprovados)
}
