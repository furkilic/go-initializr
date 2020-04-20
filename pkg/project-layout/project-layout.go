package projectlayout

import (
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
	"log"
	"reflect"
)

var FOLDERS = []string{
	"api",
	"assets",
	"build",
	"cmd",
	"configs",
	"deployments",
	"docs",
	"examples",
	"githooks",
	"init",
	"internal",
	"pkg",
	"scripts",
	"test",
	"third_party",
	"tools",
	"vendor",
	"web",
	"website",
}
var EmptyFolders = []string{
	"api",
	"configs",
	"examples",
	"githooks",
	"init",
	"test",
	"third_party",
	"tools",
	"web",
	"website",
}

var NonEmptyFolders = []IFolder{
	Assets{},
	Build{},
	Cmd{},
	Deployments{},
	Docs{},
	Internal{},
	Pkg{},
	Scripts{},

	GoMod{},
	Gow{},
}

type IFolder interface {
	generate(id string, app application.Application) error
}

func GenerateProject(app application.Application) (string, error) {
	s, e := files.GenerateId()
	if e != nil {
		log.Fatal(e)
		return s, e
	}
	createAllEmptyFolders(s)
	createAllNotEmptyFolders(s, app)
	return s, nil
}

func createAllEmptyFolders(id string) error {
	for _, folder := range EmptyFolders {
		_, err := files.CreateFolderToKeep(id, folder)
		if err != nil {
			log.Fatalf("Error While Creating Folder \"%s\" : %s", folder, err)
			return err
		}
	}
	return nil
}

func createAllNotEmptyFolders(id string, app application.Application) error {
	for _, folder := range NonEmptyFolders {
		err := folder.generate(id, app)
		if err != nil {
			log.Fatalf("Error While CreateAllNotEmptyFolders \"%s\" : %s", reflect.TypeOf(folder).Name(), err)
			return err
		}
	}
	return nil
}
