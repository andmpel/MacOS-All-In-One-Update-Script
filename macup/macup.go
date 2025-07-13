package macup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/AlecAivazis/survey/v2"
)

// Terminal color codes and constants
const (
	k_green      = "\033[32m"               // Green text
	k_red        = "\033[31m"               // Red text
	k_yellow     = "\033[33m"               // Yellow text
	k_clear      = "\033[0m"                // Reset color
	k_timeout    = 5 * time.Second          // Timeout for HTTP requests
	k_testURL    = "https://www.google.com" // URL to test internet connection
	k_gemCmdPath = "/usr/bin/gem"           // Path to the gem command
	k_configFile = ".macup.json"           // Configuration file name
)

// Update represents a single update function.
type Update struct {
	Name        string
	Description string
	Run         func(writer io.Writer)
}

// Config represents the user's selections.
type Config struct {
	SelectedUpdates []string `json:"selected_updates"`
}

// printlnGreen prints a message in green color with a newline.
func printlnGreen(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "\n%s%s%s\n", k_green, msg, k_clear)
}

// printlnYellow prints a message in yellow color (no newline).
func printlnYellow(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "%s%s%s", k_yellow, msg, k_clear)
}

// checkCommand checks if a command exists in `PATH`, print warning if not.
func checkCommand(writer io.Writer, cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		printlnYellow(writer, cmd+" is not installed.")
		return false
	}
	return true
}

// runCommand runs a shell command and directs its output to writer.
func runCommand(writer io.Writer, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer
	cmd.Run()
}

// UpdateBrew updates Homebrew formulas and perform diagnostics.
func UpdateBrew(writer io.Writer) {
	printlnGreen(writer, "Updating Brew Formulas")
	if checkCommand(writer, "brew") {
		runCommand(writer, "brew", "update")
		runCommand(writer, "brew", "upgrade")
		runCommand(writer, "brew", "cleanup", "-s")
		printlnGreen(writer, "Brew Diagnostics")
		runCommand(writer, "brew", "doctor")
		runCommand(writer, "brew", "missing")
	}
}

// UpdateVSCodeExt updates VSCode extensions.
func UpdateVSCodeExt(writer io.Writer) {
	printlnGreen(writer, "Updating VSCode Extensions")
	if checkCommand(writer, "code") {
		runCommand(writer, "code", "--update-extensions")
	}
}

// UpdateGem updates Ruby gems and clean up.
func UpdateGem(writer io.Writer) {
	printlnGreen(writer, "Updating Gems")
	gemPath, err := exec.LookPath("gem")
	if err != nil || gemPath == k_gemCmdPath {
		printlnYellow(writer, "gem is not installed.")
		return
	}
	runCommand(writer, "gem", "update", "--user-install")
	runCommand(writer, "gem", "cleanup", "--user-install")
}

// UpdateNodePkg updates global Node.js, npm, and Yarn packages.
func UpdateNodePkg(writer io.Writer) {
	printlnGreen(writer, "Updating Node Packages")
	if checkCommand(writer, "node") {
		printlnGreen(writer, "Updating Npm Packages")
		if checkCommand(writer, "npm") {
			runCommand(writer, "npm", "update", "-g")
		}

		printlnGreen(writer, "Updating Yarn Packages")
		if checkCommand(writer, "yarn") {
			runCommand(writer, "yarn", "global", "upgrade", "--latest")
		}
	}
}

// UpdateCargo updates Rust Cargo crates by reinstalling each listed crate.
func UpdateCargo(writer io.Writer) {
	printlnGreen(writer, "Updating Rust Cargo Crates")
	if checkCommand(writer, "cargo") {
		out, _ := exec.Command("cargo", "install", "--list").Output()
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				name := fields[0]
				runCommand(writer, "cargo", "install", name)
			}
		}
	}
}

// UpdateAppStore updates Mac App Store applications.
func UpdateAppStore(writer io.Writer) {
	printlnGreen(writer, "Updating App Store Applications")
	if checkCommand(writer, "mas") {
		runCommand(writer, "mas", "upgrade")
	}
}

// UpdateMacOS updates macOS system software.
func UpdateMacOS(writer io.Writer) {
	printlnGreen(writer, "Updating MacOS")
	runCommand(writer, "softwareupdate", "-i", "-a")
}

// CheckInternet checks for internet connectivity by making an HTTP request.
func CheckInternet() bool {
	client := http.Client{
		Timeout: k_timeout,
	}

	resp, err := client.Get(k_testURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%s%s%s\n", k_red, "⚠️ No Internet Connection!!!", k_clear)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

// homeDir returns the user's home directory.
func homeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	return home
}

// configPath returns the full path to the configuration file.
func configPath() string {
	return filepath.Join(homeDir(), k_configFile)
}

// LoadConfig loads the user's selections from the configuration file.
func LoadConfig() (*Config, error) {
	path := configPath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{}, nil
		}
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// SaveConfig saves the user's selections to the configuration file.
func (c *Config) SaveConfig() error {
	path := configPath()
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
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

// Run runs the selected update functions.
func Run(writer io.Writer, selectedUpdates []string) {
	var wg sync.WaitGroup
	for _, updateName := range selectedUpdates {
		for _, u := range Updates {
			if u.Name == updateName {
				wg.Add(1)
				go func(u Update) {
					defer wg.Done()
					u.Run(writer)
				}(u)
			}
		}
	}
	wg.Wait()
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
