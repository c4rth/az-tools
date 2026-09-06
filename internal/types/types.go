package types

const ConfigFileDir = "/.config/k8switch/"
const ConfigFilename = "config.yaml"

type Rancher struct {
	Name string `yaml:"name"`
}

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
	Ranchers      []Rancher      `yaml:"rancher,omitempty"`
	Subscriptions []Subscription `yaml:"azure,omitempty"`
}

type k8sKind int

const (
	KindRancher k8sKind = iota
	KindAks
)

type NodeReference struct {
	Subscription  string
	ResourceGroup string
	Cluster       string
	Kind          k8sKind
}