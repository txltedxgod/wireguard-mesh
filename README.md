# wireguard-mesh

> Automated **WireGuard full-mesh network topology** & configuration generator for multi-cloud node interconnects written in **Go**.

[![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=flat-square&logo=go)](https://golang.org)
[![WireGuard](https://img.shields.io/badge/VPN-WireGuard-88171A?style=flat-square&logo=wireguard)](https://wireguard.com)
[![Curve25519](https://img.shields.io/badge/Crypto-Curve25519-green?style=flat-square)](https://pkg.go.dev/golang.org/x/crypto/curve25519)
[![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)](LICENSE)

`#wireguard` `#mesh-vpn` `#networking` `#devops` `#multi-cloud` `#security` `#golang`

---

## Features

- **Full-Mesh Connectivity:** Automatically computes peer configurations for all nodes in the mesh topology.
- **Built-in Key Generator:** Generates cryptographic Curve25519 private/public keys without needing `wg genkey` pre-installed.
- **Multi-Cloud Ready:** Connect AWS, GCP, Hetzner, and on-premise bare-metal nodes over an encrypted tunnel.

## Quick Start

```bash
# 1. Generate keypairs for your nodes
go run cmd/wgmesh/main.go genkey

# 2. Configure mesh.yaml with your server IPs and keys

# 3. Generate all wg0.conf files
go run cmd/wgmesh/main.go generate -config=mesh.example.yaml -out=./configs

# 4. Copy each config to respective server (/etc/wireguard/wg0.conf) and start
wg-quick up wg0
```
