// TODO - Need to rewrite tests

package macup

import (
	"bytes"
	"testing"
)

// Helper to capture output
func captureOutput(f func(io.)) string {
	buf := &bytes.Buffer{}
	f(buf)
	return buf.String()
}

func TestPrintlnGreen(t *testing.T) {
	out := captureOutput(func(w io.) { printlnGreen(w, "Hello") })
	if !bytes.Contains([]byte(out), []byte("Hello")) {
		t.Errorf("Expected output to contain 'Hello', got: %s", out)
	}
}

func TestPrintlnYellow(t *testing.T) {
	buf := &bytes.Buffer{}
	printlnYellow(buf, "Warning")
	if !bytes.Contains(buf.Bytes(), []byte("Warning")) {
		t.Errorf("Expected output to contain 'Warning', got: %s", buf.String())
	}
}

func TestCheckCommand(t *testing.T) {
	buf := &bytes.Buffer{}
	if !checkCommand(buf, "echo") {
		t.Error("Expected 'echo' to exist")
	}
	if checkCommand(buf, "nonexistentcommand") {
		t.Error("Expected 'nonexistentcommand' to not exist")
	}
}

func TestRunCommand(t *testing.T) {
	buf := &bytes.Buffer{}
	runCommand(buf, "echo", "test")
	if !bytes.Contains(buf.Bytes(), []byte("test")) {
		t.Errorf("Expected output to contain 'test', got: %s", buf.String())
	}
}

func TestUpdateBrew(t *testing.T) {
	buf := &bytes.Buffer{}
	UpdateBrew(buf) // Should not panic, may print warning if brew not installed
}

func TestUpdateVSCodeExt(t *testing.T) {
	buf := &bytes.Buffer{}
	UpdateVSCodeExt(buf)
}

func TestUpdateGem(t *testing.T) {
	buf := &bytes.Buffer{}
	UpdateGem(buf)
}

func TestUpdateNodePkg(t *testing.T) {
	buf := &bytes.Buffer{}
	UpdateNodePkg(buf)
}

func TestUpdateCargo(t *testing.T) {
	buf := &bytes.Buffer{}
	UpdateCargo(buf)
}

func TestUpdateAppStore(t *testing.T) {
	buf := &bytes.Buffer{}
	UpdateAppStore(buf)
}

func TestUpdateMacOS(t *testing.T) {
	buf := &bytes.Buffer{}
	UpdateMacOS(buf)
}

func TestCheckInternet(t *testing.T) {
	_ = CheckInternet() // Should not panic
}
