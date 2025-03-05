package main

import (
	"api-go-rest-gin/database"
	"api-go-rest-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
