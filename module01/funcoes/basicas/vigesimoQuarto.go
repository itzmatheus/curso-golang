package main

import "fmt"

func f1() {
	fmt.Println("Hello World!")
}

func f2(nome string, idade int) {
	fmt.Printf("Seu nome é: %s e você tem %d anos!\n", nome, idade)
}

func f3(x, y int) int {
	return x * y
}
func f4() (string, string) {
	return "Matheus", "Leite"
}
func main() {
	f1()
	f2("Matheus", 20)
	xVEZESy := f3(2, 4)
	fmt.Println(xVEZESy)
	nome, sobrenome := f4()
	fmt.Println(nome, sobrenome)
}
