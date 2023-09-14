package pwt

import (
	"encoding/binary"
	"errors"
	"github.com/snowmerak/pwt/gen/grpc/model/token"
	"google.golang.org/protobuf/proto"
	"math"
	"reflect"
)

func SetPayload[T any](pwt *PWT, key string, value T) error {
	switch v := any(value).(type) {
	case int:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_INT,
			Value: buf,
		}
	case int8:
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_INT8,
			Value: []byte{byte(v)},
		}
	case int16:
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_INT16,
			Value: buf,
		}
	case int32:
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_INT32,
			Value: buf,
		}
	case int64:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_INT64,
			Value: buf,
		}
	case uint:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_UINT,
			Value: buf,
		}
	case uint8:
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_UINT8,
			Value: []byte{byte(v)},
		}
	case uint16:
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_UINT16,
			Value: buf,
		}
	case uint32:
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_UINT32,
			Value: buf,
		}
	case uint64:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_UINT64,
			Value: buf,
		}
	case float32:
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, math.Float32bits(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_FLOAT32,
			Value: buf,
		}
	case float64:
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, math.Float64bits(v))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_FLOAT64,
			Value: buf,
		}
	case string:
		buf := make([]byte, 4, 4+len(v))
		binary.BigEndian.PutUint32(buf, uint32(len(v)))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_STRING,
			Value: append(buf, []byte(v)...),
		}
	case []byte:
		buf := make([]byte, 4, 4+len(v))
		binary.BigEndian.PutUint32(buf, uint32(len(v)))
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_BYTES,
			Value: append(buf, v...),
		}
	case bool:
		a := byte(0)
		if v {
			a = 1
		}
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_BOOL,
			Value: []byte{byte(a)},
		}
	case proto.Message:
		buf, err := proto.Marshal(v)
		if err != nil {
			return err
		}
		pwt.token.Payloads[key] = &token.Payload{
			Type:  token.PayloadType_MESSAGE,
			Value: buf,
		}
	default:
		rf := reflect.ValueOf(v)
		switch rf.Kind() {
		case reflect.Int:
			if err := SetPayload(pwt, key, int(rf.Int())); err != nil {
				return err
			}
		case reflect.Int8:
			if err := SetPayload(pwt, key, int8(rf.Int())); err != nil {
				return err
			}
		case reflect.Int16:
			if err := SetPayload(pwt, key, int16(rf.Int())); err != nil {
				return err
			}
		case reflect.Int32:
			if err := SetPayload(pwt, key, int32(rf.Int())); err != nil {
				return err
			}
		case reflect.Int64:
			if err := SetPayload(pwt, key, rf.Int()); err != nil {
				return err
			}
		case reflect.Uint:
			if err := SetPayload(pwt, key, uint(rf.Uint())); err != nil {
				return err
			}
		case reflect.Uint8:
			if err := SetPayload(pwt, key, uint8(rf.Uint())); err != nil {
				return err
			}
		case reflect.Uint16:
			if err := SetPayload(pwt, key, uint16(rf.Uint())); err != nil {
				return err
			}
		case reflect.Uint32:
			if err := SetPayload(pwt, key, uint32(rf.Uint())); err != nil {
				return err
			}
		case reflect.Uint64:
			if err := SetPayload(pwt, key, rf.Uint()); err != nil {
				return err
			}
		case reflect.Float32:
			if err := SetPayload(pwt, key, float32(rf.Float())); err != nil {
				return err
			}
		case reflect.Float64:
			if err := SetPayload(pwt, key, rf.Float()); err != nil {
				return err
			}
		case reflect.String:
			if err := SetPayload(pwt, key, rf.String()); err != nil {
				return err
			}
		case reflect.Slice:
			switch rf.CanConvert(reflect.SliceOf(reflect.TypeOf(byte(0)))) {
			case true:
				if err := SetPayload(pwt, key, rf.Bytes()); err != nil {
					return err
				}
			default:
				return errInvalidPayloadType
			}
		case reflect.Bool:
			if err := SetPayload(pwt, key, rf.Bool()); err != nil {
				return err
			}
		default:
			return errInvalidPayloadType
		}
	}

	return nil
}

