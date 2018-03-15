package tmux_test

import (
	"os/exec"
	"testing"

	"github.com/bradylove/jkl/pkg/tmux"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestCommandExecution(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Spec("setup a new window for a project", func(t *testing.T) {
		cr := &spyCommandRunner{}
		tm := tmux.New("/tmp/tmux-socket", tmux.WithCommandRunner(cr))

		err := tm.CreateWindow("one", "basepath")
		Expect(t, err).To(Not(HaveOccurred()))
		Expect(t, cr.runCmds).To(HaveLen(1))

		cmd := cr.runCmds[0]
		Expect(t, cmd.Path).To(Equal("/bin/bash"))

		// Example of how to send multiple commands
		// fmt.Sprintf("tmux -S %s new-window -c %s -n %s \\; split-window -h \\; split-window -h \\; split-window -h \\; select-layout main-vertical",

		Expect(t, cmd.Args).To(Equal([]string{
			"bash",
			"-c",
			"tmux -S /tmp/tmux-socket new-window -n one -c basepath",
		}))
	})
}

type spyCommandRunner struct {
	runCmds  []*exec.Cmd
	runError error
}

func (r *spyCommandRunner) Run(cmd *exec.Cmd) error {
	r.runCmds = append(r.runCmds, cmd)
	return r.runError
}
