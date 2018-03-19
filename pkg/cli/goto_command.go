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

func gotoCommand(cmd *cli.Cmd) {
	log := log.New(os.Stderr, "", 0)
	project := cmd.StringArg("PROJECT", "", "names or aliases of project to open")

	cmd.Action = func() {
		tmuxVar := os.Getenv("TMUX")
		if tmuxVar == "" {
			log.Fatalln("jkl goto must be ran in TMUX")
		}

		m, err := manifest.Load(filepath.Join(os.Getenv("HOME"), ".jkl"))
		if err != nil {
			log.Fatalf("failed to read jkl manifest: %s", err)
		}

		tm := tmux.New(strings.Split(tmuxVar, ",")[0])
		p, err := findProject(*project, m.Projects)
		if err != nil {
			log.Fatalf("%s", err)
		}

		err = tm.ChangeDirectory(p.BasePath)
		if err != nil {
			log.Printf("failed to open project '%s': %s", p.Name, err)
		}
	}
}
