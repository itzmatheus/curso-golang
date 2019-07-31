package main

import "fmt"

func main() {

	arrayI := [3]int{1, 2, 3}
	var sliceI []int

	// arrayI = append(arrayI, 4, 5, 6 Nao é possível)
	sliceI = append(sliceI, 4, 5, 6)
	fmt.Println(arrayI, sliceI)

	slice2 := make([]int, 2)
	copy(slice2, sliceI)
	fmt.Println(slice2)
}
