package main

import (
	"fmt"
	"strings"
)

//Produto
type Produto struct {
	Nome     string
	Preco    float32
	Desconto float32
}

// Metodo funcao com receiver (RECEPTOR)
func (p Produto) precoComDesconto() float32 {
	return p.Preco * (1 - p.Desconto)
}

type Item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (p Pedido) valorTotal() float64 {

	total := 0.0
	for _, item := range p.itens {
		total += (item.preco * float64(item.quantidade))
	}
	return total

}

type Pessoa struct {
	Nome      string
	Sobrenome string
}

func (p Pessoa) getNome() string {
	return p.Nome + " " + p.Sobrenome
}

func (p *Pessoa) setNomeCompleto(novoNome string) {
	partes := strings.Split(novoNome, " ")
	p.Nome = partes[0]
	p.Sobrenome = partes[1]
}

func main() {

	var produto1 *Produto
	produto1 = &Produto{
		Nome:     "Dell Inspiron 3000",
		Preco:    3150.99,
		Desconto: 0.05,
	}

	fmt.Println(*produto1, produto1.precoComDesconto())

	produto2 := Produto{"Headset Logitech G935", 879.99, 0.05}

	fmt.Println(produto2, produto2.precoComDesconto())

	fmt.Println("==============")

	pedido := Pedido{
		userID: 1,
		itens: []Item{
			Item{1, 2, 30.00},
			Item{2, 3, 99.90},
			Item{3, 1, 278.90},
		},
	}

	fmt.Println(pedido)
	fmt.Printf("Valor total: R$ %.2f\n", pedido.valorTotal())

	p1 := Pessoa{Nome: "Matheus", Sobrenome: "Leite"}
	fmt.Println(p1.getNome())

	// Por baixo dos panos o Go vai pegar a referencia que não é um ponteiro e passar para a função
	// com receiver o endereço dela para ser alterada de forma encapsulada
	p1.setNomeCompleto("Leite Jose")
	fmt.Println(p1.getNome())

}
