package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

func Run(args []string) {
	log := log.New(os.Stderr, "", 0)

	app := cli.App("jkl", "project management life improver")
	app.Command("open", "opens one or more projects", func(cmd *cli.Cmd) {
		projects := cmd.StringsArg("PROJECTS", nil, "names of projects to open")

		cmd.Spec = "PROJECTS..."

		cmd.Action = func() {
			tmux := os.Getenv("TMUX")
			if tmux == "" {
				log.Fatalln("jkl must be ran in TMUX")
			}

			m, err := manifest.Load("/home/brady/.jkl.yml")
			if err != nil {
				log.Fatalf("failed to read jkl manifest: %s", err)
			}

			socket := strings.Split(tmux, ",")[0]
			_, _ = m, socket

			for _, name := range *projects {
				p, err := findProject(name, m.Projects)
				if err != nil {
					log.Println(err)
					continue
				}

				launchProject(socket, p)
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

func launchProject(socket string, p manifest.Project) {
	cmd := exec.Command(os.Getenv("SHELL"), "-c",
		fmt.Sprintf("tmux -S %s new-window -c %s -n %s ; split-window -h -c '#{pane_current_path}'",
			socket, p.BasePath, p.Name,
		),
	)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to execute tmux command: %s", err)
	}
}
