package application

import (
	"github.com/stoewer/go-strcase"
	"strings"
)

type Application struct {
	name string
}

func NewApplication(name string) Application {
	return Application{name}
}

func (an Application) GetName() string {
	return an.name
}
func (an Application) GetAppName() string {
	return strcase.KebabCase(an.name)
}
func (an Application) GetPkgName() string {
	return strings.ToLower(strcase.LowerCamelCase(an.name))
}
