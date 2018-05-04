package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) GetNomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) SetNomeSobrenome(nomeCompleto string) {
	resultado := strings.Split(nomeCompleto, " ")
	p.nome = resultado[0]
	p.sobrenome = resultado[1]
}

func main() {
	p1 := pessoa{"Luiz", "Gustavo"}
	fmt.Println(p1.GetNomeCompleto())

	p1.SetNomeSobrenome("Gabriel Caciatori")
	fmt.Println(p1.GetNomeCompleto())
}
