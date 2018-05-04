package main

import (
	"fmt"
)

type Item struct {
	produtoID int
	qtd       float64
	preco     float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (p Pedido) ValorTotal() float64 {
	total := 0.00

	for _, item := range p.itens {
		total += item.preco * item.qtd
	}

	return total
}

func main() {
	pedido := Pedido{
		userID: 1,
		itens: []Item{
			{produtoID: 1, qtd: 2, preco: 7.99},
			{produtoID: 2, qtd: 3, preco: 4.55},
			{produtoID: 3, qtd: 3, preco: 6.08},
		},
	}

	fmt.Printf("Valor total do Pedido: %.2f", pedido.ValorTotal())
}
