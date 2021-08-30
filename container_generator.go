package ioc

type ContainerGenerator interface {
	Generate(func(StructContext) interface{}, map[string]interface{}) (interface{}, error)
}
