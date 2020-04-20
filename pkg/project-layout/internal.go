package projectlayout

import (
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
)

type Internal struct{}

type InternalTemplate struct {
	Package string
}

func (f Internal) generate(id string, app application.Application) error {
	return files.CreateFileFromTemplateWithData(id, files.Join("internal", "private"), "hello.go",
		InternalTemplate{"private"})
}
