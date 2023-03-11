/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"docker-essential-aliases/util"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans up temp files associated with docker-essential-aliases",
	Long: `Cleans up tiles in /tmp/'

	Example: dea clean

	Removes the following files:
		- /tmp/dea-temp.yml
	
`,
	Run: func(cmd *cobra.Command, args []string) {
		util.DeleteTempFile()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
