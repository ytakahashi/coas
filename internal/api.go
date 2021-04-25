package internal

import (
	"bytes"
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
	ParameterEnums   []string
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

// GetPathParameters returns path parameters of an api.
func (api *API) GetPathParameters() []Parameter {
	return api.filterParameters("path")
}

// GetQueryParameters returns query parameters of an api.
func (api *API) GetQueryParameters() []Parameter {
	return api.filterParameters("query")
}

func (api *API) filterParameters(target string) []Parameter {
	result := make([]Parameter, 0)
	for _, p := range api.Parameters {
		if p.In == target {
			result = append(result, p)
		}
	}
	return result
}

// PrintParameters prints parameter information
func (api *API) PrintParameters() string {
	parametersText := ""
	if api.Parameters != nil {
		var buffer bytes.Buffer
		buffer.WriteString("\n\nParameters:\n")
		for _, p := range api.Parameters {
			buffer.WriteString(fmt.Sprintf("- %s\n", p.ToText()))
		}
		parametersText += buffer.String()

	}
	return parametersText
}

// ToText returns text which represents an api.
func (parameter *Parameter) ToText() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s: %s", parameter.In, parameter.Name))
	if parameter.Required {
		buffer.WriteString(" (required)")
	}
	if parameter.ParameterType != "" {
		buffer.WriteString(fmt.Sprintf("\n  - type: %s", parameter.ParameterType))
	}
	if parameter.ParameterFormat != "" {
		buffer.WriteString(fmt.Sprintf("\n  - format: %s", parameter.ParameterFormat))
	}
	if parameter.ParameterPattern != "" {
		buffer.WriteString(fmt.Sprintf("\n  - pattern: `%s`", parameter.ParameterPattern))
	}
	if len(parameter.ParameterEnums) > 0 {
		buffer.WriteString("\n  - enums:")
		for _, e := range parameter.ParameterEnums {
			buffer.WriteString(fmt.Sprintf("\n    - %s", e))
		}
	}
	return buffer.String()
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
			enums = append(enums, fmt.Sprintf("%v", e))
		}
		variable := ServerVariable{
			Name:        k,
			Enum:        enums,
			Default:     fmt.Sprintf("%v", v.Default),
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
	enums := make([]string, 0)
	for _, e := range param.Schema.Value.Enum {
		enums = append(enums, e.(string))
	}
	return Parameter{
		Name:             param.Name,
		In:               param.In,
		Required:         param.Required,
		ParameterType:    param.Schema.Value.Type,
		ParameterFormat:  param.Schema.Value.Format,
		ParameterPattern: param.Schema.Value.Pattern,
		ParameterEnums:   enums,
	}
}
