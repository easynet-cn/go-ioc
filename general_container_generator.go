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

func NewGeneralContainerGenerator(containerConfig ContainerConfig) ContainerGenerator {
	return &GeneralContainerGenerator{containerConfig: containerConfig}
}

func (s *GeneralContainerGenerator) Generate(convert func(StructContext) interface{}, argsBuild func(args map[string][]string) string) (interface{}, error) {
	structInfoes, err1 := s.load(s.containerConfig.InputDirectory)

	if err1 != nil {
		return nil, err1
	}

	outputFilename := path.Join(s.containerConfig.OutputDirectory, s.containerConfig.OutputFilename)

	if err := s.remove(outputFilename); err != nil {
		return nil, err
	}

	f, err2 := os.OpenFile(outputFilename, os.O_CREATE|os.O_WRONLY, 0666)

	defer f.Close()

	if err2 != nil {
		return nil, err2
	}

	_, templateFilename := path.Split(s.containerConfig.TemplateFile)

	funcMap := template.FuncMap{"unescaped": Unescaped, "buildArgs": argsBuild}
	tpl, err3 := template.New(templateFilename).Funcs(funcMap).ParseFiles(s.containerConfig.TemplateFile)

	if err3 != nil {
		return nil, err3
	}

	if err := tpl.Execute(f, convert(StructContext{StructInfoes: structInfoes})); err != nil {
		return nil, err
	}

	return structInfoes, nil
}

func (s *GeneralContainerGenerator) load(inputDirectory string) ([]StructInfo, error) {
	if files, err := os.ReadDir(s.containerConfig.InputDirectory); err != nil {
		return nil, err
	} else {
		structInfoes := make([]StructInfo, 0)

		for _, file := range files {
			filename := file.Name()

			if s.isGeneralFile(filename) {
				structName := UpperCamelCase(filename[:len(filename)-3])

				txt, err2 := ReadFile(path.Join(s.containerConfig.InputDirectory, filename))

				if err2 != nil {
					return nil, err2
				}

				requestMappingComment := s.containerConfig.RequestMappingComment

				if requestMappingComment == "" {
					requestMappingComment = "@RequestMapping"
				}

				structParser := NewStructParser(structName, requestMappingComment)

				structInfo := structParser.Parse(txt).(StructInfo)

				structInfoes = append(structInfoes, structInfo)
			}
		}

		return structInfoes, nil
	}
}

func (s *GeneralContainerGenerator) isGeneralFile(filename string) bool {
	return !strings.HasSuffix(filename, "_test.go") && !strings.HasSuffix(filename, "_container.go")
}

func (s *GeneralContainerGenerator) remove(outputFilename string) error {
	if f, err := os.Stat(outputFilename); err != nil && !os.IsNotExist(err) {
		return err

	} else if f != nil {
		if err := os.Remove(outputFilename); err != nil {
			return err
		}
	}

	return nil
}
