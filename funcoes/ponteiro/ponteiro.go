package main

import (
	"fmt"
)

func fun1(n int) {
	n++
}

// Todo ponteiro é repsentado por *
func fun2(n *int) {
	// * é usado para desreferenciar a varia ou seja acessar seu valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	fun1(n) // Passando por copia do valor
	fmt.Println(n)

	// & é usado para pegar o endereço da variavel
	fun2(&n) // Enviando a referencia da variavel
	fmt.Println(n)
}
