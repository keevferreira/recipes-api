package database

import (
	"database/sql"

	postgres "github.com/keevferreira/recipes-api/internal/database/postgres"
	"github.com/keevferreira/recipes-api/internal/utils"
)

type Database interface {
	Connect(connectionString string) (*sql.DB, error)
	Disconnect(db *sql.DB) error
}

var DB *sql.DB

func Connect(connectionString string) {
	var err error
	DB, err = postgres.ConnectToPostgresDB(connectionString)
	utils.TreatNilObjectError(err, "Não foi possível conectar no banco de dados")
}

func Disconnect(db *sql.DB) {
	postgres.DisconnectPostgresDB(DB)
}
