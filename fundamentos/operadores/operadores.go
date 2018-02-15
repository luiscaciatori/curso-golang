package main

import (
	"fmt"
)

func cromprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV42 := trab1 && trab2
	comprarTV32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTV42, comprarTV32, comprarSorvete
}

func main() {
	t1, t2, sorvete := cromprar(true, true)
	fmt.Printf("TV 42: %t. TV 32: %t. Comprar Sorvete: %t. Saudável: %t\n", t1, t2, sorvete, !sorvete)
}
