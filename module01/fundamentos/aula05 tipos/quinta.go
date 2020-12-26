package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Tipos numericos

	// Inteiros
	fmt.Println(1, 2, 1000)

	// Diz o tipo da variavel quando passar como parametro
	fmt.Println("Literal inteiro é: ", reflect.TypeOf(3))

	// Sem sinal apenas valores positivos
	// uint8 uint16 uint32 uint64
	// o byte é um alias (apelido) para o uint8
	var a byte = 255 // Se passar de 255 ele da erro de overflows byte
	fmt.Println("Byte uint8: ", a)
	fmt.Println("Tipo dele: ", reflect.TypeOf(a))

	// com sinais... int8 int16 int32 int64

	i1 := math.MaxInt64
	fmt.Println("Max int64: ", i1)
	fmt.Println("Seu tipoe é: ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	// var i2 rune = 'a' nao precisa declara o rune, ja esta explicito pelo uso das aspas simples
	fmt.Println("Valor de i: ", i2)
	fmt.Println("Seu tipo é: ", reflect.TypeOf(i2))

	// numeros reais float32 float64
	var x float32 = 49.99
	var b = float32(49.99)
	fmt.Println("Valor de x: ", x)
	fmt.Println("Seu tipo é: ", reflect.TypeOf(x))
	// Por padrao quando passa um literal do tipo float, fica subentendido que seja um tipo float64
	fmt.Println("Seu tipo é: ", reflect.TypeOf(44.98))
	fmt.Println("Nova declaracao: ", b)
	fmt.Println("Nova declaracao: ", reflect.TypeOf(b))

	// tipos booleanos

	var verdadeiro = true
	fmt.Println("Valor de verdadeiro: ", verdadeiro)
	fmt.Println("Valor: ", reflect.TypeOf(verdadeiro))
	fmt.Println("Valor de verdadeiro negado: ", !verdadeiro)

	// tipo string
	// Delimitado apenas por "" e ``
	s1 := "Matheus"
	fmt.Println("Nome: ", s1)
	fmt.Println("NSeu tipoe ", reflect.TypeOf(s1))
	fmt.Println("Nome concatenado: ", s1+" Leite")
	fmt.Println("Tamanho do nome ", s1, "é: ", len(s1))

	// string com multiplas linhas
	fmt.Println()

	campoGrande := `
	Nome completo:
	Idade:
	Curso:
	`
	fmt.Println(campoGrande)
	fmt.Println("Tipo: ", reflect.TypeOf(campoGrande))

	// tipo char
	char := 'a'
	fmt.Println(char)
	fmt.Println("tipo char: ", reflect.TypeOf(char))
}
