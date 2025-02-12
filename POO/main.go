package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
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
}
