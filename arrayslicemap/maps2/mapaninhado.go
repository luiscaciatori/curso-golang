package main

import (
	"fmt"
)

func main() {
	funcSalarios := map[string]map[string]float64{
		"G": {
			"Gustavo": 1000.00,
		},
		"E": {
			"Erica":  1500.00,
			"Eliana": 3000.00,
		},
		"P": {
			"Paulo": 400.00,
		},
	}

	for _, funcionarios := range funcSalarios {
		for nome, salario := range funcionarios {
			fmt.Println(nome, salario)
		}
	}
}
