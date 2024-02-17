package main

import (
	"github.com/keevferreira/recipes-api/config"
	"github.com/keevferreira/recipes-api/internal/api"
	"github.com/keevferreira/recipes-api/internal/database"
)

var GlobalENVConfig *config.Config

func main() {
	GlobalENVConfig = config.LoadConfig()
	databaseConnectionString := config.GetConnectionString(GlobalENVConfig)
	database.Connect(databaseConnectionString)
	api.InitializeServer(GlobalENVConfig.SERVER_PORT)
}
