package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/models"
	"github.com/keevferreira/recipes-api/internal/utils"
)

// RecipeHandler é uma estrutura para manipulação de receitas.
type RecipeHandler struct {
	// Aqui você pode incluir dependências, como um serviço de banco de dados.
	// Por simplicidade, vamos apenas lidar com manipulação de receitas diretamente.
}

// NewRecipeHandler cria uma nova instância de RecipeHandler.
func NewRecipeHandler() *RecipeHandler {
	return &RecipeHandler{}
}

// CreateRecipe cria uma nova receita.
func (rh *RecipeHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	// Decodifica o corpo da solicitação em um objeto Recipe
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Aqui você pode validar a entrada, como garantir que campos obrigatórios estejam presentes, etc.

	// Salve a receita no banco de dados ou onde quer que você esteja armazenando.
	// Suponha que haja uma função SaveRecipe no modelo de dados que manipula a persistência.
	_, err = models.CreateRecipe(recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna um status de sucesso
	w.WriteHeader(http.StatusCreated)
}

// GetRecipes recupera todas as receitas.
func (rh *RecipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	// Recupere as receitas do banco de dados ou de onde quer que você esteja armazenando.
	recipes, err := models.GetAllRecipes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serializa as receitas para JSON e envia a resposta.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (rh *RecipeHandler) GetRecipeByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da receita dos parâmetros da URL
	recipeID := mux.Vars(r)["id"]

	// Converte o ID da receita de string para inteiro
	id := utils.StringToInt(recipeID)

	// Aqui, estamos simulando a busca de uma receita em um banco de dados.
	recipe, err := models.GetRecipeByID(id)
	if err != nil {
		// Se ocorrer um erro ao buscar a receita, retorna um erro interno do servidor
		http.Error(w, "Erro ao buscar a receita", http.StatusInternalServerError)
		return
	}

	// Se a receita não for encontrada, retorna um erro de não encontrado
	if recipe.ID == 0 {
		http.Error(w, "Receita não encontrada", http.StatusNotFound)
		return
	}

	// Se a receita for encontrada, retorne-a como resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func (rh *RecipeHandler) UpdateRecipeByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da receita dos parâmetros da URL
	recipeID := mux.Vars(r)["id"]

	// Decodifica o corpo da solicitação em um objeto Recipe
	var updatedRecipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&updatedRecipe)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Supondo que você tenha uma função que atualize a receita com base no ID
	// Aqui, estamos simulando a atualização de uma receita em um banco de dados.
	// Você precisaria implementar essa função de acordo com sua lógica de negócios e banco de dados.
	err = models.UpdateRecipeByID(utils.StringToInt(recipeID), updatedRecipe)
	if err != nil {
		// Se ocorrer um erro ao atualizar a receita, retorna um erro interno do servidor
		http.Error(w, "Erro ao atualizar a receita", http.StatusInternalServerError)
		return
	}

	// Retorna a receita atualizada como resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedRecipe)
}

func (rh *RecipeHandler) DeleteRecipeByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da receita dos parâmetros da URL
	recipeID := mux.Vars(r)["id"]

	// Supondo que você tenha uma função que delete a receita com base no ID
	// Aqui, estamos simulando a exclusão de uma receita em um banco de dados.
	// Você precisaria implementar essa função de acordo com sua lógica de negócios e banco de dados.
	err := models.DeleteRecipeByID(utils.StringToInt(recipeID))
	if err != nil {
		// Se ocorrer um erro ao deletar a receita, retorna um erro interno do servidor
		http.Error(w, "Erro ao deletar a receita", http.StatusInternalServerError)
		return
	}

	// Se a receita foi deletada com sucesso, retorne um status OK
	w.WriteHeader(http.StatusOK)
}
