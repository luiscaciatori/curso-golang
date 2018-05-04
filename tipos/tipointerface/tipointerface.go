package main

import (
	"fmt"
)

func main() {
	var coisa interface{}

	coisa = 3
	fmt.Println(coisa)

	coisa = "nova coisa"
	fmt.Println(coisa)
}
