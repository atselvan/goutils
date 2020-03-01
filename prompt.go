package utils

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strings"
)

const (
	promptConfirmErrMsg = "input can be either y/n"
	promptSelectMoreMsg = "Do you want to select one more ?"
)

// PromptString prompts a input menu on the console
func PromptString(name string, validateFunc func(input string) error) (string, error) {

	template := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | faint }} ",
	}

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("%s :", name),
		Templates: template,
		Validate:  validateFunc,
	}

	return prompt.Run()
}

// PromptSelect prompts a select menu on the console
func PromptSelect(name string, items []string) (int, string, error) {

	template := &promptui.SelectTemplates{
		Label:    "{{ . }} :",
		Active:   "\U0001F680 {{ . | bold }}",
		Inactive: "  {{ . | faint }}",
		Selected: "\U0001F680 {{ . | green | bold }}",
		Details:  "",
	}

	prompt := promptui.Select{
		Label:     name,
		Items:     items,
		Size:      len(items),
		Templates: template,
	}

	return prompt.Run()
}

// PromptConfirm prompts a confirmation menu on the console
func PromptConfirm(name string) (string, error) {

	template := &promptui.PromptTemplates{
		Prompt:  " : {{ . }}",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | faint }} ",
	}

	prompt := promptui.Prompt{
		Label:     name,
		Templates: template,
		Validate: func(input string) error {
			if strings.ToLower(input) == "y" || strings.ToLower(input) == "n" {
				return nil
			}
			return errors.New(promptConfirmErrMsg)
		},
	}

	return prompt.Run()
}

// PromptMultiSelect prompts a multi-select menu on the console
// This is a combination of the PromptSelect and PromptConfirm modules
func PromptMultiSelect(name string, items []string) ([]string, error) {

	var repeat = "y"
	var out []string
	for repeat != "n" {
		_, v, _ := PromptSelect(name, items)
		out = append(out, v)
		items = RemoveEntryFromSlice(items, v)
		if len(items) > 0 {
			repeat, _ = PromptConfirm(promptSelectMoreMsg)
		} else {
			repeat = "n"
		}
	}
	return out, nil
}
