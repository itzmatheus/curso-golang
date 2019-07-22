package main

import "fmt"

func main() {

	// Homogeneas (mesmo tipo de dados) | estatica (fixo)
	var notas [3]float32
	notas[0], notas[1], notas[2] = 3.0, 5.2, 3.3
	fmt.Println(notas)

	var total float32
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float32(len(notas))
	fmt.Printf("Media do Array é: %.2f\n", media)

	// var nomes [2]string
	// Criando um array de strings e lendo no console para adicionar os nomes.
	// for index := range nomes {
	// 	fmt.Scanf("%s", &nomes[index])
	// }
	// fmt.Println(nomes)

	// [...] = É um array que possui um numero fixo, mas só poderá contar a partir da quantida de elementos que
	// eu colocar no array. DETALHE: Se remover os 3 pontinhos, vira um Slice.
	numeros := [...]int{1, 2, 3, 5, 6} // Compilador conta

	for i, num := range numeros {
		fmt.Printf("%d) %d\n", i, num)
	}
	for _, num := range numeros {
		fmt.Printf("%d ", num)
	}
}
