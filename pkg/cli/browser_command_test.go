package cli_test

import (
	"testing"

	"github.com/bradylove/jkl/pkg/cli"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestBrowserCommand(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) (*testing.T, string) {
		return t, tempManifest()
	})

	o.Spec("xdg-open the project in the browser", func(t *testing.T, manifest string) {
		cr := &cmdRunner{}

		cli.Run(nopLogger{}, cr, manifest, []string{"jkl", "browser", "jkl"},
			cli.WithRuntimeOS("linux"),
		)

		Expect(t, cr.commands).To(HaveLen(1))

		cmd := cr.commands[0]
		Expect(t, cmd.Path).To(Equal("/usr/bin/xdg-open"))
		Expect(t, cmd.Args).To(Equal([]string{
			"xdg-open",
			"https://github.com/bradylove/jkl.git",
		}))
	})

	o.Spec("open the project in the browser on darwin", func(t *testing.T, manifest string) {
		cr := &cmdRunner{}

		cli.Run(nopLogger{}, cr, manifest, []string{"jkl", "browser", "jkl"},
			cli.WithRuntimeOS("darwin"),
		)

		Expect(t, cr.commands).To(HaveLen(1))

		cmd := cr.commands[0]
		Expect(t, cmd.Path).To(Equal("/bin/open"))
		Expect(t, cmd.Args).To(Equal([]string{
			"open",
			"https://github.com/bradylove/jkl.git",
		}))
	})
}
