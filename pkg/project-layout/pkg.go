package projectlayout

import (
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
)

type Pkg struct{}

type PkgTemplate struct {
	Package string
}

func (f Pkg) generate(id string, app application.Application) error {
	return files.CreateFileFromTemplateWithData(id, files.Join("pkg", "public"), "hello.go",
		PkgTemplate{"public"})
}
