package config

import (
	"gopkg.in/yaml.v3"
    "os"
	"time"
)

type Config struct {
	Db          struct {
					Url      string `yaml:"url"`
					Name 	 string `yaml:"name"`
				} `yaml:"db"`
	Auth        struct {
					AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL" yaml:"accessTokenTTL"`
					RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL" yaml:"refreshTokenTTL"`
					Secret      string `yaml:"secret"`
				} `yaml:"auth"`
}


func Init(configsPath string) (*Config, error) {
	var cfg Config

	yamlFile, err := os.ReadFile(configsPath)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &cfg)

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
