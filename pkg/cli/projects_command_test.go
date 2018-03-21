package cli_test

import (
	"bytes"
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
}
