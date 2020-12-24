// Package checkout defines the CLI behaviours for changing branches.

package checkout

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/integralist/go-gitbranch/internal/pkg/git"
)

type Flags struct {
	Branch string
}

// ParseFlags defines and parses flags for the create subcommand.
func ParseFlags(args []string) Flags {
	fs := flag.NewFlagSet("checkout", flag.ExitOnError)
	branch := fs.String("branch", "", "branch to checkout")
	fs.Parse(args)

	return Flags{
		Branch: *branch,
	}
}

// Process executes the underlying git command.
func Process(flags Flags) {
	git.Validation()

	if flags.Branch != "" {
		err := git.CheckoutBranch(flags.Branch)
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

	fmt.Printf("\nwhich branch would you like to checkout? (type its number)\n\n")
	var selected string
	fmt.Scanln(&selected)

	for _, branch := range filtered {
		if strings.HasPrefix(branch, selected+".") {
			selected = strings.TrimPrefix(branch, selected+". ")
			break
		}
	}

	err = git.CheckoutBranch(selected)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
