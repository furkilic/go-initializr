package projectlayout

import (
	"github.com/furkilic/go-initializr/pkg/application"
	"github.com/furkilic/go-initializr/pkg/files"
)

type Build struct{}

type BuildTemplate struct {
	Name string
}

func (f Build) generate(id string, app application.Application) error {
	return files.CreateFileFromTemplateWithData(id, files.Join("build", "package"), "Dockerfile",
		BuildTemplate{app.GetAppName()})

}
