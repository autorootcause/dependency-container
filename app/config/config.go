package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type PluginType string

const (
	ExternalPlugin PluginType = "external"
	ShellCommand   PluginType = "command"
	ShellScript    PluginType = "script"
)

// Config defines the format of the config we expect
type Config struct {
	Plugins       []Plugin `yaml:"plugins"`
	PushReportURI string   `yaml:"push_report_uri"`        // this is where the dependency container pushes its report
	TriggerHook   bool     `yaml:"trigger_hook,omitempty"` // only accept IAM sigv4 auth currently
	// TriggerHook defines whether this program can accept triggers
}

type Plugin struct {
	Name        string     `yaml:"name"`
	Type        PluginType `yaml:"type"`
	Image       string     `yaml:"image,omitempty"`
	ShellScript string     `yaml:"shell_script,omitempty"`
	ShellFile   string     `yaml:"shell_file,omitempty"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	replaced := os.ExpandEnv(string(file))
	cfg := &Config{}
	err = yaml.Unmarshal([]byte(replaced), cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
