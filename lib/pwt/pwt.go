package pwt

import (
	"github.com/snowmerak/pwt/gen/grpc/model/token"
)

type PWT struct {
	token *token.Token
}

func New() *PWT {
	return &PWT{token: &token.Token{
		Header: &token.Header{
			Algorithm: &token.Algorithm{},
		},
		Payloads:  make(map[string]*token.Payload),
		Signature: make([]byte, 0),
	}}
}

func From(token *token.Token) *PWT {
	return &PWT{token: token}
}
