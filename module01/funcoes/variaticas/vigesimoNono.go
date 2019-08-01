// Funcoes que possuem uma quantidade de parametros variados
package main

import (
	"fmt"
)

func media(numeros ...float64) float64 {
	total := 0.0
	for _, x := range numeros {
		total += x
	}
	return total / float64(len(numeros))
}

func aprovados(lista ...string) {
	for index, nome := range lista {
		fmt.Printf("%d) %s\n", index+1, nome)
	}
}

func main() {

	fmt.Printf("Média de 9.9 | 6.9 | 10.0 = %.2f\n", media(9.9, 6.9, 10.0))
	lista := []string{"Matheus", "GoLang", "Python", "Java", "C", "Javascript"}
	// Passar slice para uma funcao variatica. PS: Nao funciona com array
	aprovados(lista...)

}
