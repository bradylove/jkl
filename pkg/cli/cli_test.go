package cli_test

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

var (
	manifestTemplate = `---
editor: code
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

type cmdRunner struct {
	commands     []*exec.Cmd
	commandError error
}

func (r *cmdRunner) Run(cmd *exec.Cmd) error {
	r.commands = append(r.commands, cmd)
	return r.commandError
}

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
