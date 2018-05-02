package main

import (
	"fmt"
)

func media(valores ...float64) float64 {
	if len(valores) == 0 {
		return 0.0
	}

	total := 0.0

	for _, num := range valores {
		total += num
	}

	return total / float64(len(valores))
}

func main() {
	fmt.Printf("Média: %.2f \n", media(3.4, 5.7, 8.8, 9.9, 10.0))
	fmt.Printf("Média: %.2f \n", media())
}
