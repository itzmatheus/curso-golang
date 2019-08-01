package main

import "fmt"

func obterValor(n int) int {
	defer fmt.Println("FIM")
	if n > 5000 {
		fmt.Println("Valor alto")
		return 5000
	}
	fmt.Println("Valor baixo")
	return n
}

func main() {
	// defer atrasar uma execução da sentença de código até o momento antes do metodo sair (chamar o return)
	// Ex: abrir conexao com banco de dados, na linha seguinte defer + funcao para fechar conexao
	fmt.Println("Valor: ", obterValor(6000))
	// output
	// Valor alto
	// FIM
	// Valor
}
