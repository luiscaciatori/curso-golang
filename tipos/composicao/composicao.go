package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	abrirTetoSolar()
}

type esportivoluxuoso interface {
	esportivo
	luxuoso
}

type bmwM4 struct{}

func (b bmwM4) ligarTurbo() {
	fmt.Println("Turbo ligado")
}

func (b bmwM4) abrirTetoSolar() {
	fmt.Println("Teto solar aberto")
}

func main() {
	var bmw bmwM4 = bmwM4{}
	bmw.abrirTetoSolar()
	bmw.ligarTurbo()
}
