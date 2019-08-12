package main

import (
	"fmt"
)

type Carro struct{
	nome string
	velocidadeAtual int
}

type Ferrari struct{
	Carro // [COMPOSICAO] Campo anonimo do tipo carro:, tudo que tiver em Carro vai ficar disponivel em Ferrari.
	turboLigado bool
}

func (f Ferrari) velocidadeTotal() int {
	if f.turboLigado{
		return f.velocidadeAtual + 80
	}
	return f.velocidadeAtual
}
 
func main(){

	f := Ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 300
	f.turboLigado = true

	fmt.Printf("Nome: %s\nVelocidade com turbo: %d\n", f.nome, f.velocidadeTotal())

}