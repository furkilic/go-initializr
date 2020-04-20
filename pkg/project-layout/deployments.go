package projectlayout

import (
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
)

type Deployments struct{}

type DeploymentsTemplate struct {
	Name string
}

func (f Deployments) generate(id string, app application.Application) error {
	return files.CreateFileFromTemplateWithData(id, "deployments", "docker-compose.yml",
		DeploymentsTemplate{app.GetAppName()})
}
