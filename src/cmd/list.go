package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/masamerc/ooda/src/utils"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List AWS profiles command.",
	Aliases: []string{"l"},
	Long:    "This lists all your AWS profiles.",
	Run: func(cmd *cobra.Command, args []string) {
		err := runProfileLister()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runProfileLister() error {
	profiles := utils.GetProfiles()
	for _, p := range profiles {
		fmt.Println(p)
	}
	return nil
}
