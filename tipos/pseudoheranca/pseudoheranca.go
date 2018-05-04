package main

import (
	"fmt"
)

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {

	modelo := ferrari{}
	modelo.nome = "458"
	modelo.velocidadeAtual = 100
	modelo.turboLigado = true

	fmt.Printf("A Ferrari %s está com o turbo ligado? %v", modelo.nome, modelo.turboLigado)

}
