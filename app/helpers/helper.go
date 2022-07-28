package helpers

import (
	"github.com/ahmadateya/flotta-webapp-backend/config"
	"gopkg.in/yaml.v3"
	"os"
)

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*config.Config, error) {
	c := &config.Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&c); err != nil {
		return nil, err
	}

	return c, nil
}
