package main

import (
	"POO/clientes"
	"POO/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteDoNeberson := clientes.Titular{Nome: "Neberson", CPF: "123.456.789-00", Profissao: "Desenvolvedor"}

	contaDoNeberson := contas.ContaCorrente{
		Titular:       clienteDoNeberson,
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}

	contaDoNeberson.Depositar(125.50)

	clienteDaBruna := clientes.Titular{Nome: "Bruna", CPF: "123.456.789-00", Profissao: "Desenvolvedora"}

	contaDaBruna := contas.ContaCorrente{
		Titular:       clienteDaBruna,
		NumeroAgencia: 222,
		NumeroConta:   111222,
	}

	contaDaBruna.Depositar(150)

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

	clienteDoDenis := clientes.Titular{Nome: "Denis", CPF: "987.654.321-00", Profissao: "Programador"}

	contaDoDenis := contas.ContaPoupanca{
		Titular:       clienteDoDenis,
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}

	fmt.Println(contaDoDenis)

	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis)

	PagarBoleto(&contaDoNeberson, 50)
	fmt.Println(contaDoNeberson)
}
