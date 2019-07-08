package main

import "fmt"

func imprimirResultado(nota float64) {

	if nota >= 7 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}

}

func notaConceito(nota float64) string {

	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else {
		return "C"
	}

}

func main() {

	nota := 9.5
	imprimirResultado(nota)
	conceito := notaConceito(nota)
	fmt.Println(conceito)
}
