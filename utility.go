package ioc

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func Unescaped(txt string) interface{} { return template.HTML(txt) }

func LittleCamelCase(txt string) string {
	return fmt.Sprintf("%s%s", strings.ToLower(string(txt[0])), txt[1:])
}

func UpperCamelCase(txt string) string {
	sb := new(strings.Builder)

	strs := strings.Split(txt, "_")

	for _, str := range strs {
		sb.WriteString(strings.ToUpper(string(str[0])))
		sb.WriteString(str[1:])
	}

	return sb.String()
}

func ReadFile(filename string) (string, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)

	defer f.Close()

	if err != nil {
		return "", err
	}

	if bytes, err := ioutil.ReadAll(f); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
