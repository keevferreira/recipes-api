package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/keevferreira/recipes-api/internal/database"
)

type Recipe struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Ingredients []Ingredient `json:"ingredients"`
	Category    Category     `json:"category"`
	PrepTime    int          `json:"prep_time"`
	Difficulty  string       `json:"difficulty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type Recipes []Recipe

// GetRecipeByID retrieves a recipe by its ID from the database.
func GetRecipeByID(id int) (Recipe, error) {
	var recipe Recipe

	err := database.DB.QueryRow("SELECT id, title, description, prep_time, difficulty, created_at, updated_at FROM recipes WHERE id = ?", id).Scan(
		&recipe.ID, &recipe.Title, &recipe.Description, &recipe.PrepTime, &recipe.Difficulty, &recipe.CreatedAt, &recipe.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return Recipe{}, fmt.Errorf("recipe with ID %d not found", id)
	case err != nil:
		return Recipe{}, err
	}

	// Fetch ingredients and category from the database
	recipe.Ingredients, err = GetIngredientsByRecipeID(id)
	if err != nil {
		return Recipe{}, err
	}

	recipe.Category, err = GetCategoryByRecipeID(id)
	if err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

// UpdateRecipeByID updates a recipe by its ID in the database.
func UpdateRecipeByID(id int, updatedRecipe Recipe) error {
	// Update the recipe in the database
	_, err := database.DB.Exec("UPDATE recipes SET title=?, description=?, prep_time=?, difficulty=?, updated_at=? WHERE id=?",
		updatedRecipe.Title, updatedRecipe.Description, updatedRecipe.PrepTime, updatedRecipe.Difficulty, time.Now(), id)
	if err != nil {
		return err
	}

	// Update ingredients and category (if necessary)
	err = UpdateIngredientsByRecipeID(id, updatedRecipe.Ingredients)
	if err != nil {
		return err
	}

	err = UpdateCategoryByRecipeID(id, updatedRecipe.Category)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecipeByID deletes a recipe by its ID from the database.
func DeleteRecipeByID(id int) error {
	// Delete the recipe from the database
	_, err := database.DB.Exec("DELETE FROM recipes WHERE id=?", id)
	if err != nil {
		return err
	}

	// Delete associated ingredients and category
	err = DeleteIngredientsByRecipeID(id)
	if err != nil {
		return err
	}

	err = DeleteCategoryByRecipeID(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllRecipes retrieves all recipes from the database.
func GetAllRecipes() ([]Recipe, error) {
	var recipes []Recipe

	rows, err := database.DB.Query("SELECT id, title, description, prep_time, difficulty, created_at, updated_at FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var recipe Recipe
		err := rows.Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.PrepTime, &recipe.Difficulty, &recipe.CreatedAt, &recipe.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Fetch ingredients and category for each recipe
		recipe.Ingredients, err = GetIngredientsByRecipeID(recipe.ID)
		if err != nil {
			return nil, err
		}

		recipe.Category, err = GetCategoryByRecipeID(recipe.ID)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

// InsertIngredientsAndCategory insere ingredientes e associa uma categoria a uma receita pelo ID da receita no banco de dados.
func InsertIngredientsAndCategory(recipeID int64, ingredients []Ingredient, category Category) error {
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

	// Insere os ingredientes associados à receita
	stmtIngredient, err := tx.Prepare("INSERT INTO ingredients (recipe_id, name, quantity, unit, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtIngredient.Close()

	for _, ingredient := range ingredients {
		_, err := stmtIngredient.Exec(recipeID, ingredient.Name, ingredient.Quantity, ingredient.Unit, time.Now(), time.Now())
		if err != nil {
			return err
		}
	}

	// Atualiza a categoria associada à receita
	_, err = tx.Exec("UPDATE recipes SET category_id = ? WHERE id = ?", category.ID, recipeID)
	if err != nil {
		return err
	}

	// Se tudo correr bem, faz commit da transação
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// CreateRecipe creates a new recipe in the database.
func CreateRecipe(recipe Recipe) (int, error) {
	// Insert the recipe into the database
	result, err := database.DB.Exec("INSERT INTO recipes (title, description, prep_time, difficulty, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		recipe.Title, recipe.Description, recipe.PrepTime, recipe.Difficulty, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	// Retrieve the ID of the newly inserted recipe
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Insert ingredients and category
	err = InsertIngredientsAndCategory(id, recipe.Ingredients, recipe.Category)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
