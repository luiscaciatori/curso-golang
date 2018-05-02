package main

import (
	"fmt"
)

func main() {
	var arr [3]float64

	fmt.Println(arr)

	arr[0], arr[1], arr[2] = 7.9, 10.2, 4.5
	fmt.Println(arr)

	total := 0.0
	for i := 0; i < len(arr); i++ {
		total += arr[i]
	}

	resultado := total / float64(len(arr))

	fmt.Printf("Média: %.2f", resultado)
}
