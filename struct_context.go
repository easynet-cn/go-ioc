package ioc

type StructContext struct {
	StructInfoes []StructInfo
}

func (m *StructContext) Properties() map[string][]string {
	typeMap := make(map[string][]string)
	tempTypeMap := make(map[string]map[string]string)

	for _, controller := range m.StructInfoes {
		for _, property := range controller.Properties {
			if _, ok := tempTypeMap[property.TypeName]; !ok {
				tempTypeMap[property.TypeName] = make(map[string]string)
				typeMap[property.TypeName] = make([]string, 0)
			}

			if _, ok := tempTypeMap[property.TypeName][property.Name]; !ok {
				tempTypeMap[property.TypeName][property.Name] = property.Name
				typeMap[property.TypeName] = append(typeMap[property.TypeName], property.Name)
			}
		}
	}

	return typeMap
}

func (m *StructContext) Args() map[string][]string {
	typeMap := make(map[string][]string)
	tempTypeMap := make(map[string]map[string]string)

	for _, structInfo := range m.StructInfoes {
		for _, arg := range structInfo.Args {
			if _, ok := tempTypeMap[arg.TypeName]; !ok {
				tempTypeMap[arg.TypeName] = make(map[string]string)
				typeMap[arg.TypeName] = make([]string, 0)
			}

			if _, ok := tempTypeMap[arg.TypeName][arg.Name]; !ok {
				tempTypeMap[arg.TypeName][arg.Name] = arg.Name
				typeMap[arg.TypeName] = append(typeMap[arg.TypeName], arg.Name)
			}
		}
	}

	return typeMap
}
