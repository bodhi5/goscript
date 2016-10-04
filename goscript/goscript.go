package goscript

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/tools/imports"
)

// GoScript type
type GoScript struct {
	source   string
	filePath string
	*exec.Cmd
}

// FilePath returns the current GoScipt's file path
func (gs GoScript) FilePath() string {
	return gs.filePath
}

// Clean deletes the generated goscript tempfile
func (gs GoScript) Clean() error {
	if gs.source == "" {
		return fmt.Errorf("error: no goscript tempfile to remove")
	}
	return os.Remove(gs.filePath)
}

// writeTempFile wraps a sting of go source code and adds
// required imports to make it a buildable go file
func (gs *GoScript) writeTempFile() error {
	curDir, err := os.Getwd()
	if err != nil {
		return err
	}
	tmpDir := os.TempDir()

	// Get source string sha
	shaBytes := sha1.Sum([]byte(gs.source))
	sha := hex.EncodeToString(shaBytes[:])

	gs.filePath = tmpDir + sha + ".goscript.go"

	// Wrapper for go source string
	fileSource := []byte("// +build goscript\n// Source SHA " + sha + "\npackage main\nfunc init(){\n\tos.Chdir(\"" + curDir + "\")\n}\nfunc main() {\n\t" + gs.source + "\n}")

	// Use cached file if nothing has changed
	if _, err := os.Stat(gs.filePath); !os.IsNotExist(err) {
		return nil
	}

	// Create file if it does not exist
	file, err := os.Create(gs.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// fileSource holds the go file re-written with imports added.
	fileSource, err = imports.Process(gs.filePath, fileSource, &imports.Options{})
	if err != nil {
		return err
	}

	if _, err := file.Write(fileSource); err != nil {
		return err
	}

	return nil
}

// NewFromString takes in a string of golang source and converts into a GoScript type with an embedded exec.Cmd
func NewFromString(source string, args ...string) (*GoScript, error) {
	gs := GoScript{source: source}

	if err := gs.writeTempFile(); err != nil {
		return nil, err
	}
	// Append stdin args
	cmdArgs := []string{"run", "-tags", "goscript", gs.filePath}
	cmdArgs = append(cmdArgs, args...)

	gs.Cmd = exec.Command("go", cmdArgs...)
	return &gs, nil
}

// NewFromFile takes in a filePath of golang source and converts it into a GoScript type with an embedded exec.Cmd
func NewFromFile(sourcePath string, args ...string) (*GoScript, error) {
	gs := GoScript{filePath: sourcePath}

	// Append stdin args
	cmdArgs := []string{"run", "-tags", "goscript", gs.filePath}
	cmdArgs = append(cmdArgs, args...)
	return &gs, nil
}
