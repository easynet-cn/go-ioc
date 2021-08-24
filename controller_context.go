package ioc

type ControllerContext struct {
	Controllers []Controller
}

func (m *ControllerContext) Properties() map[string][]string {
	typeMap := make(map[string][]string)
	tempTypeMap := make(map[string]map[string]string)

	for _, controller := range m.Controllers {
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
