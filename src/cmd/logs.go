package cmd

import (
	"docker-essential-aliases/util"
	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Shows the logs for a container. (default is to follow logs)",
	Long: `Shows the logs for a container. (default is to follow logs)
	Example: dea logs <container name or id> # (Ctrl + C to stop following logs)
	Example: dea logs -nf <container name or id> # do not follow logs
`,
	Run: func(cmd *cobra.Command, args []string) {

		// get container name or id, if args are passed
		containerIdOrName := ""
		if len(args) > 0 {
			containerIdOrName = args[0]
		}
		container := util.GetContainer(containerIdOrName)

		// get all flags
		noFollow, _ := cmd.Flags().GetBool("no-follow")
		verbose, _ := cmd.Flags().GetBool("verbose")

		// create an array string to hold the arguments
		baseCmd := []string{"docker", "logs"}

		// add the container name or id
		baseCmd = append(baseCmd, container.Accessor)

		// add the arguments
		if !noFollow {
			baseCmd = append(baseCmd, "-f")
		}

		// get the container id, before running the command
		util.CaptureContainer(container.Accessor)

		// run the command
		util.RunCommand(baseCmd, verbose)

	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().BoolP("no-follow", "f", false, "Do not follow the logs")
	logsCmd.Flags().BoolP("verbose", "v", false, "Show verbose output, including the original command")
}
