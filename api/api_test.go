package api

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

func Test_Parameter_ToText_WithSchema(t *testing.T) {
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
