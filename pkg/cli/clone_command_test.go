package cli_test

import (
	"testing"

	"github.com/bradylove/jkl/pkg/cli"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestCloneCommand(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Spec("execute git clone", func(t *testing.T) {
		cr := &cmdRunner{}
		cli.Run(&stubLogger{}, cr, tempManifest(), []string{"jkl", "clone", "jkl"})

		Expect(t, cr.commands[0].Args).To(Equal([]string{
			"bash", "-c",
			"git clone git@github.com:bradylove/jkl.git /tmp/jkl",
		}))
	})
}
