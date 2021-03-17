package internal

import "github.com/ytakahashi/coas/api"

func createSelectItems(parameter api.Parameter) (items []string) {
	if !parameter.Required {
		items = append(items, "")
	}
	for _, e := range parameter.ParameterEnums {
		items = append(items, e)
	}
	return items
}
