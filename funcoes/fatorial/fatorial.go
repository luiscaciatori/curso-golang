package main

import (
	"fmt"
)

func fatorial(numero int) (int, error) {
	switch {
	case numero < 0:
		return -1, fmt.Errorf("Número inválido: %d ", numero)
	case numero == 0:
		return 1, nil
	default:
		fatorialAnterior, nil := fatorial(numero - 1)
		return numero * fatorialAnterior, nil
	}
}

func main() {
	total, _ := fatorial(5)
	fmt.Println(total)
}
