package config

import (
	"fmt"
	"os"
)

// Config contém as configurações da aplicação
type Config struct {
	SERVER_PORT string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
}

// loadEnvVar lê uma variável de ambiente e a atualiza no Config se não for vazia
func loadEnvVar(envVarName string, configField *string) {
	envVarValue := os.Getenv(envVarName)
	if envVarValue != "" {
		*configField = envVarValue
	}
}

// LoadConfig carrega as configurações da aplicação a partir do arquivo de configuração
func LoadConfig() *Config {
	// Cria uma configuração padrão
	config := &Config{
		SERVER_PORT: "8080",
		DB_NAME:     "database",
		DB_HOST:     "localhost",
		DB_PORT:     "5432",
		DB_USER:     "user",
		DB_PASSWORD: "password",
	}

	// Carrega as variáveis de ambiente usando a função loadEnvVar
	loadEnvVar("SERVER_PORT", &config.SERVER_PORT)
	loadEnvVar("DB_NAME", &config.DB_NAME)
	loadEnvVar("DB_HOST", &config.DB_HOST)
	loadEnvVar("DB_PORT", &config.DB_PORT)
	loadEnvVar("DB_USER", &config.DB_USER)
	loadEnvVar("DB_PASSWORD", &config.DB_PASSWORD)

	return config
}

// GetConnectionString irá gerar a connection string
func GetConnectionString(config *Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)
}
