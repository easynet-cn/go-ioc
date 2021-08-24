package ioc

type ContainerGenerator interface {
	Generate(context StructContext) error
}
