package handler

import (
	"net/http"

	"github.com/jufry09/shop/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductHandler struct {
	db *gorm.DB
}

func NewProductHandler(db *gorm.DB) ProductHandler {
	return ProductHandler{
		db: db,
	}
}

func (h ProductHandler) Create(c echo.Context) error {
	var product entity.Product

	e := validateRequest(c, &product)
	if len(e) > 0 {
		err := map[string]interface{}{"validationError": e}
		return c.JSON(http.StatusBadRequest, err)
	}

	h.db.Create(&product)
	h.db.Preload("Category").First(&product)

	return c.JSON(http.StatusCreated, product)
}

func (h ProductHandler) Index(c echo.Context) error {
	var products []entity.Product

	h.db.Preload("Category").Find(&products)

	return c.JSON(http.StatusCreated, products)
}
