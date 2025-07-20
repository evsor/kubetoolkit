package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Repository string `yaml:"repository"`
	Image      string `yaml:"image"`
	Tag        string `yaml:"tag"`
}

const configFileName = "kubetoolkit.yaml"

func ConfigFilePath() (string, error) {
	configDir := xdg.ConfigHome
	if configDir == "" {
		return "", fmt.Errorf("XDG_CONFIG_HOME not set")
	}
	return filepath.Join(configDir, "kubetoolkit", configFileName), nil
}

func Load() (*Config, error) {
	path, err := ConfigFilePath()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func Save(cfg *Config) error {
	path, err := ConfigFilePath()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewEncoder(f).Encode(cfg)
}
