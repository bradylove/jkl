package manifest

import (
	"errors"
	"net/url"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	// ErrProjectNotFound indicates the Manifest was unable to find a given
	// project.
	ErrProjectNotFound = errors.New("project not found")
)

// Manifest defines the structure for the jkl YAML configuration file.
type Manifest struct {
	Path     string
	Editor   string    `yaml:"editor"`
	Projects []Project `yaml:"projects,flow"`
}

// Load reads the jkl YAML configuration file from the specified path and
// decodes the manifest into a Manifest struct. Editor will default to the
// EDITOR environment variable, if an editor is specified in the manifest, that
// value will be used.
func Load(path string) (Manifest, error) {
	f, err := os.Open(path)
	if err != nil {
		return Manifest{}, err
	}
	defer f.Close()

	m := Manifest{
		Path:   path,
		Editor: os.Getenv("EDITOR"),
	}
	err = yaml.NewDecoder(f).Decode(&m)
	if err != nil {
		return Manifest{}, err
	}

	for i, p := range m.Projects {
		if strings.HasPrefix(p.Path, "~/") {
			p.Path = strings.Replace(p.Path, "~/", os.Getenv("HOME")+"/", 1)
			m.Projects[i] = p
		}
	}

	return m, nil
}

// FindProject will return the first project found with a matching name or alias.
func (m Manifest) FindProject(name string) (Project, error) {
	for _, p := range m.Projects {
		if p.Name == name || p.Alias == name {
			return p, nil
		}
	}

	return Project{}, ErrProjectNotFound
}

// Project defines the YAML configuration for a single project.
type Project struct {
	Name        string `yaml:"name"`
	Alias       string `yaml:"alias"`
	Path        string `yaml:"path"`
	WorkingPath string `yaml:"working_path"`
	Layout      string `yaml:"layout"`
	Repository  string `yaml:"repository"`
}

// BrowserURL makes a best effort to build a browser URL from the Repository.
func (p Project) BrowserURL() (string, error) {
	s := strings.Replace(p.Repository, ":", "/", 1)
	s = strings.Replace(s, "git@", "", 1)

	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	u.Scheme = "https"

	return u.String(), nil
}
