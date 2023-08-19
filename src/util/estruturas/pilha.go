package main

import "fmt"

type Pilha struct {
	Calda *Node
}

func (pilha *Pilha) Adicionar(senha Senhas) {
	node := &Node{Value: senha}
	if pilha.Calda != nil {
		node.Next = pilha.Calda
	}
	pilha.Calda = node
}
func (pilha *Pilha) GetPilha() {
	node := pilha.Calda
	for node != nil {
		fmt.Println(node.Value.Senha)
		node = node.Next
	}
}
func (pilha *Pilha) Remove() Senhas {
	node := pilha.Calda
	pilha.Calda = node.Next
	return Senhas{}
}
