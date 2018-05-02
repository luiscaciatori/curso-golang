package main

import (
	"fmt"
)

func troca(p1, p2 int) (primeiro, segundo int) {
	primeiro = p2
	segundo = p1
	return
}

func main() {
	p1, p2 := troca(2, 3)
	fmt.Println(p1, p2)

	primeiro, segundo := troca(7, 4)
	fmt.Println(primeiro, segundo)
}
