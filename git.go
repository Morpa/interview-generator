package main

import (
	"os"
	"os/exec"
)

func cloneRepo(url, folder string) error {
	if _, err := os.Stat(folder); err == nil {
		return nil
	}

	cmd := exec.Command("git", "clone", url, folder)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
