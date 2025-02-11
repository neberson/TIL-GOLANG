/*
Exercício 2: Classificação de Notas
Desenvolva um programa que peça ao usuário para inserir uma nota (de 0 a 10) e, em seguida, classifique a nota em "Aprovado", "Reprovado" ou "Recuperação". Considere as seguintes regras:

Nota maior ou igual a 7: "Aprovado"
Nota menor que 7 e maior ou igual a 4: "Recuperação"
Nota menor que 4: "Reprovado"
*/

package main

import "fmt"

func main() {
	var nota int = 0

	fmt.Println("Digite a nota do aluno: ")
	fmt.Scan(&nota)

	if nota >= 7 {
		fmt.Println("Aluno Aprovado")
	} else if nota < 7 && nota >= 4 {
		fmt.Println("Aluno em recuperação")
	} else {
		fmt.Println("Aluno reprovado")
	}
}
