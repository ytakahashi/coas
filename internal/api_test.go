package internal

import (
	"fmt"
	"testing"
)

func Test_API_ToText(t *testing.T) {
	api := API{
		Path:        "/api/test",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "An api to get test value.",
	}
	actual := api.ToText()
	expected := "GET /api/test (getTestValue)"
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func Test_API_BuildDetailedDescription(t *testing.T) {
	api := API{
		Path:        "/api/test",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "An api to get test value.",
	}
	actual := api.BuildDetailedDescription(150)
	expected := `OperationID: getTestValue

An api to get test value.`
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func Test_API_BuildDetailedDescription_WithParameters(t *testing.T) {
	api := API{
		Path:        "/api/test/{id}",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "An api to get test value.",
		Parameters: []Parameter{
			{
				In:   "path",
				Name: "id",
			},
			{
				In:   "query",
				Name: "sort",
			},
		},
	}
	actual := api.BuildDetailedDescription(150)
	expected := `OperationID: getTestValue

An api to get test value.

Parameters:
- path: id
- query: sort
`
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func Test_API_BuildDetailedDescription_LongDescription(t *testing.T) {
	api := API{
		Path:        "/api/test",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
	actual := api.BuildDetailedDescription(150)
	expected := `OperationID: getTestValue

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
minim veniam, quis nostrud exercitation ullamco laboris nisi ut
aliquip ex ea commodo consequat. Duis aute irure dolor in
reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
culpa qui officia deserunt mollit anim id est laborum.`
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func Test_Parameter_ToText(t *testing.T) {
	parameter := Parameter{
		In:   "path",
		Name: "id",
	}
	actual := parameter.ToText()
	expected := "path: id"
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func Test_Parameter_ToText_WithFormat(t *testing.T) {
	parameter := Parameter{
		In:              "query",
		Name:            "offset",
		ParameterType:   "integer",
		ParameterFormat: "int32",
	}
	actual := parameter.ToText()
	expected := `query: offset
  - type: integer
  - format: int32`
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func Test_Parameter_ToText_WithSchemaPattern(t *testing.T) {
	parameter := Parameter{
		In:               "query",
		Name:             "sampleKey",
		ParameterType:    "string",
		ParameterPattern: "^[0-9A-F]{9}$",
	}
	actual := parameter.ToText()
	expected := `query: sampleKey
  - type: string
  - pattern: ` + "`" + `^[0-9A-F]{9}$` + "`"
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func Test_Parameter_ToText_WithSchemaEnums(t *testing.T) {
	parameter := Parameter{
		In:             "query",
		Name:           "sort",
		ParameterType:  "string",
		ParameterEnums: []string{"createdAt", "updatedAt"},
	}
	actual := parameter.ToText()
	expected := `query: sort
  - type: string
  - enums:
    - createdAt
    - updatedAt`
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}
