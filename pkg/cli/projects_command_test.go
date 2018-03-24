package cli_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/bradylove/jkl/pkg/cli"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestProjectsCommand(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.BeforeEach(func(t *testing.T) (*testing.T, string) {
		return t, tempManifest()
	})

	o.Spec("print a list of known projects", func(t *testing.T, manifest string) {
		buf := bytes.NewBuffer(nil)
		cli.Run(nopLogger{}, nil, manifest, []string{"jkl", "projects"},
			cli.WithErrorWriter(buf),
		)

		Expect(t, strings.Split(buf.String(), "\n")).To(Equal([]string{
			"NAME                     ALIAS  PATH",
			"jkl                             /tmp/jkl",
			"non-existent             ne     /tmp/non-existent-directory",
			"simple-file-server       sfs    /tmp/sfs",
			"thing-with-working-path  wp     /tmp/wp",
			"",
		}))
	})

	o.Spec("fatally log when loading manifest fails", func(t *testing.T, _ string) {
		defer func() {
			err := recover()
			Expect(t, fmt.Sprint(err)).To(Equal("failed to read jkl manifest: open /tmp/unknown: no such file or directory"))
		}()

		loggr := &stubLogger{}
		cli.Run(loggr, nil, "/tmp/unknown", []string{"jkl", "projects"})
	})
}
