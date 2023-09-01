package pwt

import "github.com/snowmerak/pwt/gen/grpc/model/token"

func (p *PWT) SetAlgorithm(sign token.SignatureAlgorithm, hash token.HashAlgorithm) *PWT {
	p.token.Header.Algorithm = &token.Algorithm{
		Signature: sign,
		Hash:      hash,
	}

	return p
}

func (p *PWT) SetType(tokenType token.Type) *PWT {
	p.token.Header.Type = tokenType

	return p
}
