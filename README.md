jkl
====================================

jkl is a [tmux][tmux] and [Git][git] focused project navigation tool.

**jkl only supports Linux and OSX.**

## Why is it named jkl?

No reason really other than it is short, quick and easy to type.

## Installation

### From Source

```
go get -u github.com/bradylove/jkl/cmd/jkl
```

## Manifest

The manifest for jkl is written in `yaml`. The manifest should be placed in your
users home directory with the filename `.jkl` (`$HOME/.jkl`).

### Root

| Field          | Description                                                                           |
|---------------:|---------------------------------------------------------------------------------------|
| **`projects`** | Array of [projects](#project)                                                         |
| **`editor`**   | (Optional) Overrides the editor to use. Defaults to the `EDITOR` environment variable |

### Project

| Field              | Description                                                                                        |
|-------------------:|----------------------------------------------------------------------------------------------------|
| **`name`**         | The full name of the project. This name can be used as the `PROJECT` argument for any jkl commands |
| **`path`**         | The filepath of the project                                                                        |
| **`respository`**  | The remote git repository for the project                                                          |
| **`alias`**        | (Optional) Alias for the project. This can be used as the `PROJECT` argument for any jkl commands  |
| **`layout`**       | (Optional) tmux layout                                                                             |
| **`working_path`** | (Optional) Path to open editor if different from `path`. Relative to `path`                        |

### Example

``` yaml
---
editor: code
projects:
- name: jkl
  repository: git@github.com:bradylove/jkl.git
  path: ~/gocode/src/github.com/bradylove/jkl
  alias: jk
  layout: main-vertical
  working_path: cmd/jkl
```

## Usage

```
$ jkl --help

Usage: jkl COMMAND [arg...]

developer project management life improver

Commands:
  browser, b    open the projects page in the browser
  clone, c      clone the project to the projects path
  edit, e       open the jkl manifest for editing
  goto, cd      change the current directory of the current tmux pane to the project directory
  open, o       open one or more projects
  projects, p   list known projects

Run 'jkl COMMAND --help' for more information on a command.
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
