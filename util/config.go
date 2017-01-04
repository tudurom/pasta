package util

import (
	"io/ioutil"

	"github.com/Sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	DatabasePath  string
	StoragePath   string
	PasteEndpoint string
}

var GlobalConfig Config

func (c *Config) ReadConfig(fpath string) {
	logrus.Printf("Reading config file from %s\n", fpath)
	f, err := ioutil.ReadFile(fpath)
	Crash("Couldn't read config file", err)

	yaml.Unmarshal(f, c)
}
