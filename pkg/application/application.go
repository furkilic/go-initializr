package application

import (
	"github.com/stoewer/go-strcase"
	"strings"
)

type Application struct {
	Repo  string
	Owner string
	Name  string
}

func (an Application) GetName() string {
	return an.Name
}
func (an Application) GetAppName() string {
	return strcase.KebabCase(an.Name)
}
func (an Application) GetPkgName() string {
	return strings.ToLower(strcase.LowerCamelCase(an.Name))
}
