/*
Crie um programa em Go que declare três variáveis de diferentes tipos (um inteiro, um float e uma string).
Utilize a inferência de tipos para declarar essas variáveis.
Em seguida, imprima o valor e o tipo de cada uma delas usando a função TypeOf do pacote reflect.
*/

package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Neberson Andrade"
	idade := 29
	versao := 2.1

	fmt.Println("Seu nome é ", nome, "você tem", idade, "anos", " versão do sistema:", versao)
	fmt.Println(" nome é tipo", reflect.TypeOf(nome), " idade é tipo", reflect.TypeOf(idade), " versao é tipo", reflect.TypeOf(versao))
}
