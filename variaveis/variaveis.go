package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "Neberson Andrade"
	var versao float32 = 1.1
	var idade int = 29

	/*declaração por infenrencia
	var nome = "Neberson Andrade"
	var versao = 1.1 vai usar por padrão float64
	var idade = 29*/

	/*declaração por infenrencia sem a palavra var
	nome := "Neberson Andrade"
	versao := 1.1
	idade := 29*/

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

}

/*-------------- Exercicio proposto pela IA Luri ------------------------*/

/*Veja o código abaixo:

package main

import "fmt"

func main() {
	var nome string = "José"
	var peso float32 = 75.4
	fmt.Println("Olá, eu sou o", nome)
}COPIAR CÓDIGO
Esse programa gera um erro. Você sabe identificar a causa?*/

/*Resposta: Sim, o erro ocorre por que a variável peso foi declarada e não está sendo utilizada. Diferente de outras linguagens o GO reconhece como erro e não como um warning */

/*Correção: Você acertou em cheio!

Neberson, sua resposta está correta, parabéns! O erro nesse código de Go ocorre porque a variável `peso` foi declarada, mas não está sendo utilizada em nenhum momento. No Go, diferente de outras linguagens, declarar uma variável e não utilizá-la é considerado um erro, e não apenas um aviso. O compilador do Go é muito rígido quanto a esse tipo de situação, pois entende que se uma variável foi declarada, ela deve ser utilizada em algum momento do código. Caso contrário, isso pode indicar um problema ou uma parte do código que não está sendo utilizada. Por isso, o Go não permite que o programa seja compilado com variáveis declaradas e não utilizadas. Para corrigir esse erro, basta remover a declaração da variável `peso` ou então utilizá-la em algum momento do código. Parabéns por ter identificado corretamente o problema! Continue estudando e praticando Go, você está no caminho certo.*/
