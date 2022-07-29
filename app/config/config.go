package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server Server
}

type Server struct {
	Host string
	Port string
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	c := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	return c, nil
}
