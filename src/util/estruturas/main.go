package main

import "fmt"

type Senhas struct {
	Tamanho     int
	Senha       string
	Dificuldade string
}

func main() {
	lista := List{}

	senha1 := Senhas{12, "123412341234", "Baixa"}
	senha2 := Senhas{8, "casa1234", "Baixa"}

	lista.Adicionar(senha1)
	lista.Adicionar(senha2)

	lista.getLista()
	fmt.Println("|------------------------|")
	senha3 := lista.Busca("casa1234")
	fmt.Println(senha3)

	fmt.Println("|------------------------|")
	lista.Remove("casa1234")
	lista.getLista()
}
