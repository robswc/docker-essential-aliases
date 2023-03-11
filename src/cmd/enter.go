package cmd

import (
	"docker-essential-aliases/util"
	"github.com/spf13/cobra"
)

// enterCmd represents the enter command
var enterCmd = &cobra.Command{
	Use:   "enter",
	Short: "Enters a container.",
	Long: `Enters a docker container.  By default it will enter with a bash shell.  You can specify a different shell with the --shell flag.

	Example: dea enter <container name or id>
	Example: dea enter <container name or id> -s /bin/sh
`,
	Run: func(cmd *cobra.Command, args []string) {
		// get container name or id, if args are passed
		containerIdOrName := ""
		if len(args) > 0 {
			containerIdOrName = args[0]
		}

		container := util.GetContainer(containerIdOrName)

		// get all flags
		shell, _ := cmd.Flags().GetString("shell")
		verbose, _ := cmd.Flags().GetBool("verbose")

		baseCmd := []string{"docker", "exec", "-it", container.Accessor, shell}

		// store the container id, before running the command
		util.CaptureContainer(container.Accessor)
		util.RunCommand(baseCmd, verbose)

	},
}

func init() {
	rootCmd.AddCommand(enterCmd)
	enterCmd.Flags().StringP("shell", "s", "/bin/bash", "Specify a shell to enter the container with")
	enterCmd.Flags().BoolP("verbose", "v", false, "Show verbose output, including the original command")
}
