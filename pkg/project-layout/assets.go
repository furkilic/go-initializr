package projectlayout

import (
	"github.com/furkilic/go-initializr/pkg/application"
	"github.com/furkilic/go-initializr/pkg/files"
)

type Assets struct{}

func (f Assets) generate(id string, app application.Application) error {
	return files.CopyFileFromExisting(id, "assets", "icon.png")
}
