package cli

import (
	"os/exec"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

// BrowserCommand is the executor for the top level github command.
func BrowserCommand(
	log Logger,
	cr CommandRunner,
	m manifest.Manifest,
	runtimeOS string,
) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		project := cmd.StringArg("PROJECT", "", "name or alias of project to open in browser")

		cmd.Action = func() {
			p, err := m.FindProject(*project)
			if err != nil {
				log.Fatalf("%s", err)
			}

			u, err := p.BrowserURL()
			if err != nil {
				log.Fatalf("failed to build browser URL: %s", err)
			}

			bin := "xdg-open"
			if runtimeOS == "darwin" {
				bin = "open"
			}

			c := exec.Command(bin, u)
			err = cr.Run(c)
			if err != nil {
				log.Fatalf("failed to open project in browser: %s", err)
			}
		}
	}
}
