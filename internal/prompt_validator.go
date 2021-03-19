package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func createValidator(parameter Parameter, factory validatorFactory) func(input string) error {
	var validators = []func(input string) error{}
	if parameter.ParameterType != "" {
		typeValidator := factory.createTypeValidator(parameter.ParameterType, parameter.Required)
		validators = append(validators, typeValidator)
	}

	if parameter.ParameterPattern != "" {
		patternValidator := factory.createPatternValidator(parameter.ParameterPattern, parameter.Required)
		validators = append(validators, patternValidator)
	}

	return func(input string) error {
		if len(validators) == 0 {
			return nil
		}
		for _, validator := range validators {
			if err := validator(input); err != nil {
				return err
			}
		}
		return nil
	}
}

func (v *inputValidator) createTypeValidator(valueType string, isRequired bool) func(input string) error {
	return func(input string) error {
		if len(input) == 0 {
			if isRequired {
				return errors.New("parameter required")
			}
			return nil
		}
		if valueType == "number" {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return errors.New("invalid value")
			}
		} else if valueType == "integer" {
			_, err := strconv.ParseInt(input, 10, 32)
			if err != nil {
				return errors.New("invalid value")
			}
		} else if valueType == "boolean" {
			if input != "true" && input != "false" {
				return errors.New("invalid value")
			}
		}
		return nil
	}
}

func (v *inputValidator) createPatternValidator(pattern string, isRequired bool) func(input string) error {
	re := regexp.MustCompile(pattern)
	return func(input string) error {
		if len(input) == 0 {
			if isRequired {
				return errors.New("parameter required")
			}
			return nil
		}
		if !re.MatchString(input) {
			return fmt.Errorf("parameter does not match a pattern `%s`", pattern)
		}
		return nil
	}
}
