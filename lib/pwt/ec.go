package pwt

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func getECPrivateKeyFromPem(secret []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(secret)
	if block == nil {
		return nil, errors.New("invalid secret")
	}

	if block.Type != "PRIVATE KEY" {
		return nil, errors.New("secret is not private key")
	}

	parsedKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey, nil
}

func makeECPrivateKeyToPem(pk *ecdsa.PrivateKey) ([]byte, error) {
	parsed, err := x509.MarshalECPrivateKey(pk)
	if err != nil {
		return nil, err
	}

	rs := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: parsed,
	})

	return rs, nil
}

func getECPublicKeyFromPem(secret []byte) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode(secret)
	if block == nil {
		return nil, errors.New("invalid secret")
	}

	if block.Type != "PUBLIC KEY" {
		return nil, errors.New("secret is not public key")
	}

	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey.(*ecdsa.PublicKey), nil
}

func makeECPublicKeyToPem(pk *ecdsa.PublicKey) ([]byte, error) {
	block, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		return nil, err
	}

	rs := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: block,
	})

	return rs, nil
}
