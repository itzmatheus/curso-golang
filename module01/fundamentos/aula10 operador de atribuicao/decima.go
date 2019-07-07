package main

import "fmt"

func main() {

	var b byte = 3
	fmt.Println(b)

	i := 3 // inferencia de tipo
	fmt.Println(i)
	i += 10
	fmt.Println(i)
	i++
	fmt.Println(i)
	i -= 4
	fmt.Println(i)
	i /= 2
	fmt.Println(i)
	i *= 2
	fmt.Println(i)
	i %= 2
	fmt.Println(i)

	x, y := 1, 2.2

	fmt.Println(x, y)

}
