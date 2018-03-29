package cli_test

import (
	"fmt"
	"os"
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

	o.Spec("fatally log if TMUX environment variable is not set", func(t *testing.T) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal("jkl edit must be ran in tmux"))
		}()

		cli.Run(&stubLogger{}, &cmdRunner{}, tempManifest(), []string{"jkl", "edit"},
			cli.WithTmuxSocket(""),
		)
	})

	o.Spec("create file with template if not exist", func(t *testing.T) {
		filePath := "/tmp/jkl-temp-manifest.yml"

		defer func() {
			os.Remove(filePath)
		}()

		cr := &cmdRunner{}
		cli.Run(&stubLogger{}, cr, filePath, []string{"jkl", "edit"},
			cli.WithTmuxSocket("/tmp/tmux"),
		)

		_, err := os.Stat(filePath)
		Expect(t, err).To(Not(HaveOccurred()))
	})
}
