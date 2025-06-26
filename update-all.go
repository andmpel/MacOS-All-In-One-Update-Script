package main

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

func updateBrew() {
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

func updateVSCode() {
	printlnGreen("Updating VSCode Extensions")
	if checkCommand("code") {
		runCommand("code", "--install-extension")
	}
}

func updateGem() {
	printlnGreen("Updating Gems")
	gemPath, err := exec.LookPath("gem")
	if err != nil || gemPath == "/usr/bin/gem" {
		printlnRed("gem is not installed.")
		return
	}
	runCommand("gem", "update", "--user-install")
	runCommand("gem", "cleanup", "--user-install")
}

func updateNpm() {
	printlnGreen("Updating Npm Packages")
	if checkCommand("npm") {
		runCommand("npm", "update", "-g")
	}
}

func updateYarn() {
	printlnGreen("Updating Yarn Packages")
	if checkCommand("yarn") {
		runCommand("yarn", "upgrade", "--latest")
	}
}

func updateCargo() {
	printlnGreen("Updating Rust Cargo Crates")
	if checkCommand("cargo") {
		out, _ := exec.Command("cargo", "install", "--list").Output()
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 0 {
				name := fields[0]
				runCommand("cargo", "install", name)
			}
		}
	}
}

func updateAppStore() {
	printlnGreen("Updating App Store Applications")
	if checkCommand("mas") {
		runCommand("mas", "upgrade")
	}
}

func updateMacOS() {
	printlnGreen("Updating MacOS")
	runCommand("softwareupdate", "-i", "-a")
}

func checkInternet() bool {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func main() {
	if !checkInternet() {
		os.Exit(1)
	}
	updateBrew()
	updateVSCode()
	updateGem()
	updateNpm()
	updateYarn()
	updateCargo()
	updateAppStore()
	updateMacOS()
}
