package util

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func RunCommand(cmd *cobra.Command, baseCmd []string, verbose bool) {

	// if verbose, print the command
	HandleVerbose(verbose, baseCmd)

	dockerCmd := exec.Command(baseCmd[0], baseCmd[1:]...)
	dockerCmd.Stdout = cmd.OutOrStdout()
	dockerCmd.Stderr = cmd.OutOrStderr()
	err := dockerCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
