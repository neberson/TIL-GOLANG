package main

import (
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "João de Andrade", Historia: "Historia 1"},
		{Nome: "Maria de Andrade", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
