// Package create defines the CLI behaviours for creating a branch.

package create

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/go-gitbranch/internal/pkg/git"
)

type Flags struct {
	Branch string
}

// ParseFlags defines and parses flags for the create subcommand.
func ParseFlags(args []string) Flags {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	branch := fs.String("branch", "", "branch to create")
	fs.Parse(args)

	return Flags{
		Branch: *branch,
	}
}

// Process executes the underlying git command.
func Process(flags Flags) {
	git.Validation()

	branch := fmt.Sprintf("%s%s", git.BranchPrefix(), git.BranchNormalize(flags.Branch))

	err := git.CreateBranch(branch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
