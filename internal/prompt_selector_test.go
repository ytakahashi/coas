package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ytakahashi/burl/api"
)

func TestCreateSelectItems_requitedParameter(t *testing.T) {
	parameter := api.Parameter{
		In:             "query",
		Name:           "sort",
		Required:       true,
		ParameterType:  "string",
		ParameterEnums: []string{"foo", "bar", "baz"},
	}

	actual := createSelectItems(parameter)
	expected := []string{"foo", "bar", "baz"}
	assert.Equal(t, expected, actual)
}

func TestCreateSelectItems_optionalParameter(t *testing.T) {
	parameter := api.Parameter{
		In:             "query",
		Name:           "sort",
		Required:       false,
		ParameterType:  "string",
		ParameterEnums: []string{"foo", "bar", "baz"},
	}

	actual := createSelectItems(parameter)
	expected := []string{"", "foo", "bar", "baz"}
	assert.Equal(t, expected, actual)
}
