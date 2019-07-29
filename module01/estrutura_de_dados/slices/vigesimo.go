package main

import (
	"fmt"
	"reflect"
)

//Pessoa ...
type Pessoa struct {
	Nome  string
	Idade int
}

//ListaPessoas ...
type ListaPessoas struct {
	Pessoas []Pessoa
}

func main() {

	arraY := [3]int{3, 2, 3} //Array
	slicE := []int{1, 4, 2}  //Slice

	fmt.Println(arraY, slicE)
	fmt.Println(reflect.TypeOf(arraY), reflect.TypeOf(slicE))

	a2 := [5]int{1, 2, 3, 5, 6}

	// Slice nao é um array! Define um pedaço de um array.
	s2 := a2[2:4]

	fmt.Println(s2)

	pessoasArray := []Pessoa{{Nome: "Leite", Idade: 19}, {Nome: "Jose", Idade: 20}}
	fmt.Println(pessoasArray)
	pessoas := &ListaPessoas{}
	pessoas.Pessoas = pessoasArray
	fmt.Println(pessoas)

	novaPessoa := &Pessoa{Nome: "Matheus", Idade: 19}
	pessoas.Pessoas = append(pessoas.Pessoas, *novaPessoa)
	fmt.Println(pessoas)

}
