package main

import (
	"bytes"
	"flag"
	"fmt"
	"macup/macup"
	"os"
)

// Entry point of the update script
func main() {
	// Define and parse the --yes flag
	yes := flag.Bool("yes", false, "Use previous selections without prompting")
	flag.Parse()

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
	var selectedUpdates []string
	if *yes && len(config.SelectedUpdates) > 0 {
		selectedUpdates = config.SelectedUpdates
	} else {
		if *yes {
			println("No previous updates selected. Prompting for selection.")
		}
		var err error
		selectedUpdates, err = macup.SelectUpdates(config)
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
	}

	// Run the selected updates
	buffers := make(map[string]*bytes.Buffer)
	for _, updateName := range selectedUpdates {
		buffers[updateName] = &bytes.Buffer{}
	}
	for _, updateName := range selectedUpdates {
		for _, u := range macup.Updates {
			if u.Name == updateName {
				u.Run()
			}
		}
	}

	// Print the output of each update function
	for _, updateName := range selectedUpdates {
		if buffer, ok := buffers[updateName]; ok {
			fmt.Println(buffer.String())
		}
	}
}
