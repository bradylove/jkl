package cli_test

import "fmt"

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

type stubLogger struct {
	printfMessages []string
}

func (s *stubLogger) Printf(f string, a ...interface{}) {
	s.printfMessages = append(s.printfMessages, fmt.Sprintf(f, a...))
}

func (s *stubLogger) Fatalf(f string, a ...interface{}) {
	panic(fmt.Sprintf(f, a...))
}

type nopLogger struct{}

func (nopLogger) Printf(string, ...interface{}) {
	panic("not implemented")
}

func (nopLogger) Fatalf(string, ...interface{}) {
	panic("not implemented")
}
