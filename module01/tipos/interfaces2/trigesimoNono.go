package main

import (
	"fmt"
	"reflect"
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

func init() {
	fmt.Println("Inicio do Programa")
}

func main() {

	carro1 := ferrari{"F40", false, 0}
	carro1.ligarTurbo()
	fmt.Println(carro1)
	fmt.Println(reflect.TypeOf(carro1))

	// Trabalhar com interface implementando um contrato
	// faz-se necessário adicionar o & comercial para
	// o GoLang atribuir o ponteiro
	var carro2 esportivo = &ferrari{"F50", false, 0}
	carro2.ligarTurbo()
	fmt.Println(carro2)
	fmt.Println(reflect.TypeOf(carro2))

}
