package create

import (
	"flag"
	"fmt"

	"github.com/integralist/go-gitbranch/internal/pkg/shared"
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
	shared.Validation()
	fmt.Println("name:", flags.Name)

	// TODO:
	//
	// 1. normalize branch name by replacing hyphen with underscores
	// 2. prefix branch name with 'integralist/'
	// 3. suffix branch name with '<yyyy_mm_dd> + <normalized input>'
	// 4. execute 'git checkout -b <branch>'
	//
	// https://gobyexample.com/spawning-processes
}
