package manifest

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Manifest defines the structure for the jkl YAML configuration file.
type Manifest struct {
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
		Editor: os.Getenv("EDITOR"),
	}
	err = yaml.NewDecoder(f).Decode(&m)
	if err != nil {
		return Manifest{}, err
	}

	return m, nil
}

// Project defines the YAML configuration for a single project.
type Project struct {
	Name        string `yaml:"name"`
	Alias       string `yaml:"alias"`
	BasePath    string `yaml:"base_path"`
	WorkingPath string `yaml:"working_path"`
	Layout      string `yaml:"layout"`
	Repository  string `yaml:"repository"`
	Submodules  bool   `yaml:"submodules"`
}
