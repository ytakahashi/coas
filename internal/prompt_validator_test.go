package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedInputValidator struct {
	mock.Mock
}

func (v *mockedInputValidator) createTypeValidator(valueType string, isRequired bool) func(input string) error {
	v.Called(valueType, isRequired)
	return func(input string) error {
		return nil
	}
}

func (v *mockedInputValidator) createPatternValidator(pattern string, isRequired bool) func(input string) error {
	v.Called(pattern, isRequired)
	return func(input string) error {
		return nil
	}
}

func TestCreateValidator_WithType(t *testing.T) {
	parameter := Parameter{
		In:               "path",
		Name:             "id",
		ParameterType:    "string",
		ParameterPattern: "^abc$",
	}

	mock := new(mockedInputValidator)
	mock.On("createTypeValidator", "string", false)
	mock.On("createPatternValidator", "^abc$", false)
	createValidator(parameter, mock)
	mock.AssertCalled(t, "createTypeValidator", "string", false)
	mock.AssertCalled(t, "createPatternValidator", "^abc$", false)
}

func TestInputValidator_createTypeValidator_notRequired(t *testing.T) {
	sut := new(inputValidator)
	actual := sut.createTypeValidator("number", false)

	assert.EqualError(t, actual("a"), "invalid value")
	assert.Nil(t, actual("123"))
	assert.Nil(t, actual(""))
}

func TestInputValidator_createTypeValidator_required(t *testing.T) {
	sut := new(inputValidator)
	actual := sut.createTypeValidator("boolean", true)

	assert.EqualError(t, actual(""), "parameter required")
	assert.EqualError(t, actual("a"), "invalid value")
	assert.Nil(t, actual("true"))
}

func TestInputValidator_createPatternValidator_notRequired(t *testing.T) {
	sut := new(inputValidator)
	actual := sut.createPatternValidator("^abc$", false)

	assert.EqualError(t, actual("a"), "parameter does not match a pattern `^abc$`")
	assert.EqualError(t, actual("ac"), "parameter does not match a pattern `^abc$`")
	assert.Nil(t, actual("abc"))
	assert.Nil(t, actual(""))
}

func TestInputValidator_createPatternValidator_required(t *testing.T) {
	sut := new(inputValidator)
	actual := sut.createPatternValidator("^0[0-9]$", true)

	assert.EqualError(t, actual(""), "parameter required")
	assert.EqualError(t, actual("a"), "parameter does not match a pattern `^0[0-9]$`")
	assert.EqualError(t, actual("0"), "parameter does not match a pattern `^0[0-9]$`")
	assert.Nil(t, actual("01"))
}
