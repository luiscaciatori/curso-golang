package main

import (
	"fmt"
)

func main() {
	a := 1

	var p *int = nil
	p = &a // pegando o endereço da variável
	*p++   // pegando o valor associado ao ponteiro
	a++

	// Go não tem aritmética de ponteiros

	fmt.Printf("%v, %v, %v, %v", *p, a, &a, p)
}
