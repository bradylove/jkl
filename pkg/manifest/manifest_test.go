package manifest_test

import (
	"testing"

	"github.com/bradylove/jkl/pkg/manifest"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestManifest(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Group("FindProject", func() {
		o.Spec("return project with matching name", func(t *testing.T) {
			m := manifest.Manifest{
				Projects: []manifest.Project{
					{Name: "one"},
					{Name: "two"},
					{Name: "three"},
					{Name: "four"},
				},
			}

			p, err := m.FindProject("three")
			Expect(t, err).To(Not(HaveOccurred()))
			Expect(t, p.Name).To(Equal("three"))
		})

		o.Spec("return project with matching alias", func(t *testing.T) {
			m := manifest.Manifest{
				Projects: []manifest.Project{
					{Name: "one", Alias: "o"},
					{Name: "two", Alias: "tw"},
					{Name: "three", Alias: "th"},
					{Name: "four", Alias: "f"},
				},
			}

			p, err := m.FindProject("tw")
			Expect(t, err).To(Not(HaveOccurred()))
			Expect(t, p.Name).To(Equal("two"))
		})

		o.Spec("return error when no match is found", func(t *testing.T) {
			m := manifest.Manifest{}

			p, err := m.FindProject("unknokwn")
			Expect(t, err).To(Equal(manifest.ErrProjectNotFound))
			Expect(t, p).To(Equal(manifest.Project{}))
		})
	})
}
