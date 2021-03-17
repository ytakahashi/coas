package internal

import "github.com/ytakahashi/coas/api"

func createSelectItems(parameter api.Parameter) []string {
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
