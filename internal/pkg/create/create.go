package create

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

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

	prefix := "integralist"
	date := time.Now().Format("20060102")
	normalize := strings.ReplaceAll(flags.Name, "-", "_")
	branch := fmt.Sprintf("%s/%s_%s", prefix, date, normalize)

	err := git.CreateBranch(branch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
