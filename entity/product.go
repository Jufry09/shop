package entity

import (
	"time"

	"github.com/thedevsaddam/govalidator"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	Category    Category  `json:"category"`
}

func (Product) Rules() govalidator.MapData {
	return govalidator.MapData{
		"name":        []string{"required"},
		"price":       []string{"required"},
		"stock":       []string{"required"},
		"category_id": []string{"required"},
	}
}
