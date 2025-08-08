package macup

import (
	"io"
	"os/exec"
	"strings"
)

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
		lines := strings.SplitSeq(string(out), "\n")
		for line := range lines {
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
