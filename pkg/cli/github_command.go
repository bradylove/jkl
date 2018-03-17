package cli

import (
	"log"
	"os"

	cli "github.com/jawher/mow.cli"
)

// githubCommand is the executor for the top level github command.
func githubCommand(cmd *cli.Cmd) {
	log := log.New(os.Stderr, "", 0)

	cmd.Command("issues", "list github issues for a given project", notImplementedPlan)
	cmd.Action = func() { log.Fatal("not implemented") }
}
