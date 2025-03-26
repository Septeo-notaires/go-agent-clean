package internal

import (
	"errors"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Agents []struct {
		Name    string
		Service string
		Path    string
	}
}

func DecodeFile(blob string) Config {
	var config Config
	if _, err := toml.Decode(blob, &config); err != nil {
		panic(err)
	}

	if isEmptyConf(&config) {
		err := errors.New("fields cannot be empty or null in config file")
		panic(err)
	}
	return config
}

func isEmptyConf(config *Config) bool {
	for _, agent := range config.Agents {
		if agent.Name == "" {
			return true
		}

		if agent.Path == "" {
			return true
		}

		if agent.Service == "" {
			return true
		}
	}
	return false
}
