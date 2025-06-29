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
	k_green      = "\033[32m"
	k_red        = "\033[31m"
	k_yellow     = "\033[33m"
	k_clear      = "\033[0m"
	k_timeout    = 5 * time.Second          // Timeout for HTTP requests
	k_testURL    = "https://www.google.com" // URL to test internet connection
	k_gemCmdPath = "/usr/bin/gem"           // Path to the gem command
)

func printlnGreen(msg string) {
	fmt.Printf("\n%s%s%s\n", k_green, msg, k_clear)
}

func printlnYellow(msg string) {
	fmt.Printf("\n%s%s%s\n", k_yellow, msg, k_clear)
}

func printlnRed(msg string) {
	fmt.Fprintf(os.Stderr, "\n%s%s%s\n", k_red, msg, k_clear)
}

func checkCommand(cmd string) bool {
	_, err := exec.LookPath(cmd)

	if err != nil {
		printlnYellow(cmd + " is not installed.")
		return false
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
		runCommand("code", "--update-extensions")
	}
}

func UpdateGem() {
	printlnGreen("Updating Gems")
	gemPath, err := exec.LookPath("gem")
	if err != nil || gemPath == k_gemCmdPath {
		printlnRed("gem is not installed.")
		return
	}
	runCommand("gem", "update", "--user-install")
	runCommand("gem", "cleanup", "--user-install")
}

func UpdateNodePkg() {
	printlnGreen("Updating Node Packages")
	if checkCommand("node") {
		if checkCommand("npm") {
			runCommand("npm", "update", "-g")
		}

		if checkCommand("yarn") {
			runCommand("yarn", "upgrade", "--latest")
		}
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
		Timeout: k_timeout,
	}

	resp, err := client.Get(k_testURL)

	if err != nil {
		printlnRed("⚠️ No Internet Connection!!!")
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
