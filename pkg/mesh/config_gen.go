package mesh

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateMeshConfigs(mesh *MeshConfig, outDir string) error {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	for _, localNode := range mesh.Nodes {
		var sb strings.Builder

		// [Interface]
		sb.WriteString(fmt.Sprintf("# WireGuard Mesh Config for node: %s\n", localNode.Name))
		sb.WriteString("[Interface]\n")
		sb.WriteString(fmt.Sprintf("Address = %s\n", localNode.MeshIP))
		sb.WriteString(fmt.Sprintf("PrivateKey = %s\n", localNode.PrivateKey))
		if localNode.ListenPort > 0 {
			sb.WriteString(fmt.Sprintf("ListenPort = %d\n", localNode.ListenPort))
		}
		sb.WriteString("\n")

		// [Peer] for every other node in the mesh
		for _, peerNode := range mesh.Nodes {
			if peerNode.Name == localNode.Name {
				continue
			}

			sb.WriteString(fmt.Sprintf("# Peer: %s\n", peerNode.Name))
			sb.WriteString("[Peer]\n")
			sb.WriteString(fmt.Sprintf("PublicKey = %s\n", peerNode.PublicKey))
			sb.WriteString(fmt.Sprintf("AllowedIPs = %s\n", peerNode.MeshIP))
			if peerNode.Endpoint != "" {
				sb.WriteString(fmt.Sprintf("Endpoint = %s\n", peerNode.Endpoint))
				sb.WriteString("PersistentKeepalive = 25\n")
			}
			sb.WriteString("\n")
		}

		filePath := filepath.Join(outDir, fmt.Sprintf("%s.conf", localNode.Name))
		if err := os.WriteFile(filePath, []byte(sb.String()), 0600); err != nil {
			return err
		}
	}

	return nil
}
