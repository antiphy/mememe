package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/antiphy/mememe/dal/consts"
	"github.com/antiphy/mememe/dal/dbactions"
	"github.com/antiphy/mememe/dal/models"
	"github.com/antiphy/mememe/utils"
	"github.com/dgrijalva/jwt-go"
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
	params := models.QueryParams{Page: page, PageSize: 10}
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
	a := getAccount(c)
	if a == nil {
		data["message"] = "login required"
		return c.Render(http.StatusOK, "message.html", data)
	}
	data["title"] = "create blog article"
	return c.Render(http.StatusOK, "blog/create_blog.html", data)
}

func BlogCreateArticlePOST(c echo.Context) error {
	res := make(map[string]interface{})
	a := getAccount(c)
	if a == nil {
		res["code"] = 1
		res["msg"] = "login required"
		return c.JSON(http.StatusOK, res)
	}
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
	a := getAccount(c)
	if a != nil {
		return c.Redirect(http.StatusOK, "/")
	}
	data["title"] = "login"
	return c.Render(http.StatusOK, "blog/login.html", data)
}

func BlogLoginPOST(c echo.Context) error {
	ip := strings.Split(c.Request().RemoteAddr, ",")[0]
	res := make(map[string]interface{})
	if blocker.IsBlocked(ip) {
		res["code"] = 1
		res["msg"] = "blocked"
		return c.JSON(http.StatusOK, res)
	}
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
		blocker.Incr(ip)
		return c.JSON(http.StatusOK, res)
	}
	appName := cache.GET("app_name").SettingValue
	claims := models.JWTClaims{UID: account.ID, Name: account.Name, ExpireTS: time.Now().Unix() + 30*86400, Issuer: appName}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString([]byte(consts.WebSecretKey))
	c.SetCookie(&http.Cookie{Name: appName, Value: ts, Domain: cache.GET("app_domain").SettingValue, Path: "/"})
	res["code"] = 0
	return c.JSON(http.StatusOK, res)
}
