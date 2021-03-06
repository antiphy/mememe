package controllers

import (
	"net/http"

	"github.com/antiphy/mememe/dal/dbactions"
	"github.com/antiphy/mememe/dal/models"
	"github.com/labstack/echo/v4"
)

var (
	cache   models.Cache
	blocker models.Blocker
)

func newBaseData() map[string]interface{} {
	data := map[string]interface{}{
		"app_name":   cache.GET("app_name").SettingValue,
		"app_domain": cache.GET("app_domain").SettingType,
		"app_desc":   cache.GET("app_desc").SettingValue,
		"email":      cache.GET("email").SettingValue,
		"phone":      cache.GET("phone").SettingValue,
		"address":    cache.GET("address").SettingValue,
		"github":     cache.GET("github").SettingValue,
		"twitter":    cache.GET("twitter").SettingValue,
		"admin":      cache.GET("admin").SettingValue,
	}
	return data
}

func getAccount(c echo.Context) *models.Account {
	a, ok := c.Get("account").(*models.Account)
	if !ok {
		return nil
	}
	return a
}

func Message(c echo.Context) error {
	data := newBaseData()
	data["title"] = "hint"
	return c.Render(http.StatusOK, "message.html", data)
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
	blocker = models.NewBlocker()
}
