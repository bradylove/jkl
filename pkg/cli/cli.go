package cli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

// GOOS is a variable that defaults to the runtime.GOOS constant, however it can
// be replaced at runtime.
var GOOS = runtime.GOOS

// Run initializes and executes the CLI.
func Run(
	log Logger,
	cr CommandRunner,
	manifest string,
	args []string,
	opts ...RunOption,
) {
	cfg := runConfig{
		runtimeOS: runtime.GOOS,
	}

	for _, o := range opts {
		o(&cfg)
	}

	app := cli.App("jkl", "developer project management life improver")

	commands := []command{
		{
			name:        "browser",
			description: "open the projects page in the browser",
			cmd:         BrowserCommand(log, cr, manifest, cfg.runtimeOS),
		},
		{
			name:        "edit",
			description: "opens the jkl manifest for editing",
			cmd:         editCommand,
		},
		{
			name:        "goto",
			description: "change the current directory of the current tmux pane to the project directory",
			cmd:         gotoCommand,
		},
		{
			name:        "open",
			description: "(limited) opens one or more projects",
			cmd:         openCommand,
		},
		{
			name:        "init",
			description: "(not implemented) initializes the jkl config file",
			cmd:         notImplementedPlan,
		},
	}

	for _, cmd := range commands {
		app.Command(cmd.name, cmd.description, cmd.cmd)
	}

	app.Run(args)
}

// CommandRunner is used to execute commands.
type CommandRunner interface {
	Run(cmd *exec.Cmd) error
}

// Logger is the interface used for all output by the CLI.
type Logger interface {
	Printf(string, ...interface{})
	Fatalf(string, ...interface{})
}

// RunOption is a function that can be used to configure optional run properties.
type RunOption func(cfg *runConfig)

// WithRuntimeOS is a RunOption that can be used to set the runtimeOS on the
// run config. Default value is runtime.GOOS.
func WithRuntimeOS(os string) RunOption {
	return func(cfg *runConfig) {
		cfg.runtimeOS = os
	}
}

type command struct {
	name        string
	description string
	cmd         func(cmd *cli.Cmd)
}

type runConfig struct {
	runtimeOS string
}

func notImplementedPlan(cmd *cli.Cmd) {
	log := log.New(os.Stderr, "", 0)

	cmd.Action = func() { log.Fatal("not implemented") }
}

func findProject(name string, projects []manifest.Project) (manifest.Project, error) {
	for _, p := range projects {
		if p.Name == name || p.Alias == name {
			return p, nil
		}
	}

	return manifest.Project{}, fmt.Errorf("project named %s not found in manifest", name)
}
