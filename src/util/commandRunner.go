package util

import (
	"os"
	"os/exec"
)

func RunCommand(baseCmd []string, verbose bool) {

	// if verbose, print the command
	HandleVerbose(verbose, baseCmd)

	dockerCmd := exec.Command(baseCmd[0], baseCmd[1:]...)
	dockerCmd.Stdin = os.Stdin
	dockerCmd.Stdout = os.Stdout
	dockerCmd.Stderr = os.Stderr

	err := dockerCmd.Run()
	if err != nil {
		panic(err)
	}

}
