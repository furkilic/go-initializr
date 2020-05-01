package configs

import (
	"github.com/furkilic/go-boot-config/pkg/go-boot-config"
	"github.com/furkilic/go-initializr/pkg/application"
	"log"
)

func Init() application.Application {
	gobootconfig.Load()
	name, err := gobootconfig.GetString("app.name")
	if err != nil {
		log.Fatalf("Error while loading App Name: %v", err)
	}
	owner, err := gobootconfig.GetString("app.owner")
	if err != nil {
		log.Fatalf("Error while loading App Owner: %v", err)
	}
	repo := gobootconfig.GetStringWithDefault("app.repo", "github.com")
	app := application.Application{
		Name:  name,
		Owner: owner,
		Repo:  repo,
	}
	return app
}
