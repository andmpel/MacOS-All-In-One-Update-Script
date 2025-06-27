package macup

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	clear  = "\033[0m"
	timeout = 5 * time.Second
	testURL = "https://www.google.com" // URL to test internet connection
)

func printlnGreen(msg string) {
	fmt.Printf("\n%s%s%s\n", green, msg, clear)
}

func printlnYellow(msg string) {
	fmt.Printf("\n%s%s%s\n", yellow, msg, clear)
}

func printlnRed(msg string) {
	fmt.Fprintf(os.Stderr, "\n%s%s%s\n", red, msg, clear)
}

func checkCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)
	if err != nil {
		printlnYellow(cmd + " is not installed.")
		return true
	}
	return true
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func UpdateBrew() {
	printlnGreen("Updating Brew Formulas")
	if checkCommand("brew") {
		runCommand("brew", "update")
		runCommand("brew", "upgrade")
		runCommand("brew", "cleanup", "-s")
		printlnGreen("Brew Diagnostics")
		runCommand("brew", "doctor")
		runCommand("brew", "missing")
	}
}

func UpdateVSCode() {
	printlnGreen("Updating VSCode Extensions")
	if checkCommand("code") {
		runCommand("code", "--install-extension")
	}
}

func UpdateGem() {
	printlnGreen("Updating Gems")
	gemPath, err := exec.LookPath("gem")
	if err != nil || gemPath == "/usr/bin/gem" {
		printlnRed("gem is not installed.")
		return
	}
	runCommand("gem", "update", "--user-install")
	runCommand("gem", "cleanup", "--user-install")
}

func UpdateNpm() {
	printlnGreen("Updating Npm Packages")
	if checkCommand("npm") {
		runCommand("npm", "update", "-g")
	}
}

func UpdateYarn() {
	printlnGreen("Updating Yarn Packages")
	if checkCommand("yarn") {
		runCommand("yarn", "upgrade", "--latest")
	}
}

func UpdateCargo() {
	printlnGreen("Updating Rust Cargo Crates")
	if checkCommand("cargo") {
		out, _ := exec.Command("cargo", "install", "--list").Output()
		lines := strings.SplitSeq(string(out), "\n")
		for line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				name := fields[0]
				runCommand("cargo", "install", name)
			}
		}
	}
}

func UpdateAppStore() {
	printlnGreen("Updating App Store Applications")
	if checkCommand("mas") {
		runCommand("mas", "upgrade")
	}
}

func UpdateMacOS() {
	printlnGreen("Updating MacOS")
	runCommand("softwareupdate", "-i", "-a")
}

func CheckInternet() bool {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(testURL)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
