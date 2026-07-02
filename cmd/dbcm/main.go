package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var configName = "config"
var extensions = []string{ "yaml", "yml", "toml", "json"}

var configDirPaths = []string{
	os.ExpandEnv("$HOME/.config/dbcm"),
	"./configs",
	".",
}

func main() {
	fmt.Printf("/dbcm\n")

	err := LoadConfig()

	if err != nil {
		if !configExists() {
			// There is no config file, let's create one based on the template conf and try loading again.

			configDirPath := configDirPaths[0]
			defaultConf := readDefaultConfig()
			configFilePath := filepath.Join(configDirPath, configName + ".yaml")

			createConfigDir(configDirPath)
			createConfigFile(configFilePath, defaultConf)
			err := LoadConfig()

			if err != nil {
				panic(err)
			}
		}
	}
}

func createConfigFile(configPath string, content []byte) {
	err := os.WriteFile(configPath, content, 0600)
	if err != nil {
		panic(err)
	}
}

func readDefaultConfig() []byte {
	defaultConf, err := os.ReadFile("./configs/default.yaml")
	if err != nil {
		panic(err)
	}
	return defaultConf
}

func createConfigDir(path string) {
	err := os.MkdirAll(path, 0700)
	if err != nil {
		panic(err)
	}
}

func configExists() bool {
	for _, dirPath := range configDirPaths {
		for _, extension := range extensions {
			filename := fmt.Sprintf("%s.%s", configName, extension)
			_, err := os.Stat(filepath.Join(dirPath, filename))

			if err != nil {
				continue
			}

			return true
		}
	}
	return false
}

func LoadConfig() error {
	v := viper.New()
	v.SetConfigName(configName)

	for _, path := range configDirPaths {
		v.AddConfigPath(path)
	}

	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Could not read config file %w", err)
	}

	return nil
}
