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

// ChangeDirectory will execute `cd {path}` in the current active Tmux pane.
func (t Tmux) ChangeDirectory(path string) error {
	return t.Execute(fmt.Sprintf("cd %s", path))
}

// Execute will execute any given command in the current active Tmux pane.
func (t Tmux) Execute(cmdStr string) error {
	args := []string{
		fmt.Sprintf("tmux -S %s send-keys '%s' Enter", t.socket, cmdStr),
	}

	cmd := exec.Command("bash", "-c", strings.Join(args, " \\; "))

	return t.commandRunner.Run(cmd)
}

// Valid returns a boolean indicating whether the socket that tmux.Tmux was
// configured with is valid.
func (t Tmux) Valid() bool {
	if t.socket == "" {
		return false
	}

	return true
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

// CommandRunner is used to execute commands on the Tmux session.
type CommandRunner interface {
	Run(cmd *exec.Cmd) error
}

// Runner is the default CommandRunner for Tmux.
type Runner struct{}

// Run executes the given command .
func (r Runner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
