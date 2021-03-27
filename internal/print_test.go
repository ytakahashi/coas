package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockedPrinter struct {
	printed string
}

func (p *mockedPrinter) Print(text string) {
	p.printed = text
}

func TestPrintAPIDetails(t *testing.T) {
	api := API{
		Path:        "/api/test",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "An api to get test value.",
	}

	mock := &mockedPrinter{}
	PrintAPIDetails(api, mock)
	expected := `GET /api/test
OperationID: getTestValue
An api to get test value.`
	assert.Equal(t, expected, mock.printed)
}

func TestPrintAPIDetails_with_parameters(t *testing.T) {
	api := API{
		Path:        "/api/test",
		Method:      "GET",
		OperationID: "getTestValue",
		Description: "An api to get test value.",
		Parameters: []Parameter{
			{
				In:              "query",
				Name:            "offset",
				ParameterType:   "integer",
				ParameterFormat: "int32",
			},
		},
	}

	mock := &mockedPrinter{}
	PrintAPIDetails(api, mock)
	expected := `GET /api/test
OperationID: getTestValue
An api to get test value.

Parameters:
- query: offset
  - type: integer
  - format: int32
`
	assert.Equal(t, expected, mock.printed)
}
