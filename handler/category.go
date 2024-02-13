package handler

import (
	"net/http"

	"github.com/jufry09/shop/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) CategoryHandler {
	return CategoryHandler{
		db: db,
	}
}

func (h CategoryHandler) Index(c echo.Context) error {
	var categories []entity.Category

	h.db.Find(&categories)
	return c.JSON(http.StatusCreated, categories)
}

func (h CategoryHandler) Create(c echo.Context) error {
	var category entity.Category
	c.Bind(&category)

	h.db.Create(&category)

	return c.JSON(http.StatusCreated, category)
}
