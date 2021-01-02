package api

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// Parameter represents an api parameter
type Parameter struct {
	In               string
	Name             string
	Required         bool
	ParameterType    string
	ParameterFormat  string
	ParameterPattern string
}

// API is an API
type API struct {
	Path        string
	Method      string
	OperationID string
	Description string
	Parameters  []Parameter
}

// ParsePaths parses openapi3.Paths model to a slice of API model defined in this package
func ParsePaths(paths openapi3.Paths) []API {
	var apis []API
	for k, v := range paths {
		apisPerPath := parsePath(k, *v)
		apis = append(apis, apisPerPath...)
	}
	return apis
}

func parsePath(path string, pathItem openapi3.PathItem) []API {
	var apis []API
	operations := pathItem.Operations()
	for method, operation := range operations {
		if operation == nil {
			continue
		}

		api := API{
			Path:        path,
			Method:      method,
			OperationID: operation.OperationID,
			Description: operation.Description,
			Parameters:  buildParameters(operation.Parameters),
		}
		apis = append(apis, api)

	}
	return apis
}

func buildParameters(params openapi3.Parameters) []Parameter {
	var parameters []Parameter
	for _, p := range params {
		parameters = append(parameters, newParameter(*p.Value))
	}
	return parameters
}

func newParameter(param openapi3.Parameter) Parameter {
	return Parameter{
		Name:             param.Name,
		In:               param.In,
		Required:         param.Required,
		ParameterType:    param.Schema.Value.Type,
		ParameterFormat:  param.Schema.Value.Format,
		ParameterPattern: param.Schema.Value.Pattern,
	}
}
