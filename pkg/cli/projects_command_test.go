package cli_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
		f, err := ioutil.TempFile("", "")
		Expect(t, err).To(Not(HaveOccurred()))

		_, err = f.Write([]byte(manifestTemplate))
		Expect(t, err).To(Not(HaveOccurred()))

		err = f.Close()
		Expect(t, err).To(Not(HaveOccurred()))

		return t, f.Name()
	})

	o.Spec("print a list of known projects", func(t *testing.T, manifest string) {
		buf := bytes.NewBuffer(nil)
		cli.Run(nopLogger{}, nil, manifest, []string{"jkl", "projects"},
			cli.WithErrorWriter(buf),
		)

		Expect(t, strings.Split(buf.String(), "\n")).To(Equal([]string{
			"NAME                ALIAS  PATH",
			"simple-file-server  sfs    ~/gocode/src/github.com/bradylove/sfs",
			"jkl                        ~/gocode/src/github.com/bradylove/jkl",
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
