package main

import (
	"fmt"
)

func exibirResultadoAprovacao(nota float64) {

	if nota >= 7.0 {
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}

}

func main() {
	exibirResultadoAprovacao(7.2)
	exibirResultadoAprovacao(5.1)
}
