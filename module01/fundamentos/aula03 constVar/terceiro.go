package main

import (
	"fmt"
	_ "log"  // importando a funcao log sem utilizar
	m "math" //chamando o import e referenciando em m
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	//forma reduzida de criar uma variavel em go
	// (preferencia) declara a variavel e inicializa tudo numa construcao reduzida
	area := PI * m.Pow(raio, 2)

	fmt.Printf("Area: %.2f\n", area)

	//bloco de construcao de constantes
	const (
		a = 1
		b = 2
	)
	// bloco de construcao de variaveis
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println("true =", e, "e false =", f)

	x, y := "Matheus", "Leite"
	fmt.Println(x, y)
}
