package cli

import (
	"fmt"
	"io"
	"sort"
	"text/tabwriter"

	"github.com/bradylove/jkl/pkg/manifest"
	cli "github.com/jawher/mow.cli"
)

// ProjectsCommand will sort and print all known projects.
func ProjectsCommand(log Logger, w io.Writer, m manifest.Manifest) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			projects := m.Projects
			sort.Sort(SortableProjects(projects))

			tw := tabwriter.NewWriter(w, 0, 2, 2, ' ', 0)
			fmt.Fprintf(tw, "NAME\tALIAS\tPATH\n")

			for _, p := range projects {
				fmt.Fprintf(tw, "%s\t%s\t%s\n",
					p.Name,
					p.Alias,
					p.Path,
				)
			}

			tw.Flush()
		}
	}
}

// SortableProjects satisfies the sort.Interface for sorting by Project name.
type SortableProjects []manifest.Project

func (s SortableProjects) Len() int               { return len(s) }
func (s SortableProjects) Less(i int, j int) bool { return s[i].Name < s[j].Name }
func (s SortableProjects) Swap(i int, j int)      { s[i], s[j] = s[j], s[i] }
