package projectlayout

import (
	"fmt"
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
)

type Docs struct{}

type DocsTemplate struct {
	Title, Description string
}

func (f Docs) generate(id string, app application.Application) error {
	title := fmt.Sprintf("%s Documentation", app.GetName())
	description := fmt.Sprintf("Here comes the documentation for \"%s\"", app.GetName())
	err := files.CreateFileFromTemplateWithData(id, "docs", "README.md",
		DocsTemplate{title, description})

	if err != nil {
		return err
	}
	return files.CreateFileFromTemplateWithData(id, "", "README.md",
		DocsTemplate{title, description})

}
