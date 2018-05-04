package main

import (
	"fmt"
)

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {

	switch {
	case n.entre(9, 10):
		return "A"

	case n.entre(7, 8.9):
		return "B"

	case n.entre(5, 6.99):
		return "C"

	case n.entre(3, 4.99):
		return "B"

	default:
		return "E"

	}

}

func main() {
	fmt.Println(notaParaConceito(6.0))
}
