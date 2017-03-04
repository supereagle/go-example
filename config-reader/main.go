package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

const (
	defaultPort    = 8080
	requestTimeout = 5 * time.Second

	// Config source supported
	ENV  string = "env"
	ETCD        = "etcd"
	JSON        = "json"
	YAML        = "yaml"
)

func main() {
	// Create the command line app
	app := cli.NewApp()
	app.Name = "ConfigReader"
	app.Usage = "Tool to read config from environment variables, etcd, and json or yaml files"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.HelpFlag,
		cli.VersionFlag,
		cli.StringFlag{
			Name:  "source, s",
			Usage: "config source, supported env, etcd, json or yaml, default is env",
			Value: "env",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "config file path, needed when config source is json or yaml",
		},
		cli.StringFlag{
			Name:  "endpoints, e",
			Usage: "urls of ETCD cluster separated by comma, needed when config source is etcd",
		},
	}
	app.Action = func(c *cli.Context) {
		// Read the config
		source := c.String("source")

		var err error
		cfg := &Config{}

		switch source {
		case ENV:
			cfg, err = ReadEnv()
		case JSON, YAML:
			path := c.String("config")
			if len(path) == 0 {
				err = fmt.Errorf("The config file path must be provided when config source is JSON or YAML.")
				break
			}
			if source == JSON {
				cfg, err = ReadJson(path)
			} else {
				cfg, err = ReadYaml(path)
			}
		case ETCD:
			endpointsStr := c.String("endpoints")
			if len(endpointsStr) == 0 {
				err = fmt.Errorf("The ETCD endpoints must be provided when config source is ETCD.")
				break
			}
			endpoints := strings.Split(endpointsStr, ",")
			cfg, err = ReadEtcd(endpoints)
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
	JenkinsServer     string `json:"jenkins_server,omitempty" yaml:"jenkins_server,omitempty"`
	JenkinsUser       string `json:"jenkins_user,omitempty" yaml:"jenkins_user,omitempty"`
	JenkinsPassword   string `json:"jenkins_password,omitempty" yaml:"jenkins_password,omitempty"`
	JenkinsCredential string `json:"jenkins_credential,omitempty" yaml:"jenkins_credential,omitempty"`
	Port              int    `json:"port,omitempty" yaml:"port,omitempty"`
}

const (
	JENKINS_SERVER     = "/dev/JENKINS_SERVER"
	JENKINS_USER       = "/dev/JENKINS_USER"
	JENKINS_PASSWORD   = "/dev/JENKINS_PASSWORD"
	JENKINS_CREDENTIAL = "/dev/JENKINS_CREDENTIAL"
	PORT               = "/dev/PORT"
)

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

func ReadEtcd(endpoints []string) (*Config, error) {
	c, err := client.New(client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("Fail to new the ETCD client as %s", err.Error())
	}

	cfg := &Config{}
	kapi := client.NewKeysAPI(c)

	// Get the config from ETCD one by one. Any better way?
	resp, err := kapi.Get(context.Background(), JENKINS_SERVER, nil)
	if err != nil {
		return nil, fmt.Errorf("Fail to get the config from ETCD as %s", err.Error())
	}
	cfg.JenkinsServer = resp.Node.Value

	resp, err = kapi.Get(context.Background(), JENKINS_USER, nil)
	if err != nil {
		return nil, fmt.Errorf("Fail to get the config from ETCD as %s", err.Error())
	}
	cfg.JenkinsUser = resp.Node.Value

	resp, err = kapi.Get(context.Background(), JENKINS_PASSWORD, nil)
	if err != nil {
		return nil, fmt.Errorf("Fail to get the config from ETCD as %s", err.Error())
	}
	cfg.JenkinsPassword = resp.Node.Value

	resp, err = kapi.Get(context.Background(), JENKINS_CREDENTIAL, nil)
	if err != nil {
		return nil, fmt.Errorf("Fail to get the config from ETCD as %s", err.Error())
	}
	cfg.JenkinsCredential = resp.Node.Value

	resp, err = kapi.Get(context.Background(), PORT, nil)
	if err != nil {
		return nil, fmt.Errorf("Fail to get the config from ETCD as %s", err.Error())
	}
	port, err := strconv.Atoi(resp.Node.Value)
	if err != nil {
		return nil, fmt.Errorf("Fail to convert the port from string to int as %s", err.Error())
	}
	cfg.Port = port

	return cfg, nil
}

func ReadEnv() (*Config, error) {
	cfg := &Config{}

	cfg.JenkinsServer, cfg.JenkinsUser, cfg.JenkinsPassword, cfg.JenkinsCredential =
		os.Getenv("JENKINS_SERVER"), os.Getenv("JENKINS_USER"), os.Getenv("JENKINS_PASSWORD"), os.Getenv("JENKINS_CREDENTIAL")

	portEnv := os.Getenv("PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return nil, fmt.Errorf("Fail to get the port as the environment variable PORT is not set or not corret.")
	}
	cfg.Port = port

	return cfg, nil
}
