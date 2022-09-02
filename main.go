package main

import (
	"Alura-API-Gin_go/database"
	"Alura-API-Gin_go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
