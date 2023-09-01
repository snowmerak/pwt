package pwt

import (
	"encoding/base64"
	"github.com/snowmerak/pwt/gen/grpc/model/token"
	"google.golang.org/protobuf/proto"
)

func (p *PWT) Export() (string, error) {
	v, err := proto.Marshal(p.token)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(v), nil
}

func Import(data string) (*PWT, error) {
	v, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	
	t := &token.Token{}
	if err := proto.Unmarshal(v, t); err != nil {
		return nil, err
	}

	return From(t), nil
}
