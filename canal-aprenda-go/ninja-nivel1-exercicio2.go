/*
Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
 1. Identificador "x" deverá ter tipo int
 2. Identificador "y" deverá ter tipo string
 3. Identificador "z" deverá ter tipo bool

Na função main:
 1. Demonstre os valores de cada identificador
 2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam? Em golang os valores para cada váriavel não inicializada, são chamadas de valor "0" cada tipo.

Exemplo: "valor zero de int", "valor zero de string"...
*/
package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%v\n%v\n%v", x, y, z)
}