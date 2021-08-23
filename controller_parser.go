package ioc

import "strings"

type ControllerParser struct {
	structParser          StructParser
	requestMappingComment string
}

func NewControllerParser(structName, requestMappingComment string) Parser {
	return &ControllerParser{structParser: StructParser{structName: structName}, requestMappingComment: requestMappingComment}
}

func (s *ControllerParser) Parse(txt string) interface{} {
	return Controller{
		StructInfo:     s.structParser.Parse(txt).(StructInfo),
		RequestMapping: s.parseRequestMapping(txt),
	}
}

func (s *ControllerParser) parseRequestMapping(txt string) string {
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
