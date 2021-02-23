package internal

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/ytakahashi/burl/api"
)

// PromptUI is ui
type PromptUI struct{}

func (ui *PromptUI) readInput(parameter api.Parameter) string {
	prompt := promptui.Prompt{
		Label:    parameter.Name,
		Validate: createValidator(parameter),
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func createValidator(parameter api.Parameter) func(input string) error {
	if parameter.Required {
		return requiredValueValidator
	}
	return noopValidator
}

func noopValidator(input string) error {
	return nil
}

func requiredValueValidator(input string) error {
	if len(input) == 0 {
		return errors.New("Parameter required")
	}
	return nil
}
