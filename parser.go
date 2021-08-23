package ioc

type Parser interface {
	Parse(txt string) interface{}
}
