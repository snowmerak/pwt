package pwt

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/snowmerak/pwt/gen/grpc/model/token"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
	"google.golang.org/protobuf/proto"
	"hash"
	"lukechampine.com/blake3"
	"sort"
)

func (p *PWT) Sign(secret []byte) (string, error) {
	hashed, err := p.makeWithoutSignature()
	if err != nil {
		return "", err
	}

	switch p.token.Header.Algorithm.Signature {
	case token.SignatureAlgorithm_NONE:
		p.token.Signature = hashed
	case token.SignatureAlgorithm_HMAC:
		p.token.Signature, err = p.signHMac(hashed, secret)
		if err != nil {
			return "", err
		}
	case token.SignatureAlgorithm_ECDSA:
		parsedKey, err := getECPrivateKeyFromPem(secret)
		if err != nil {
			return "", err
		}

		signature, err := ecdsa.SignASN1(rand.Reader, parsedKey, hashed)
		if err != nil {
			return "", err
		}

		p.token.Signature = signature
	case token.SignatureAlgorithm_ED25519:
		parsedKey, err := getEd25519PrivateKeyFromPem(secret)
		if err != nil {
			return "", err
		}

		p.token.Signature = ed25519.Sign(parsedKey, hashed)
	}

	buf, err := proto.Marshal(p.token)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(buf), nil
}

func (p *PWT) Verify(secret []byte) error {
	hashed, err := p.makeWithoutSignature()
	if err != nil {
		return err
	}

	switch p.token.Header.Algorithm.Signature {
	case token.SignatureAlgorithm_NONE:
		if !bytes.Equal(p.token.Signature, hashed) {
			return errors.New("invalid signature")
		}
	case token.SignatureAlgorithm_HMAC:
		signature, err := p.signHMac(hashed, secret)
		if err != nil {
			return err
		}

		if !hmac.Equal(p.token.Signature, signature) {
			return errors.New("invalid signature")
		}
	case token.SignatureAlgorithm_ECDSA:
		parsedKey, err := getECPublicKeyFromPem(secret)
		if err != nil {
			return err
		}

		if !ecdsa.VerifyASN1(parsedKey, hashed, p.token.Signature) {
			return errors.New("invalid signature")
		}
	case token.SignatureAlgorithm_ED25519:
		parsedKey, err := getEd25519PublicKeyFromPem(secret)
		if err != nil {
			return err
		}

		if !ed25519.Verify(parsedKey, hashed, p.token.Signature) {
			return err
		}
	}

	return nil
}

func (p *PWT) makeWithoutSignature() ([]byte, error) {
	if p.token == nil && p.token.Header == nil && p.token.Header.Algorithm == nil {
		return nil, errors.New("header is nil")
	}

	buffer := bytes.NewBuffer(nil)
	data, err := proto.Marshal(p.token.Header)
	if err != nil {
		return nil, err
	}
	buffer.Write(data)

	payloadKeys := make([]string, 0, len(p.token.Payloads))
	for k := range p.token.Payloads {
		payloadKeys = append(payloadKeys, k)
	}
	sort.Strings(payloadKeys)

	for _, k := range payloadKeys {
		buffer.Write([]byte(k))
		buffer.Write(p.token.Payloads[k].Value)
	}

	payloadsBuffer := buffer.Bytes()

	hashed := make([]byte, 0)

	switch p.token.Header.Algorithm.Hash {
	case token.HashAlgorithm_SHA3_256:
		b := sha3.Sum256(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_SHA3_384:
		b := sha3.Sum384(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_SHA3_512:
		b := sha3.Sum512(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE2B_256:
		b := blake2b.Sum256(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE2B_384:
		b := blake2b.Sum384(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE2B_512:
		b := blake2b.Sum512(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE2S_256:
		b := blake2b.Sum256(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE2S_384:
		b := blake2b.Sum384(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE2S_512:
		b := blake2b.Sum512(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE3_256:
		b := blake3.Sum256(payloadsBuffer)
		hashed = b[:]
	case token.HashAlgorithm_BLAKE3_512:
		b := blake3.Sum512(payloadsBuffer)
		hashed = b[:]
	default:
		return nil, errors.New("unknown hash algorithm")
	}

	return hashed, nil
}

func (p *PWT) signHMac(hashed []byte, secret []byte) ([]byte, error) {
	hs := hmac.New(func() hash.Hash {
		switch p.token.Header.Algorithm.Hash {
		case token.HashAlgorithm_SHA3_256:
			return sha3.New256()
		case token.HashAlgorithm_SHA3_384:
			return sha3.New384()
		case token.HashAlgorithm_SHA3_512:
			return sha3.New512()
		case token.HashAlgorithm_BLAKE2B_256:
			h, _ := blake2b.New256(nil)
			return h
		case token.HashAlgorithm_BLAKE2B_384:
			h, _ := blake2b.New384(nil)
			return h
		case token.HashAlgorithm_BLAKE2B_512:
			h, _ := blake2b.New512(nil)
			return h
		case token.HashAlgorithm_BLAKE2S_256:
			h, _ := blake2b.New256(nil)
			return h
		case token.HashAlgorithm_BLAKE2S_384:
			h, _ := blake2b.New384(nil)
			return h
		case token.HashAlgorithm_BLAKE2S_512:
			h, _ := blake2b.New512(nil)
			return h
		case token.HashAlgorithm_BLAKE3_256:
			return blake3.New(32, nil)
		case token.HashAlgorithm_BLAKE3_512:
			return blake3.New(64, nil)
		default:
			return nil
		}
	}, secret)

	if _, err := hs.Write(hashed); err != nil {
		return nil, err
	}

	return hs.Sum(nil), nil
}
