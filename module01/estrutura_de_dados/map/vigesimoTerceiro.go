package main

import "fmt"

func main() {

	// var aprovados map[int]string
	// Maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[10] = "Matheus"
	aprovados[20] = "Leite"
	aprovados[30] = "José"

	fmt.Println(aprovados)

	for key, value := range aprovados {
		fmt.Println(key, "-", value)
	}

	delete(aprovados, 30)
	fmt.Println(aprovados)

	listaSalarios := map[string]float32{
		"Matheus":     9999.90,
		"Jose Emslon": 981.89,
		"Pedro":       898.25,
	}

	fmt.Println(listaSalarios)

	listaSalarios["Eu, Voce"] = 14525.00
	fmt.Println(listaSalarios)
	delete(listaSalarios, "INexiste")

	funcionarios := map[string]map[string]float64{
		"M": {
			"Matheus Leite":    1025.32,
			"Mariana da Silva": 5481.05,
		},
		"A": {
			"Ana Silva": 25858.32,
			"Aabraao":   1.1,
		},
	}
	fmt.Println(funcionarios)
}
