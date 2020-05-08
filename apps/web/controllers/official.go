package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func Products(c echo.Context) error {
	return c.Render(http.StatusOK, "products.html", nil)
}

func About(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", nil)
}
