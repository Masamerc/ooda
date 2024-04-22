package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/masamerc/ooda/src/utils"
)

var rootCmd = &cobra.Command{
	Use:   "ooda",
	Short: "ooda - switch between AWS profiles.",
	Long:  "ooda - switch between AWS profiles.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runProfileSwitcher(); err != nil {
			log.Fatal(err)
		}
	},
}

// Entry point for the CLI tool
func Execute() {
	if shouldRunDirectProfileSwitch() {
		profile := os.Args[1]
		if err := directProfileSwitch(profile); err != nil {
			log.Fatal(err)
		}
		return
	}
	runRootCmd()
}

func runRootCmd() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runProfileSwitcher() error {
	profiles := utils.GetProfiles()
	fmt.Printf(utils.NoticeColor, "AWS Profile Switcher\n")
	profile, err := getProfileFromPrompt(profiles)
	if err != nil {
		return err
	}
	fmt.Printf(utils.PromptColor, "Choose a profile")
	fmt.Printf(utils.NoticeColor, "? ")
	fmt.Printf(utils.CyanColor, profile)
	fmt.Println()
	return utils.WriteFile(profile, utils.GetHomeDir())
}

func shouldRunDirectProfileSwitch() bool {
	invalidProfiles := []string{"l", "list", "completion", "help", "--help", "-h", "v", "version"}
	return len(os.Args) > 1 && !utils.Contains(invalidProfiles, os.Args[1])
}

func directProfileSwitch(desiredProfile string) error {
	profiles := utils.GetProfiles()
	if utils.Contains(profiles, desiredProfile) {
		printColoredMessage("Profile ", utils.PromptColor)
		printColoredMessage(desiredProfile, utils.CyanColor)
		printColoredMessage(" set.\n", utils.PromptColor)
		return utils.WriteFile(desiredProfile, utils.GetHomeDir())
	}
	printColoredMessage("WARNING: Profile ", utils.NoticeColor)
	printColoredMessage(desiredProfile, utils.CyanColor)
	printColoredMessage(" does not exist.\n", utils.PromptColor)
	return nil
}

func getProfileFromPrompt(profiles []string) (string, error) {
	prompt := promptui.Select{
		Label:        fmt.Sprintf(utils.PromptColor, "Choose a profile"),
		Items:        profiles,
		HideHelp:     true,
		HideSelected: true,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   fmt.Sprintf("%s {{ . | cyan }}", promptui.IconSelect),
			Inactive: "  {{.}}",
			Selected: "  {{ . | cyan }}",
		},
		Searcher:          utils.NewPromptUISearcher(profiles),
		StartInSearchMode: true,
		Stdout:            &utils.BellSkipper{},
	}

	_, result, err := prompt.Run()
	if err != nil {
		utils.CheckError(err)
		return "", nil
	}
	return result, nil
}

func printColoredMessage(msg, color string) {
	fmt.Printf(color, msg)
}
