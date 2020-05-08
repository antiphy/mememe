package controllers

import (
	"net/http"
	"strconv"

	"github.com/antiphy/mememe/dal/consts"
	"github.com/antiphy/mememe/dal/dbactions"
	"github.com/antiphy/mememe/dal/models"
	"github.com/antiphy/mememe/utils"
	"github.com/labstack/echo/v4"
)

func BlogIndex(c echo.Context) error {
	var (
		page int
		err  error
	)
	page, err = strconv.Atoi(c.Param(":page"))
	if err != nil {
		page = 1
	}
	params := models.QueryParams{Page: page}
	articles, err := dbactions.QueryBlogArticles(&params)
	if err != nil {
		// TODO:
	}
	data := map[string]interface{}{
		"title":    consts.AppName,
		"articles": articles,
		"page":     params.Page,
		"total":    params.Total,
	}
	return c.Render(http.StatusOK, "blog/index.html", data)
}

func BlogDetail(c echo.Context) error {
	var (
		id  int
		err error
	)
	id, err = strconv.Atoi(c.Param(":id"))
	if err != nil {
		// TODO:
	}
	params := models.QueryParams{ID: id}
	article, err := dbactions.QueryBlogArticle(&params)
	if err != nil {
		// TODO:
	}
	data := map[string]interface{}{
		"title":   article.Title + consts.AppName,
		"article": article,
	}
	return c.Render(http.StatusOK, "blog/article.html", data)
}

func BlogCreateArticle(c echo.Context) error {
	return c.Render(http.StatusOK, "blog/create_blog.html", nil)
}

func BlogCreateArticlePOST(c echo.Context) error {
	res := make(map[string]interface{})
	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		res["code"] = 1
		res["msg"] = "data bind err:" + err.Error()
		return c.JSON(http.StatusOK, res)
	}
	err = dbactions.CreateBlogArticle(&article)
	if err != nil {
		res["code"] = 1
		res["msg"] = "create article failed:" + err.Error()
		return c.JSON(http.StatusOK, res)
	}
	res["code"] = 0
	return c.JSON(http.StatusOK, res)
}

func BlogLoginGET(c echo.Context) error {
	return c.Render(http.StatusOK, "blog/login.html", nil)
}

func BlogLoginPOST(c echo.Context) error {
	res := make(map[string]interface{})
	var account models.Account
	err := c.Bind(&account)
	if err != nil {
		res["code"] = 1
		res["msg"] = "data bind err:" + err.Error()
		return c.JSON(http.StatusOK, res)
	}
	password := account.Password
	err = dbactions.QueryAccount(&account)
	if err != nil {
		res["code"] = 1
		res["msg"] = "query account err:" + err.Error()
		return c.JSON(http.StatusOK, res)
	}
	if utils.MD5(password) != account.Password {
		res["code"] = 1
		res["msg"] = "incorrect password"
		// TODO:
		return c.JSON(http.StatusOK, res)
	}
	res["code"] = 0
	return c.JSON(http.StatusOK, res)
}
