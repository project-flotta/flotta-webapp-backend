package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server  Server
	Storage Storage
}

type Server struct {
	Host string
	Port string
}

type Storage struct {
	S3 S3
}

type S3 struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	Region          string `yaml:"region"`
	Bucket          string `yaml:"bucket"`
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
