package handler

import (
	"net/http"

	"github.com/jufry09/shop/entity"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) UserHandler {
	return UserHandler{
		db: db,
	}
}

func (h UserHandler) Index(c echo.Context) error {
	var users []entity.User
	h.db.Find(&users)
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}

func (h UserHandler) Create(c echo.Context) error {
	var user entity.User
	c.Bind(&user)

	user.Password, _ = HashPassword(user.Password)
	err := h.db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, "gagal insert")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": user})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
