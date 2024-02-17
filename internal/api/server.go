package api

import (
	"fmt"
	"log"
	"net/http"
)

func InitializeServer(port string) {
	fmt.Printf("Servidor escutando em %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
