package config

import (
	"os"
)

// Config contém as configurações da aplicação
type Config struct {
	ServerPort         string
	DBConnectionString string
}

// LoadConfig carrega as configurações da aplicação a partir do arquivo de configuração
func LoadConfig() *Config {
	// Cria uma configuração padrão
	config := &Config{
		ServerPort:         "8080",
		DBConnectionString: "postgres://user:password@localhost:5432/recipes_db?sslmode=disable",
	}

	// Verifica se há uma variável de ambiente para a porta do servidor
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort != "" {
		config.ServerPort = serverPort
	}

	// Verifica se há uma variável de ambiente para a string de conexão com o banco de dados
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString != "" {
		config.DBConnectionString = dbConnectionString
	}

	return config
}
