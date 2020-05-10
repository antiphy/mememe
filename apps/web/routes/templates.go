package routes

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	// register template filter
	_ "github.com/antiphy/mememe/apps/web/viewfuncs"
	"github.com/antiphy/mememe/dal/consts"

	"github.com/flosch/pongo2"
	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo/v4"
)

type template struct {
	tmplMap map[string]*pongo2.Template
}

func (t *template) Render(w io.Writer, templateName string, data interface{}, c echo.Context) error {
	templateName = consts.GetViewsDirPath() + templateName
	dataMap := data.(map[string]interface{})
	template, exist := t.tmplMap[templateName]
	if !exist {
		return errors.New("template " + templateName + " not found")
	}
	return template.ExecuteWriter(dataMap, w)
}

func preCompile(dir string) *template {
	tplMap := make(map[string]*pongo2.Template)
	tpl := template{tmplMap: tplMap}
	go fswatch(dir, &tpl)

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			innerFileInfos, _ := ioutil.ReadDir(path.Join(dir, fileInfo.Name()))
			for _, innerFileInfo := range innerFileInfos {
				pathkey := path.Join(dir, fileInfo.Name(), innerFileInfo.Name())
				t, err := pongo2.FromFile(pathkey)
				if err != nil {
					log.Fatalf("\"%s\": %v", pathkey, err)
				}
				tplMap[pathkey] = t
			}
		} else {
			pathkey := path.Join(dir, fileInfo.Name())
			t, err := pongo2.FromFile(pathkey)
			if err != nil {
				log.Fatalf("\"%s\": %v", pathkey, err)
			}
			tplMap[pathkey] = t
		}

	}

	return &tpl
}

func fswatch(dir string, tpl *template) {
	watcher, _ := fsnotify.NewWatcher()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			path, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			err = watcher.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	go func() {
		for {
			select {
			case <-watcher.Events:
				{
					//reload tpl
					tplMap := make(map[string]*pongo2.Template)
					fileInfos, _ := ioutil.ReadDir(dir)
					for _, fileInfo := range fileInfos {
						if fileInfo.IsDir() {
							innerFileInfos, _ := ioutil.ReadDir(path.Join(dir, fileInfo.Name()))
							for _, innerFileInfo := range innerFileInfos {
								pathkey := path.Join(dir, fileInfo.Name(), innerFileInfo.Name())
								t, err := pongo2.FromFile(pathkey)
								if err != nil {
									log.Fatalf("\"%s\": %v", pathkey, err)
								}
								tplMap[pathkey] = t
							}
						} else {
							pathkey := path.Join(dir, fileInfo.Name())
							t, err := pongo2.FromFile(pathkey)
							if err != nil {
								log.Fatalf("\"%s\": %v", pathkey, err)
							}
							tplMap[pathkey] = t
						}

					}
					tpl.tmplMap = tplMap
					fmt.Println("templates reloaded")
				}
			case err := <-watcher.Errors:
				{
					fmt.Println("error : ", err)
					return
				}
			}
		}
	}()
	select {}
}
