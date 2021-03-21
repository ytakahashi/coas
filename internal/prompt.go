package internal

import (
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/manifoldco/promptui"
)

// PromptType is type of prompt
type PromptType string

const (
	Validator   PromptType = "Prompt"
	Select      PromptType = "Select"
	FuzzySelect PromptType = "FuzzySelect"
)

// PromptRunner is an interface of prompt
type PromptRunner interface {
	Run() (string, error)
}

// InputValidator is an instance of PromptRunner to validate input
type InputValidator struct {
	label     string
	validator func(string) error
}

// ItemSelector is an instance of PromptRunner to select an item
type ItemSelector struct {
	label string
	items []string
}

// FuzzySelector is an instance of PromptRunner to fuzzy select an item
type FuzzySelector struct {
	Items []string
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

func (p *FuzzySelector) Run() (string, error) {
	index, err := fuzzyfinder.Find(
		p.Items,
		func(i int) string {
			return p.Items[i]
		},
	)
	return p.Items[index], err
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
	create(promptType PromptType, content PromptRunnerFactoryContext) PromptRunner
}

type PromptRunnerFactory struct{}

func (f PromptRunnerFactory) create(
	promptType PromptType, context PromptRunnerFactoryContext) PromptRunner {
	switch promptType {
	case Validator:
		return &InputValidator{
			label:     context.label,
			validator: context.validator,
		}
	case Select:
		return &ItemSelector{
			label: context.label,
			items: context.items,
		}
	case FuzzySelect:
		return &FuzzySelector{
			Items: context.items,
		}
	default:
		panic("error")
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
