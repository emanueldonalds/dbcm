package config

import (
	"fmt"
	"os"
	"path"

	"github.com/emanueldonalds/dbcm/configs"
	"github.com/spf13/viper"
)

var configName = "config"
var configDirPaths = []string{
	".",
	"./configs",
	os.ExpandEnv("$XGD_CONFIG_HOME/dbcm"),
	os.ExpandEnv("$HOME/.config/dbcm"),
	os.ExpandEnv("$HOME/.dbcm"),
	os.ExpandEnv("/etc/dbcm"),
}

type Config struct {
	Global      Global       `json:"global"`
	Connections []Connection `json:"connections"`
}

type Global struct {
	SshKey string `json:"ssh_key"`
}

type Connection struct {
	Name string `json:"name"`
}

type Loader struct {
	SearchPaths []string
}

func NewLoader() *Loader {
	return &Loader{
		SearchPaths: configDirPaths,
	}
}

func (l *Loader) Load() (*Config, error) {
	viper.SetConfigName(configName)

	for _, path := range configDirPaths {
		viper.AddConfigPath(path)
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// There is no config file, let's create one based on the template conf and try loading one more time.
			if err := l.createDefaultConfigFile(); err != nil {
				return nil, fmt.Errorf("Error: failed when trying to create default config file: %w", err)
			}
			if err := viper.ReadInConfig(); err != nil {
				return nil, fmt.Errorf("Error: failed when reading config: %w", err)
			}
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("Error: failed when unmarshalling config: %w", err)
	}
	return &cfg, nil
}

func (l *Loader) createDefaultConfigFile() error {
	configDirPath := l.SearchPaths[0]

	if err := os.MkdirAll(configDirPath, 0700); err != nil {
		return fmt.Errorf("could not create config dir: %w", err)
	}

	configPath := path.Join(configDirPath, configName+".yaml")
	if err := os.WriteFile(configPath, configs.DefaultConfig, 0600); err != nil {
		return fmt.Errorf("could not create default config file: %w", err)
	}

	return nil
}
