package macup

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// Update represents a single update function.
type Update struct {
	Name        string
	Description string
	Run         func()
}

// Config represents the user's selections.
type Config struct {
	SelectedUpdates []string `json:"selected_updates"`
}

// Updates is a list of all available update functions.
var Updates = []Update{
	{"brew", "Update Homebrew packages", UpdateBrew},
	{"vscode", "Update VSCode extensions", UpdateVSCodeExt},
	{"gem", "Update Ruby gems", UpdateGem},
	{"node", "Update Node.js packages", UpdateNodePkg},
	{"cargo", "Update Rust packages", UpdateCargo},
	{"appstore", "Update Mac App Store apps", UpdateAppStore},
	{"macos", "Update macOS system", UpdateMacOS},
}

func Run(selectedUpdates []string) {
	for _, updateName := range selectedUpdates {
		for _, u := range Updates {
			if u.Name == updateName {
				go func(u Update) {
					fmt.Printf("Starting: %s\n", u.Description)
					u.Run()
					fmt.Printf("Finished: %s\n", u.Description)
				}(u)
			}
		}
	}
}

// SelectUpdates prompts the user to select which updates to run.
func SelectUpdates(config *Config) ([]string, error) {
	options := make([]string, len(Updates))
	for i, u := range Updates {
		options[i] = u.Description
	}

	defaults := []string{}
	for _, s := range config.SelectedUpdates {
		for _, u := range Updates {
			if u.Name == s {
				defaults = append(defaults, u.Description)
			}
		}
	}

	prompt := &survey.MultiSelect{
		Message: "Select the updates you want to run:",
		Options: options,
		Default: defaults,
	}

	var selectedDescriptions []string
	if err := survey.AskOne(prompt, &selectedDescriptions); err != nil {
		return nil, err
	}

	selectedNames := []string{}
	for _, desc := range selectedDescriptions {
		for _, u := range Updates {
			if u.Description == desc {
				selectedNames = append(selectedNames, u.Name)
			}
		}
	}

	return selectedNames, nil
}
