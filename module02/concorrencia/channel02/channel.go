package main

import (
	"fmt"
	"time"
)

// Channel (canal) - forma de comunicacao entre as goroutines
// channel é um tipo chan

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados

	time.Sleep(2 * time.Second)
	c <- 3 * base

	time.Sleep(time.Second * 3)
	c <- 4 * base

}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo os dados do canal

	fmt.Println(a, b)

	fmt.Println(<-c)

}
