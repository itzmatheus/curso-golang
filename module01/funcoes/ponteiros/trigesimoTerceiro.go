package main

import "fmt"

func func1(n int) {
	n++
}
func func2(n *int) {
	*n++
}
func main() {

	n := 1
	func1(n) // Passagem por valor
	fmt.Println(n)
	func2(&n) // Passagem por referencia
	fmt.Println(n)

}
