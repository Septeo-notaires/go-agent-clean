package internal

import "github.com/BurntSushi/toml"

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
	return config
}
