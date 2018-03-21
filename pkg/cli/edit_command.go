package cli

import (
	"fmt"

	"github.com/bradylove/jkl/pkg/manifest"
	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

// EditCommand launches the jkl manifest in the users editor for editing.
func EditCommand(
	log Logger,
	tm tmux.Tmux,
	manifestPath string,
) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			if !tm.Valid() {
				log.Fatalf("jkl edit must be ran in tmux")
			}

			m, err := manifest.Load(manifestPath)
			if err != nil {
				log.Fatalf("failed to read jkl manifest: %s", err)
			}

			err = tm.Execute(fmt.Sprintf("%s %s", m.Editor, manifestPath))
			if err != nil {
				log.Fatalf("failed to open editor: %s", err)
			}
		}
	}
}
