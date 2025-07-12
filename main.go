package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"macup/macup"
	"os"
	"sync"
)

// Entry point of the update script
func main() {
	// Check for internet connectivity before proceeding
	if !macup.CheckInternet() {
		os.Exit(1)
	}

	// Define flags for each update function
	brewFlag := flag.Bool("brew", false, "Update Homebrew packages")
	vscodeFlag := flag.Bool("vscode", false, "Update VSCode extensions")
	gemFlag := flag.Bool("gem", false, "Update Ruby gems")
	nodeFlag := flag.Bool("node", false, "Update Node.js packages")
	cargoFlag := flag.Bool("cargo", false, "Update Rust packages")
	appstoreFlag := flag.Bool("appstore", false, "Update Mac App Store apps")
	macosFlag := flag.Bool("macos", false, "Update macOS system")
	flag.Parse()

	// List of update functions to run in parallel
	updateFuncs := []struct {
		fn   func(writer io.Writer)
		flag bool
	}{
		{macup.UpdateBrew, *brewFlag},
		{macup.UpdateVSCodeExt, *vscodeFlag},
		{macup.UpdateGem, *gemFlag},
		{macup.UpdateNodePkg, *nodeFlag},
		{macup.UpdateCargo, *cargoFlag},
		{macup.UpdateAppStore, *appstoreFlag},
		{macup.UpdateMacOS, *macosFlag},
	}

	runAll := flag.NFlag() == 0

	var wg sync.WaitGroup
	buffers := make([]*bytes.Buffer, len(updateFuncs))

	// Run each update function in a separate goroutine
	for i, u := range updateFuncs {
		if runAll || u.flag {
			wg.Add(1)
			buffers[i] = &bytes.Buffer{}
			go func(fn func(writer io.Writer), buffer *bytes.Buffer) {
				defer wg.Done()
				fn(buffer) // Execute update and write output to buffer
			}(u.fn, buffers[i])
		}
	}

	wg.Wait() // Wait for all updates to finish

	// Print the output of each update function
	for i, u := range updateFuncs {
		if (runAll || u.flag) && buffers[i] != nil {
			fmt.Println(buffers[i].String())
		}
	}
}