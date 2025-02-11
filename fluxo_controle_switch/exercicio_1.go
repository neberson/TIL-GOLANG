/*Mário foi contratado por um restaurante para desenvolver um programa que solicita a escolha do cardápio. Ele implementou na linguagem Go e seu código está listado abaixo.*/

package main

import "fmt"

func main() {
	var prato string
	fmt.Println("Digite seu prato preferido...")
	fmt.Println("P - Pizza")
	fmt.Println("H - Hambúrguer")
	fmt.Println("B - Bife com fritas")
	fmt.Println("S - Salada Caesar")
	fmt.Println("F - Salada de Frutas")
	fmt.Println("E - Estrogonofe")
	fmt.Println("O - Outros")
	fmt.Scan(&prato)

	switch prato {
	case "B":
		fmt.Println("Com batatas Palito ou Noisete?")
	case "H":
		fmt.Println("Com Queijo ou com Ovo?")
	case "P":
		fmt.Println("Calabresa ou Napolitana?")
	case "S":
		fmt.Println("Alface ou Rúcula?")
	case "F":
		fmt.Println("Kiwi ou Frango?")
	case "E":
		fmt.Println("Carne ou Frango?")
	case "O":
		fmt.Println("Não gostou de nosso cardápio?")
	default:
		fmt.Println("Não entendi seu paladar...")
	}
}

/*Resposta: Quando o cliente digitar b vai ser impresso "Não entendi seu paladar...", pois o Switch está analisando "B" em caixa alta, não encontrando o valor no cardápio.*/

/*Correção: Você acertou em cheio!
Neberson, sua resposta está correta, parabéns! O motivo de quando o cliente digitar "b" ser impresso "Não entendi seu paladar..." é que o código do Mário está verificando os valores do cardápio em letras maiúsculas, ou seja, "B" em vez de "b". Portanto, quando o cliente digitar "b", o switch não encontrará esse valor no cardápio e entrará no caso default, imprimindo a mensagem "Não entendi seu paladar...". Você demonstrou um bom entendimento do código e da lógica envolvida. Continue assim, estudando e praticando, você está no caminho certo para se tornar um excelente programador!*/
