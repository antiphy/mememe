package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	data := newBaseData()
	data["title"] = cache.GET("app_name").SettingValue
	return c.Render(http.StatusOK, "index.html", data)
}

func Products(c echo.Context) error {
	data := newBaseData()
	data["title"] = "Products"
	return c.Render(http.StatusOK, "products.html", data)
}

func About(c echo.Context) error {
	data := newBaseData()
	data["title"] = "About"
	return c.Render(http.StatusOK, "about.html", data)
}

func Contact(c echo.Context) error {
	data := newBaseData()
	data["title"] = "Contact"
	return c.Render(http.StatusOK, "contact.html", data)
}
