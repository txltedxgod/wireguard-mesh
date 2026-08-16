package mesh

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Node struct {
	Name       string `yaml:"name"`
	MeshIP     string `yaml:"mesh_ip"`     // e.g. "10.100.0.1/24"
	Endpoint   string `yaml:"endpoint"`    // e.g. "203.0.113.10:51820"
	ListenPort int    `yaml:"listen_port"` // e.g. 51820
	PrivateKey string `yaml:"private_key"`
	PublicKey  string `yaml:"public_key"`
}

type MeshConfig struct {
	NetworkName string `yaml:"network_name"`
	Subnet      string `yaml:"subnet"` // e.g. "10.100.0.0/24"
	Nodes       []Node `yaml:"nodes"`
}

func LoadMeshConfig(path string) (*MeshConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg MeshConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
