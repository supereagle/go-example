package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	defaultPort = 8080

	// defaultExpiration Default expiration time for token.
	defaultExpiration = 900
)

type Config struct {
	Port       int `json:"port,omitempty"`
	Expiration int `json:"expiration,omitempty"`
}

func Read(path string) (*Config, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Fail to read the config file %s", path)
	}

	cfg := &Config{}
	err = json.Unmarshal(contents, cfg)
	if err != nil {
		return nil, fmt.Errorf("Fail to unmarshal a JSON object from the config file %s", path)
	}

	// Set the default config for configures not specified
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}

	if cfg.Expiration == 0 {
		cfg.Expiration = defaultExpiration
	}

	return cfg, nil
}
