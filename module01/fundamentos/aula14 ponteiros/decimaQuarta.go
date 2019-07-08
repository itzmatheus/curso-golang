package main

import "fmt"

func alterarPoint(i *int) {
	*i = 10
}

func main() {

	i := 1
	var x *int
	fmt.Println(x)
	x = &i // Pega o endereco da variavel I

	// Go nao tem aritmetica de ponteiros
	// x++

	fmt.Println(x)
	alterarPoint(&i)
	fmt.Println(*x)
	*x++
	fmt.Println(*x)
	fmt.Println(i)
	fmt.Println(x == &i) // Comparando o endereco armazenado em X e de fato igual ao endereco de I

}
