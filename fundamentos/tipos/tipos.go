package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 10, 100, 1000)
	fmt.Println("Literal de um Inteiro é:", reflect.TypeOf(1000))

	// números sem sinal ou seja apenas positivos podem ser uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// números com sinal são int8, int16, int32, int64
	iMax64 := math.MaxInt64
	iMax32 := math.MaxInt32

	fmt.Println("Max int64", iMax64)
	fmt.Println("Max int32", iMax32)

	iMin64 := math.MinInt64
	iMin32 := math.MinInt32
	fmt.Println("Min int64", iMin64)
	fmt.Println("Min int32", iMin32)

	var code rune = 'a'
	fmt.Println("Rune é", reflect.TypeOf(code))

	bo := true
	fmt.Println("Boolean é", reflect.TypeOf(bo))

	txt := "Meu tipo é tal"
	fmt.Println("Reflection da frase", reflect.TypeOf(txt))
	fmt.Println("Tamanho de st", len(txt))

	txtComCrase := `Meu
	tipo
	é
	tal`
	fmt.Println("Reflection da", reflect.TypeOf(txtComCrase))
	fmt.Println(len(txtComCrase))

}
