package cli

import (
	"os/exec"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

// BrowserCommand is the executor for the top level github command.
func BrowserCommand(log Logger, cr CommandRunner, manifestPath string) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Command("issues", "list github issues for a given project", notImplementedPlan)

		project := cmd.StringArg("PROJECT", "", "name or alias of project to open in browser")

		cmd.Action = func() {
			m, err := manifest.Load(manifestPath)
			if err != nil {
				log.Fatalf("failed to read jkl manifest: %s", err)
			}

			p, err := findProject(*project, m.Projects)
			if err != nil {
				log.Fatalf("%s", err)
			}

			u, err := p.BrowserURL()
			if err != nil {
				log.Fatalf("failed to build browser URL: %s", err)
			}

			c := exec.Command("xdg-open", u)
			err = cr.Run(c)
			if err != nil {
				log.Fatalf("failed to open project in browser: %s", err)
			}
		}
	}
}
