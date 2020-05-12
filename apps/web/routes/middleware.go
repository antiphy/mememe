package routes

import (
	"errors"
	"strings"

	"github.com/antiphy/mememe/dal/consts"
	"github.com/antiphy/mememe/dal/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func middleware() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Path(), "/public") {
				return h(c)
			}
			if cookie, err := c.Cookie("mememe"); err == nil {
				token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, errors.New("invalid cookie")
					}
					return []byte(consts.WebSecretKey), nil
				})
				if err == nil {
					if claims := token.Claims.(jwt.MapClaims); token.Valid {
						account := models.Account{
							ID:   int(claims["UID"].(float64)),
							Name: claims["Name"].(string),
						}
						c.Set("account", account)
					}
				}
			}
			return h(c)
		}
	}
}
