package cli_test

import (
	"fmt"
	"testing"

	"github.com/bradylove/jkl/pkg/cli"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestOpenCommand(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Spec("launch one project in tmux", func(t *testing.T) {
		cr := &cmdRunner{}

		cli.Run(&stubLogger{}, cr, tempManifest(), []string{"jkl", "open", "jkl"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)

		Expect(t, cr.commands).To(HaveLen(2))
		Expect(t, cr.commands[0].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux new-window -n jkl -c /tmp/jkl",
		}))

		Expect(t, cr.commands[1].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux send-keys 'code .' Enter",
		}))
	})

	o.Spec("launch more than one project in tmux", func(t *testing.T) {
		cr := &cmdRunner{}

		cli.Run(&stubLogger{}, cr, tempManifest(), []string{"jkl", "open", "jkl", "sfs"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)

		Expect(t, cr.commands).To(HaveLen(4))
		Expect(t, cr.commands[0].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux new-window -n jkl -c /tmp/jkl",
		}))

		Expect(t, cr.commands[1].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux send-keys 'code .' Enter",
		}))

		Expect(t, cr.commands[2].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux new-window -n simple-file-server -c /tmp/sfs",
		}))

		Expect(t, cr.commands[3].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux send-keys 'code .' Enter",
		}))
	})

	o.Spec("fatally exits if the project directory does not exist", func(t *testing.T) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal(
				"project directory for non-existent does not exist",
			))
		}()

		cli.Run(&stubLogger{}, &cmdRunner{}, tempManifest(), []string{"jkl", "open", "ne"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)
	})

	o.Spec("do not open editor with no-edit flag", func(t *testing.T) {
		cr := &cmdRunner{}

		cli.Run(&stubLogger{}, cr, tempManifest(), []string{"jkl", "open", "--no-edit", "jkl"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)

		Expect(t, cr.commands).To(HaveLen(1))
		Expect(t, cr.commands[0].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux new-window -n jkl -c /tmp/jkl",
		}))

	})

	o.Spec("open a split if working path and layout are not empty", func(t *testing.T) {
		cr := &cmdRunner{}

		cli.Run(&stubLogger{}, cr, tempManifest(), []string{"jkl", "open", "wp"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)

		Expect(t, cr.commands[0].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux new-window -n thing-with-working-path -c /tmp/wp \\; " +
				"split-window -h -c /tmp/wp/wp \\; " +
				"select-layout main-vertical",
		}))
	})

	o.Spec("fatally log if tmux socket is not valid", func(t *testing.T) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal("jkl open must be ran in tmux"))
		}()

		cli.Run(&stubLogger{}, &cmdRunner{}, tempManifest(), []string{"jkl", "open", "jkl", "sfs"},
			cli.WithTmuxSocket(""),
		)
	})
}
