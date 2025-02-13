package main

import (
	"POO/contas"
	"fmt"
)

func main() {

	contaDoNeberson := contas.ContaCorrente{
		Titular:       "Neberson Andrade",
		NumeroAgencia: 589,
		NumeroConta:   123456,
		Saldo:         125.5,
	}

	contaDaBruna := contas.ContaCorrente{
		Titular:       "Bruna Andrade",
		NumeroAgencia: 222,
		NumeroConta:   111222,
		Saldo:         200,
	}

	fmt.Println(contaDoNeberson)
	fmt.Println(contaDaBruna)

	status := contaDoNeberson.Transferir(20, &contaDaBruna)

	fmt.Println(status)

	valorDoSaque := 50.0
	fmt.Println(contaDoNeberson.Sacar(valorDoSaque))

	fmt.Println(contaDoNeberson)

	saldoAtual, mensagem := contaDoNeberson.Depositar(100)

	fmt.Println(mensagem, "Saldo atual:", saldoAtual)

	fmt.Println(contaDoNeberson)
	fmt.Println(contaDaBruna)
}
