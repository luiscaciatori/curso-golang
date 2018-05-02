package main

import (
	"fmt"
)

func fatorial(numero uint) uint {
	if numero == 0 {
		return 1
	}

	return numero * fatorial(numero-1)
}

func main() {
	total := fatorial(5)
	fmt.Println(total)
}
