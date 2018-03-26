package manifest_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/bradylove/jkl/pkg/manifest"

	"github.com/apoydence/onpar"
	. "github.com/apoydence/onpar/expect"
	. "github.com/apoydence/onpar/matchers"
)

func TestManifest(t *testing.T) {
	o := onpar.New()
	defer o.Run(t)

	o.Group("Load", func() {
		o.Spec("replace ~/ with users home directory", func(t *testing.T) {
			m, err := manifest.Load(tempManifest())
			Expect(t, err).To(Not(HaveOccurred()))

			var p manifest.Project
			for _, project := range m.Projects {
				if project.Name == "jkl" {
					p = project
					break
				}
			}

			Expect(t, p.Name).To(Equal("jkl"))
			Expect(t, p.Path).To(Equal(os.Getenv("HOME") + "/jkl"))
		})
	})

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

var (
	manifestTemplate = `---
editor: code
projects:
- name: simple-file-server
  alias: sfs
  repository: git@github.com:bradylove/sfs.git
  path: /tmp/sfs
- name: jkl
  repository: git@github.com:bradylove/jkl.git
  path: ~/jkl
`
)

func tempManifest() string {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(manifestTemplate))
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	return f.Name()
}
