package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {
	var nota = int(n)

	// SWITCH Nao tem uma selecao default para false.
	switch nota {
	case 10:
		// Palavra reservada para o case continuar descendo.
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida!"
	}
}

func tipo(i interface{}) string {
	switch i.(type) {

	case int:
		return "Inteiro"
	case string:
		return "String"
	case float32, float64:
		return "Real"
	case func():
		return "Função"
	case bool:
		return "Boolean"
	default:
		return "Sei não, viss."
	}
}

func main() {

	fmt.Println("Nota 9 = ", notaParaConceito(9))
	fmt.Println("Nota 12 = ", notaParaConceito(12))
	fmt.Println("Nota 5 = ", notaParaConceito(5))

	tempo := time.Now()

	// Quando nao coloca uma variavel, ele procurará algum case que seja true e irá executar.
	switch { // switch True
	case tempo.Hour() < 12:
		fmt.Println("Bom dia!")
	case tempo.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa Noite!")
	}

	fmt.Println(tipo(3.2))
	fmt.Println(tipo(3))
	fmt.Println(tipo(true))
	fmt.Println(tipo("Matheus"))
	fmt.Println(tipo(time.Hour))
	fmt.Println(tipo(func() {}))

}
