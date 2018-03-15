package cli

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bradylove/jkl/pkg/manifest"
	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

func openCommand(cmd *cli.Cmd) {
	log := log.New(os.Stderr, "", 0)
	projects := cmd.StringsArg("PROJECTS", nil, "names or aliases of projects to open")

	cmd.Spec = "PROJECTS..."

	cmd.Action = func() {
		tmuxVar := os.Getenv("TMUX")
		if tmuxVar == "" {
			log.Fatalln("jkl must be ran in TMUX")
		}

		m, err := manifest.Load(filepath.Join(os.Getenv("HOME"), ".jkl"))
		if err != nil {
			log.Fatalf("failed to read jkl manifest: %s", err)
		}

		tm := tmux.New(strings.Split(tmuxVar, ",")[0])
		for _, name := range *projects {
			p, err := findProject(name, m.Projects)
			if err != nil {
				log.Println(err)
				continue
			}

			var opts []tmux.CreateWindowOption
			if p.WorkingPath != "" {
				opts = append(opts, tmux.WithVerticalSplitPath(p.WorkingPath))
			}

			if p.Layout != "" {
				opts = append(opts, tmux.WithLayout(p.Layout))
			}

			err = tm.CreateWindow(p.Name, p.BasePath, opts...)
			if err != nil {
				log.Printf("failed to open project '%s': %s", p.Name, err)
			}
		}
	}
}
