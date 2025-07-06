package main

import (
	"github.com/luacarol/api-go-gin/database"
	"github.com/luacarol/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
