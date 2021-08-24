package ioc

type ArgsBuilder interface {
	Build(args map[string][]string) string
}
