package api

import (
	"encoding/json"
	"net/http"
	"recipes-app/internal/database"
	"recipes-app/internal/model"
)

// Handler representa os manipuladores da API
type Handler struct {
	DB database.Database
}

// NewHandler cria uma nova instância do manipulador da API
func NewHandler(db database.Database) *Handler {
	return &Handler{
		DB: db,
	}
}

// GetRecipeHandler manipula a requisição para obter uma receita por ID
func (h *Handler) GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da receita da URL
	// Aqui vou implementar a lógica para extrair o ID da requisição,
	// parseá-lo e verificar se é válido.
	// Por exemplo:
	// recipeID := parseRecipeIDFromURL(r)

	// Obtem a receita do banco de dados usando o ID
	recipe, err := h.DB.GetRecipeByID(recipeID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar a receita"))
		return
	}

	// Codifica a receita em JSON e envia como resposta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

// CreateRecipeHandler manipula a requisição para criar uma nova receita
func (h *Handler) CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// Decodifica a receita do corpo da requisição
	var recipe model.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao decodificar o corpo da requisição"))
		return
	}

	// Insere a receita no banco de dados
	err = h.DB.CreateRecipe(&recipe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao criar a receita"))
		return
	}

	// Retorna o ID da receita recém-criada como resposta
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(recipe.ID))
}

// Outros manipuladores de requisições HTTP, como UpdateRecipeHandler, DeleteRecipeHandler, etc., devo implementar da mesma forma.
