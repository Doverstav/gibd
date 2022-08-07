package helpers

import (
	"fmt"
	"strings"
)

func GetBranchNames(branches []string) []string {
	branchNames := []string{}

	for _, branchString := range branches {
		splitString := strings.Split(branchString, " ")
		fmt.Println(splitString)
		branchRef := splitString[0]

		branchName := strings.TrimPrefix(branchRef, "refs/heads/")
		branchNames = append(branchNames, branchName)
	}

	return branchNames
}
