package main

import (
	"log"
	"net/http"
	"recipes-api/config"
	"recipes-api/internal/api"
	"recipes-api/internal/database/postgres"
)

func main() {
	// Carrega as configurações da aplicação
	cfg := config.LoadConfig()

	// Inicializa a conexão com o banco de dados PostgreSQL
	db, err := postgres.NewPostgresDB(cfg.DBConnectionString)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Inicializa os manipuladores da API
	handler := api.NewHandler(db)

	// Registra os manipuladores e middlewares HTTP
	http.Handle("/", api.MiddlewareLogging(handler))

	// Inicia o servidor HTTP
	log.Printf("Servidor escutando em %s", cfg.ServerPort)
	err = http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
