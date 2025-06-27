package main

import (
	"macup/macup"
	"os"
)

func main() {
	if !macup.CheckInternet() {
		os.Exit(1)
	}

	macup.UpdateBrew()
	macup.UpdateVSCode()
	macup.UpdateGem()
	macup.UpdateNpm()
	macup.UpdateYarn()
	macup.UpdateCargo()
	macup.UpdateAppStore()
	macup.UpdateMacOS()
}
