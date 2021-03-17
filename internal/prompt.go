package internal

import (
	"os"

	"github.com/manifoldco/promptui"
)

// PromptRunner is an interface of prompt
type PromptRunner interface {
	Run() (string, error)
}

// ItemSelector is an instance of PromptRunner to select an item
type ItemSelector struct {
	label string
	items []string
}

// InputValidator is an instance of PromptRunner to validate input
type InputValidator struct {
	label     string
	validator func(string) error
}

func (p *ItemSelector) Run() (string, error) {
	prompt := promptui.Select{
		Label:  p.label,
		Items:  p.items,
		Stdout: &bellSkipper{},
	}
	_, result, err := prompt.Run()
	return result, err
}

func (p *InputValidator) Run() (string, error) {
	prompt := promptui.Prompt{
		Label:    p.label,
		Validate: p.validator,
	}
	return prompt.Run()
}

type PromptRunnerFactoryContext struct {
	label     string
	items     []string
	validator func(string) error
}

type IPromptRunnerFactory interface {
	create(content PromptRunnerFactoryContext) PromptRunner
}

type PromptRunnerFactory struct{}

func (f PromptRunnerFactory) create(context PromptRunnerFactoryContext) PromptRunner {
	if len(context.items) > 0 {
		return &ItemSelector{
			label: context.label,
			items: context.items,
		}
	}
	return &InputValidator{
		label:     context.label,
		validator: context.validator,
	}
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
