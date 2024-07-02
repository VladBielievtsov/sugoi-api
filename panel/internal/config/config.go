package config

import (
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App  appConfig
	Site siteConfig
}

type appConfig struct {
	Port int `yaml:"port"`
}

type siteConfig struct {
	Title string `yaml:"title"`
}

var cfg *Config

func New(confDir string) (*Config, error) {
	var app appConfig
	var site siteConfig

	if err := cleanenv.ReadConfig(filepath.Join(confDir, "app.yaml"), &app); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(filepath.Join(confDir, "site.yaml"), &site); err != nil {
		return nil, err
	}

	cfg = &Config{
		App:  app,
		Site: site,
	}

	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}
