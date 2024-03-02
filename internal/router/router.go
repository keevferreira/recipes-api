package router

import (
	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/api"
	"github.com/keevferreira/recipes-api/internal/router/routes"
)

// Cria um novo gerenciador de rotas
func CreateNewRouter() *mux.Router {
	return mux.NewRouter()
}

// ConfigureRoutes configura todas as rotas da API.
func ConfigureRoutes(routerControler *mux.Router) {
	//Middleware
	routerControler.Use(api.LoggingMiddleware)
	//Rotas
	routes.RecipesConfigureRoutes(routerControler)
	routes.IngredientsConfigureRoutes(routerControler)
	routes.CategoryConfigureRoutes(routerControler)
}
