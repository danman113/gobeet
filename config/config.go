package config

import (
	"encoding/json"
	"github.com/danman113/gobeet/site"
	"io/ioutil"
)

type EmailConfig struct {
	Address    string   `json: address`
	Server     string   `json: server`
	Port       string   `json: port`
	Template   string   `json: template`
	Recipients []string `json: recipients`
}
type Config struct {
	Sites []site.Website `json: sites`
	Email EmailConfig    `json: email`
}

func ParseConfigFile(filename string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := Config{}
	if err := json.Unmarshal(configBytes, &c); err != nil {
		panic(err)
	}
	return &c, nil

}
