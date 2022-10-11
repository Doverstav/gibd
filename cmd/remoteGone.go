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
	Short: "Delete branches where the remote is gone",
	Long: `Delete branches where the remote is gone.

You might want to prune your remote branches first 
(or use the -p flag) to ensure that upstream status 
are up to date.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get all flag values that we want
		remoteName := cmd.Flag("remote").Value.String()
		defaultBranchName := cmd.Flag("default-branch").Value.String()
		defaultBranchNameSet := cmd.Flag("default-branch").Changed
		includeDefault, _ := strconv.ParseBool(cmd.Flag("include-default").Value.String())
		forceDelete, _ := strconv.ParseBool(cmd.Flag("force").Value.String())
		prune, _ := strconv.ParseBool(cmd.Flag("prune").Value.String())

		if prune {
			output, err := helpers.PruneRemote(remoteName)
			if err != nil {
				fmt.Printf("Could not prune remote \"%s\". Error given was:\n%v", remoteName, output)
			}
		}

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

			return bName != dName && !strings.Contains(bName, dName)
		}
		if !includeDefault {
			branches = helpers.Filter(branches, testFunc)
		}

		branchNamesWithRemoteGone := helpers.GetBranchNamesWithRemoteGone(branches)

		if len(branchNamesWithRemoteGone) == 0 {
			fmt.Println("No branches to delete")
		}

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
