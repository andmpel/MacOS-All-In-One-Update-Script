package main

import (
	"bytes"
	"fmt"
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

	// Load user's previous selections
	config, err := macup.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Prompt the user to select updates
	selectedUpdates, err := macup.SelectUpdates(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error selecting updates: %v\n", err)
		os.Exit(1)
	}

	// Save the user's selections
	config.SelectedUpdates = selectedUpdates
	if err := config.SaveConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
		os.Exit(1)
	}

	// Run the selected updates
	var wg sync.WaitGroup
	buffers := make(map[string]*bytes.Buffer)
	for _, updateName := range selectedUpdates {
		buffers[updateName] = &bytes.Buffer{}
	}

	for _, updateName := range selectedUpdates {
		for _, u := range macup.Updates {
			if u.Name == updateName {
				wg.Add(1)
				go func(u macup.Update) {
					defer wg.Done()
					u.Run(buffers[u.Name])
				}(u)
			}
		}
	}

	wg.Wait()

	// Print the output of each update function
	for _, updateName := range selectedUpdates {
		if buffer, ok := buffers[updateName]; ok {
			fmt.Println(buffer.String())
		}
	}
}