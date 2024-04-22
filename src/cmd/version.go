package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "v0.0.9"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "ooda version command",
	Aliases: []string{"v"},
	Long:    "Returns the current version of ooda",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ooda version:", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
