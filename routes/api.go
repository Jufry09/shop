package routes

import (
	"github.com/jufry09/shop/handler"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoute(e *echo.Echo, db *gorm.DB) {
	userHandler := handler.NewUserHandler(db)

	e.GET("/user", userHandler.Index)
	e.POST("/user", userHandler.Create)

}
