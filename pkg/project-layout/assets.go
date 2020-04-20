package projectlayout

import (
	"go-initializr/pkg/application"
	"go-initializr/pkg/files"
)

type Assets struct{}

func (f Assets) generate(id string, app application.Application) error {
	return files.CopyFileFromExisting(id, "assets", "icon.png")
}
