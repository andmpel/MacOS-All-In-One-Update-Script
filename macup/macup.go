package macup

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
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
)

// Print a message in green color with a newline
func printlnGreen(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "\n%s%s%s\n", k_green, msg, k_clear)
}

// Print a message in yellow color (no newline)
func printlnYellow(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "%s%s%s", k_yellow, msg, k_clear)
}

// Check if a command exists in `PATH`, print warning if not
func checkCommand(writer io.Writer, cmd string) bool {
	_, err := exec.LookPath(cmd)

	if err != nil {
		printlnYellow(writer, cmd+" is not installed.")
		return false
	}

	return true
}

// Run a shell command and direct its output to writer
func runCommand(writer io.Writer, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer
	cmd.Run()
}

// Update Homebrew formulas and perform diagnostics
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

// Update VSCode extensions
func UpdateVSCodeExt(writer io.Writer) {
	printlnGreen(writer, "Updating VSCode Extensions")
	if checkCommand(writer, "code") {
		runCommand(writer, "code", "--update-extensions")
	}
}

// Update Ruby gems and clean up
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

// Update global Node.js, npm, and Yarn packages
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

// Update Rust Cargo crates by reinstalling each listed crate
func UpdateCargo(writer io.Writer) {
	printlnGreen(writer, "Updating Rust Cargo Crates")
	if checkCommand(writer, "cargo") {
		out, _ := exec.Command("cargo", "install", "--list").Output()
		lines := strings.SplitSeq(string(out), "\n")
		for line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				name := fields[0]
				runCommand(writer, "cargo", "install", name)
			}
		}
	}
}

// Update Mac App Store applications
func UpdateAppStore(writer io.Writer) {
	printlnGreen(writer, "Updating App Store Applications")
	if checkCommand(writer, "mas") {
		runCommand(writer, "mas", "upgrade")
	}
}

// Update macOS system software
func UpdateMacOS(writer io.Writer) {
	printlnGreen(writer, "Updating MacOS")
	runCommand(writer, "softwareupdate", "-i", "-a")
}

// Check for internet connectivity by making an HTTP request
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
