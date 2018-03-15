package cli

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bradylove/jkl/pkg/manifest"
	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

func Run(args []string) {
	log := log.New(os.Stderr, "", 0)

	app := cli.App("jkl", "project management life improver")
	app.Command("open", "opens one or more projects", func(cmd *cli.Cmd) {
		projects := cmd.StringsArg("PROJECTS", nil, "names of projects to open")

		cmd.Spec = "PROJECTS..."

		cmd.Action = func() {
			tmuxVar := os.Getenv("TMUX")
			if tmuxVar == "" {
				log.Fatalln("jkl must be ran in TMUX")
			}

			m, err := manifest.Load("/home/brady/.jkl.yml")
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

				err = tm.CreateWindow(p.Name, p.BasePath)
				if err != nil {
					log.Printf("failed to open project '%s': %s", p.Name, err)
				}
			}
		}
	})

	app.Run(args)
}

func findProject(name string, projects []manifest.Project) (manifest.Project, error) {
	for _, p := range projects {
		if p.Name == name || p.Alias == name {
			return p, nil
		}
	}

	return manifest.Project{}, fmt.Errorf("project named %s not found in manifest", name)
}
