package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "João de Andrade", Historia: "Historia 1"},
		{Id: 2, Nome: "Maria de Andrade", Historia: "Historia 2"},
	}

	database.ConectaComBancoDeDados()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
