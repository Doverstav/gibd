package helpers

import (
	"os"
	"os/exec"
	"strings"
)

func GetBranchesWithRemoteStatus() ([]string, error) {
	gitCmd := exec.Command("git", "for-each-ref", "--format=%(refname) %(upstream:track)", "refs/heads")
	gitCmd.Env = append(os.Environ(), "LANG=en_US", "LC_ALL=en_US", "LANGUAGE=en_US") // Does nothing?

	out, err := gitCmd.CombinedOutput()

	if err != nil {
		return nil, err
	}

	outString := strings.TrimSpace(string(out))
	outArray := strings.Split(outString, "\n")

	return outArray, nil
}

// func DeleteBranch(branchName string) (error) {
// 	gitCmd := exec.Command("git", "branch", "-d", branchName)
// }
