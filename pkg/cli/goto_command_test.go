package cli_test

import (
	"fmt"
	"testing"

	"github.com/bradylove/jkl/pkg/cli"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestGoToCommand(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Spec("changes directory to the project path", func(t *testing.T) {
		cr := &cmdRunner{}

		cli.Run(&stubLogger{}, cr, tempManifest(), []string{"jkl", "goto", "jkl"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)

		Expect(t, cr.commands[0].Args).To(Equal([]string{
			"bash", "-c",
			"tmux -S /tmp/tmux send-keys 'cd /tmp/jkl' Enter",
		}))
	})

	o.Spec("fatally logs if tmux socket is not valid", func(t *testing.T) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal("jkl goto must be ran in tmux"))
		}()

		cli.Run(&stubLogger{}, &cmdRunner{}, tempManifest(), []string{"jkl", "goto", "jkl"},
			cli.WithTmuxSocket(""),
		)
	})

	o.Spec("fatally log if the project is not found", func(t *testing.T) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal("project not found"))
		}()

		cli.Run(&stubLogger{}, &cmdRunner{}, tempManifest(), []string{"jkl", "goto", "unknown"})
	})
}
