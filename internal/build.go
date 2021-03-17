package internal

import (
	"fmt"
	"os"
	"strings"
)

// InputHandler handles user input
type InputHandler interface {
	readInput(parameter Parameter) string
}

// BuildURL builds url from an api
func BuildURL(target API, handler InputHandler) string {
	fmt.Println(target.ToText())
	path := target.Path
	for _, p := range target.GetPathParameters() {
		result := handler.readInput(p)
		path = replacePathParam(path, p, result)
	}

	query := ""
	for _, p := range target.GetQueryParameters() {
		result := handler.readInput(p)
		query = appendQueryString(query, p, result)
	}

	var url string
	if query == "" {
		url = path
	} else {
		url = fmt.Sprintf("%s?%s", path, query)
	}
	return url
}

func replacePathParam(path string, param Parameter, value string) string {
	target := fmt.Sprintf("{%s}", param.Name)
	return strings.Replace(path, target, value, 1)
}

func appendQueryString(query string, param Parameter, value string) string {
	if value == "" {
		return query
	}
	q := fmt.Sprintf("%s=%s", param.Name, value)
	if query == "" {
		return q
	}
	return fmt.Sprintf("%s&%s", query, q)
}

type validatorFactory interface {
	createTypeValidator(valueType string, isRequired bool) func(input string) error
	createPatternValidator(pattern string, isRequired bool) func(input string) error
}

type inputValidator struct{}

// PromptUI is ui
type PromptUI struct {
	promptRunnerFactory PromptRunnerFactory
}

func (ui *PromptUI) readInput(parameter Parameter) (result string) {
	prompt := ui.promptRunnerFactory.create(
		PromptRunnerFactoryContext{
			label:     parameter.Name,
			items:     createSelectItems(parameter),
			validator: createValidator(parameter, new(inputValidator)),
		},
	)
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Failed to read input. %v\n", err)
		os.Exit(1)
	}
	return
}
