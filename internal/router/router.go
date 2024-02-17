package router

import (
	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/api"
	"github.com/keevferreira/recipes-api/internal/router/routes"
)

// ConfigureRoutes configura todas as rotas da API.
func ConfigureRoutes() {
	Router := mux.NewRouter()

	//Middleware
	Router.Use(api.LoggingMiddleware)

	//Rotas
	routes.RecipesConfigureRoutes(Router)
	routes.IngredientsConfigureRoutes(Router)
	routes.CategoryConfigureRoutes(Router)
}
