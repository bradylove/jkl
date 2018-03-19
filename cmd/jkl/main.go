package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/bradylove/jkl/pkg/cli"
)

func main() {
	log := log.New(os.Stderr, "", 0)

	cli.Run(
		log,
		CommandRunner{},
		filepath.Join(os.Getenv("HOME"), ".jkl"),
		os.Args,
	)
}

// CommandRunner is used to run exec.Cmd.
type CommandRunner struct{}

// Run executes the given command .
func (r CommandRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
