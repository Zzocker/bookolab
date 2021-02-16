package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Load will read yaml config file
func Load(filename string) (*ApplicationConf, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var conf ApplicationConf
	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
