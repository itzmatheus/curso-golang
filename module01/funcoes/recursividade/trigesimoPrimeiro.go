package main

import "fmt"

func fatorial(n int) (int, error) {

	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}

}
func fatorial2(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial2(n-1)
	}

}
func main() {

	a, _ := fatorial(3)
	fmt.Println(a)
	_, err := fatorial(-1)
	fmt.Println(err)
	b := fatorial2(3)
	fmt.Println(b)

}
