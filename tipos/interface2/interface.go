package main

import (
	"fmt"
)

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {

	f40 := ferrari{"F40", false, 0}
	f40.ligarTurbo()

	var f430 esportivo = &ferrari{"F430", false, 0}
	f430.ligarTurbo()

	fmt.Println(f40, f430)

}
