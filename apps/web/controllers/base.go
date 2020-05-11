package controllers

import (
	"github.com/antiphy/mememe/dal/dbactions"
	"github.com/antiphy/mememe/dal/models"
)

var cache models.Cache

func newBaseData() map[string]interface{} {
	data := map[string]interface{}{
		"app_name": cache.GET("app_name").SettingValue,
		"app_desc": cache.GET("app_desc").SettingValue,
		"email":    cache.GET("email").SettingValue,
		"phone":    cache.GET("phone").SettingValue,
		"address":  cache.GET("address").SettingValue,
	}
	return data
}

func init() {
	cache = models.NewCache()
	settings, err := dbactions.QuerySettings()
	if err != nil {
		panic(err)
	}
	for i := range settings {
		cache.SET(settings[i].SettingKey, &settings[i])
	}
}
