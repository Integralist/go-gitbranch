package main

import (
	"fmt"
	"os"

	"github.com/integralist/go-gitbranch/internal/pkg/checkout"
	"github.com/integralist/go-gitbranch/internal/pkg/create"
	"github.com/integralist/go-gitbranch/internal/pkg/delete"
	"github.com/integralist/go-gitbranch/internal/pkg/rename"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("no subcommand provided")
		os.Exit(1)
	}

	// TODO:
	// add package descriptions
	//
	// NOTE:
	// look at fastly/cli app.Run for how to pass args for testing purposes.

	switch args[0] {
	case "create":
		flags := create.ParseFlags(args[1:])
		create.Process(flags)
	case "rename":
		flags := rename.ParseFlags(args[1:])
		rename.Process(flags)
	case "checkout":
		flags := checkout.ParseFlags(args[1:])
		checkout.Process(flags)
	case "delete":
		flags := delete.ParseFlags(args[1:])
		delete.Process(flags)
	}
}
