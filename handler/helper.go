package handler

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
)

type HasRules interface {
	Rules() govalidator.MapData
}

func validateRequest(c echo.Context, data HasRules) url.Values {
	opts := govalidator.Options{
		Request: c.Request(), // request object
		Rules:   data.Rules(),
		Data:    data,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()

	return e
}
