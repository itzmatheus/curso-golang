package main

import "fmt"

type Imprimivel interface {
	toString() string
}
type Pessoa struct {
	nome, sobrenome string
}

type Produto struct {
	nome  string
	preco float64
}

// interfaces sao implementadas implicitamente

func (p Pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p Produto) toString() string {
	return fmt.Sprintf("%s - %.2f", p.nome, p.preco)
}

func imprimir(x Imprimivel) {
	fmt.Println(x.toString())
}

func main() {

	var coisa Imprimivel = Pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	// Polimorfismo
	coisa = Produto{"Carro", 110000.99}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	imprimir(Produto{"Carro SUV", 110000.99})

}
