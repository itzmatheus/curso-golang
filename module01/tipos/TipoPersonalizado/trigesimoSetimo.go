package main

import (
	"fmt"
)

type Nota float64

func (n Nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n Nota) string {
	if n.entre(9, 10){
		return "A"
	}else if n.entre(7,8.99){
		return "B"
	}else if n.entre(3, 6.99){
		return "C"
	}
	return "D"
}

func main(){

	fmt.Println("Nota: ", notaParaConceito(9.8))
	fmt.Println("Nota: ", notaParaConceito(7.8))
	fmt.Println("Nota: ", notaParaConceito(3.8))


}