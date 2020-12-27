package main

import "fmt"

func secret(nome string) (retorno string) {
	retorno = fmt.Sprintf("Atenção caro %s.", nome) + "Eu sou Batman!"
	return
}
func exec(funcao func(string) string, usuario, senha string) (retorno string) {
	if usuario == "matheus" && senha == "123456" {
		retorno = funcao(usuario)
	} else {
		retorno = "Usuario Incorreto"
	}
	return
}
func main() {
	var usuario, senha string
	fmt.Println("Digite seu usuario")
	fmt.Scanf("%s", &usuario)
	fmt.Println("Digite sua senha")
	fmt.Scanf("%s", &senha)
	resultado := exec(secret, usuario, senha)
	fmt.Println(resultado)
}
