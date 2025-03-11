// Echo1 exibe seus argumentos de linha de comando.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	tempoInicio := time.Now()
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("Argumento na posição: ", i, " valor: ", os.Args[i])
	}

	fmt.Println(s)

	tempoExecucao := time.Now().Sub(tempoInicio)
	fmt.Println("Tempo de execução:", tempoExecucao)

}
