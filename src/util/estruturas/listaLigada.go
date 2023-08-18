package main

import "fmt"

type List struct {
	Cabeca *Node
	Calda  *Node
}

func (lista *List) Adicionar(senha Senhas) {
	node := &Node{Value: senha}
	if lista.Cabeca == nil {
		lista.Cabeca = node
	}
	if lista.Calda != nil {
		lista.Calda.Next = node

	}
	lista.Calda = node
}

func (lista *List) getLista() {
	node := lista.Cabeca
	for node != nil {
		fmt.Println(node.Value.Senha)
		node = node.Next
	}
}
func (lista *List) Busca(senha string) Senhas {
	node := lista.Cabeca
	for node != nil {
		if node.Value.Senha == senha {
			break
		}
		node = node.Next
	}
	if node != nil {
		return node.Value
	}
	return Senhas{}
}

func (lista *List) Remove(senha string) Senhas {
	if lista.Cabeca.Value.Senha == senha {
		lista.Cabeca = lista.Cabeca.Next
		return Senhas{}
	}
	anterior := lista.Cabeca
	node := lista.Cabeca.Next
	for node != nil {
		if node.Value.Senha == senha {
			anterior.Next = node.Next
			break
		}
		anterior = node
		node = node.Next
	}
	if lista.Calda == node {
		lista.Calda = anterior
	}
	return Senhas{}
}

type Node struct {
	Value Senhas
	Next  *Node
}
