package pwt

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/snowmerak/pwt/lib/seeder"
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

func MakeECPrivateKeyToPem(pk *ecdsa.PrivateKey) ([]byte, error) {
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

func MakeECPublicKeyToPem(pk *ecdsa.PublicKey) ([]byte, error) {
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

func getEd25519PrivateKeyFromPem(secret []byte) (ed25519.PrivateKey, error) {
	block, _ := pem.Decode(secret)
	if block == nil {
		return nil, errors.New("invalid secret")
	}

	if block.Type != "PRIVATE KEY" {
		return nil, errors.New("secret is not private key")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey.(ed25519.PrivateKey), nil
}

func MakeEd25519PrivateKeyToPem(pk ed25519.PrivateKey) ([]byte, error) {
	block, err := x509.MarshalPKCS8PrivateKey(pk)
	if err != nil {
		return nil, err
	}

	rs := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: block,
	})

	return rs, nil
}

func getEd25519PublicKeyFromPem(secret []byte) (ed25519.PublicKey, error) {
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

	return parsedKey.(ed25519.PublicKey), nil
}

func MakeEd25519PublicKeyToPem(pk ed25519.PublicKey) ([]byte, error) {
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

func NewEcdsaKey(curve elliptic.Curve) (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(curve, rand.Reader)
}

func NewEcdsaKeyFromSeed(curve elliptic.Curve, seed []byte) (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(curve, seeder.New(seed))
}

func NewEd25519Key() (ed25519.PublicKey, ed25519.PrivateKey, error) {
	return ed25519.GenerateKey(rand.Reader)
}

func NewEd25519KeyFromSeed(seed []byte) (ed25519.PublicKey, ed25519.PrivateKey, error) {
	privateKey := ed25519.NewKeyFromSeed(seed)
	return privateKey.Public().(ed25519.PublicKey), privateKey, nil
}
