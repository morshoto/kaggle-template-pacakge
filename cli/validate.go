package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go run scripts/check_notebook.go <notebook.ipynb>")
		os.Exit(2)
	}

	nbPath := os.Args[1]
	if filepath.Ext(nbPath) != ".ipynb" {
		fmt.Fprintf(os.Stderr, "not a notebook: %s\n", nbPath)
		os.Exit(2)
	}

	if _, err := os.Stat(nbPath); err != nil {
		fmt.Fprintf(os.Stderr, "cannot access notebook: %v\n", err)
		os.Exit(1)
	}

	tmpDir, err := os.MkdirTemp("", "nbcheck-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create temp dir: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tmpDir)

	outName := filepath.Base(nbPath)

	cmd := exec.Command(
		"uv", "run", "--env-file", ".env", "--",
		"jupyter", "nbconvert",
		"--to", "notebook",
		"--execute", nbPath,
		"--output", outName,
		"--output-dir", tmpDir,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "."

	fmt.Printf("Checking notebook execution: %s\n", nbPath)
	fmt.Printf("Temporary output dir: %s\n", tmpDir)

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\nnotebook execution failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("notebook execution succeeded")
}
