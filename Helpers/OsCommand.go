package Helpers

import (
	"os"
	"os/exec"
)

func OsCommand(commandName string, args []string, outputToStdOut bool) error {
	cmd := exec.Command(commandName, args...)

	if outputToStdOut {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
