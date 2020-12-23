package rename

import (
	"flag"
	"fmt"
	"os"

	"github.com/integralist/go-gitbranch/internal/pkg/git"
)

type Flags struct {
	Name   string
	Prefix bool
}

// ParseFlags defines and parses flags for the create subcommand.
func ParseFlags(args []string) Flags {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	name := fs.String("name", "", "name of the branch to create")
	prefix := fs.Bool("prefix", false, "whether to prefix integralist/ to the branch name")
	fs.Parse(args)

	// TODO: prefix should come from an environment variable rather than be
	// hardcoded to my own username (for open-source reusability)

	return Flags{
		Name:   *name,
		Prefix: *prefix,
	}
}

// Process executes the underlying git command.
func Process(flags Flags) {
	git.Validation()
	fmt.Println("name:", flags.Name)
	fmt.Println("prefix:", flags.Prefix)

	branches, err := git.GetBranches()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filtered := git.FilterBranches(branches)
	fmt.Print(filtered)

	// TODO:
	//
	// 1. shortcircuit if name flag provided
	// 2. execute 'git branch -m <old> <new>'
	// 3. otherwise print all branches except master/main (prefix each with incrementing number)
	// 4. read user input for selected branch
	// 5. read user input for new name
	// 6. execute 'git branch -m <old> <new>' (add prefix if necessary)
	//
	// https://gobyexample.com/spawning-processes
}
