package routes

import (
	"github.com/antiphy/mememe/apps/web/controllers"
	"github.com/antiphy/mememe/dal/consts"
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
