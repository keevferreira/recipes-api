package main

import (
	"github.com/keevferreira/recipes-api/config"
	"github.com/keevferreira/recipes-api/internal/api"
	"github.com/keevferreira/recipes-api/internal/database"
	"github.com/keevferreira/recipes-api/internal/router"
)

var GlobalENVConfig *config.Config

func main() {
	//Carrega as variáveis do arquivo .env para o OS
	GlobalENVConfig = config.LoadConfig()
	databaseConnectionString := config.GetConnectionString(GlobalENVConfig)
	database.Connect(databaseConnectionString)
	routerControler := router.CreateNewRouter()
	router.ConfigureRoutes(routerControler)
	api.InitializeServer(GlobalENVConfig.SERVER_PORT, routerControler)

}
