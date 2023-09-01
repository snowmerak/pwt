package pwt

import (
	"github.com/snowmerak/pwt/gen/grpc/model/token"
)

type PWT struct {
	token token.Token
}

func New() *PWT {
	return &PWT{token: token.Token{
		Header:    &token.Header{},
		Payloads:  make(map[string]*token.Payload),
		Signature: make([]byte, 0),
	}}
}
