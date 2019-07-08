package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {

	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2
	comprarSorvete := trab1 || trab2

	return comprarTV50, comprarTV32, comprarSorvete

}

func main() {

	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf(`
	Situacao para o fim de semana:
	tv50: %t
	tv32: %t
	Sorvete com a Mulher: %t
	Saudável: %t
	`, tv50, tv32, sorvete, !sorvete)

}
