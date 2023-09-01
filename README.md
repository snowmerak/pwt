# pwt

PWT is a short for "Protobuf Web Token". It is a simple token format based on [Protobuf](https://developers.google.com/protocol-buffers/) and [JWT](https://jwt.io/).

## install

```bash
go get github.com/snowmerak/pwt
```

## usage

### hmac with blake3-256

```go
package main

import (
	"fmt"
	"github.com/snowmerak/pwt/gen/grpc/model/token"
	"github.com/snowmerak/pwt/lib/pwt"
)

func main() {
	t := pwt.New().SetType(token.Type_REFRESH).SetAlgorithm(token.SignatureAlgorithm_HMAC, token.HashAlgorithm_BLAKE3_256)
	if err := pwt.SetPayload[int64](t, "test", 123); err != nil {
		panic(err)
	}

	secret := []byte("test")

	if _, err := t.Sign(secret); err != nil {
		panic(err)
	}

	wt, err := t.Export()
	if err != nil {
		panic(err)
	}

	fmt.Println(wt)

	at, err := pwt.Import(wt)
	if err != nil {
		panic(err)
	}

	if err := at.Verify(secret); err != nil {
		panic(err)
	}
	
	v, err := pwt.GetPayload[int64](at, "test")
	if err != nil {
        panic(err)
    }
	
	fmt.Println(v)
}

// CggKBAgBEAYQARIUCgR0ZXN0EgwIBBIIAAAAAAAAAHsaIOzezSm6d2UCzzR82tM2yXMmYCFuiTgkO4c5PYYQ9wSE
// 123
```

## ed25519 with blake3-512

```go
package main

import (
	"fmt"
	"github.com/snowmerak/pwt/gen/grpc/model/token"
	"github.com/snowmerak/pwt/lib/pwt"
	"lukechampine.com/blake3"
)

func main() {
	t := pwt.New().SetType(token.Type_REFRESH).SetAlgorithm(token.SignatureAlgorithm_ED25519, token.HashAlgorithm_BLAKE3_512)
	if err := pwt.SetPayload(t, "test", 3425263454); err != nil {
		panic(err)
	}

	secret := blake3.Sum256([]byte("test"))

	pubKey, privKey, err := pwt.NewEd25519KeyFromSeed(secret[:])
	if err != nil {
		panic(err)
	}

	pubKeyBuf, err := pwt.MakeEd25519PublicKeyToPem(pubKey)
	if err != nil {
		panic(err)
	}

	privKeyBuf, err := pwt.MakeEd25519PrivateKeyToPem(privKey)
	if err != nil {
		panic(err)
	}

	if _, err := t.Sign(privKeyBuf); err != nil {
		panic(err)
	}

	wt, err := t.Export()
	if err != nil {
		panic(err)
	}

	fmt.Println(wt)

	at, err := pwt.Import(wt)
	if err != nil {
		panic(err)
	}

	if err := at.Verify(pubKeyBuf); err != nil {
		panic(err)
	}

	v, err := pwt.GetPayload[int](at, "test")
	if err != nil {
		panic(err)
	}

	fmt.Println(v)
}

// CggKBAgEEAcQARISCgR0ZXN0EgoSCAAAAADMKV9eGkBEwpZdesOqzH69CJhehQDsIW8BAvPmoYGUgNapBiNTSP1HtaZCgPJA6YzkHHb571385qlNPBaqtWccylwR6IwA
// 3425263454
```

### P521 with blake2b-384

```go
package main

import (
	"crypto/elliptic"
	"fmt"
	"github.com/snowmerak/pwt/gen/grpc/model/token"
	"github.com/snowmerak/pwt/lib/pwt"
)

func main() {
	t := pwt.New().SetType(token.Type_REFRESH).SetAlgorithm(token.SignatureAlgorithm_ECDSA, token.HashAlgorithm_BLAKE2B_384)
	if err := pwt.SetPayload(t, "test", []byte("hello, world!")); err != nil {
		panic(err)
	}

	secret := []byte("1234567890")

	privKey, err := pwt.NewEcdsaKeyFromSeed(elliptic.P521(), secret)
	if err != nil {
		panic(err)
	}
	pubKey := &privKey.PublicKey

	pubKeyBuf, err := pwt.MakeECPublicKeyToPem(pubKey)
	if err != nil {
		panic(err)
	}

	privKeyBuf, err := pwt.MakeECPrivateKeyToPem(privKey)
	if err != nil {
		panic(err)
	}

	if _, err := t.Sign(privKeyBuf); err != nil {
		panic(err)
	}

	wt, err := t.Export()
	if err != nil {
		panic(err)
	}

	fmt.Println(wt)

	at, err := pwt.Import(wt)
	if err != nil {
		panic(err)
	}

	if err := at.Verify(pubKeyBuf); err != nil {
		panic(err)
	}

	v, err := pwt.GetPayload[[]byte](at, "test")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(v))
}

// CggKBAgDEAIQARIdCgR0ZXN0EhUIDRIRAAAADWhlbGxvLCB3b3JsZCEaiwEwgYgCQgGH2Kiam4wJRvueFneSehAehNmX0i46agwQgx668denyR2OWghZGG1PEgVWgyvYol67tsRZ9ktuz5MlZDoJ4FXpRwJCASNPqnvMsqgYYwjKPtSnhnRg74UEVqANzXUgsTvTbyHzufv3Dt3r4n5Qalu2jClNKK_qAncsHV1CzFN8UrPfwPHS
// hello, world!
```

## types

```protobuf
enum PayloadType {
    INT = 0; // 64 bit integer, int in golang
    INT8 = 1; // 8 bit integer, int8 in golang
    INT16 = 2; // 16 bit integer, int16 in golang
    INT32 = 3; // 32 bit integer, int32 in golang
    INT64 = 4; // 64 bit integer, int64 in golang
    UINT = 5; // 64 bit unsigned integer, uint in golang
    UINT8 = 6; // 8 bit unsigned integer, uint8 in golang
    UINT16 = 7; // 16 bit unsigned integer, uint16 in golang
    UINT32 = 8; // 32 bit unsigned integer, uint32 in golang
    UINT64 = 9; // 64 bit unsigned integer, uint64 in golang
    FLOAT32 = 10; // 32 bit float, float32 in golang
    FLOAT64 = 11; // 64 bit float, float64 in golang
    STRING = 12; // string
    BYTES = 13; // []byte
    BOOL = 14; // bool
    MESSAGE = 15; // protobuf message
}
```

PWT is can set/get above types and alias types of them.
