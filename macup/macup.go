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

const (
	k_green   = "\033[32m"
	k_red     = "\033[31m"
	k_yellow  = "\033[33m"
	k_clear   = "\033[0m"
	k_timeout = 5 * time.Second          // Timeout for HTTP requests
	k_testURL = "https://www.google.com" // URL to test internet connection
	k_gemCmdPath = "/usr/bin/gem"           // Path to the gem command
)

func printlnGreen(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "\n%s%s%s\n", k_green, msg, k_clear)
}

func printlnYellow(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "\n%s%s%s\n", k_yellow, msg, k_clear)
}

func printlnRed(writer io.Writer, msg string) {
	fmt.Fprintf(writer, "\n%s%s%s\n", k_red, msg, k_clear)
}

func checkCommand(writer io.Writer, cmd string) bool {
	_, err := exec.LookPath(cmd)

	if err != nil {
		printlnYellow(writer, cmd+" is not installed.")
		return false
	}

	return true
}

func runCommand(writer io.Writer, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = writer
	cmd.Stderr = writer
	cmd.Run()
}

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

func UpdateVSCode(writer io.Writer) {
	printlnGreen(writer, "Updating VSCode Extensions")
	if checkCommand(writer, "code") {
		runCommand(writer, "code", "--update-extensions")
	}
}

func UpdateGem(writer io.Writer) {
	printlnGreen(writer, "Updating Gems")
	gemPath, err := exec.LookPath("gem")
	if err != nil || gemPath == k_gemCmdPath {
		printlnRed(writer, "gem is not installed.")
		return
	}
	runCommand(writer, "gem", "update", "--user-install")
	runCommand(writer, "gem", "cleanup", "--user-install")
}

func UpdateNodePkg(writer io.Writer) {
	printlnGreen(writer, "Updating Node Packages")
	if checkCommand(writer, "node") {
		printlnGreen(writer, "Updating Npm Packages")
		if checkCommand(writer, "npm") {
			runCommand(writer, "npm", "update", "-g")
		}

		printlnGreen(writer, "Updating Yarn Packages")
		if checkCommand(writer, "yarn") {
			runCommand(writer, "yarn", "upgrade", "--latest")
		}
	}
}

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

func UpdateAppStore(writer io.Writer) {
	printlnGreen(writer, "Updating App Store Applications")
	if checkCommand(writer, "mas") {
		runCommand(writer, "mas", "upgrade")
	}
}

func UpdateMacOS(writer io.Writer) {
	printlnGreen(writer, "Updating MacOS")
	runCommand(writer, "softwareupdate", "-i", "-a")
}

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