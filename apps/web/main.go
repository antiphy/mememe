package main

import (
	"fmt"
	"os"
	"time"

	"github.com/antiphy/mememe/apps/web/controllers"

	"github.com/antiphy/mememe/apps/web/routes"
	"github.com/antiphy/mememe/dal/consts"
	"github.com/antiphy/mememe/dal/dbactions"
	"github.com/antiphy/mememe/dal/models"
	"github.com/antiphy/mememe/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "init_db":
			err := dbactions.CreateTableSetting()
			if err != nil {
				fmt.Println("create table setting err: " + err.Error())
				return
			}
			err = dbactions.CreateTableAccount()
			if err != nil {
				fmt.Println("create table account err: " + err.Error())
				return
			}
			err = dbactions.CreateTableBlogArticle()
			if err != nil {
				fmt.Println("create table blog article err: " + err.Error())
				return
			}
			fmt.Println("success!")
			return
		case "create_user":
			if len(args) != 4 {
				fmt.Println("username and password required")
				return
			}
			ts := time.Now().Unix()
			account := models.Account{Name: args[2], Password: utils.MD5(args[3]), CreateTS: ts, UpdateTS: ts}
			err := dbactions.CreateAccount(&account)
			if err != nil {
				fmt.Println("create account err: " + err.Error())
				return
			}
			fmt.Println("success!")
		default:
			fmt.Printf("unknown command: %s\n", args[1])
		}
		return
	}
	controllers.InitSettingCache()
	httpEngine := echo.New()
	httpEngine.Pre(middleware.HTTPSRedirect())
	go func() {
		httpEngine.HideBanner = true
		httpEngine.Start(":80")
	}()
	httpsEngine := routes.NewRouter()
	httpsEngine.Logger.Fatal(httpsEngine.StartTLS(":443", consts.GetCRTFilePath(), consts.GetKEYFilePath()))
}
