package core

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var SharedConfig *config

type config struct {
	Environment string `yaml:"environment"`
	Redis       struct {
		Host string
		Port uint16
	}
	Database struct {
		Host     string
		Port     uint16
		Username string
		Password string
	}
}

func InitConfig(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		SharedLogger.Infof("ReadFile err: %s", err)
	}

	err = yaml.Unmarshal(data, &SharedConfig)
	if err != nil {
		SharedLogger.Infof("Unmarshal data: %s err: %s", data, err)
	}
	SharedLogger.Infof("Unmarshal data: %s conf: %s", data, SharedConfig)
}
