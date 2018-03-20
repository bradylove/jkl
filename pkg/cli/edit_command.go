package cli

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bradylove/jkl/pkg/manifest"
	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

func editCommand(cmd *cli.Cmd) {
	log := log.New(os.Stderr, "", 0)

	cmd.Action = func() {
		tmuxVar := os.Getenv("TMUX")
		if tmuxVar == "" {
			log.Fatalln("jkl edit must be ran in TMUX")
		}

		m, err := manifest.Load(filepath.Join(os.Getenv("HOME"), ".jkl"))
		if err != nil {
			log.Fatalf("failed to read jkl manifest: %s", err)
		}

		tm := tmux.New(strings.Split(tmuxVar, ",")[0])
		err = tm.Execute(fmt.Sprintf("%s %s", m.Editor, filepath.Join(os.Getenv("HOME"), ".jkl")))
		if err != nil {
			log.Printf("failed to open editor: %s", err)
		}
	}
}
