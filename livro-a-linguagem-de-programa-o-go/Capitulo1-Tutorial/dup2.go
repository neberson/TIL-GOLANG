//Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma vez
//na entrada-padrão. Ele lê de stdin ou de uma lista de arquivos nomeados.

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Line struct {
	count    int
	filename string
}

func main() {
	counts := make(map[string]Line)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, n.filename, line)
		}
	}
}

func countLines(f *os.File, counts map[string]Line) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = Line{count: counts[input.Text()].count + 1, filename: f.Name()}
	}
	//NOTA: ignorando erros em potencial de input.Err()
}
