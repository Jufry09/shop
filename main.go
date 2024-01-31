package main

import (
	"net/http"

	"github.com/jufry09/shop/config"
	"github.com/jufry09/shop/database"
	"github.com/jufry09/shop/entity"
	"github.com/jufry09/shop/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.GetConfig()
	db := database.NewConnection(conf)

	db.AutoMigrate(
		entity.User{},
	)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.SetupRoute(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}
