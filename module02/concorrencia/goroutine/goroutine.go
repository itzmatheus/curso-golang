package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {

	// Em série
	// fale("Matheus", "Por quê você não fala comigo ? ", 3)
	// fale("Joao", "Só posso falar depois de você", 1)

	// go routine
	// a palavra reservada go antes do método: criar uma go routine
	// go fale("Matheus", "Ei ...", 5)
	// go fale("João", "Opa ...", 5)

	// time.Sleep(time.Second * 12)
	// fmt.Println("Fim")

	go fale("Matheus", "Entendi!!!", 10)
	fale("João", "Parabéns!!!", 5)

}
