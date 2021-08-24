package ioc

import (
	"html/template"
	"os"
	"path"
	"strings"
)

type GeneralContainerGenerator struct {
	containerConfig ContainerConfig
}

func NewGeneralContainerGenerator(containerConfig ContainerConfig) *GeneralContainerGenerator {
	return &GeneralContainerGenerator{containerConfig: containerConfig}
}

func (s *GeneralContainerGenerator) Generate(argsBuilder func(args map[string][]string) string) ([]interface{}, error) {
	files, err1 := os.ReadDir(s.containerConfig.InputDirectory)

	if err1 != nil {
		return nil, err1
	}

	structInfoes := make([]StructInfo, 0)

	for _, file := range files {
		filename := file.Name()

		if s.isGeneralFile(filename) {
			structName := UpperCamelCase(filename[:len(filename)-3])

			txt, err2 := ReadFile(path.Join(s.containerConfig.InputDirectory, filename))

			if err2 != nil {
				return nil, err2
			}

			structParser := NewStructParser(structName)

			structInfo := structParser.Parse(txt).(StructInfo)

			structInfoes = append(structInfoes, structInfo)
		}
	}

	outputFilename := path.Join(s.containerConfig.OutputDirectory, s.containerConfig.OutputFilename)

	if f, err := os.Stat(outputFilename); err != nil && !os.IsNotExist(err) {
		return nil, err

	} else if f != nil {
		if err := os.Remove(outputFilename); err != nil {
			return nil, err
		}
	}

	f, err2 := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0666)

	defer f.Close()

	if err2 != nil {
		return nil, err2
	}

	_, templateFilename := path.Split(s.containerConfig.TemplateFile)

	funcMap := template.FuncMap{"unescaped": Unescaped, "buildArgs": argsBuilder}
	tpl, err3 := template.New(templateFilename).Funcs(funcMap).ParseFiles(s.containerConfig.TemplateFile)

	if err3 != nil {
		return nil, err3
	}

	if err := tpl.Execute(f, &StructContext{StructInfoes: structInfoes}); err != nil {
		return nil, err
	}

	result := make([]interface{}, len(structInfoes))

	return result, nil
}

func (s *GeneralContainerGenerator) isGeneralFile(filename string) bool {
	return !strings.HasSuffix(filename, "_test.go") && !strings.HasSuffix(filename, "_container.go")
}
