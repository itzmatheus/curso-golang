package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.\n")
	fmt.Println("Print com quebra de linha no final")
	fmt.Printf("Linguagem %s ANO: %d\n", "Go", 2019)

	x := 3.141516

	fmt.Println("Valor de X é:", x)

	// Retorna uma string a partir da funcao Sprint
	xs := fmt.Sprint(x)

	fmt.Println("Novo Valor de X é: " + xs)
	fmt.Println("Valor de X é: " + fmt.Sprint(x))

	fmt.Printf("Valor de PI É: %.2f ", x)

	a := 1
	b := 1.99999
	c := false
	d := "GoLang"

	fmt.Printf("\n%d - %.2f %t %s", a, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)

}
