package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/keevferreira/recipes-api/internal/database"
)

type Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Categories []Category

// GetCategoryByID recupera uma categoria pelo seu ID do banco de dados.
func GetCategoryByID(id int) (Category, error) {
	var category Category

	err := database.DB.QueryRow("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = ?", id).Scan(
		&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return Category{}, fmt.Errorf("category with ID %d not found", id)
	case err != nil:
		return Category{}, err
	}

	return category, nil
}

// UpdateCategoryByID atualiza uma categoria pelo seu ID no banco de dados.
func UpdateCategoryByID(id int, updatedCategory Category) error {
	// Atualiza a categoria no banco de dados
	_, err := database.DB.Exec("UPDATE categories SET name=?, description=?, updated_at=? WHERE id=?",
		updatedCategory.Name, updatedCategory.Description, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCategoryByID exclui uma categoria pelo seu ID do banco de dados.
func DeleteCategoryByID(id int) error {
	// Exclui a categoria do banco de dados
	_, err := database.DB.Exec("DELETE FROM categories WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllCategories recupera todas as categorias do banco de dados.
func GetAllCategories() ([]Category, error) {
	var categories []Category

	rows, err := database.DB.Query("SELECT id, name, description, created_at, updated_at FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

// CreateCategory cria uma nova categoria no banco de dados.
func CreateCategory(category Category) (int, error) {
	// Insere a categoria no banco de dados
	result, err := database.DB.Exec("INSERT INTO categories (name, description, created_at, updated_at) VALUES (?, ?, ?, ?)",
		category.Name, category.Description, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	// Recupera o ID da nova categoria inserida
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetCategoryByRecipeID(recipeID int) (Category, error) {
	var category Category

	err := database.DB.QueryRow("SELECT c.id, c.name, c.description, c.created_at, c.updated_at FROM categories c JOIN recipes r ON c.id = r.category_id WHERE r.id = ?", recipeID).Scan(
		&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return Category{}, fmt.Errorf("category for recipe with ID %d not found", recipeID)
	case err != nil:
		return Category{}, err
	}

	return category, nil
}

// UpdateCategoryByRecipeID atualiza a categoria associada a uma receita pelo seu ID no banco de dados.
func UpdateCategoryByRecipeID(recipeID int, updatedCategory Category) error {
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

	// Atualiza a categoria associada à receita
	_, err = tx.Exec("UPDATE recipes SET category_id = ? WHERE id = ?", updatedCategory.ID, recipeID)
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

// DeleteCategoryByRecipeID exclui a categoria associada a uma receita pelo seu ID no banco de dados.
func DeleteCategoryByRecipeID(recipeID int) error {
	// Exclui a categoria associada à receita
	_, err := database.DB.Exec("UPDATE recipes SET category_id = NULL WHERE id = ?", recipeID)
	if err != nil {
		return err
	}

	return nil
}
