package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.145
	var raio = 3.2

	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circuferencia é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 1
		d = 2
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := true, 1, "epa"

	fmt.Println(g, h, i)

}
