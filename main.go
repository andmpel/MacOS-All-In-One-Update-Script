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
	macup.UpdateNodePkg()
	macup.UpdateCargo()
	macup.UpdateAppStore()
	macup.UpdateMacOS()
}
