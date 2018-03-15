package cli

import cli "github.com/jawher/mow.cli"

// githubCommand is the executor for the top level github command.
func githubCommand(cmd *cli.Cmd) {
	cmd.Command("issues", "list github issues for a given project", notImplementedPlan)
	cmd.Action = func() { panic("not implemented") }
}
