package main

import (
	"fmt"
)

func main() {
	p1 := Ponto{x: 2.0, y: 2.0}
	p2 := Ponto{x: 2.0, y: 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
