package manifest_test

import (
	"testing"

	"github.com/bradylove/jkl/pkg/manifest"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestProject(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Spec("BrowserURL returns a link to the project repository for the browser", func(t *testing.T) {
		p := manifest.Project{
			Repository: "git@github.com:bradylove/jkl.git",
		}

		url, err := p.BrowserURL()
		Expect(t, err).To(Not(HaveOccurred()))
		Expect(t, url).To(Equal("https://github.com/bradylove/jkl.git"))
	})
}
