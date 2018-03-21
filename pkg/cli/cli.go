package cli

import (
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

// Run initializes and executes the CLI.
func Run(
	log Logger,
	cr CommandRunner,
	manifest string,
	args []string,
	opts ...RunOption,
) {
	cfg := runConfig{
		runtimeOS:   runtime.GOOS,
		errorWriter: os.Stderr,
		tmuxSocket:  defaultTmuxSocket(),
	}

	for _, o := range opts {
		o(&cfg)
	}

	tm := tmux.New(cfg.tmuxSocket,
		tmux.WithCommandRunner(cr),
	)

	app := cli.App("jkl", "developer project management life improver")

	commands := []command{
		{
			name:        "browser b",
			description: "open the projects page in the browser",
			cmd:         BrowserCommand(log, cr, manifest, cfg.runtimeOS),
		},
		{
			name:        "edit e",
			description: "open the jkl manifest for editing",
			cmd:         EditCommand(log, tm, manifest),
		},
		{
			name:        "goto cd",
			description: "change the current directory of the current tmux pane to the project directory",
			cmd:         gotoCommand,
		},
		{
			name:        "projects p",
			description: "list known projects",
			cmd:         ProjectsCommand(log, cfg.errorWriter, manifest),
		},
		{
			name:        "open o",
			description: "(limited) open one or more projects",
			cmd:         openCommand,
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

// WithErrorWriter is a RunOption that can be used to override the io.Writer for
// errors. Default is os.Stderr.
func WithErrorWriter(w io.Writer) RunOption {
	return func(cfg *runConfig) {
		cfg.errorWriter = w
	}
}

// WithTmuxSocket is a RunOption to override the default tmux socket path.
// Default is the value of the TMUX environment variable.
func WithTmuxSocket(s string) RunOption {
	return func(cfg *runConfig) {
		cfg.tmuxSocket = s
	}
}

type command struct {
	name        string
	description string
	cmd         func(cmd *cli.Cmd)
}

type runConfig struct {
	runtimeOS   string
	tmuxSocket  string
	errorWriter io.Writer
}

func defaultTmuxSocket() string {
	s := os.Getenv("TMUX")
	if s == "" {
		return s
	}

	return strings.Split(s, ",")[0]
}
