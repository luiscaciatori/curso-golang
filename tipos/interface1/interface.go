package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return fmt.Sprintf("%s %s", p.nome, p.sobrenome)
}

func (p produto) toString() string {
	return fmt.Sprintf("%s R$ %.2f", p.nome, p.preco)
}

func main() {

	var coisa imprimivel = pessoa{"Luiz", "Gustavo"}
	fmt.Println(coisa.toString())

	coisa = produto{"Veja Multiuso", 4.59}
	fmt.Println(coisa.toString())

	p1 := produto{"Nescau 2.0", 9.29}
	fmt.Println(p1.toString())
}
