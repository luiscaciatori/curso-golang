package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 3.24
	y := 2

	fmt.Println(x / float64(y))

	fmt.Println("Converto o número 192 para string resultado:", strconv.Itoa(192))

	inteiro, _ := strconv.Atoi("123")

	fmt.Println("Convertendo o string para numero com o strconv:", inteiro)

	b, _ := strconv.ParseBool("true")

	if b {
		fmt.Println("Verdadeiro lek")
	}

}
