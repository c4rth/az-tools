package model

import (
	"gopkg.in/yaml.v2"
	"os"
)

const ConfigFilename = "az-tools.yaml"

type Config struct {
	Subscriptions []struct {
		Name           string `yaml:"name"`
		ResourceGroups []struct {
			Name string `yaml:"name"`
			Aks  []struct {
				Name string `yaml:"name"`
			} `yaml:"aks"`
		} `yaml:"resource-groups"`
	} `yaml:"subscriptions"`
}

func ReadConfig() (Config, error) {
	b, err := os.ReadFile(ConfigFilename)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
