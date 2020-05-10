package consts

import (
	"os"
)

const (
	AppName = "antiphy的博客"
)

func GetCRTFilePath() string {
	info, err := os.Stat("website.crt")
	if err == nil && !info.IsDir() {
		return "website.crt"
	}
	return "/server/mememe/website.crt"
}

func GetKEYFilePath() string {
	info, err := os.Stat("website.key")
	if err == nil && !info.IsDir() {
		return "website.key"
	}
	return "/server/mememe/website.key"
}

func GetStaticDirPath() string {
	info, err := os.Stat("public")
	if err == nil && info.IsDir() {
		return "public"
	}
	return "/server/mememe/public"
}

func GetViewsDirPath() string {
	info, err := os.Stat("views/")
	if err == nil && info.IsDir() {
		return "views/"
	}
	return "/server/mememe/views/"
}
