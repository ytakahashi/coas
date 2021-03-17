package internal

import (
	"fmt"
	"testing"

	"github.com/ytakahashi/coas/api"
)

func TestBuildURL(t *testing.T) {
	path := "/v1/myResources/{resourceId}"
	param1 := api.Parameter{
		In:       "path",
		Name:     "resourceId",
		Required: true,
	}
	param2 := api.Parameter{
		In:       "query",
		Name:     "sort",
		Required: true,
	}

	api := api.API{
		Path:        path,
		Method:      "GET",
		OperationID: "getMyResource",
		Description: "get my resource",
		Parameters:  []api.Parameter{param1, param2},
	}

	actual := BuildURL(api, new(mockedInputHandler))
	expected := "/v1/myResources/foo?sort=foo"
	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func TestReplacePathParam(t *testing.T) {
	path := "/v1/myResources/{resourceId}"
	param := api.Parameter{
		In:       "path",
		Name:     "resourceId",
		Required: true,
	}
	value := "123456"
	actual := replacePathParam(path, param, value)
	expected := "/v1/myResources/123456"

	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func TestAppendQueryStringForTheFirstQuery(t *testing.T) {
	query := ""
	param := api.Parameter{
		In:       "query",
		Name:     "sort",
		Required: true,
	}
	value := "myId"
	actual := appendQueryString(query, param, value)
	expected := "sort=myId"

	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func TestAppendQueryStringForTheSecondQuery(t *testing.T) {
	query := "sort=myId"
	param := api.Parameter{
		In:       "query",
		Name:     "limit",
		Required: true,
	}
	value := "100"
	actual := appendQueryString(query, param, value)
	expected := "sort=myId&limit=100"

	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

func TestAppendQueryStringForOptionalQuery(t *testing.T) {
	query := "sort=myId"
	param := api.Parameter{
		In:       "query",
		Name:     "offset",
		Required: false,
	}
	value := ""
	actual := appendQueryString(query, param, value)
	expected := "sort=myId"

	if actual != expected {
		fmt.Println(actual)
		fmt.Println(expected)
		t.Error("Error")
	}
}

type mockedInputHandler struct{}

func (mock *mockedInputHandler) readInput(parameter api.Parameter) string {
	return "foo"
}
