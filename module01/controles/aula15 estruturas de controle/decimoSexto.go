package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

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

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {

	nota := 9.5
	imprimirResultado(nota)
	conceito := notaConceito(nota)
	fmt.Println(conceito)

	// if com inicializacao de variavel apenas dentro do escopo do IF

	if i := numeroAleatorio(); i > 5 { // suportado tambem no switch
		log.Println("Maior que 5", i)
	} else {
		log.Println("Menor que 5", i)
	}
}
