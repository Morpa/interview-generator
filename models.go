package main

import "github.com/google/uuid"

// Category representa as categorias válidas para uma questão
type Category string

const (
	CategoryJS    Category = "JS"
	CategoryReact Category = "React"
)

// IsValid verifica se a categoria é válida (JS ou React)
func (c Category) IsValid() bool {
	return c == CategoryJS || c == CategoryReact
}

type Question struct {
	ID          uuid.UUID `json:"id"`
	Question    string    `json:"question"`
	Explanation string    `json:"explanation"`
	Example     string    `json:"example"`
	Category    Category  `json:"category"`
}
