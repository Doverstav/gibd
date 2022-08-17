package helpers

import (
	"fmt"
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

func PruneRemote(remoteName string) (string, error) {
	gitCmd := exec.Command("git", "remote", "prune", remoteName)

	out, err := gitCmd.CombinedOutput()

	return string(out), err
}

func GetDefaultBranchRef(remoteName string) (string, error) {
	remoteHEADRef := fmt.Sprintf("refs/remotes/%s/HEAD", remoteName)
	gitCmd := exec.Command("git", "symbolic-ref", remoteHEADRef)

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
