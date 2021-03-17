package internal

func createSelectItems(parameter Parameter) []string {
	items := []string{}
	if len(parameter.ParameterEnums) == 0 {
		return items
	}
	if !parameter.Required {
		items = append(items, "")
	}
	items = append(items, parameter.ParameterEnums...)
	return items
}
