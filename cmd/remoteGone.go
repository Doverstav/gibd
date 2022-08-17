/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/doverstav/gibd/helpers"
	"github.com/spf13/cobra"
)

// remoteGoneCmd represents the remoteGone command
var remoteGoneCmd = &cobra.Command{
	Use:   "remote-gone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("remoteGone called")

		remoteName := cmd.Flag("remote").Value.String()                // Default origin
		defaultBranchName := cmd.Flag("default-branch").Value.String() // Default master
		defaultBranchNameSet := cmd.Flag("default-branch").Changed
		includeDefault, _ := strconv.ParseBool(cmd.Flag("include-default").Value.String())
		forceDelete, _ := strconv.ParseBool(cmd.Flag("force").Value.String())

		// Get all branch refs
		branches, err := helpers.GetBranchesWithRemoteStatus()
		if err != nil {
			return err
		}

		// Remove default branch
		defaultRef := defaultBranchName
		if !defaultBranchNameSet {
			defaultRef, err = helpers.GetDefaultBranchRef(remoteName)
			if err != nil {
				defaultRef = defaultBranchName
			}
		}
		testFunc := func(s string) bool {
			bSplit := strings.Split(s, "/")
			bName := strings.TrimSpace(bSplit[len(bSplit)-1])
			dSplit := strings.Split(defaultRef, "/")
			dName := strings.TrimSpace(dSplit[len(dSplit)-1])

			return bName != dName
		}
		if !includeDefault {
			branches = helpers.Filter(branches, testFunc)
		}

		branchNamesWithRemoteGone := helpers.GetBranchNamesWithRemoteGone(branches)

		fmt.Println(branchNamesWithRemoteGone)

		// Ask what branches to delete
		branchesToDelete := []string{}
		survey.AskOne(&survey.MultiSelect{
			Message: "Select multiple",
			Options: branchNamesWithRemoteGone,
		}, &branchesToDelete)

		if forceDelete {
			helpers.ForceDeleteBranches(branchesToDelete)
		} else {
			helpers.DeleteBranches(branchesToDelete)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(remoteGoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteGoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteGoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	remoteGoneCmd.Flags().BoolP("prune", "p", false, "Prune remote before running command")
}
