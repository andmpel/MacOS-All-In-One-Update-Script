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
	"time"
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
	k_configFile = ".macup.json"            // Configuration file name
)

// printlnGreen prints a message in green color with a newline.
func printlnGreen(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "\n%s%s%s\n", k_green, msg, k_clear)
}

// printlnRed prints a message in red color (no newline).
func printlnRed(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "%s%s%s", k_red, msg, k_clear)
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
	// Allow only specific commands
	allowedCommands := map[string]bool{
		"brew":           true,
		"code":           true,
		"gem":            true,
		"npm":            true,
		"yarn":           true,
		"cargo":          true,
		"mas":            true,
		"softwareupdate": true,
	}

	if !allowedCommands[name] {
		printlnRed(writer, "Command not allowed: "+name)
		return
	}
	// Optionally validate arguments (e.g., no special characters)
	for _, arg := range args {
		if strings.ContainsAny(arg, "&|;$><") {
			printlnRed(writer, "Invalid Argument: "+arg)
			return
		}
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer
	if err := cmd.Run(); err != nil {
		printlnRed(writer, "Error running command: "+err.Error())
	}
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
