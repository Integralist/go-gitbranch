package delete

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/integralist/go-gitbranch/internal/pkg/git"
)

type Flags struct {
	Name string
}

// ParseFlags defines and parses flags for the create subcommand.
func ParseFlags(args []string) Flags {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	name := fs.String("name", "", "branch to delete")
	fs.Parse(args)

	return Flags{
		Name: *name,
	}
}

// Process executes the underlying git command.
func Process(flags Flags) {
	git.Validation()

	if flags.Name != "" {
		err := git.DeleteBranch(flags.Name)
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

	fmt.Printf("\nwhich branch would you like to delete? (type its number)\n\n")
	var selected string
	fmt.Scanln(&selected)

	for _, branch := range filtered {
		if strings.HasPrefix(branch, selected+".") {
			selected = strings.TrimPrefix(branch, selected+". ")
			break
		}
	}

	err = git.DeleteBranch(selected)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
