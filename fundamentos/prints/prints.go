package main

import (
	"fmt"
)

func main() {
	fmt.Print("Eita")
	fmt.Print(" nois")

	fmt.Println(" Pula")
	fmt.Println("linha")

	x := 3.141516

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x:", xs)

	fmt.Printf("O valor de x: %f", x)

	fmt.Printf("\nO valor de x: %v", x)

	fmt.Printf("\nO valor de x: %.2f", x)
}
