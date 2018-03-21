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
	m manifest.Manifest,
) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			if !tm.Valid() {
				log.Fatalf("jkl edit must be ran in tmux")
			}

			err := tm.Execute(fmt.Sprintf("%s %s", m.Editor, m.Path))
			if err != nil {
				log.Fatalf("failed to open editor: %s", err)
			}
		}
	}
}
