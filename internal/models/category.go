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

func GetCategoryByID(id int) (Category, error) {
	var category Category

	err := database.DB.QueryRow("SELECT id, name, description, created_at, updated_at FROM categories WHERE id = ?", id).
		Scan(&category.ID, &category.Name, &category.Description, &category.CreatedAt, &category.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		return Category{}, fmt.Errorf("category with ID %d not found", id)
	case err != nil:
		return Category{}, err
	}

	return category, nil
}

func UpdateCategoryByID(id int, updatedCategory Category) error {
	_, err := database.DB.Exec("UPDATE categories SET name=?, description=?, updated_at=? WHERE id=?",
		updatedCategory.Name, updatedCategory.Description, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteCategoryByID(id int) error {
	_, err := database.DB.Exec("DELETE FROM categories WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

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

func CreateCategory(category Category) (int, error) {
	result, err := database.DB.Exec("INSERT INTO categories (name, description, created_at, updated_at) VALUES (?, ?, ?, ?)",
		category.Name, category.Description, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func GetCategoriesByRecipeID(recipeID int) ([]Category, error) {
	query := `
		SELECT c.id, c.name, c.description, c.created_at, c.updated_at
		FROM categories c
		INNER JOIN recipecategories rc ON c.id = rc.categoryid
		WHERE rc.recipeid = ?
	`

	rows, err := database.DB.Query(query, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
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

func UpdateCategoriesByRecipeID(recipeID int, updatedCategories []Category) error {
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

	_, err = tx.Exec("DELETE FROM recipecategories WHERE recipeid = ?", recipeID)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO recipecategories (recipeid, categoryid, updatedAt) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, category := range updatedCategories {
		_, err := stmt.Exec(recipeID, category.ID, time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteCategoryByRecipeID(recipeID int) error {
	_, err := database.DB.Exec("DELETE FROM recipecategories WHERE recipeid = ?", recipeID)
	if err != nil {
		return err
	}

	return nil
}
