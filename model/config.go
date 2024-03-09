package model

import (
	"gopkg.in/yaml.v2"
	"os"
)

const ConfigFileDir = "/.az-tools/"
const ConfigFilename = "config.yaml"

type Aks struct {
	Name string `yaml:"name"`
}

type ResourceGroup struct {
	Name string `yaml:"name"`
	Aks  []Aks  `yaml:"aks"`
}

type Subscription struct {
	Name           string          `yaml:"name"`
	ResourceGroups []ResourceGroup `yaml:"resource-groups"`
}

type Config struct {
	Subscriptions []Subscription `yaml:"subscriptions"`
}

func ReadConfig() (Config, bool, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, false, nil
	}
	fullConfigPath := homeDir + ConfigFileDir + ConfigFilename
	if isFileExists, _ := exists(fullConfigPath); !isFileExists {
		return createDummyConfig(homeDir)
	} else {
		b, err := os.ReadFile(fullConfigPath)
		if err != nil {
			return Config{}, false, err
		}
		var config Config
		err = yaml.Unmarshal(b, &config)
		if err != nil {
			return Config{}, false, err
		}
		return config, false, nil
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createDummyConfig(homeDir string) (Config, bool, error) {
	if isDirExists, _ := exists(homeDir + ConfigFileDir); !isDirExists {
		err := os.Mkdir(homeDir+ConfigFileDir, 0777)
		if err != nil {
			return Config{}, false, err
		}
	}
	dummyConfig := getDummyConfig()
	dummyB, err := yaml.Marshal(dummyConfig)
	err = os.WriteFile(homeDir+ConfigFileDir+ConfigFilename, dummyB, 0644)
	if err != nil {
		return Config{}, false, err
	}
	return dummyConfig, true, nil
}

func getDummyConfig() Config {
	return Config{
		[]Subscription{
			{
				Name: "subscription1",
				ResourceGroups: []ResourceGroup{
					{
						Name: "resource-group-1",
						Aks: []Aks{
							{Name: "aks11"},
							{Name: "aks12"},
						},
					},
					{
						Name: "resource-group-2",
						Aks: []Aks{
							{Name: "aks21"},
							{Name: "aks22"},
						},
					},
				},
			},
			{
				Name: "subscription2",
				ResourceGroups: []ResourceGroup{
					{
						Name: "resource-group-3",
						Aks: []Aks{
							{Name: "aks31"},
							{Name: "aks32"},
						},
					},
					{
						Name: "resource-group-4",
						Aks: []Aks{
							{Name: "aks41"},
							{Name: "aks42"},
						},
					},
				},
			},
		},
	}
}
