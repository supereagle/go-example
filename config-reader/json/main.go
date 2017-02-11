package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

const (
	defaultPort = 8080
)

func main() {
	// Create the command line app
	app := cli.NewApp()
	app.Name = "ConfigReader"
	app.Usage = "Tool to read config from JSON file"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.HelpFlag,
		cli.VersionFlag,
		cli.StringFlag{
			Name:  "config, c",
			Usage: "config file path",
			Value: "config.json",
		},
	}
	app.Action = func(c *cli.Context) {
		// Read the config
		path := c.String("config")
		cfg, err := Read(path)
		if err != nil {
			log.Errorln(err.Error())
			return
		}

		fmt.Printf("The config is: %v\n", cfg)
	}

	app.Run(os.Args)
}

type Config struct {
	JenkinsServer       string `json:"jenkins_server,omitempty"`
	JenkinsUser         string `json:"jenkins_user,omitempty"`
	JenkinsPassword     string `json:"jenkins_password,omitempty"`
	JenkinsCredentialId string `json:"jenkins_credential,omitempty"`
	Port                int    `json:"port,omitempty"`
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

	return cfg, nil
}
