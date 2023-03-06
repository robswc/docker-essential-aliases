package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Lists all running containers.",
	Long:  `Lists all running containers, in a similar way to the docker ps command`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flags
		all, _ := cmd.Flags().GetBool("all")
		verbose, _ := cmd.Flags().GetBool("verbose")

		// create an array string to hold the arguments
		baseCmd := []string{"docker", "ps"}

		// add the arguments
		if all {
			baseCmd = append(baseCmd, "-a")
		}

		if verbose {
			fmt.Println("Docker Command:")
			fmt.Println(baseCmd)
		}

		// create the command
		dockerCmd := exec.Command(baseCmd[0], baseCmd[1:]...)
		dockerCmd.Stdout = cmd.OutOrStdout()
		dockerCmd.Stderr = cmd.OutOrStderr()
		err := dockerCmd.Run()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Docker Command Executed Successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	psCmd.Flags().BoolP("verbose", "v", false, "Show verbose output, including the original command")
}
