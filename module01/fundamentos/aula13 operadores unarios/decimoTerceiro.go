package main

import "fmt"

func main() {

	x := 1
	y := 2
	fmt.Println(x)

	// apenas postfix
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)
	fmt.Println(y)
	// fmt.Println(x == y--) nao pode
}
