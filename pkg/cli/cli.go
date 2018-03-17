package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

// Run initializes and executes the CLI.
func Run(args []string) {
	app := cli.App("jkl", "project management life improver")

	commands := []command{
		{
			name:        "edit",
			description: "opens the jkl manifest for editing",
			cmd:         editCommand,
		},
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
	log := log.New(os.Stderr, "", 0)

	cmd.Action = func() { log.Fatal("not implemented") }
}

func findProject(name string, projects []manifest.Project) (manifest.Project, error) {
	for _, p := range projects {
		if p.Name == name || p.Alias == name {
			return p, nil
		}
	}

	return manifest.Project{}, fmt.Errorf("project named %s not found in manifest", name)
}
