/*
Exercício 1: Verificação de Idade
Crie um programa que solicite ao usuário que insira sua idade e, em seguida,
 verifique se ele é maior de idade (18 anos ou mais) ou menor de idade.
 O programa deve imprimir uma mensagem apropriada para cada caso.
*/

package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)
	if idade >= 18 {
		fmt.Println("Você é maior de idade.")
	} else {
		fmt.Println("Você é menor de idade.")
	}
}
