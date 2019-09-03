package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxoso interface {
	// composicao das outras duas interfaces
	esportivo
	luxuoso
	// Mais metodos
}

type bmw struct{}

func (b bmw) ligarTurbo() {
	fmt.Println("Turbo ligado")
}
func (b bmw) fazerBaliza() {
	fmt.Println("Baliza...")
}
func main() {

	var b esportivoLuxoso = bmw{}
	b.ligarTurbo()
	b.fazerBaliza()

}
