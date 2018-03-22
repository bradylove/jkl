package cli

import (
	"path/filepath"

	"github.com/bradylove/jkl/pkg/manifest"
	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

// OpenCommand will open one or more projects in new tmux windows.
func OpenCommand(
	log Logger,
	tm tmux.Tmux,
	m manifest.Manifest,
) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		projects := cmd.StringsArg("PROJECTS", nil, "names or aliases of projects to open")

		cmd.Spec = "PROJECTS..."
		cmd.Action = func() {
			if !tm.Valid() {
				log.Fatalf("jkl open must be ran in tmux")
			}

			for _, name := range *projects {
				p, err := m.FindProject(name)
				if err != nil {
					log.Printf("%s", err)
					continue
				}

				var opts []tmux.CreateWindowOption
				if p.WorkingPath != "" {
					opts = append(opts, tmux.WithVerticalSplitPath(
						filepath.Join(p.Path, p.WorkingPath)),
					)
				}

				if p.Layout != "" {
					opts = append(opts, tmux.WithLayout(p.Layout))
				}

				err = tm.CreateWindow(p.Name, p.Path, opts...)
				if err != nil {
					log.Printf("failed to open project '%s': %s", p.Name, err)
				}

				err = tm.Execute((m.Editor + " ."))
				if err != nil {
					log.Printf("failed to open editor: %s", err)
				}
			}
		}
	}
}
