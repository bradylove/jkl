jkl
====================================

jkl is a [Tmux][tmux] and [Git][git] focused developer project management tool.
jkl makes it easy to switch from one project in a single command.

**jkl only supports Linux and OSX.**

## Why is it named jkl?

No reason really other than it is short, quick and easy to type.

## Installation

### From Source

```
go get -u github.com/bradylove/jkl/cmd/jkl
```

## Example Configuration

``` yaml
---
projects:
- name: jkl
  alias: jk
  base_path: ~/gocode/src/github.com/bradylove/jkl
  working_path: .
  layout: main-vertical
```

## Usage

```
$ jkl --help

Usage: jkl COMMAND [arg...]

project management life improver

Commands:
  open         (limited) opens one or more projects
  edit         (not implemented) opens the jkl manifest for editing
  github       (not implemented) open the projects github page in the browser
  init         (not implemented) initializes the jkl config file

Run 'jkl COMMAND --help' for more information on a command.
```

```
$ jkl open jkl
```

## Running Tests

```
$ go get -u -t ./....
$ go test ./...
```

## License

Copyright 2018 Brady Love <love.brady at gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[git]:  https://git-scm.com/
[tmux]: https://github.com/tmux/tmux
