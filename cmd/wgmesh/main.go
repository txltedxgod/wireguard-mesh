package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/txltedxgod/wireguard-mesh/pkg/keys"
	"github.com/txltedxgod/wireguard-mesh/pkg/mesh"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "genkey":
		pair, err := keys.GenerateKeyPair()
		if err != nil {
			log.Fatalf("Error generating keypair: %v", err)
		}
		fmt.Printf("PrivateKey = %s\nPublicKey  = %s\n", pair.PrivateKey, pair.PublicKey)

	case "generate":
		genCmd := flag.NewFlagSet("generate", flag.ExitOnError)
		configPath := genCmd.String("config", "mesh.yaml", "Mesh definition yaml")
		outDir := genCmd.String("out", "configs", "Output directory for wg0.conf files")
		_ = genCmd.Parse(os.Args[2:])

		m, err := mesh.LoadMeshConfig(*configPath)
		if err != nil {
			log.Fatalf("Error loading mesh config: %v", err)
		}

		if err := mesh.GenerateMeshConfigs(m, *outDir); err != nil {
			log.Fatalf("Error generating configs: %v", err)
		}
		fmt.Printf("[+] Successfully generated %d WireGuard mesh configs in '%s/'\n", len(m.Nodes), *outDir)

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("wireguard-mesh - Multi-Cloud WireGuard Mesh Network Topology Generator")
	fmt.Println("\nCommands:")
	fmt.Println("  genkey              Generate a new Curve25519 WireGuard keypair")
	fmt.Println("  generate -config=.. Generate full mesh .conf files for all nodes")
}
