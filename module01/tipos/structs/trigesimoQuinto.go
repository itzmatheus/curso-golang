package main

import "fmt"

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

}
