package main

import "fmt"

func pesquisaBinaria(lista []int, item int) (int, int) {
	baixo := 0
	alto := len(lista) - 1
	etapas := 0

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]
		if chute == item {
			return meio, etapas
		}

		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
		etapas++

	}
	return -1, etapas
}

func main() {
	minha_lista := make([]int, 128)
	resultado, etapas := pesquisaBinaria(minha_lista, 126)

	if resultado == -1 {
		println("Não encontrado")
	}

	fmt.Printf("Encontrado na posição %d com %d etapas", resultado, etapas)

}
