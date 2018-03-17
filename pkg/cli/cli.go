package cli

import (
	"fmt"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

// Run initializes and executes the CLI.
func Run(args []string) {
	app := cli.App("jkl", "project management life improver")

	commands := []command{
		{
			name:        "goto",
			description: "change the current directory of the current tmux pane to the project directory",
			cmd:         gotoCommand,
		},
		{
			name:        "open",
			description: "(limited) opens one or more projects",
			cmd:         openCommand,
		},
		{
			name:        "edit",
			description: "(not implemented) opens the jkl manifest for editing",
			cmd:         notImplementedPlan,
		},
		{
			name:        "github",
			description: "(not implemented) open the projects github page in the browser",
			cmd:         githubCommand,
		},
		{
			name:        "init",
			description: "(not implemented) initializes the jkl config file",
			cmd:         notImplementedPlan,
		},
	}

	for _, cmd := range commands {
		app.Command(cmd.name, cmd.description, cmd.cmd)
	}

	app.Run(args)
}

type command struct {
	name        string
	description string
	cmd         func(cmd *cli.Cmd)
}

func notImplementedPlan(cmd *cli.Cmd) {
	cmd.Action = func() { panic("not implemented") }
}

func findProject(name string, projects []manifest.Project) (manifest.Project, error) {
	for _, p := range projects {
		if p.Name == name || p.Alias == name {
			return p, nil
		}
	}

	return manifest.Project{}, fmt.Errorf("project named %s not found in manifest", name)
}
