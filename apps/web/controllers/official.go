package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	data := map[string]interface{}{
		"title": "",
	}
	return c.Render(http.StatusOK, "index.html", data)
}

func Products(c echo.Context) error {
	data := map[string]interface{}{
		"title": "",
	}
	return c.Render(http.StatusOK, "products.html", data)
}

func About(c echo.Context) error {
	data := map[string]interface{}{
		"title": "",
	}
	return c.Render(http.StatusOK, "about.html", data)
}
