package helpers

import (
	"strings"
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
