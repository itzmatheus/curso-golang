package main

import "runtime/debug"

func f3() {
	// Imprimir a pilha de execucao de um programa no momento que essa funcao for chamada
	debug.PrintStack()
}
func f2() {
	f3()

}
func f1() {
	f2()
}
func main() {
	f1()
}
