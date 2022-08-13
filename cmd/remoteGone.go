/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remoteGone called")

		branches, _ := helpers.GetBranchesWithRemoteStatus()
		branchNamesWithRemoteGone := helpers.GetBranchNamesWithRemoteGone(branches)

		fmt.Println(branchNamesWithRemoteGone)

		// Ask what branches to delete
		branchesToDelete := []string{}
		survey.AskOne(&survey.MultiSelect{
			Message: "Select multiple",
			Options: branchNamesWithRemoteGone,
		}, &branchesToDelete)

		helpers.DeleteBranches(branchesToDelete)
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
}
