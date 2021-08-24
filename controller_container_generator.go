package ioc

import (
	"html/template"
	"os"
	"path"
	"strings"
)

type ControllerContainerGenerator struct {
	containerConfig ContainerConfig
}

func NewControllerContainerGenerator(containerConfig ContainerConfig) *ControllerContainerGenerator {
	return &ControllerContainerGenerator{containerConfig: containerConfig}
}

func (s *ControllerContainerGenerator) Generate(argsBuilder func(args map[string][]string) string) ([]Controller, error) {
	files, err1 := os.ReadDir(s.containerConfig.InputDirectory)

	if err1 != nil {
		return nil, err1
	}

	controllers := make([]Controller, 0)

	for _, file := range files {
		filename := file.Name()

		if s.isControllerFile(filename) {
			controllerName := UpperCamelCase(filename[:len(filename)-3])

			txt, err2 := ReadFile(path.Join(s.containerConfig.InputDirectory, filename))

			if err2 != nil {
				return nil, err2
			}

			controllerParser := NewControllerParser(controllerName, "@RequestMapping")

			controller := controllerParser.Parse(txt).(Controller)

			controllers = append(controllers, controller)
		}
	}

	outputFilename := path.Join(s.containerConfig.OutputDirectory, s.containerConfig.OutputFilename)

	if f, err := os.Stat(outputFilename); err != nil {
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

	if err := tpl.Execute(f, &ControllerContext{Controllers: controllers}); err != nil {
		return nil, err
	}

	return controllers, nil
}

func (s *ControllerContainerGenerator) isControllerFile(filename string) bool {
	return strings.HasSuffix(filename, "_controller.go") && !strings.HasSuffix(filename, "_test.go")
}
