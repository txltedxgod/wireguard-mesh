package keys

import (
	"encoding/base64"
	"testing"
)

func TestGenerateKeyPair(t *testing.T) {
	k1, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("unexpected GenerateKeyPair error: %v", err)
	}
	k2, err := GenerateKeyPair()
	if err != nil {
		t.Fatalf("unexpected GenerateKeyPair error: %v", err)
	}

	if k1.PrivateKey == "" || k1.PublicKey == "" {
		t.Fatal("expected non-empty keys")
	}
	if k1.PrivateKey == k2.PrivateKey || k1.PublicKey == k2.PublicKey {
		t.Fatal("expected unique keypairs across calls")
	}

	priv, err := base64.StdEncoding.DecodeString(k1.PrivateKey)
	if err != nil || len(priv) != 32 {
		t.Fatalf("expected 32-byte base64 private key, got len=%d err=%v", len(priv), err)
	}
	if priv[0]&7 != 0 || priv[31]&128 != 0 || priv[31]&64 == 0 {
		t.Fatalf("private key is not properly WireGuard-clamped")
	}

	pub, err := base64.StdEncoding.DecodeString(k1.PublicKey)
	if err != nil || len(pub) != 32 {
		t.Fatalf("expected 32-byte base64 public key, got len=%d err=%v", len(pub), err)
	}
}
