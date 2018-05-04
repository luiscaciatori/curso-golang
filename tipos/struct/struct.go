package main

import "fmt"

type Produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p Produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 Produto
	produto1 = Produto{
		nome:     "Lápis",
		preco:    8.50,
		desconto: 0.07,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := Produto{"Caneta", 9.99, 0.1}

	fmt.Println(produto2, produto2.precoComDesconto())
}
