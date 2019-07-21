package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0
	for i < 10 { // No while in Go
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println("")
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}
	// Par ou Impar
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d é Par\n", i)
		} else {
			fmt.Printf("%d é Impar\n", i)
		}
	}
	i = 0
	// Laço Infito
	for {
		fmt.Printf("Ao infinito e além! | x:%d \n", i)
		i++
		//Funcao que para o codigo por 1 seg
		time.Sleep(time.Second)
	}
}
