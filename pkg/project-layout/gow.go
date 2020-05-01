package projectlayout

import (
	"github.com/furkilic/go-initializr/pkg/application"
	"github.com/furkilic/go-initializr/pkg/files"
)

type Gow struct{}

type GowTemplate struct {
	Name string
}

func (f Gow) generate(id string, app application.Application) error {
	_, err := files.CreateFolderWithFile(id, files.Join(".go", "wrapper"), "go-wrapper.properties")
	if err != nil {
		return err
	}
	err = files.CopyFileFromExisting(id, "", ".gitignore")
	if err != nil {
		return err
	}
	err = files.CopyFileFromExisting(id, "", "gow")
	if err != nil {
		return err
	}
	return files.CopyFileFromExisting(id, "", "gow.cmd")
}
