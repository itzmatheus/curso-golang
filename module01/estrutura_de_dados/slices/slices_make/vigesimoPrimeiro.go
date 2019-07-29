package main

import "fmt"

func main() {

	// Criando um slice a partir da funcao MAKE
	s := make([]int, 10)
	s[9] = 1000

	fmt.Println(s)

	// Criando um slice, mas o array interno que automaticamente cria, ira possuir 20 elementos
	// Referenciando apenas 10 elementos
	s = make([]int, 10, 20)

	// Cap = Capacidade maxima do array interno utilizado pelo slice
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(s, len(s), cap(s))

	// Caso tamanho seja atingido, internamente o slice vai referenciando varios
	// arrays internos, replicando os dados
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))

}
