package cli_test

import (
	"fmt"
	"testing"

	"github.com/bradylove/jkl/pkg/cli"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestEditCommand(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Spec("open jkl manifest in editor", func(t *testing.T) {
		logger := &stubLogger{}
		cr := &cmdRunner{}
		manifestPath := tempManifest()

		cli.Run(logger, cr, manifestPath, []string{"jkl", "edit"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)

		Expect(t, cr.commands[0].Args).To(Equal([]string{
			"bash", "-c",
			fmt.Sprintf("tmux -S /tmp/tmux send-keys 'code %s' Enter", manifestPath),
		}))
	})

	o.Spec("fatally log if loading manifest fails", func(t *testing.T) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal("failed to read jkl manifest: open /tmp/unknown: no such file or directory"))
		}()

		cli.Run(&stubLogger{}, &cmdRunner{}, "/tmp/unknown", []string{"jkl", "edit"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)
	})

	o.Spec("fatally log if TMUX environment variable is not set", func(t *testing.T) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal("jkl edit must be ran in tmux"))
		}()

		cli.Run(&stubLogger{}, &cmdRunner{}, tempManifest(), []string{"jkl", "edit"},
			cli.WithTmuxSocket(""),
		)
	})
}
