package routes

import (
	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/api/handlers"
)

func CategoryConfigureRoutes(Router *mux.Router) {
	categoryHandler := handlers.NewCategoryHandler()

	/**
	ENDPOINTS /category/{id} ROUTES
	**/

	// Roteamento para a função GetCategoryByID quando a solicitação é um método GET
	Router.HandleFunc("/category/{id}", categoryHandler.GetCategoryByID).Methods("GET")

	// Roteamento para a função UpdateCategoryByID quando a solicitação é um método PUT
	Router.HandleFunc("/category/{id}", categoryHandler.UpdateCategoryByID).Methods("PUT")

	// Roteamento para a função DeleteCategoryByID quando a solicitação é um método DELETE
	Router.HandleFunc("/category/{id}", categoryHandler.DeleteCategoryByID).Methods("DELETE")

	/**
	ENDPOINTS /categories/ ROUTES
	**/

	// Roteamento para a função GetCategories quando a solicitação é um método GET
	Router.HandleFunc("/categories/", categoryHandler.GetCategories).Methods("GET")

	// Roteamento para a função CreateCategory quando a solicitação é um método POST
	Router.HandleFunc("/categories/", categoryHandler.CreateCategory).Methods("POST")
}
