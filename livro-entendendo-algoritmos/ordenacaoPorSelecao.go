package main

func main() {
	// Cria uma lista de inteiros
	minhaLista := []int{5, 3, 8, 4, 2, 7, 1, 6}

	// Ordena a lista usando o algoritmo de ordenação por seleção
	listaOrdenada := ordenacaoPorSelecao(minhaLista)

	// Imprime a lista ordenada
	for _, valor := range listaOrdenada {
		println(valor)
	}
}

func buscaMenor(lista []int) int {
	menor := lista[0]
	indiceMenor := 0

	for i := range lista {
		if lista[i] < menor {
			menor = lista[i]
			indiceMenor = i
		}
	}

	return indiceMenor
}

// ordenacaoPorSelecao ordena uma lista de inteiros usando o algoritmo de ordenação por seleção
func ordenacaoPorSelecao(lista []int) []int {
	novoArray := make([]int, len(lista))
	for i := range lista {
		menor := buscaMenor(lista)
		novoArray[i] = lista[menor]
		lista = append(lista[:menor], lista[menor+1:]...)
	}

	return novoArray
}
