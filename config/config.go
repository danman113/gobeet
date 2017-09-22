package config

import (
	"io/ioutil"
	// 	. "github.com/danman113/gobeet/re"
	"encoding/json"
	"github.com/danman113/gobeet/site"
)

type pingable interface {
}

type Config struct {
	Sites []site.Website `json: sites`
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
