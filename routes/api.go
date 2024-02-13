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
	e.PUT("/user/:id", userHandler.Update)
	e.GET("/user/:id", userHandler.Show)
	e.DELETE("/user/:id", userHandler.Delete)

	categoryHandler := handler.NewCategoryHandler(db)
	cat := e.Group("/category")
	cat.GET("", categoryHandler.Index)
	cat.POST("", categoryHandler.Create)

	productHandler := handler.NewProductHandler(db)
	product := e.Group("/product")
	product.GET("", productHandler.Index)
	product.POST("", productHandler.Create)

}
