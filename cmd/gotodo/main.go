package main

import (
	"github.com/haydenrou/gotodo/views"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/", func(c echo.Context) error {
		component := views.Index()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
