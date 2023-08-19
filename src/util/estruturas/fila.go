package main

import "fmt"

type Fila struct {
	Tamanho int
	Cabeca  *Node
}

func (fila *Fila) Adicionar(senha Senhas) {
	node := &Node{Value: senha}
	if fila.Cabeca == nil {
		fila.Cabeca = node
	} else {
		auxiliar := fila.Cabeca
		for auxiliar.Next != nil {
			auxiliar = auxiliar.Next
		}
		auxiliar.Next = node
	}

}

func (fila *Fila) GetFila() {
	node := fila.Cabeca
	for node != nil {
		fmt.Println(node.Value.Senha)
		node = node.Next
	}
}

func (fila *Fila) Remove() Senhas {
	fila.Cabeca = fila.Cabeca.Next
	return Senhas{}
}
