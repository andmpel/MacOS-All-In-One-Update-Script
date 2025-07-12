package main

import (
	"bytes"
	"fmt"
	"io"
	"macup/macup"
	"os"
	"sync"
)

func main() {
	if !macup.CheckInternet() {
		os.Exit(1)
	}

	updateFuncs := []struct {
		fn func(writer io.Writer)
	}{
		{macup.UpdateBrew},
		{macup.UpdateVSCode},
		{macup.UpdateGem},
		{macup.UpdateNodePkg},
		{macup.UpdateCargo},
		{macup.UpdateAppStore},
		{macup.UpdateMacOS},
	}

	var wg sync.WaitGroup
	buffers := make([]*bytes.Buffer, len(updateFuncs))

	for i, u := range updateFuncs {
		wg.Add(1)
		buffers[i] = &bytes.Buffer{}
		go func(fn func(writer io.Writer), buffer *bytes.Buffer) {
			defer wg.Done()
			fn(buffer)
		}(u.fn, buffers[i])
	}

	wg.Wait()

	for i := range updateFuncs {
		fmt.Println(buffers[i].String())
	}
}
