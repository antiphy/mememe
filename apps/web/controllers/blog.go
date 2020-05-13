package controllers

import (
	"net/http"
	"strconv"

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
	data := newBaseData()
	page, err = strconv.Atoi(c.Param(":page"))
	if err != nil {
		page = 1
	}
	params := models.QueryParams{Page: page}
	articles, err := dbactions.QueryBlogArticles(&params)
	if err != nil {
		data["message"] = "server error:" + err.Error()
		return c.Render(http.StatusOK, "message.html", data)
	}
	data["title"] = "blog index"
	data["articles"] = articles
	data["page"] = params.Page
	data["total"] = params.Total
	return c.Render(http.StatusOK, "blog/index.html", data)
}

func BlogDetail(c echo.Context) error {
	var (
		id  int
		err error
	)
	data := newBaseData()
	id, err = strconv.Atoi(c.Param(":id"))
	if err != nil {
		data["message"] = "invalid article id"
		return c.Render(http.StatusOK, "message.html", data)
	}
	params := models.QueryParams{ID: id}
	article, err := dbactions.QueryBlogArticle(&params)
	if err != nil {
		data["message"] = "server error:" + err.Error()
		return c.Render(http.StatusOK, "message.html", data)
	}
	data["title"] = article.Title
	data["article"] = article
	return c.Render(http.StatusOK, "blog/article.html", data)
}

func BlogCreateArticle(c echo.Context) error {
	data := newBaseData()
	data["title"] = "create blog article"
	return c.Render(http.StatusOK, "blog/create_blog.html", data)
}

func BlogCreateArticlePOST(c echo.Context) error {
	res := make(map[string]interface{})
	var article models.Article
	err := c.Bind(&article)
	if err != nil {
		res["code"] = 1
		res["msg"] = "invalid request data:" + err.Error()
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
	data := newBaseData()
	data["title"] = "login"
	return c.Render(http.StatusOK, "blog/login.html", data)
}

func BlogLoginPOST(c echo.Context) error {
	res := make(map[string]interface{})
	var account models.Account
	err := c.Bind(&account)
	if err != nil {
		res["code"] = 1
		res["msg"] = "invalid request data:" + err.Error()
		return c.JSON(http.StatusOK, res)
	}
	password := account.Password
	err = dbactions.QueryAccount(&account)
	if err != nil {
		res["code"] = 1
		res["msg"] = "server err:" + err.Error()
		return c.JSON(http.StatusOK, res)
	}
	if utils.MD5(password) != account.Password {
		res["code"] = 1
		res["msg"] = "incorrect password"
		// TODO: antispam
		return c.JSON(http.StatusOK, res)
	}
	res["code"] = 0
	return c.JSON(http.StatusOK, res)
}
