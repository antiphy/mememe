package controllers

import (
	"net/http"

	"github.com/antiphy/mememe/dal/dbactions"
	"github.com/antiphy/mememe/dal/models"
	"github.com/labstack/echo/v4"
)

var cache models.Cache

func Message(c echo.Context) error {
	data := newBaseData()
	return c.Render(http.StatusOK, "message.html", data)
}

func newBaseData() map[string]interface{} {
	data := map[string]interface{}{
		"app_name": cache.GET("app_name").SettingValue,
		"app_desc": cache.GET("app_desc").SettingValue,
		"email":    cache.GET("email").SettingValue,
		"phone":    cache.GET("phone").SettingValue,
		"address":  cache.GET("address").SettingValue,
		"github":   cache.GET("github"),
		"twitter":  cache.GET("twitter"),
	}
	return data
}

func InitSettingCache() {
	cache = models.NewCache()
	settings, err := dbactions.QuerySettings()
	if err != nil {
		panic(err)
	}
	for i := range settings {
		cache.SET(settings[i].SettingKey, &settings[i])
	}
}
