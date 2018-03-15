package tmux

import (
	"fmt"
	"os/exec"
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
func (t Tmux) CreateWindow(name, path string) error {
	cmd := exec.Command("bash", "-c",
		fmt.Sprintf("tmux -S %s new-window -n %s -c %s", t.socket, name, path),
	)

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

// Runner is the default CommandRunner for Tmux.
type Runner struct{}

// Run executes the given command .
func (r Runner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
