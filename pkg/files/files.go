package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

var TemplatesFolder = Join("assets", "templates")

func Join(names ...string) string {
	return filepath.Join(names...)
}

func GenerateId() (string, error) {
	return ioutil.TempDir("", "go-initializr-output-")
}

func CreateFolderToKeep(id string, folder string) (*os.File, error) {
	return CreateFolderWithFile(id, folder, ".keep")
}

func CreateFolderWithFile(id string, folder, fileName string) (*os.File, error) {
	err := CreateFolder(id, folder)
	if err != nil {
		return nil, err
	}
	return CreateFileWithId(id, folder, fileName)
}

func CreateFolder(id, folder string) error {
	return os.MkdirAll(Join(id, folder), os.ModePerm)
}

func CreateFileWithId(id string, folder string, fileName string) (*os.File, error) {
	return os.Create(Join(id, folder, fileName))
}
func CreateFile(folder string, fileName string) (*os.File, error) {
	return os.Create(Join(folder, fileName))
}

func CreateFileFromTemplateWithData(id, folder string, fileName string, data interface{}) error {
	return CreateFileFromTemplateWithDataAndTarget(id, folder, fileName, fileName, data)
}
func CreateFileFromTemplateWithDataAndTarget(id, folder string, sourceFileName string, targetFileName string, data interface{}) error {
	err := CreateFolder(id, folder)
	if err != nil {
		return err
	}
	file, errC := CreateFileWithId(id, folder, targetFileName)
	if errC != nil {
		return errC
	}
	templateName := fmt.Sprintf("%s.tmpl", sourceFileName)
	tmpl := template.Must(template.ParseGlob(Join(TemplatesFolder, templateName)))
	return tmpl.Execute(file, data)
}

func CopyFileFromExisting(id, folder string, fileName string) error {
	err := CreateFolder(id, folder)
	if err != nil {
		return err
	}
	input, err := ioutil.ReadFile(Join(TemplatesFolder, fileName))
	if err != nil {
		return err
	}
	return ioutil.WriteFile(Join(id, folder, fileName), input, 0644)
}
