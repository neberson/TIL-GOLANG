/*
Desenvolva um programa que solicite ao usuário que digite seu nome, idade e altura.
Utilize as funções Scanf ou Scan do pacote fmt para ler esses dados.
Após a leitura, imprima uma mensagem personalizada que inclua o nome do usuário e sua idade.
*/

package main

import (
	"fmt"
)

func main() {
	var nome string
	var idade int
	var altura int

	fmt.Println("\n Digite seu nome:")
	fmt.Scan(&nome)

	fmt.Println("\n Digite sua idade:")
	fmt.Scan(&idade)

	fmt.Println("\n Digite sua altura em centimentros:")
	fmt.Scan(&altura)

	fmt.Println("Seu nome é ", nome, "você tem", idade, "anos", " sua altura:", altura)
}
