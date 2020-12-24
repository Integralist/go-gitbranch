// Package rename defines the CLI behaviours for renaming a branch.

package rename

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/integralist/go-gitbranch/internal/pkg/git"
)

type Flags struct {
	Branch    string
	NewName   string
	Normalize bool
	Prefix    bool
}

// ParseFlags defines and parses flags for the create subcommand.
func ParseFlags(args []string) Flags {
	fs := flag.NewFlagSet("rename", flag.ExitOnError)
	branch := fs.String("branch", "", "branch to rename")
	name := fs.String("name", "", "new branch name")
	normalize := fs.Bool("normalize", false, "whether to normalize the given branch name")
	prefix := fs.Bool("prefix", false, "whether to generate a unique prefix for the branch name")
	fs.Parse(args)

	return Flags{
		Branch:    *branch,
		NewName:   *name,
		Normalize: *normalize,
		Prefix:    *prefix,
	}
}

// Process executes the underlying git command.
func Process(flags Flags) {
	git.Validation()

	if flags.Branch != "" && flags.NewName != "" {
		err := git.RenameBranch(flags.Branch, flags.NewName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	branches, err := git.GetBranches()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filtered := git.FilterBranches(branches)

	fmt.Println() // I like breathing space in my terminal output
	for _, branch := range filtered {
		fmt.Println(branch)
	}

	fmt.Printf("\nwhich branch would you like to rename? (type its number)\n\n")
	var selected string
	fmt.Scanln(&selected)

	for _, branch := range filtered {
		if strings.HasPrefix(branch, selected+".") {
			selected = strings.TrimPrefix(branch, selected+". ")
			break
		}
	}

	fmt.Printf("\nwhat's the new branch name? (remember: --prefix and --normalize)\n\n")
	var newbranch string
	fmt.Scanln(&newbranch)

	unmodified := newbranch

	if flags.Prefix {
		newbranch = fmt.Sprintf("%s%s", git.BranchPrefix(), unmodified)
	}
	if flags.Normalize {
		newbranch = fmt.Sprintf("%s%s", git.BranchPrefix(), git.BranchNormalize(unmodified))
	}

	err = git.RenameBranch(selected, newbranch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
