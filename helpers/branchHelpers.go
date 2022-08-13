package helpers

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mgutz/ansi"
)

func GetBranchNames(branches []string) []string {
	branchNames := []string{}

	for _, branchString := range branches {
		splitString := strings.Split(strings.TrimSpace(branchString), " ")

		branchRef := splitString[0]

		branchName := strings.TrimPrefix(strings.TrimSpace(branchRef), "refs/heads/")
		branchNames = append(branchNames, branchName)
	}

	return branchNames
}

func GetBranchNamesWithRemoteGone(branches []string) []string {
	const DESIRED_REMOTE_STATUS = "[gone]"
	branchNames := []string{}

	for _, branchString := range branches {
		splitString := strings.Split(strings.TrimSpace(branchString), " ")
		if len(splitString) == 2 {

			branchRef := splitString[0]
			remoteStatus := splitString[1]

			if remoteStatus == DESIRED_REMOTE_STATUS {
				branchName := strings.TrimPrefix(strings.TrimSpace(branchRef), "refs/heads/")
				branchNames = append(branchNames, branchName)
			}
		}
	}

	return branchNames
}

func DeleteBranches(branchesToDelete []string) {
	// For each branch
	for _, branch := range branchesToDelete {
		fmt.Printf("Deleting branch %s\n", branch)
		// Try to delete it
		output, err := DeleteBranch(branch)
		if err != nil {
			// If that fails, display error output
			header := ansi.Color(fmt.Sprintf("Got this error when attempting to delete branch %s:", branch), "red")
			message := ansi.Color(strings.TrimSpace(output), "yellow")
			fmt.Printf("\n%s\n%s\n\n", header, message)

			// Ask if user wants to attempt a force delete
			tryForce := false
			survey.AskOne(&survey.Confirm{
				Message: "Do you wish to try a force delete?",
			}, &tryForce)

			// If yes
			if tryForce {
				// Try to force delete
				output, err = ForceDeleteBranch(branch)
				if err != nil {
					header := ansi.Color(fmt.Sprintf("Failed to delete branch %s with error:", branch), "red")
					message := ansi.Color(strings.TrimSpace(output), "yellow")
					fmt.Printf("\n%s\n%s\n\n", header, message)
				}
			}
		}
	}
}
