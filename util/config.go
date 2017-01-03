package util

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DatabasePath string
}

var GlobalConfig Config

func (c *Config) ReadConfig(fpath string) {
	f, err := ioutil.ReadFile(fpath)
	Crash("Couldn't read config file", err)

	yaml.Unmarshal(f, c)
}
