package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/keevferreira/recipes-api/internal/database"
)

type Ingredient struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Quantity  float64   `json:"quantity"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Ingredients []Ingredient

// GetIngredientByID recupera um ingrediente pelo seu ID do banco de dados.
func GetIngredientByID(id int) (Ingredient, error) {
	var ingredient Ingredient

	err := database.DB.QueryRow("SELECT id, name, quantity, unit, created_at, updated_at FROM ingredients WHERE id = ?", id).Scan(
		&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.Unit, &ingredient.CreatedAt, &ingredient.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return Ingredient{}, fmt.Errorf("ingredient with ID %d not found", id)
	case err != nil:
		return Ingredient{}, err
	}

	return ingredient, nil
}

// UpdateIngredientByID atualiza um ingrediente pelo seu ID no banco de dados.
func UpdateIngredientByID(id int, updatedIngredient Ingredient) error {
	// Atualiza o ingrediente no banco de dados
	_, err := database.DB.Exec("UPDATE ingredients SET name=?, quantity=?, unit=?, updated_at=? WHERE id=?",
		updatedIngredient.Name, updatedIngredient.Quantity, updatedIngredient.Unit, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteIngredientByID exclui um ingrediente pelo seu ID do banco de dados.
func DeleteIngredientByID(id int) error {
	// Exclui o ingrediente do banco de dados
	_, err := database.DB.Exec("DELETE FROM ingredients WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllIngredients recupera todos os ingredientes do banco de dados.
func GetAllIngredients() ([]Ingredient, error) {
	var ingredients []Ingredient

	rows, err := database.DB.Query("SELECT id, name, quantity, unit, created_at, updated_at FROM ingredients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ingredient Ingredient
		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.Unit, &ingredient.CreatedAt, &ingredient.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ingredients, nil
}

// CreateIngredient cria um novo ingrediente no banco de dados.
func CreateIngredient(ingredient Ingredient) (int, error) {
	// Insere o ingrediente no banco de dados
	result, err := database.DB.Exec("INSERT INTO ingredients (name, quantity, unit, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		ingredient.Name, ingredient.Quantity, ingredient.Unit, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	// Recupera o ID do novo ingrediente inserido
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetIngredientsByRecipeID(recipeID int) ([]Ingredient, error) {
	var ingredients []Ingredient

	rows, err := database.DB.Query("SELECT id, name, quantity, unit, created_at, updated_at FROM ingredients WHERE recipe_id = ?", recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ingredient Ingredient
		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.Unit, &ingredient.CreatedAt, &ingredient.UpdatedAt)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ingredients, nil
}

func UpdateIngredientsByRecipeID(recipeID int, updatedIngredients []Ingredient) error {
	// Inicia uma transação
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			// Se houver um erro, faz rollback da transação
			tx.Rollback()
			return
		}
	}()

	// Primeiro, exclua os ingredientes existentes para a receita
	_, err = tx.Exec("DELETE FROM ingredients WHERE recipe_id = ?", recipeID)
	if err != nil {
		return err
	}

	// Agora, insira os ingredientes atualizados
	stmt, err := tx.Prepare("INSERT INTO ingredients (recipe_id, name, quantity, unit, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, ingredient := range updatedIngredients {
		_, err := stmt.Exec(recipeID, ingredient.Name, ingredient.Quantity, ingredient.Unit, time.Now(), time.Now())
		if err != nil {
			return err
		}
	}

	// Se tudo correr bem, faz commit da transação
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// DeleteIngredientsByRecipeID exclui os ingredientes associados a uma receita pelo seu ID no banco de dados.
func DeleteIngredientsByRecipeID(recipeID int) error {
	// Exclui os ingredientes associados à receita
	_, err := database.DB.Exec("DELETE FROM ingredients WHERE recipe_id = ?", recipeID)
	if err != nil {
		return err
	}

	return nil
}
