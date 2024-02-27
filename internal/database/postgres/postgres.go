package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/keevferreira/recipes-api/internal/utils"
	_ "github.com/lib/pq"
)

func ConnectToPostgresDB(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	utils.TreatNilObjectError(err, "Falha ao conectar ao banco de dados: %v")
	log.Println("Conexão com o banco de dados PostgreSQL estabelecida")
	return db, err
}

func DisconnectPostgresDB(DB *sql.DB) {
	if DB != nil {
		DB.Close()
		fmt.Print("Disconnected from the database")
	}
}
