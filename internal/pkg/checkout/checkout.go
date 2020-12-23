package checkout

import (
	"flag"
	"fmt"

	"github.com/integralist/go-gitbranch/internal/pkg/git"
)

type Flags struct {
	Name string
}

// ParseFlags defines and parses flags for the create subcommand.
func ParseFlags(args []string) Flags {
	fs := flag.NewFlagSet("create", flag.ExitOnError)
	name := fs.String("name", "", "name of the branch to create")
	fs.Parse(args)

	return Flags{
		Name: *name,
	}
}

// Process executes the underlying git command.
func Process(flags Flags) {
	git.Validation()
	fmt.Println("name:", flags.Name)

	// TODO:
	//
	// 1. shortcircuit if name flag provided
	// 2. execute 'git checkout <name>'
	// 3. otherwise print all branches except master/main (prefix each with incrementing number)
	// 4. read user input for selected branch
	// 5. execute 'git checkout -b <branch>'
	//
	// https://gobyexample.com/spawning-processes
}
