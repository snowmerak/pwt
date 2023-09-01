package pwt

import (
	"github.com/snowmerak/pwt/gen/grpc/model/token"
	"google.golang.org/protobuf/proto"
)

// type PayloadTypes interface {
//	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | []byte | bool | proto.Message
// }

func AccordPayloadType[T any](record token.PayloadType, value T) bool {
	switch any(value).(type) {
	case int:
		return record == token.PayloadType_INT
	case int8:
		return record == token.PayloadType_INT8
	case int16:
		return record == token.PayloadType_INT16
	case int32:
		return record == token.PayloadType_INT32
	case int64:
		return record == token.PayloadType_INT64
	case uint:
		return record == token.PayloadType_UINT
	case uint8:
		return record == token.PayloadType_UINT8
	case uint16:
		return record == token.PayloadType_UINT16
	case uint32:
		return record == token.PayloadType_UINT32
	case uint64:
		return record == token.PayloadType_UINT64
	case float32:
		return record == token.PayloadType_FLOAT32
	case float64:
		return record == token.PayloadType_FLOAT64
	case string:
		return record == token.PayloadType_STRING
	case []byte:
		return record == token.PayloadType_BYTES
	case bool:
		return record == token.PayloadType_BOOL
	case proto.Message:
		return record == token.PayloadType_MESSAGE
	}

	return false
}
