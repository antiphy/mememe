package routes

import (
	"github.com/antiphy/mememe/apps/web/controllers"
	"github.com/antiphy/mememe/dal/consts"
	"github.com/antiphy/mememe/dal/models"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.HideBanner = true
	templates := preCompile(consts.GetViewsDirPath())
	e.Renderer = templates
	e.Static("/public", consts.GetStaticDirPath())

	// static pages
	e.GET("", controllers.Index)
	e.GET("/products", controllers.Products)
	e.GET("/about", controllers.About)
	e.GET("/setting", controllers.Setting)
	e.POST("/setting", controllers.CreateOrUpdateSettings)

	// blog
	blog := e.Group("/blog")
	blog.GET("", controllers.BlogIndex)
	blog.GET("/p/:page", controllers.BlogIndex)
	blog.GET("/a/:id/:title", controllers.BlogDetail)
	blog.GET("/login", controllers.BlogLoginGET)
	blog.POST("/login", controllers.BlogLoginPOST)
	blog.GET("/create_article", controllers.BlogCreateArticle)
	blog.POST("/create_article", controllers.BlogCreateArticlePOST)

	return e
}

func adminAccess(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		account, ok := c.Get("account").(*models.Account)
		if ok && account.IsAdmin() {
			return h(c)
		}
		c.Set("msg", "access denied")
		return controllers.Message(c)
	}
}
