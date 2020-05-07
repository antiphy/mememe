package main

import (
	"github.com/antiphy/mememe/apps/web/routes"
)

func main() {
	httpsEngine := routes.NewRouter()
	httpsEngine.Logger.Fatal(httpsEngine.StartTLS(":443", "crt.filepath", "key.filepath"))
}
