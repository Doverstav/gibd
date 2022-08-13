/*
Copyright © 2022 Pontus Doverstav <doverstav@gmail.com>

*/
package cmd

import (
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/doverstav/gibd/helpers"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gibd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get all branch refs
		branches, err := helpers.GetBranchesWithRemoteStatus()
		if err != nil {
			return err
		}

		// Remove default branch
		defaultRef, err := helpers.GetDefaultBranchRef()
		if err != nil {
			// TODO Support user supplied default branch & remote
			defaultRef = "refs/remotes/origin/master"
		}
		testFunc := func(s string) bool {
			bSplit := strings.Split(s, "/")
			bName := strings.TrimSpace(bSplit[len(bSplit)-1])
			dSplit := strings.Split(defaultRef, "/")
			dName := strings.TrimSpace(dSplit[len(dSplit)-1])

			return bName != dName
		}
		branches = helpers.Filter(branches, testFunc)

		// Extract all names
		branchNames := helpers.GetBranchNames(branches)

		// Ask what branches to delete
		branchesToDelete := []string{}
		survey.AskOne(&survey.MultiSelect{
			Message: "Select multiple",
			Options: branchNames,
		}, &branchesToDelete)

		helpers.DeleteBranches(branchesToDelete)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gibd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
