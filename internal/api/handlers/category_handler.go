package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/keevferreira/recipes-api/internal/models"
	"github.com/keevferreira/recipes-api/internal/utils"
)

// CategoryHandler é uma estrutura para manipulação de categorias de receitas.
type CategoryHandler struct {
	// Aqui você pode incluir dependências, como um serviço de banco de dados.
	// Por simplicidade, vamos apenas lidar com manipulação de receitas diretamente.
}

// NewCategoryHandler cria uma nova instância de CategoryHandler.
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

// CreateCategory cria uma nova receita.
func (ch *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	// Decodifica o corpo da solicitação em um objeto Category
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Aqui você pode validar a entrada, como garantir que campos obrigatórios estejam presentes, etc.

	// Salve a receita no banco de dados ou onde quer que você esteja armazenando.
	// Suponha que haja uma função SaveRecipe no modelo de dados que manipula a persistência.
	_, err = models.CreateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retorna um status de sucesso
	w.WriteHeader(http.StatusCreated)
}

// GetCategories recupera todas as categorias de receitas.
func (ch *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	// Recupere as categorias de receitas do banco de dados ou de onde quer que você esteja armazenando.
	recipes, err := models.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serializa as categorias de receitas para JSON e envia a resposta.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (ch *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da categoria dos parâmetros da URL
	categoryID := mux.Vars(r)["id"]

	// Converte o ID da categoria de string para inteiro
	id := utils.StringToInt(categoryID)

	// Aqui, estamos simulando a busca de uma categoria em um banco de dados.
	category, err := models.GetCategoryByID(id)
	if err != nil {
		// Se ocorrer um erro ao buscar a categoria, retorna um erro interno do servidor
		http.Error(w, "Erro ao buscar a categoria", http.StatusInternalServerError)
		return
	}

	// Se a categoria não for encontrada, retorna um erro de não encontrado
	if category.ID == 0 {
		http.Error(w, "Categoria não encontrada", http.StatusNotFound)
		return
	}

	// Se a receita for encontrada, retorne-a como resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (ch *CategoryHandler) UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da categoria dos parâmetros da URL
	categoryID := mux.Vars(r)["id"]

	// Decodifica o corpo da solicitação em um objeto Category
	var updatedCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da solicitação", http.StatusBadRequest)
		return
	}

	// Supondo que você tenha uma função que atualize a categoria com base no ID
	// Aqui, estamos simulando a atualização de uma categoria em um banco de dados.
	// Você precisaria implementar essa função de acordo com sua lógica de negócios e banco de dados.
	err = models.UpdateCategoryByID(utils.StringToInt(categoryID), updatedCategory)
	if err != nil {
		// Se ocorrer um erro ao atualizar a categoria, retorna um erro interno do servidor
		http.Error(w, "Erro ao atualizar a categoria", http.StatusInternalServerError)
		return
	}

	// Retorna a categoria atualizada como resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCategory)
}

func (ch *CategoryHandler) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID da categoria dos parâmetros da URL
	categoryID := mux.Vars(r)["id"]

	// Supondo que você tenha uma função que delete a categoria com base no ID
	// Aqui, estamos simulando a exclusão de uma categoria em um banco de dados.
	// Você precisaria implementar essa função de acordo com sua lógica de negócios e banco de dados.
	err := models.DeleteCategoryByID(utils.StringToInt(categoryID))
	if err != nil {
		// Se ocorrer um erro ao deletar a categoria, retorna um erro interno do servidor
		http.Error(w, "Erro ao deletar a categoria", http.StatusInternalServerError)
		return
	}

	// Se a categoria foi deletada com sucesso, retorne um status OK
	w.WriteHeader(http.StatusOK)
}
