package controllers

import (
	"net/http"

	"github.com/antiphy/mememe/dal/dbactions"

	"github.com/antiphy/mememe/dal/models"
	"github.com/labstack/echo/v4"
)

func Setting(c echo.Context) error {
	data := newBaseData()
	return c.Render(http.StatusOK, "setting.html", data)
}

func CreateOrUpdateSettings(c echo.Context) error {
	res := make(map[string]interface{})
	var settings []models.Setting
	err := c.Bind(&settings)
	if err != nil {
		res["code"] = 1
		res["msg"] = "invalid request params:" + err.Error()
		return c.JSON(http.StatusOK, res)
	}
	err = dbactions.InsertOrUpdateSettings(settings)
	if err != nil {
		res["code"] = 1
		res["msg"] = err.Error()
		return c.JSON(http.StatusOK, res)
	}

	res["code"] = 0
	return c.JSON(http.StatusOK, res)
}
