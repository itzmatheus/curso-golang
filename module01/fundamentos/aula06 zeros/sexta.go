package main

import "fmt"

func main() {

	var a int
	var b float32
	var c bool
	var d string
	var e *int

	fmt.Printf(`
	int: %v
	float32: %v
	bool: %v
	string: %q
	*int: %v
	`, a, b, c, d, e)

}
