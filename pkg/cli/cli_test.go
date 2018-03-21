package cli_test

var (
	manifestTemplate = `---
projects:
- name: simple-file-server
  alias: sfs
  repository: git@github.com:bradylove/sfs.git
  path: ~/gocode/src/github.com/bradylove/sfs
- name: jkl
  repository: git@github.com:bradylove/jkl.git
  path: ~/gocode/src/github.com/bradylove/jkl
`
)

type nopLogger struct{}

func (nopLogger) Printf(string, ...interface{}) {
	panic("not implemented")
}

func (nopLogger) Fatalf(string, ...interface{}) {
	panic("not implemented")
}
