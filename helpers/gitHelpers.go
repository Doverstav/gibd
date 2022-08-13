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

func GetDefaultBranchRef() (string, error) {
	// TODO Allow user to specify remote (is hardcoded to origin now)
	gitCmd := exec.Command("git", "symbolic-ref", "refs/remotes/origin/HEAD")

	out, err := gitCmd.CombinedOutput()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func DeleteBranch(branchName string) (string, error) {
	gitCmd := exec.Command("git", "branch", "-d", branchName)

	out, err := gitCmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}

	return string(out), nil
}

func ForceDeleteBranch(branchName string) (string, error) {
	gitCmd := exec.Command("git", "branch", "-D", branchName)

	out, err := gitCmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}

	return string(out), nil
}
