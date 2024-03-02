package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/models"
	"github.com/keevferreira/recipes-api/internal/utils"
)

// RecipeHandler é uma estrutura para manipulação de ingredientes.
type IngredientHandler struct {
	// Aqui você pode incluir dependências, como um serviço de banco de dados.
	// Por simplicidade, vamos apenas lidar com manipulação de ingredientes diretamente.
}

// NewIngredientHandler cria uma nova instância de IngredientHandler.
func NewIngredientHandler() *IngredientHandler {
	return &IngredientHandler{}
}

// CreateIngredient cria uma nova receita.
func (ih *IngredientHandler) CreateIngredient(w http.ResponseWriter, r *http.Request) {
	// Decodifica o corpo da solicitação em um objeto Ingredient
	var ingredient models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Aqui você pode validar a entrada, como garantir que campos obrigatórios estejam presentes, etc.

	// Salve a receita no banco de dados ou onde quer que você esteja armazenando.
	// Suponha que haja uma função SaveRecipe no modelo de dados que manipula a persistência.
	_, err = models.CreateIngredient(ingredient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna um status de sucesso
	w.WriteHeader(http.StatusCreated)
}

// GetIngredients recupera todos os ingredientes.
func (ih *IngredientHandler) GetIngredients(w http.ResponseWriter, r *http.Request) {
	// Recupere os ingredientes do banco de dados ou de onde quer que você esteja armazenando.
	ingredients, err := models.GetAllIngredients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serializa as receitas para JSON e envia a resposta.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}

func (ih *IngredientHandler) GetIngredientByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID do ingrediente dos parâmetros da URL
	ingredientID := mux.Vars(r)["id"]

	// Converte o ID do ingrediente de string para inteiro
	id := utils.StringToInt(ingredientID)

	// Aqui, estamos simulando a busca de um ingrediente em um banco de dados.
	ingredient, err := models.GetIngredientByID(id)
	if err != nil {
		// Se ocorrer um erro ao buscar o ingrediente, retorna um erro interno do servidor
		http.Error(w, "Erro ao buscar o ingrediente", http.StatusInternalServerError)
		return
	}

	// Se o ingrediente não for encontrado, retorna um erro de não encontrado
	if ingredient.ID == 0 {
		http.Error(w, "Ingrediente não encontrado", http.StatusNotFound)
		return
	}

	// Se o ingrediente for encontrado, retorne-a como resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredient)
}

func (ih *IngredientHandler) UpdateIngredientByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID do ingrediente dos parâmetros da URL
	ingredientID := mux.Vars(r)["id"]

	// Decodifica o corpo da solicitação em um objeto Ingredient
	var updatedIngredient models.Ingredient
	err := json.NewDecoder(r.Body).Decode(&updatedIngredient)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Supondo que você tenha uma função que atualize o ingrediente com base no ID
	// Aqui, estamos simulando a atualização de um ingrediente em um banco de dados.
	// Você precisaria implementar essa função de acordo com sua lógica de negócios e banco de dados.
	err = models.UpdateIngredientByID(utils.StringToInt(ingredientID), updatedIngredient)
	if err != nil {
		// Se ocorrer um erro ao atualizar o ingrediente, retorna um erro interno do servidor
		http.Error(w, "Erro ao atualizar o ingrediente", http.StatusInternalServerError)
		return
	}

	// Retorna a receita atualizada como resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedIngredient)
}

func (ri *IngredientHandler) DeleteIngredientByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID do ingrediente dos parâmetros da URL
	ingredientID := mux.Vars(r)["id"]

	// Supondo que você tenha uma função que delete o ingrediente com base no ID
	// Aqui, estamos simulando a exclusão de um ingrediente em um banco de dados.
	// Você precisaria implementar essa função de acordo com sua lógica de negócios e banco de dados.
	err := models.DeleteIngredientByID(utils.StringToInt(ingredientID))
	if err != nil {
		// Se ocorrer um erro ao deletar o ingrediente, retorna um erro interno do servidor
		http.Error(w, "Erro ao deletar o ingredinte", http.StatusInternalServerError)
		return
	}

	// Se o ingrediente foi deletada com sucesso, retorne um status OK
	w.WriteHeader(http.StatusOK)
}
