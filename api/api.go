package api

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

// ServerVariable is a struct to hold server variable.
type ServerVariable struct {
	Name        string
	Enum        []string
	Default     string
	Description string
}

// Server represents server information defined on oas.
type Server struct {
	URL       string
	Variables []ServerVariable
}

// Parameter represents an api parameter.
type Parameter struct {
	In               string
	Name             string
	Required         bool
	ParameterType    string
	ParameterFormat  string
	ParameterPattern string
}

// API is an API.
type API struct {
	Path        string
	Method      string
	OperationID string
	Description string
	Parameters  []Parameter
}

// ToText returns text which represents an api.
func (api *API) ToText() string {
	return fmt.Sprintf("%s %s (%s)", api.Method, api.Path, api.OperationID)
}

// ToText returns text which represents an api.
func (parameter *Parameter) ToText() string {
	txt := fmt.Sprintf("%s: %s", parameter.In, parameter.Name)
	if parameter.ParameterFormat != "" {
		txt += fmt.Sprintf("format: %s", parameter.ParameterFormat)
	}
	if parameter.ParameterPattern != "" {
		txt += fmt.Sprintf("pattern: %s", parameter.ParameterPattern)
	}
	if parameter.Required {
		txt += " (required)"
	}

	return txt
}

// ParseServers parses openapi3.Servers model to a slice of API model defined in this package.
func ParseServers(values openapi3.Servers) []Server {
	var servers []Server
	for _, s := range values {
		server := parseServer(*s)
		servers = append(servers, server)
	}
	return servers
}

func parseServer(s openapi3.Server) Server {
	var variables []ServerVariable
	for k, v := range s.Variables {

		enums := []string{}
		for _, e := range v.Enum {
			enums = append(enums, e.(string))
		}
		variable := ServerVariable{
			Name:        k,
			Enum:        enums,
			Default:     v.Default.(string),
			Description: v.Description,
		}
		variables = append(variables, variable)
	}
	return Server{
		URL:       s.URL,
		Variables: variables,
	}
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
