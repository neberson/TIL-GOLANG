package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque

		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) depositar(valorDoDeposito float64) (float64, string) {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return c.saldo, "Depósito realizado com sucesso"
	} else {
		return c.saldo, "Valor do depósito menor que zero"
	}
}

func main() {

	contaDoNeberson := ContaCorrente{
		titular:       "Neberson Andrade",
		numeroAgencia: 589,
		numeroConta:   123456,
		saldo:         125.5,
	}

	contaDaBruna := ContaCorrente{"Bruna Andrade", 222, 111222, 200}

	fmt.Println(contaDoNeberson)
	fmt.Println(contaDaBruna)

	valorDoSaque := 50.0
	fmt.Println(contaDoNeberson.sacar(valorDoSaque))

	fmt.Println(contaDoNeberson)

	saldoAtual, mensagem := contaDoNeberson.depositar(100)

	fmt.Println(mensagem, "Saldo atual:", saldoAtual)
}
