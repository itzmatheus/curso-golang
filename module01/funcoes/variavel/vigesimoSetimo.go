package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {

	fmt.Println(soma(1, 3))

	sub := func(a, b int) int { return a - b }

	fmt.Println(sub(3, 2))
}
