package mesh

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadMeshConfig_ValidAndErrors(t *testing.T) {
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "mesh.yaml")
	yamlContent := "network_name: prod\nsubnet: 10.100.0.0/24\nnodes:\n  - name: gw\n    mesh_ip: 10.100.0.1/24\n    listen_port: 51820\n    private_key: priv1\n    public_key: pub1\n"
	if err := os.WriteFile(cfgPath, []byte(yamlContent), 0o600); err != nil {
		t.Fatalf("failed to write yaml: %v", err)
	}

	cfg, err := LoadMeshConfig(cfgPath)
	if err != nil {
		t.Fatalf("unexpected LoadMeshConfig error: %v", err)
	}
	if cfg.NetworkName != "prod" || len(cfg.Nodes) != 1 || cfg.Nodes[0].Name != "gw" {
		t.Fatalf("unexpected loaded mesh config: %+v", cfg)
	}

	if _, err := LoadMeshConfig(filepath.Join(dir, "missing.yaml")); err == nil {
		t.Fatal("expected error for missing config file")
	}
}

func TestGenerateMeshConfigs_CreatesConfigsAndRejectsPathTraversal(t *testing.T) {
	outDir := t.TempDir()
	m := &MeshConfig{
		NetworkName: "test",
		Nodes: []Node{
			{
				Name:       "alpha",
				MeshIP:     "10.100.0.1/24",
				Endpoint:   "203.0.113.10:51820",
				ListenPort: 51820,
				PrivateKey: "privA",
				PublicKey:  "pubA",
			},
			{
				Name:       "beta",
				MeshIP:     "10.100.0.2/24",
				PrivateKey: "privB",
				PublicKey:  "pubB",
			},
		},
	}

	if err := GenerateMeshConfigs(m, outDir); err != nil {
		t.Fatalf("unexpected GenerateMeshConfigs error: %v", err)
	}

	alphaBytes, err := os.ReadFile(filepath.Join(outDir, "alpha.conf"))
	if err != nil {
		t.Fatalf("failed reading alpha.conf: %v", err)
	}
	alphaCfg := string(alphaBytes)
	if !strings.Contains(alphaCfg, "PrivateKey = privA") || !strings.Contains(alphaCfg, "PublicKey = pubB") {
		t.Fatalf("alpha.conf missing expected keys:\n%s", alphaCfg)
	}
	if strings.Contains(alphaCfg, "PublicKey = pubA") {
		t.Fatalf("alpha.conf should not include itself as a peer:\n%s", alphaCfg)
	}

	betaBytes, err := os.ReadFile(filepath.Join(outDir, "beta.conf"))
	if err != nil {
		t.Fatalf("failed reading beta.conf: %v", err)
	}
	betaCfg := string(betaBytes)
	if !strings.Contains(betaCfg, "Endpoint = 203.0.113.10:51820") || !strings.Contains(betaCfg, "PersistentKeepalive = 25") {
		t.Fatalf("beta.conf missing endpoint or keepalive for peer alpha:\n%s", betaCfg)
	}

	bad := &MeshConfig{Nodes: []Node{{Name: "../evil", MeshIP: "10.0.0.1/32"}}}
	if err := GenerateMeshConfigs(bad, outDir); err == nil {
		t.Fatal("expected path traversal node name to be rejected")
	}
}
