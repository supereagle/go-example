package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
)

const (
	defaultPort = 8080
)

func main() {
	// Create the command line app
	app := cli.NewApp()
	app.Name = "ConfigReader"
	app.Usage = "Tool to read config from etcd, json and yaml files"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.HelpFlag,
		cli.VersionFlag,
		cli.StringFlag{
			Name:  "source, s",
			Usage: "source of config, supported etcd, json and yaml files",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "config file path",
			Value: "config.json",
		},
	}
	app.Action = func(c *cli.Context) {
		// Read the config
		source := c.String("source")
		path := c.String("config")

		var err error
		cfg := &Config{}

		switch source {
		case "json":
			cfg, err = ReadJson(path)
		case "yaml":
			cfg, err = ReadYaml(path)
		case "etcd":
			log.Infoln("Will be supported")
		default:
			log.Errorf("The source type of config %s is not supported, only supported etcd, json and yaml.", source)
		}

		if err != nil {
			log.Errorln(err.Error())
			return
		}

		fmt.Printf("The config is: %v\n", cfg)
	}

	app.Run(os.Args)
}

type Config struct {
	JenkinsServer       string `json:"jenkins_server,omitempty" yaml:"jenkins_server,omitempty"`
	JenkinsUser         string `json:"jenkins_user,omitempty" yaml:"jenkins_user,omitempty"`
	JenkinsPassword     string `json:"jenkins_password,omitempty" yaml:"jenkins_password,omitempty"`
	JenkinsCredentialId string `json:"jenkins_credential,omitempty" yaml:"jenkins_credential,omitempty"`
	Port                int    `json:"port,omitempty" yaml:"port,omitempty"`
}

func ReadJson(path string) (*Config, error) {
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

func ReadYaml(path string) (*Config, error) {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Fail to read the config file %s", path)
	}

	cfg := &Config{}
	err = yaml.Unmarshal(contents, cfg)
	if err != nil {
		return nil, fmt.Errorf("Fail to unmarshal a JSON object from the config file %s", path)
	}

	// Set the default config for configures not specified
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}

	return cfg, nil
}
