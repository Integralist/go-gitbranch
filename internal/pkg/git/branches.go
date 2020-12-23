package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const ErrGitBranch = "error executing 'git branch': %w\n"
const NotGitDir = "you're not inside of a git directory."

// IsGitDir indicates whether the user is inside of a git directory.
//
// We  do not want to execute any git commands if outside of a git directory
// as those operations will otherwise always fail.
func IsGitDir() (bool, error) {
	var buf bytes.Buffer

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		return false, fmt.Errorf(buf.String())
	}

	return true, nil
}

// Validation will exit the command process if the user is not inside of a git
// directory, otherwise it'll succeed quietly and allow the caller to continue.
func Validation() {
	ok, err := IsGitDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if !ok {
		fmt.Println(NotGitDir)
		os.Exit(1)
	}
}

func GetBranches() (bytes.Buffer, error) {
	var err bytes.Buffer
	var out bytes.Buffer

	cmd := exec.Command("git", "branch")
	cmd.Stdout = &out
	cmd.Stderr = &err

	runErr := cmd.Run()
	if runErr != nil {
		return err, fmt.Errorf(err.String())
	}

	return out, nil
}

// FilterBranches returns all git branches (except master/main) and numerically
// prefixes them.
func FilterBranches(branches bytes.Buffer) []string {
	index := 0
	filtered := []string{}

	bs := strings.Split(branches.String(), "\n")
	for _, branch := range bs {
		branch = strings.TrimPrefix(branch, "*")
		branch = strings.TrimSpace(branch)
		if branch != "master" && branch != "main" && branch != "" {
			index++
			filtered = append(filtered, fmt.Sprintf("%d. %s", index, branch))
		}
	}

	return filtered
}

// CreateBranch creates a new git branch.
func CreateBranch(branch string) error {
	var buf bytes.Buffer

	cmd := exec.Command("git", "checkout", "-b", branch)
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(buf.String())
	}

	return nil
}

// DeleteBranch deletes a specified git branch.
func DeleteBranch(branch string) error {
	var buf bytes.Buffer

	cmd := exec.Command("git", "branch", "-D", branch)
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(buf.String())
	}

	return nil
}
