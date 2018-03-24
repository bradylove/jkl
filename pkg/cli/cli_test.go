package cli_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

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
  path: /tmp/jkl
- name: thing-with-working-path
  alias: wp
  repository: git@github.com:bradylove/thing-with-working-path.git
  path: /tmp/wp
  working_path: wp
  layout: main-vertical
- name: non-existent
  alias: ne
  repository: git@github.com:bradylove/non-existent.git
  path: /tmp/non-existent-directory
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

func (nopLogger) Printf(string, ...interface{}) {}
func (nopLogger) Fatalf(string, ...interface{}) {}

type cmdRunner struct {
	commands     []*exec.Cmd
	commandError error
}

func (r *cmdRunner) Run(cmd *exec.Cmd) error {
	r.commands = append(r.commands, cmd)
	return r.commandError
}

func tempManifest() string {
	os.Mkdir("/tmp/jkl", os.ModePerm)
	os.Mkdir("/tmp/sfs", os.ModePerm)
	os.Mkdir("/tmp/wp", os.ModePerm)

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
