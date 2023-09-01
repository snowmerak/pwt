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
}

// CggKBAgBEAYQARIUCgR0ZXN0EgwIBBIIAAAAAAAAAHsaIOzezSm6d2UCzzR82tM2yXMmYCFuiTgkO4c5PYYQ9wSE
```
