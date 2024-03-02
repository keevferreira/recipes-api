package routes

import (
	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/api/handlers"
)

func IngredientsConfigureRoutes(Router *mux.Router) {
	ingredientHandler := handlers.NewIngredientHandler()

	/**
	ENDPOINTS /ingredient/{id} ROUTES
	**/

	// Roteamento para a função GetIngredientByID quando a solicitação é um método GET
	Router.HandleFunc("/ingredient/{id}", ingredientHandler.GetIngredientByID).Methods("GET")

	// Roteamento para a função UpdateIngredientByID quando a solicitação é um método PUT
	Router.HandleFunc("/ingredient/{id}", ingredientHandler.UpdateIngredientByID).Methods("PUT")

	// Roteamento para a função DeleteIngredientByID quando a solicitação é um método DELETE
	Router.HandleFunc("/ingredient/{id}", ingredientHandler.DeleteIngredientByID).Methods("DELETE")

	/**
	ENDPOINTS /ingredients/ ROUTES
	**/

	// Roteamento para a função GetIngredients quando a solicitação é um método GET
	Router.HandleFunc("/igredients/", ingredientHandler.GetIngredients).Methods("GET")

	// Roteamento para a função CreateIngredient quando a solicitação é um método POST
	Router.HandleFunc("/igredients/", ingredientHandler.CreateIngredient).Methods("POST")
}
