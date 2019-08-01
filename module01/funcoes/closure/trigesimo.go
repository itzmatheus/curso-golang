package main

import (
	"fmt"
)

func closure() func() {
	x := 10
	var funcao = func() {
		var x = 102
		fmt.Println(x)
	}
	fmt.Println(x)
	return funcao
}

func main() {
	// Fechamento, algo que envolve alguma coisa. Uma funcao que envolve outra funcao, fazendo com que a funcao interna
	// tenha memoria do local onde foi declarada. A funcao externa funciona como um Closure (fechamento) como se
	// o contexto da funcao. Em go uma func é um Closure.
	x := 20
	fmt.Println(x)
	imprimeX := closure()
	imprimeX()

	// output
	// 20
	// 10
	// 102
}
