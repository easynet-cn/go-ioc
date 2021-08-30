package ioc

type ContainerGenerator interface {
	Generate(convert func(StructContext) interface{}, argsBuild func(map[string][]string) string) (interface{}, error)
}
