package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia2"},
	}
	database.ConectaComBancoDeDados()
	println("INICIANDO SERVIDOR REST GO")
	routes.HandleRequest()
}
