package main

import (
	"fmt"
)

func main() {
	funcSalarios := map[string]float64{
		"Luiz Gustavo": 3000.00,
		"Erica":        4500.00,
		"Roberto":      21000.00,
	}

	funcSalarios["Rafael"] = 3400.00

	for nome, salario := range funcSalarios {
		fmt.Println(nome, salario)
	}

}
