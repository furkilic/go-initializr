package projectlayout

import (
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
)

type Cmd struct{}

type CmdTemplate struct {
	Name, FolderName string
}

func (f Cmd) generate(id string, app application.Application) error {
	cmdFiles := []string{"main.go", "main_test.go"}
	for _, cmdFile := range cmdFiles {
		err := files.CreateFileFromTemplateWithData(id, files.Join("cmd", app.GetAppName()), cmdFile,
			CmdTemplate{app.GetName(), app.GetAppName()})
		if err != nil {
			return err
		}
	}
	return nil
}
