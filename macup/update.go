package macup

import (
	"os/exec"
	"strings"
)

// UpdateBrew updates Homebrew formulas and perform diagnostics.
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

// UpdateVSCodeExt updates VSCode extensions.
func UpdateVSCodeExt() {
	printlnGreen("Updating VSCode Extensions")
	if checkCommand("code") {
		runCommand("code", "--update-extensions")
	}
}

// UpdateGem updates Ruby gems and clean up.
func UpdateGem() {
	printlnGreen("Updating Gems")
	gemPath, err := exec.LookPath("gem")
	if err != nil || gemPath == k_gemCmdPath {
		printlnYellow("gem is not installed.")
		return
	}
	runCommand("gem", "update", "--user-install")
	runCommand("gem", "cleanup", "--user-install")
}

// UpdateNodePkg updates global Node.js, npm, and Yarn packages.
func UpdateNodePkg() {
	printlnGreen("Updating Node Packages")
	if checkCommand("node") {
		printlnGreen("Updating Npm Packages")
		if checkCommand("npm") {
			runCommand("npm", "update", "-g")
		}

		printlnGreen("Updating Yarn Packages")
		if checkCommand("yarn") {
			runCommand("yarn", "global", "upgrade", "--latest")
		}
	}
}

// UpdateCargo updates Rust Cargo crates by reinstalling each listed crate.
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

// UpdateAppStore updates Mac App Store applications.
func UpdateAppStore() {
	printlnGreen("Updating App Store Applications")
	if checkCommand("mas") {
		runCommand("mas", "upgrade")
	}
}

// UpdateMacOS updates macOS system software.
func UpdateMacOS() {
	printlnGreen("Updating MacOS")
	runCommand("softwareupdate", "-i", "-a")
}
