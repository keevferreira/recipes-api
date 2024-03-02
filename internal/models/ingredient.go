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

func GetIngredientByID(id int) (Ingredient, error) {
	var ingredient Ingredient

	err := database.DB.QueryRow("SELECT id, name, quantity, unit, created_at, updated_at FROM ingredients WHERE id = ?", id).
		Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.Unit, &ingredient.CreatedAt, &ingredient.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return Ingredient{}, fmt.Errorf("ingredient with ID %d not found", id)
	case err != nil:
		return Ingredient{}, err
	}

	return ingredient, nil
}

func UpdateIngredientByID(id int, updatedIngredient Ingredient) error {
	_, err := database.DB.Exec("UPDATE ingredients SET name=?, quantity=?, unit=?, updated_at=? WHERE id=?",
		updatedIngredient.Name, updatedIngredient.Quantity, updatedIngredient.Unit, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteIngredientByID(id int) error {
	_, err := database.DB.Exec("DELETE FROM ingredients WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

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

// CreateIngredient creates a new ingredient in the database.
func CreateIngredient(ingredient Ingredient) (int, error) {
	result, err := database.DB.Exec("INSERT INTO ingredients (name, quantity, unit, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		ingredient.Name, ingredient.Quantity, ingredient.Unit, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetIngredientsByRecipeID retrieves ingredients associated with a recipe from the database.
func GetIngredientsByRecipeID(recipeID int) ([]Ingredient, error) {
	var ingredients []Ingredient

	query := `
		SELECT i.id, i.name, i.quantity, i.unit, i.created_at, i.updated_at
		FROM ingredients i
		INNER JOIN recipeingredients ri ON i.id = ri.ingredient_id
		WHERE ri.recipe_id = ?
	`

	rows, err := database.DB.Query(query, recipeID)
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

// UpdateIngredientsByRecipeID updates ingredients associated with a recipe in the database.
func UpdateIngredientsByRecipeID(recipeID int, updatedIngredients []Ingredient) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// First, delete existing ingredients associated with the recipe
	_, err = tx.Exec("DELETE FROM recipeingredients WHERE recipe_id = ?", recipeID)
	if err != nil {
		return err
	}

	// Now, insert updated ingredients into recipeingredients
	stmt, err := tx.Prepare("INSERT INTO recipeingredients (recipe_id, ingredient_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, ingredient := range updatedIngredients {
		_, err := stmt.Exec(recipeID, ingredient.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteIngredientsByRecipeID deletes ingredients associated with a recipe from the database.
func DeleteIngredientsByRecipeID(recipeID int) error {
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	// Delete ingredients from the recipeingredients table
	_, err = tx.Exec("DELETE FROM recipeingredients WHERE recipe_id = ?", recipeID)
	if err != nil {
		return err
	}

	// Delete ingredients from the ingredients table
	_, err = tx.Exec("DELETE FROM ingredients WHERE id IN (SELECT ingredient_id FROM recipeingredients WHERE recipe_id = ?)", recipeID)
	if err != nil {
		return err
	}

	return nil
}
