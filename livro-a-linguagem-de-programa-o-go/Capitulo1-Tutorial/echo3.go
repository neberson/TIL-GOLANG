package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	tempoInicio := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	tempoExecucao := time.Now().Sub(tempoInicio)
	fmt.Println("Tempo de execução:", tempoExecucao)
}
