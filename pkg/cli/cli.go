package cli

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/bradylove/jkl/pkg/manifest"
	"github.com/bradylove/jkl/pkg/tmux"
	cli "github.com/jawher/mow.cli"
)

// Run initializes and executes the CLI.
func Run(
	log Logger,
	cr CommandRunner,
	manifestPath string,
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

	if fileNotExists(manifestPath) {
		err := ioutil.WriteFile(manifestPath, []byte(manifestTemplate), 0664)
		if err != nil {
			log.Fatalf("failed to create jkl manifest: %s", err)
		}

		log.Printf("example manifest created at %s", manifestPath)
	}

	m, err := manifest.Load(manifestPath)
	if err != nil {
		log.Fatalf("failed to read jkl manifest: %s", err)
	}

	app := cli.App("jkl", "developer project management life improver")
	project := app.StringArg("PROJECT", "", "name or alias of project to goto")
	app.Spec = "[PROJECT]"

	app.Action = func() {
		if *project == "" {
			app.PrintHelp()
			cli.Exit(1)
		}

		if !tm.Valid() {
			log.Fatalf("jkl goto must be ran in tmux")
		}

		p, err := m.FindProject(*project)
		if err != nil {
			log.Fatalf("%s", err)
		}

		err = tm.ChangeDirectory(p.Path)
		if err != nil {
			log.Fatalf("failed to open project '%s': %s", p.Name, err)
		}
	}

	commands := []command{
		{
			name:        "browser b",
			description: "open the projects page in the browser",
			cmd:         BrowserCommand(log, cr, m, cfg.runtimeOS),
		},
		{
			name:        "clone c",
			description: "clone the project to the projects path",
			cmd:         CloneCommand(cr, m),
		},
		{
			name:        "edit e",
			description: "open the jkl manifest for editing",
			cmd:         EditCommand(log, tm, m),
		},
		{
			name:        "goto cd",
			description: "change the current directory of the current tmux pane to the project directory",
			cmd:         GoToCommand(log, tm, m),
		},
		{
			name:        "open o",
			description: "open one or more projects",
			cmd:         OpenCommand(log, tm, m),
		},
		{
			name:        "projects p",
			description: "list known projects",
			cmd:         ProjectsCommand(log, cfg.errorWriter, m),
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

var manifestTemplate = `---
# editor: vim
# projects:
# - name: jkl
#   path: ~/gocode/src/github.com/bradylove/jkl
#   repository: git@github.com:bradylove/jkl.git
`

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
