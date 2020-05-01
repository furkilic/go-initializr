package main

import (
	"fmt"
	"github.com/furkilic/go-initializr/internal/configs"
	"github.com/furkilic/go-initializr/pkg/files"
	"github.com/furkilic/go-initializr/pkg/project-layout"
	"log"
)

func main() {
	app := configs.Init()
	path, err := projectlayout.GenerateProject(app)
	if err != nil {
		log.Fatal(err)
	}
	file, err := files.CreateFile("", fmt.Sprintf("%s.zip", app.GetAppName()))
	if err != nil {
		log.Fatal(err)
	}
	err = files.Zip(path, file)
	if err != nil {
		log.Fatal(err)
	}
}
