package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// cuidado
	fmt.Println("Idade é: " + string(97)) // Ire retornar o inteiro baseado na tabela ASCII

	// int para string
	fmt.Println("Idade é: " + strconv.Itoa(10))
	fmt.Println("Idade é: " + fmt.Sprint(10))

	// String to int
	idade, erro := strconv.Atoi("123")
	if erro != nil {
		fmt.Println("Erro na idade")
	}
	fmt.Println(idade + 10)

	// Boolean

	b, erro := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	} else {
		fmt.Println("Falso")
	}
}
