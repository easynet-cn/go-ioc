package ioc

type ContainerGenerator interface {
	Generate(argsBuilder func(args map[string][]string) string) ([]interface{}, error)
}
