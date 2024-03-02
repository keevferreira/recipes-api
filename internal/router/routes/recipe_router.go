package routes

import (
	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/api/handlers"
)

func RecipesConfigureRoutes(Router *mux.Router) {
	recipeHandler := handlers.NewRecipeHandler()

	/**
	ENDPOINTS /recipe/{id} ROUTES
	**/

	// Roteamento para a função GetRecipeByID quando a solicitação é um método GET
	Router.HandleFunc("/recipe/{id}", recipeHandler.GetRecipeByID).Methods("GET")

	// Roteamento para a função UpdateRecipeByID quando a solicitação é um método PUT
	Router.HandleFunc("/recipe/{id}", recipeHandler.UpdateRecipeByID).Methods("PUT")

	// Roteamento para a função DeleteRecipeByID quando a solicitação é um método DELETE
	Router.HandleFunc("/recipe/{id}", recipeHandler.DeleteRecipeByID).Methods("DELETE")

	/**
	ENDPOINTS /recipes/ ROUTES
	**/

	// Roteamento para a função GetRecipes quando a solicitação é um método GET
	Router.HandleFunc("/recipes/", recipeHandler.GetRecipes).Methods("GET")

	// Roteamento para a função CreateRecipe quando a solicitação é um método POST
	Router.HandleFunc("/recipes/", recipeHandler.CreateRecipe).Methods("POST")
}
