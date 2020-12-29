package main

import "math"

// Inicializando com letra maiuscula é PÚBLICO (visibilidade dentro e fora do pacote)
// Inicializar com letra minúscula é PRIVADO (visivel apenas dentro do pacote)

// Ponto representa uma coordenada do plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia é reponsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
