package manifest

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Projects []Project `yaml:"projects,flow"`
}

func Load(path string) (Manifest, error) {
	f, err := os.Open(path)
	if err != nil {
		return Manifest{}, err
	}
	defer f.Close()

	var m Manifest
	err = yaml.NewDecoder(f).Decode(&m)
	if err != nil {
		return Manifest{}, err
	}

	return m, nil
}

type Project struct {
	Name        string `yaml:"name"`
	Alias       string `yaml:"alias"`
	BasePath    string `yaml:"base_path"`
	WorkingPath string `yaml:"working_path"`
	Repository  string `yaml:"repository"`
	Submodules  bool   `yaml:"submodules"`
}
