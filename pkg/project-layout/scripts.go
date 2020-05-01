package projectlayout

import (
	"github.com/furkilic/go-initializr/pkg/application"
	"github.com/furkilic/go-initializr/pkg/files"
)

type Scripts struct{}

type ScriptsTemplate struct {
	Name string
}

var templatedScripts = []string{
	"build.sh",
	"build.cmd",
	"docker-build.sh",
	"docker-build.cmd",
}
var copyScripts = []string{
	"common.sh",
	"test.sh",
	"bench.sh",
	"analysis.sh",
	"docker-compose-up.sh",
	"fmt.sh",
	"test.cmd",
	"bench.cmd",
	"analysis.cmd",
	"docker-compose-up.cmd",
	"fmt.cmd",
}

func (f Scripts) generate(id string, app application.Application) error {
	err := files.CopyFileFromExisting(id, "", "Makefile")
	if err != nil {
		return err
	}
	for _, script := range templatedScripts {
		err := files.CreateFileFromTemplateWithData(id, "scripts", script,
			ScriptsTemplate{app.GetAppName()})
		if err != nil {
			return err
		}
	}
	for _, script := range copyScripts {
		err := files.CopyFileFromExisting(id, "scripts", script)
		if err != nil {
			return err
		}
	}
	return nil
}
