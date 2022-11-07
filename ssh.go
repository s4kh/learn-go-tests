package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
)

type keyPair struct {
	publicKey  string
	privateKey string
}

func generateKeyPair() (*keyPair, error) {
	bitSize := 4096

	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	sshPrivateKey := string(privBytes)

	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}

	sshPubKey := string(ssh.MarshalAuthorizedKey(pub))

	return &keyPair{sshPubKey, sshPrivateKey}, nil
}
