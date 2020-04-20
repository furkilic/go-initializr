package projectlayout

import (
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
)

type GoMod struct{}

type GoModTemplate struct {
	Name string
}

func (f GoMod) generate(id string, app application.Application) error {
	return files.CreateFileFromTemplateWithData(id, "", "go.mod",
		GoModTemplate{app.GetAppName()})
}