func GetPayload[T any](pwt *PWT, key string) (T, error) {
	if pwt == nil || pwt.token == nil || pwt.token.Payloads == nil {
		return any(*new(T)).(T), errors.New("invalid pwt")
	}

	r := any(*new(T))
	payload, ok := pwt.token.Payloads[key]
	if !ok {
		return r.(T), errors.New("invalid payload key")
	}
	switch r.(type) {
	case int:
		if payload.Type != token.PayloadType_INT {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 8 {
			return r.(T), errInvalidPayloadLength
		}
		r = int(binary.BigEndian.Uint64(payload.Value))
	case int8:
		if payload.Type != token.PayloadType_INT8 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 1 {
			return r.(T), errInvalidPayloadLength
		}
		r = int8(payload.Value[0])
	case int16:
		if payload.Type != token.PayloadType_INT16 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 2 {
			return r.(T), errInvalidPayloadLength
		}
		r = int16(binary.BigEndian.Uint16(payload.Value))
	case int32:
		if payload.Type != token.PayloadType_INT32 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 4 {
			return r.(T), errInvalidPayloadLength
		}
		r = int32(binary.BigEndian.Uint32(payload.Value))
	case int64:
		if payload.Type != token.PayloadType_INT64 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 8 {
			return r.(T), errInvalidPayloadLength
		}
		r = int64(binary.BigEndian.Uint64(payload.Value))
	case uint:
		if payload.Type != token.PayloadType_UINT {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 8 {
			return r.(T), errInvalidPayloadLength
		}
		r = uint(binary.BigEndian.Uint64(payload.Value))
	case uint8:
		if payload.Type != token.PayloadType_UINT8 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 1 {
			return r.(T), errInvalidPayloadLength
		}
		r = uint8(payload.Value[0])
	case uint16:
		if payload.Type != token.PayloadType_UINT16 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 2 {
			return r.(T), errInvalidPayloadLength
		}
		r = uint16(binary.BigEndian.Uint16(payload.Value))
	case uint32:
		if payload.Type != token.PayloadType_UINT32 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 4 {
			return r.(T), errInvalidPayloadLength
		}
		r = uint32(binary.BigEndian.Uint32(payload.Value))
	case uint64:
		if payload.Type != token.PayloadType_UINT64 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 8 {
			return r.(T), errInvalidPayloadLength
		}
		r = uint64(binary.BigEndian.Uint64(payload.Value))
	case float32:
		if payload.Type != token.PayloadType_FLOAT32 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 4 {
			return r.(T), errInvalidPayloadLength
		}
		r = math.Float32frombits(binary.BigEndian.Uint32(payload.Value))
	case float64:
		if payload.Type != token.PayloadType_FLOAT64 {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 8 {
			return r.(T), errInvalidPayloadLength
		}
		r = math.Float64frombits(binary.BigEndian.Uint64(payload.Value))
	case string:
		if payload.Type != token.PayloadType_STRING {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) < 4 {
			return r.(T), errInvalidPayloadLength
		}
		length := binary.BigEndian.Uint32(payload.Value)
		if len(payload.Value) != int(length)+4 {
			return r.(T), errInvalidPayloadLength
		}
		r = string(payload.Value[4:])
	case []byte:
		if payload.Type != token.PayloadType_BYTES {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) < 4 {
			return r.(T), errInvalidPayloadLength
		}
		length := binary.BigEndian.Uint32(payload.Value)
		if len(payload.Value) != int(length)+4 {
			return r.(T), errInvalidPayloadLength
		}
		r = payload.Value[4:]
	case bool:
		if payload.Type != token.PayloadType_BOOL {
			return r.(T), errInvalidPayloadType
		}
		if len(payload.Value) != 1 {
			return r.(T), errInvalidPayloadLength
		}
		r = payload.Value[0] != 0
	case proto.Message:
		if payload.Type != token.PayloadType_MESSAGE {
			return r.(T), errInvalidPayloadType
		}
		m := any(r).(proto.Message)
		if err := proto.Unmarshal(payload.Value, m); err != nil {
			return r.(T), err
		}
		r = m
	default:
		rf := reflect.ValueOf(r)
		switch rf.Kind() {
		case reflect.Int:
			v, err := GetPayload[int](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Int8:
			v, err := GetPayload[int8](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Int16:
			v, err := GetPayload[int16](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Int32:
			v, err := GetPayload[int32](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Int64:
			v, err := GetPayload[int64](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Uint:
			v, err := GetPayload[uint](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Uint8:
			v, err := GetPayload[uint8](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Uint16:
			v, err := GetPayload[uint16](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Uint32:
			v, err := GetPayload[uint32](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Uint64:
			v, err := GetPayload[uint64](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Float32:
			v, err := GetPayload[float32](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Float64:
			v, err := GetPayload[float64](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.String:
			v, err := GetPayload[string](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		case reflect.Slice:
			switch rf.CanConvert(reflect.SliceOf(reflect.TypeOf(byte(0)))) {
			case true:
				v, err := GetPayload[[]byte](pwt, key)
				if err != nil {
					return r.(T), err
				}
				r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
			default:
				return r.(T), errInvalidPayloadType
			}
		case reflect.Bool:
			v, err := GetPayload[bool](pwt, key)
			if err != nil {
				return r.(T), err
			}
			r = reflect.ValueOf(v).Convert(rf.Type()).Interface()
		default:
			return r.(T), errInvalidPayloadType
		}
	}

	return r.(T), nil
}
