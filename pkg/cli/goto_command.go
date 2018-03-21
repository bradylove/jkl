package cli

import (
	"github.com/bradylove/jkl/pkg/manifest"
	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

// GoToCommand will cd in to the projects path in the current tmux pane.
func GoToCommand(log Logger, tm tmux.Tmux, m manifest.Manifest) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		project := cmd.StringArg("PROJECT", "", "names or aliases of project to open")

		cmd.Action = func() {
			if !tm.Valid() {
				log.Fatalf("jkl goto must be ran in tmux")
			}

			p, err := m.FindProject(*project)
			if err != nil {
				log.Fatalf("%s", err)
			}

			err = tm.ChangeDirectory(p.Path)
			if err != nil {
				log.Fatalf("failed to open project '%s': %s", p.Name, err)
			}
		}
	}
}
