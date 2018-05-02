package main

import (
	"fmt"
	"math"
)

func main() {

	a := 4
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)

	c := 3.0
	d := 2.0

	fmt.Println(math.Max(c, d))
	fmt.Println(math.Min(c, d))
	fmt.Println(math.Pow(c, d))

	fmt.Println("E =>", 11&10)
	fmt.Println("OU =>", 11|10)
	fmt.Println("XOR =>", 11^10)

}
