package main

import (
	"github.com/luacarol/api-go-gin/models"
	"github.com/luacarol/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Luana dos Anjos", CPF: "00000000000", RG: "000000000"},
		{Nome: "Matheus dos Anjos", CPF: "11111111111", RG: "111111111"},
	}
	routes.HandleRequests()
}
