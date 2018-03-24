package cli

import (
	"log"
	"os"
	"os/exec"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

// CloneCommand is used to clone a project repository to the projects path.
func CloneCommand(cr CommandRunner, m manifest.Manifest) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		project := cmd.StringArg("PROJECT", "", "name of project to clone")

		cmd.Action = func() {
			p, err := m.FindProject(*project)
			if err != nil {
				log.Fatalf("%s", err)
			}

			c := exec.Command("git", "clone", p.Repository, p.Path)
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			c.Stdin = os.Stdin

			err = cr.Run(c)
			if err != nil {
				log.Fatalf("failed to clone project: %s", err)
			}
		}
	}
}
