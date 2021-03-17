package internal

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/ytakahashi/coas/api"
)

// PromptUI is ui
type PromptUI struct{}

type validatorFactory interface {
	createTypeValidator(valueType string, isRequired bool) func(input string) error
	createPatternValidator(pattern string, isRequired bool) func(input string) error
}

type inputValidator struct{}

func (ui *PromptUI) readInput(parameter api.Parameter) (result string) {
	var err error
	if len(parameter.ParameterEnums) > 0 {
		prompt := promptui.Select{
			Label:  parameter.Name,
			Items:  createSelectItems(parameter),
			Stdout: &bellSkipper{},
		}
		_, result, err = prompt.Run()
	} else {
		prompt := promptui.Prompt{
			Label:    parameter.Name,
			Validate: createValidator(parameter, new(inputValidator)),
		}
		result, err = prompt.Run()
	}
	if err != nil {
		fmt.Printf("Failed to read input. %v\n", err)
		os.Exit(1)
	}
	return
}

// https://github.com/manifoldco/promptui/issues/49#issuecomment-573814976
type bellSkipper struct{}

func (bs *bellSkipper) Write(b []byte) (int, error) {
	const charBell = 7
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

func (bs *bellSkipper) Close() error {
	return os.Stderr.Close()
}
