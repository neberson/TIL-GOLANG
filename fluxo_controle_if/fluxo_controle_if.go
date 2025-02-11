package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Neberson Andrade"
	var versao float32 = 1.1
	var idade int = 29

	fmt.Println("Olá, sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("O tipo da váriavel nome é:", reflect.TypeOf(nome))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O endereço da váriavel comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço este comando")
	}
}
