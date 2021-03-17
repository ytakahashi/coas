package internal

import (
	"fmt"
	"testing"

	"github.com/ytakahashi/coas/api"
)

func TestBuildPreviewText(t *testing.T) {
	api := api.API{
		Path:        "/api/test",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "An api to get test value.",
	}
	actual := buildPreviewText(api, 150)
	expected := `OperationID: getTestValue

An api to get test value.`
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func TestBuildPreviewText_WithParameters(t *testing.T) {
	api := api.API{
		Path:        "/api/test/{id}",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "An api to get test value.",
		Parameters: []api.Parameter{
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
	actual := buildPreviewText(api, 150)
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

func TestBuildPreviewText_LongDescription(t *testing.T) {
	api := api.API{
		Path:        "/api/test",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
	actual := buildPreviewText(api, 150)
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
