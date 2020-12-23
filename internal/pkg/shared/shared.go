package shared

import (
	"fmt"
	"os"
	"os/exec"
)

const NotGitDir = "you're not inside of a git directory."
const ErrGitDir = "error executing 'git rev-parse --show-toplevel' to determine if you are in a git directory: %s\n"

// isGitDir indicates whether the user is inside of a git directory.
//
// We  do not want to execute any git commands if outside of a git directory
// as those operations will otherwise always fail.
func isGitDir() (bool, error) {
	_, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return false, err
	}

	return true, nil
}

// Validation will exit the command process if the user is not inside of a git
// directory, otherwise it'll succeed quietly and allow the caller to continue.
func Validation() {
	ok, err := isGitDir()
	if err != nil {
		fmt.Printf(ErrGitDir, err)
		os.Exit(1)
	}

	if !ok {
		fmt.Println(NotGitDir)
		os.Exit(1)
	}
}
