package cmd

import (
	"docker-essential-aliases/util"
	"github.com/spf13/cobra"
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

		// run the command
		util.RunCommand(cmd, baseCmd, verbose)
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	psCmd.Flags().BoolP("verbose", "v", false, "Show verbose output, including the original command")
}
