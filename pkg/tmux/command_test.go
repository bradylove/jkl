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

	o.Group("Valid", func() {
		o.Spec("return true with a valid socket", func(t *testing.T) {
			tm := tmux.New("/tmp/tmux-socket")
			Expect(t, tm.Valid()).To(BeTrue())
		})

		o.Spec("return false with a valid socket", func(t *testing.T) {
			tm := tmux.New("")
			Expect(t, tm.Valid()).To(BeFalse())
		})
	})

	o.Group("CreateWindow", func() {
		o.Spec("setup a new window for a project", func(t *testing.T) {
			cr := &spyCommandRunner{}
			tm := tmux.New("/tmp/tmux-socket", tmux.WithCommandRunner(cr))

			err := tm.CreateWindow("one", "basepath")
			Expect(t, err).To(Not(HaveOccurred()))
			Expect(t, cr.runCmds).To(HaveLen(1))

			cmd := cr.runCmds[0]
			Expect(t, cmd.Path).To(Equal("/bin/bash"))

			Expect(t, cmd.Args).To(Equal([]string{
				"bash",
				"-c",
				"tmux -S /tmp/tmux-socket new-window -n one -c basepath",
			}))
		})

		o.Spec("setup a new window with a split with path", func(t *testing.T) {
			cr := &spyCommandRunner{}
			tm := tmux.New("/tmp/tmux-socket", tmux.WithCommandRunner(cr))

			err := tm.CreateWindow("one", "basepath",
				tmux.WithVerticalSplitPath("codepath"),
			)
			Expect(t, err).To(Not(HaveOccurred()))
			Expect(t, cr.runCmds).To(HaveLen(1))

			cmd := cr.runCmds[0]
			Expect(t, cmd.Args).To(Equal([]string{
				"bash",
				"-c",
				"tmux -S /tmp/tmux-socket new-window -n one -c basepath \\; split-window -h -c codepath",
			}))
		})

		o.Spec("setup new window with layout", func(t *testing.T) {
			cr := &spyCommandRunner{}
			tm := tmux.New("/tmp/tmux-socket", tmux.WithCommandRunner(cr))

			err := tm.CreateWindow("one", "basepath",
				tmux.WithLayout("main-horizontal"),
			)
			Expect(t, err).To(Not(HaveOccurred()))
			Expect(t, cr.runCmds).To(HaveLen(1))

			cmd := cr.runCmds[0]
			Expect(t, cmd.Args).To(Equal([]string{
				"bash",
				"-c",
				"tmux -S /tmp/tmux-socket new-window -n one -c basepath \\; select-layout main-horizontal",
			}))
		})
	})

	o.Group("ChangeDirectory", func() {
		o.Spec("send keys to change directory to tmux", func(t *testing.T) {
			cr := &spyCommandRunner{}
			tm := tmux.New("/tmp/tmux-socket", tmux.WithCommandRunner(cr))

			err := tm.ChangeDirectory("~/")
			Expect(t, err).To(Not(HaveOccurred()))

			cmd := cr.runCmds[0]
			Expect(t, cmd.Args).To(Equal([]string{
				"bash",
				"-c",
				"tmux -S /tmp/tmux-socket send-keys 'cd ~/' Enter",
			}))
		})
	})

	o.Group("Execute", func() {
		o.Spec("execute arbitrary commands in tmux pane", func(t *testing.T) {
			cr := &spyCommandRunner{}
			tm := tmux.New("/tmp/tmux-socket", tmux.WithCommandRunner(cr))

			err := tm.Execute("echo Hello!")
			Expect(t, err).To(Not(HaveOccurred()))

			cmd := cr.runCmds[0]
			Expect(t, cmd.Args).To(Equal([]string{
				"bash",
				"-c",
				"tmux -S /tmp/tmux-socket send-keys 'echo Hello!' Enter",
			}))
		})
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
