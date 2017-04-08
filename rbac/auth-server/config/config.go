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
	Port  int    `json:"port,omitempty"`
	Token *Token `json:"token,omitempty"`
}

type Token struct {
	Expiration int64  `json:"expiration,omitempty"`
	Algorithm  string `json:"algorithm,omitempty"`
	Key        string `yaml:"key,omitempty"`
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

	return cfg, nil
}

func Verify(cfg *Config) error {
	// Verify the token config
	if cfg.Token == nil {
		return fmt.Errorf("Token is not configured.")
	}

	if cfg.Token.Key == "" {
		return fmt.Errorf("The key of token is not configured.")
	}

	if cfg.Token.Algorithm == "" {
		return fmt.Errorf("The algorithm of token is not configured.")
	}

	// Set the default config for configures not specified
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}

	if cfg.Token.Expiration == 0 {
		cfg.Token.Expiration = defaultExpiration
	}

	return nil
}
