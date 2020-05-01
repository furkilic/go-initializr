package projectlayout

import (
	"github.com/furkilic/go-initializr/pkg/application"
	"github.com/furkilic/go-initializr/pkg/files"
)

type GoMod struct{}

type GoModTemplate struct {
	Name  string
	Owner string
	Repo  string
}

func (f GoMod) generate(id string, app application.Application) error {
	return files.CreateFileFromTemplateWithData(id, "", "go.mod",
		GoModTemplate{app.GetAppName(), app.Owner, app.Repo})
}
