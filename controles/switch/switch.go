package main

import (
	"fmt"
)

func classificaNota(notaFloat float64) string {

	nota := int(notaFloat)

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}

}

func main() {
	fmt.Println(classificaNota(0))
	fmt.Println(classificaNota(6.9))
	fmt.Println(classificaNota(8.9))
}
