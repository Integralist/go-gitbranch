package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

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
		fmt.Println("you're not inside of a git directory.")
		os.Exit(1)
	}
}

// BranchPrefix generates a custom string to use as a prefix to a branch.
func BranchPrefix() string {
	prefix := os.Getenv("GITBRANCH_PREFIX")

	if prefix == "" {
		prefix = "integralist"
	}

	date := time.Now().Format("20060102")
	return fmt.Sprintf("%s/%s_", prefix, date)
}

// BranchNormalize replaces hyphens with underscores.
func BranchNormalize(branch string) string {
	return strings.ReplaceAll(branch, "-", "_")
}

// GetBranches retrieves all git branches.
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

// RenameBranch changes the name of a specified git branch.
func RenameBranch(o, n string) error {
	var buf bytes.Buffer

	cmd := exec.Command("git", "branch", "-m", o, n)
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(buf.String())
	}

	return nil
}

// CheckoutBranch checks out the specified git branch.
func CheckoutBranch(branch string) error {
	var buf bytes.Buffer

	cmd := exec.Command("git", "checkout", branch)
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf(buf.String())
	}

	return nil
}
