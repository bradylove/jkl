package tmux

import (
	"fmt"
	"os/exec"
	"strings"
)

// Tmux handles all interactions with Tmux.
type Tmux struct {
	socket        string
	commandRunner CommandRunner
}

// New initializes and returns a new Tmux.
func New(socket string, opts ...Option) Tmux {
	t := Tmux{
		socket:        socket,
		commandRunner: Runner{},
	}

	for _, o := range opts {
		o(&t)
	}

	return t
}

// CreateWindow creates a new tmux window with the given name.
func (t Tmux) CreateWindow(name, path string, opts ...CreateWindowOption) error {
	args := []string{
		fmt.Sprintf("tmux -S %s new-window -n %s -c %s", t.socket, name, path),
	}

	for _, o := range opts {
		o(&args)
	}

	cmd := exec.Command("bash", "-c", strings.Join(args, " \\; "))

	return t.commandRunner.Run(cmd)
}

// CommandRunner is used to execute commands on the Tmux session.
type CommandRunner interface {
	Run(cmd *exec.Cmd) error
}

// Option can be used to configure optional settings on a Tmux struct during
// initialization.
type Option func(t *Tmux)

// WithCommandRunner overrides the default command runner.
func WithCommandRunner(r CommandRunner) Option {
	return func(t *Tmux) {
		t.commandRunner = r
	}
}

// CreateWindowOption can be used to configure optional settings when creating
// a new Tmux window
type CreateWindowOption func(args *[]string)

// WithVerticalSplitPath adds a vertical split with the given path when creating
// a new window.
func WithVerticalSplitPath(path string) CreateWindowOption {
	return func(args *[]string) {
		*args = append(*args, fmt.Sprintf("split-window -h -c %s", path))
	}
}

// WithLayout sets the layout to apply to a new window on creation.
func WithLayout(layout string) CreateWindowOption {
	return func(args *[]string) {
		*args = append(*args, fmt.Sprintf("select-layout %s", layout))
	}
}

// Runner is the default CommandRunner for Tmux.
type Runner struct{}

// Run executes the given command .
func (r Runner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
