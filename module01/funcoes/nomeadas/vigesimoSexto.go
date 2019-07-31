package main

import "fmt"

func troca(x, y int) (segundo, primeiro int) {
	segundo = y
	primeiro = x
	return // retorno limpo
}

func main() {
	primeiro, segundo := troca(5, 10)
	fmt.Println(primeiro, segundo)
}
