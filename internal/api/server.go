package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeServer(port string, routerControler *mux.Router) {
	fmt.Printf("Servidor escutando em %s", port)
	err := http.ListenAndServe(":"+port, routerControler)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
