package ioc

import (
	"fmt"
	"regexp"
	"strings"
)

type StructParser struct {
	structName            string
	requestMappingComment string
}

func NewStructParser(structName string, requestMappingComment string) Parser {
	return &StructParser{structName: structName, requestMappingComment: requestMappingComment}
}

func (s *StructParser) Parse(txt string) interface{} {
	return StructInfo{
		Name:           s.structName,
		Properties:     s.parseProperties(txt),
		Args:           s.parseArgs(txt),
		RequestMapping: s.parseRequestMapping(txt),
	}
}

func (s *StructParser) parseProperties(txt string) []Arg {
	props := make([]Arg, 0)
	startTxt := fmt.Sprintf("type %s struct {", s.structName)

	if firstIndex := strings.Index(txt, startTxt) + len(startTxt); firstIndex > len(startTxt)-1 {
		lastIndex := firstIndex + strings.Index(txt[firstIndex:], "}")

		strs := strings.Split(txt[firstIndex:lastIndex], "\n")

		for _, str := range strs {
			argTxt := strings.TrimSpace(str)

			if len(argTxt) > 0 {
				props = append(props, s.parseArg(argTxt))
			}
		}
	}

	return props
}

func (s *StructParser) parseArgs(txt string) []Arg {
	args := make([]Arg, 0)
	startTxt := fmt.Sprintf("New%s(", s.structName)

	if firstIndex := strings.Index(txt, startTxt) + len(startTxt); firstIndex > len(startTxt)-1 {
		lastIndex := firstIndex + strings.Index(txt[firstIndex:], ")")

		reg := regexp.MustCompile(`\s+`)

		strs := strings.Split(reg.ReplaceAllString(txt[firstIndex:lastIndex], " "), ",")

		for _, str := range strs {
			argTxt := strings.TrimSpace(str)

			if len(argTxt) > 0 {
				args = append(args, s.parseArg(argTxt))
			}
		}
	}

	return args
}

func (s *StructParser) parseRequestMapping(txt string) string {
	sb := new(strings.Builder)

	sb.WriteString("// ")

	if strings.HasPrefix(s.requestMappingComment, "//") {
		sb.WriteString(strings.TrimSpace(s.requestMappingComment[2:]))
	} else {
		sb.WriteString(s.requestMappingComment)
	}

	sb.WriteString("(")

	startTxt := sb.String()

	if firstIndex := strings.Index(txt, startTxt) + len(startTxt); firstIndex > len(startTxt)-1 {
		lastIndex := firstIndex + strings.Index(txt[firstIndex:], ")")

		return strings.TrimSpace(txt[firstIndex:lastIndex])
	}

	return ""
}

func (s *StructParser) parseArg(txt string) Arg {
	strs := strings.Split(txt, " ")

	return Arg{Name: strings.TrimSpace(strs[0]), TypeName: strings.TrimSpace(strs[1])}
}
