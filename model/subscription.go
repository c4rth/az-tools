package model

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Aks struct {
	Name          string `yaml:"name"`
	ResourceGroup string `yaml:"resource-group"`
}

type Subscription struct {
	Subscription string `yaml:"subscription"`
	Aks          []Aks
}

func ReadSubscriptions(filename string) ([]Subscription, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Subscription{}, err
	}
	var items []Subscription
	err = yaml.Unmarshal(b, &items)
	if err != nil {
		return []Subscription{}, err
	}
	return items, nil
}
