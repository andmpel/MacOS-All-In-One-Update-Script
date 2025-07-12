package main

import (
	"bytes"
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

	// List of update functions to run in parallel
	updateFuncs := []struct {
		fn func(writer io.Writer)
	}{
		{macup.UpdateBrew},      // Update Homebrew packages
		{macup.UpdateVSCodeExt}, // Update VSCode extensions
		{macup.UpdateGem},       // Update Ruby gems
		{macup.UpdateNodePkg},   // Update Node.js packages
		{macup.UpdateCargo},     // Update Rust packages
		{macup.UpdateAppStore},  // Update Mac App Store apps
		{macup.UpdateMacOS},     // Update macOS system
	}

	var wg sync.WaitGroup
	buffers := make([]*bytes.Buffer, len(updateFuncs))

	// Run each update function in a separate goroutine
	for i, u := range updateFuncs {
		wg.Add(1)
		buffers[i] = &bytes.Buffer{}
		go func(fn func(writer io.Writer), buffer *bytes.Buffer) {
			defer wg.Done()
			fn(buffer) // Execute update and write output to buffer
		}(u.fn, buffers[i])
	}

	wg.Wait() // Wait for all updates to finish

	// Print the output of each update function
	for i := range updateFuncs {
		fmt.Println(buffers[i].String())
	}
}
