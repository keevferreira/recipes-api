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
	Categories  []Category   `json:"category"`
	PrepTime    int          `json:"prep_time"`
	Difficulty  string       `json:"difficulty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// Recipes represents a collection of recipes.
type Recipes []Recipe

// GetRecipeByID retrieves a recipe by its ID from the database.
func GetRecipeByID(id int) (Recipe, error) {
	var recipe Recipe

	err := database.DB.QueryRow("SELECT id, title, description, prep_time, difficulty, created_at, updated_at FROM recipes WHERE id = ?", id).
		Scan(&recipe.ID, &recipe.Title, &recipe.Description, &recipe.PrepTime, &recipe.Difficulty, &recipe.CreatedAt, &recipe.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return Recipe{}, fmt.Errorf("recipe with ID %d not found", id)
	case err != nil:
		return Recipe{}, err
	}

	// Fetch ingredients and categories from the database
	recipe.Ingredients, err = GetIngredientsByRecipeID(id)
	if err != nil {
		return Recipe{}, err
	}

	recipe.Categories, err = GetCategoriesByRecipeID(id)
	if err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

// UpdateRecipeByID updates a recipe by its ID in the database.
func UpdateRecipeByID(id int, updatedRecipe Recipe) error {
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

	_, err = tx.Exec("UPDATE recipes SET title=?, description=?, prep_time=?, difficulty=?, updated_at=? WHERE id=?",
		updatedRecipe.Title, updatedRecipe.Description, updatedRecipe.PrepTime, updatedRecipe.Difficulty, time.Now(), id)
	if err != nil {
		return err
	}

	err = UpdateIngredientsByRecipeID(id, updatedRecipe.Ingredients)
	if err != nil {
		return err
	}

	err = UpdateCategoriesByRecipeID(id, updatedRecipe.Categories)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecipeByID deletes a recipe by its ID from the database.
func DeleteRecipeByID(id int) error {
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

	_, err = tx.Exec("DELETE FROM recipes WHERE id=?", id)
	if err != nil {
		return err
	}

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

		recipe.Ingredients, err = GetIngredientsByRecipeID(recipe.ID)
		if err != nil {
			return nil, err
		}

		recipe.Categories, err = GetCategoriesByRecipeID(recipe.ID)
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

// CreateRecipe creates a new recipe in the database.
func CreateRecipe(recipe Recipe) (int, error) {
	tx, err := database.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	result, err := tx.Exec("INSERT INTO recipes (title, description, prep_time, difficulty, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		recipe.Title, recipe.Description, recipe.PrepTime, recipe.Difficulty, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	err = InsertIngredientsAndCategories(int(id), recipe.Ingredients, recipe.Categories)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// InsertIngredientsAndCategories inserts ingredients and associates categories with a recipe by its ID in the database.
func InsertIngredientsAndCategories(recipeID int, ingredients []Ingredient, categories []Category) error {
	// Begin a transaction
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			// Rollback the transaction if an error occurs
			tx.Rollback()
			return
		}
	}()

	// Insert ingredients and associate them with the recipe
	for _, ingredient := range ingredients {
		_, err := tx.Exec("INSERT INTO recipeingredients (recipeid, ingredientid) VALUES (?, ?)", recipeID, ingredient.ID)
		if err != nil {
			// Rollback the transaction if an error occurs
			return err
		}
	}

	// Associate categories with the recipe
	for _, category := range categories {
		_, err := tx.Exec("INSERT INTO recipecategories (recipeid, categoryid) VALUES (?, ?)", recipeID, category.ID)
		if err != nil {
			// Rollback the transaction if an error occurs
			return err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		// Rollback the transaction if an error occurs during commit
		tx.Rollback()
		return err
	}

	return nil
}
