package main

import (
	"fmt"
)

func main() {

	fmt.Println("Strings == : ", "Banana" == "Banana")
	fmt.Println("Strings != : ", "Banana" != "Banana")
	fmt.Println("3 < 2 = ", 3 < 2)
	fmt.Println("3 > 2 = ", 3 > 2)
	fmt.Println("3 <= 2 = ", 3 <= 2)
	fmt.Println("3 >= 2 = ", 3 >= 2)

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Matheus"}
	p2 := Pessoa{"Matheus"}
	fmt.Println("Pessoas: ", p1 == p2)
}
