package main

import (
	"fmt"
	m "math"
)

func main() {

	a := 3
	b := 2

	fmt.Println("SOMA => ", a+b)
	fmt.Println("SUBTRAÇÃO => ", a-b)
	fmt.Println("DIVISÃO => ", a/b)
	fmt.Println("MULTIPLICAÇÃO => ", a*b)
	fmt.Println("MÓDULO => ", a%b)

	// bitwise

	fmt.Println("E >=", a&b)   // 11 & 10 = 10
	fmt.Println("OU >=", a|b)  // 11 | 10 = 11
	fmt.Println("XOR >=", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// Operacoes com math

	fmt.Println("Maior valor", m.Max(c, d))
	fmt.Println("Menor valor", m.Min(c, d))
	fmt.Println("Exponenciação", m.Pow(c, d))

}
