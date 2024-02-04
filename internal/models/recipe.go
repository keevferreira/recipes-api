package models

import "time"

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
