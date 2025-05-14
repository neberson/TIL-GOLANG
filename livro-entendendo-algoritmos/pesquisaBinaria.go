package main

import "fmt"

func main() {

	minhaLista := []int{1, 3, 5, 7, 9, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98}
	valorEncontrado, execucoes := pesquiaBinaria(minhaLista, 98)

	if valorEncontrado == -1 {
		fmt.Println("Valor não encontrado")
	} else {
		fmt.Printf("Valor encontrado na posição %d\n", valorEncontrado)
		fmt.Printf("Número de execuções: %d\n", execucoes)
	}
}

func pesquiaBinaria(lista []int, item int) (int, int) {
	baixo := 0
	alto := len(lista) - 1
	contaExecucoes := 0

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]
		contaExecucoes++

		if chute == item {
			return meio, contaExecucoes
		}

		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}

	}

	return -1, contaExecucoes
}
