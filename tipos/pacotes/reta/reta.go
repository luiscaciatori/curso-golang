package main

import (
	"math"
)

// Iniciando com a letra Maiuscula é Publico (Fica visível dentro e fora do pacote)
// Iniciando com a letra Minuscula é Privado (Fica visível apenas dentro do pacote)

// Ponto -> representa uma cordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(a.x - b.x)
	cy = math.Abs(a.y - b.y)
	return
}

// Distancia -> Reponsável por calcular distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt((math.Pow(cx, 2)) + math.Pow(cy, 2))
}
