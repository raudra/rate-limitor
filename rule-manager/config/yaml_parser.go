package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type YAMLParser struct {
	filePath   string
	dataStruct interface{}
}

func (self *YAMLParser) Parse() error {
	yamlfile, err := ioutil.ReadFile(self.filePath)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlfile, &self.dataStruct)

	if err != nil {
		return err
	}

	return nil
}
