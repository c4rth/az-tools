package config

import (
	"gopkg.in/yaml.v2"
	"os"

	"k8switch/internal/types"
)

func ReadConfig() (types.Config, bool, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return types.Config{}, false, nil
	}
	fullConfigPath := homeDir + types.ConfigFileDir + types.ConfigFilename
	if isFileExists, _ := exists(fullConfigPath); !isFileExists {
		return createDummyConfig(homeDir)
	} else {
		b, err := os.ReadFile(fullConfigPath)
		if err != nil {
			return types.Config{}, false, err
		}
		var config types.Config
		err = yaml.Unmarshal(b, &config)
		if err != nil {
			return types.Config{}, false, err
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

func createDummyConfig(homeDir string) (types.Config, bool, error) {
	if isDirExists, _ := exists(homeDir + types.ConfigFileDir); !isDirExists {
		err := os.Mkdir(homeDir+types.ConfigFileDir, 0777)
		if err != nil {
			return types.Config{}, false, err
		}
	}
	dummyConfig := getDummyConfig()
	dummyB, err := yaml.Marshal(dummyConfig)
	err = os.WriteFile(homeDir+types.ConfigFileDir+types.ConfigFilename, dummyB, 0644)
	if err != nil {
		return types.Config{}, false, err
	}
	return dummyConfig, true, nil
}

func getDummyConfig() types.Config {
	return types.Config{
		Ranchers: []types.Rancher{
			{Name: "rancher1"},
			{Name: "rancher2"},
			{Name: "rancher3"},
		},
		Subscriptions: []types.Subscription{
			{
				Name: "subscription1",
				ResourceGroups: []types.ResourceGroup{
					{
						Name: "resource-group-1",
						Aks: []types.Aks{
							{Name: "aks11"},
							{Name: "aks12"},
						},
					},
					{
						Name: "resource-group-2",
						Aks: []types.Aks{
							{Name: "aks21"},
							{Name: "aks22"},
						},
					},
				},
			},
			{
				Name: "subscription2",
				ResourceGroups: []types.ResourceGroup{
					{
						Name: "resource-group-3",
						Aks: []types.Aks{
							{Name: "aks31"},
							{Name: "aks32"},
						},
					},
					{
						Name: "resource-group-4",
						Aks: []types.Aks{
							{Name: "aks41"},
							{Name: "aks42"},
						},
					},
				},
			},
		},
	}
}
