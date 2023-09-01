// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: model/token/token.proto

package token

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignatureAlgorithm int32

const (
	SignatureAlgorithm_NONE    SignatureAlgorithm = 0
	SignatureAlgorithm_HMAC    SignatureAlgorithm = 1
	SignatureAlgorithm_RSA     SignatureAlgorithm = 2
	SignatureAlgorithm_ECDSA   SignatureAlgorithm = 3
	SignatureAlgorithm_ED25519 SignatureAlgorithm = 4
)

// Enum value maps for SignatureAlgorithm.
var (
	SignatureAlgorithm_name = map[int32]string{
		0: "NONE",
		1: "HMAC",
		2: "RSA",
		3: "ECDSA",
		4: "ED25519",
	}
	SignatureAlgorithm_value = map[string]int32{
		"NONE":    0,
		"HMAC":    1,
		"RSA":     2,
		"ECDSA":   3,
		"ED25519": 4,
	}
)

func (x SignatureAlgorithm) Enum() *SignatureAlgorithm {
	p := new(SignatureAlgorithm)
	*p = x
	return p
}

func (x SignatureAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignatureAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_model_token_token_proto_enumTypes[0].Descriptor()
}

func (SignatureAlgorithm) Type() protoreflect.EnumType {
	return &file_model_token_token_proto_enumTypes[0]
}

func (x SignatureAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignatureAlgorithm.Descriptor instead.
func (SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{0}
}

type HashAlgorithm int32

const (
	HashAlgorithm_SHA3_384    HashAlgorithm = 0
	HashAlgorithm_SHA3_512    HashAlgorithm = 1
	HashAlgorithm_BLAKE2B_384 HashAlgorithm = 2
	HashAlgorithm_BLAKE2B_512 HashAlgorithm = 3
	HashAlgorithm_BLAKE2S_384 HashAlgorithm = 4
	HashAlgorithm_BLAKE2S_512 HashAlgorithm = 5
	HashAlgorithm_BLAKE3_256  HashAlgorithm = 6
	HashAlgorithm_BLAKE3_512  HashAlgorithm = 7
)

// Enum value maps for HashAlgorithm.
var (
	HashAlgorithm_name = map[int32]string{
		0: "SHA3_384",
		1: "SHA3_512",
		2: "BLAKE2B_384",
		3: "BLAKE2B_512",
		4: "BLAKE2S_384",
		5: "BLAKE2S_512",
		6: "BLAKE3_256",
		7: "BLAKE3_512",
	}
	HashAlgorithm_value = map[string]int32{
		"SHA3_384":    0,
		"SHA3_512":    1,
		"BLAKE2B_384": 2,
		"BLAKE2B_512": 3,
		"BLAKE2S_384": 4,
		"BLAKE2S_512": 5,
		"BLAKE3_256":  6,
		"BLAKE3_512":  7,
	}
)

func (x HashAlgorithm) Enum() *HashAlgorithm {
	p := new(HashAlgorithm)
	*p = x
	return p
}

func (x HashAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_model_token_token_proto_enumTypes[1].Descriptor()
}

func (HashAlgorithm) Type() protoreflect.EnumType {
	return &file_model_token_token_proto_enumTypes[1]
}

func (x HashAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashAlgorithm.Descriptor instead.
func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{1}
}

type Type int32

const (
	Type_ACCESS             Type = 0
	Type_REFRESH            Type = 1
	Type_RESET_PASSWORD     Type = 2
	Type_VERIFY_EMAIL       Type = 3
	Type_VERIFY_PHONE       Type = 4
	Type_VERIFY_IDENTITY    Type = 5
	Type_VERIFY_ADDRESS     Type = 6
	Type_VERIFY_DOCUMENT    Type = 7
	Type_VERIFY_FACE        Type = 8
	Type_VERIFY_FINGERPRINT Type = 9
	Type_VERIFY_VOICE       Type = 10
	Type_VERIFY_EYE         Type = 11
	Type_VERIFY_SIGNATURE   Type = 12
	Type_VERIFY_BEHAVIOR    Type = 13
	Type_VERIFY_LOCATION    Type = 14
	Type_VERIFY_DEVICE      Type = 15
	Type_VERIFY_BIOMETRIC   Type = 16
	Type_VERIFY_OTP         Type = 17
	Type_VERIFY_PIN         Type = 18
	Type_VERIFY_PATTERN     Type = 19
	Type_VERIFY_PASSWORD    Type = 20
	Type_VERIFY_ANSWER      Type = 21
	Type_VERIFY_QUESTION    Type = 22
	Type_VERIFY_TOKEN       Type = 23
	Type_VERIFY_CODE        Type = 24
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0:  "ACCESS",
		1:  "REFRESH",
		2:  "RESET_PASSWORD",
		3:  "VERIFY_EMAIL",
		4:  "VERIFY_PHONE",
		5:  "VERIFY_IDENTITY",
		6:  "VERIFY_ADDRESS",
		7:  "VERIFY_DOCUMENT",
		8:  "VERIFY_FACE",
		9:  "VERIFY_FINGERPRINT",
		10: "VERIFY_VOICE",
		11: "VERIFY_EYE",
		12: "VERIFY_SIGNATURE",
		13: "VERIFY_BEHAVIOR",
		14: "VERIFY_LOCATION",
		15: "VERIFY_DEVICE",
		16: "VERIFY_BIOMETRIC",
		17: "VERIFY_OTP",
		18: "VERIFY_PIN",
		19: "VERIFY_PATTERN",
		20: "VERIFY_PASSWORD",
		21: "VERIFY_ANSWER",
		22: "VERIFY_QUESTION",
		23: "VERIFY_TOKEN",
		24: "VERIFY_CODE",
	}
	Type_value = map[string]int32{
		"ACCESS":             0,
		"REFRESH":            1,
		"RESET_PASSWORD":     2,
		"VERIFY_EMAIL":       3,
		"VERIFY_PHONE":       4,
		"VERIFY_IDENTITY":    5,
		"VERIFY_ADDRESS":     6,
		"VERIFY_DOCUMENT":    7,
		"VERIFY_FACE":        8,
		"VERIFY_FINGERPRINT": 9,
		"VERIFY_VOICE":       10,
		"VERIFY_EYE":         11,
		"VERIFY_SIGNATURE":   12,
		"VERIFY_BEHAVIOR":    13,
		"VERIFY_LOCATION":    14,
		"VERIFY_DEVICE":      15,
		"VERIFY_BIOMETRIC":   16,
		"VERIFY_OTP":         17,
		"VERIFY_PIN":         18,
		"VERIFY_PATTERN":     19,
		"VERIFY_PASSWORD":    20,
		"VERIFY_ANSWER":      21,
		"VERIFY_QUESTION":    22,
		"VERIFY_TOKEN":       23,
		"VERIFY_CODE":        24,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_model_token_token_proto_enumTypes[2].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_model_token_token_proto_enumTypes[2]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{2}
}

type PayloadType int32

const (
	PayloadType_INT     PayloadType = 0
	PayloadType_INT8    PayloadType = 1
	PayloadType_INT16   PayloadType = 2
	PayloadType_INT32   PayloadType = 3
	PayloadType_INT64   PayloadType = 4
	PayloadType_UINT    PayloadType = 5
	PayloadType_UINT8   PayloadType = 6
	PayloadType_UINT16  PayloadType = 7
	PayloadType_UINT32  PayloadType = 8
	PayloadType_UINT64  PayloadType = 9
	PayloadType_FLOAT32 PayloadType = 10
	PayloadType_FLOAT64 PayloadType = 11
	PayloadType_STRING  PayloadType = 12
	PayloadType_BYTES   PayloadType = 13
	PayloadType_BOOL    PayloadType = 14
	PayloadType_MESSAGE PayloadType = 15
)

// Enum value maps for PayloadType.
var (
	PayloadType_name = map[int32]string{
		0:  "INT",
		1:  "INT8",
		2:  "INT16",
		3:  "INT32",
		4:  "INT64",
		5:  "UINT",
		6:  "UINT8",
		7:  "UINT16",
		8:  "UINT32",
		9:  "UINT64",
		10: "FLOAT32",
		11: "FLOAT64",
		12: "STRING",
		13: "BYTES",
		14: "BOOL",
		15: "MESSAGE",
	}
	PayloadType_value = map[string]int32{
		"INT":     0,
		"INT8":    1,
		"INT16":   2,
		"INT32":   3,
		"INT64":   4,
		"UINT":    5,
		"UINT8":   6,
		"UINT16":  7,
		"UINT32":  8,
		"UINT64":  9,
		"FLOAT32": 10,
		"FLOAT64": 11,
		"STRING":  12,
		"BYTES":   13,
		"BOOL":    14,
		"MESSAGE": 15,
	}
)

func (x PayloadType) Enum() *PayloadType {
	p := new(PayloadType)
	*p = x
	return p
}

func (x PayloadType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PayloadType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_token_token_proto_enumTypes[3].Descriptor()
}

func (PayloadType) Type() protoreflect.EnumType {
	return &file_model_token_token_proto_enumTypes[3]
}

func (x PayloadType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PayloadType.Descriptor instead.
func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{3}
}

type Algorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature SignatureAlgorithm `protobuf:"varint,1,opt,name=signature,proto3,enum=token.SignatureAlgorithm" json:"signature,omitempty"`
	Hash      HashAlgorithm      `protobuf:"varint,2,opt,name=hash,proto3,enum=token.HashAlgorithm" json:"hash,omitempty"`
}

func (x *Algorithm) Reset() {
	*x = Algorithm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_token_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Algorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Algorithm) ProtoMessage() {}

func (x *Algorithm) ProtoReflect() protoreflect.Message {
	mi := &file_model_token_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Algorithm.ProtoReflect.Descriptor instead.
func (*Algorithm) Descriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{0}
}

func (x *Algorithm) GetSignature() SignatureAlgorithm {
	if x != nil {
		return x.Signature
	}
	return SignatureAlgorithm_NONE
}

func (x *Algorithm) GetHash() HashAlgorithm {
	if x != nil {
		return x.Hash
	}
	return HashAlgorithm_SHA3_384
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=token.PayloadType" json:"type,omitempty"`
	Value []byte      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_token_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_model_token_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{1}
}

func (x *Payload) GetType() PayloadType {
	if x != nil {
		return x.Type
	}
	return PayloadType_INT
}

func (x *Payload) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm *Algorithm `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Type      Type       `protobuf:"varint,2,opt,name=type,proto3,enum=token.Type" json:"type,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_token_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_model_token_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{2}
}

func (x *Header) GetAlgorithm() *Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return nil
}

func (x *Header) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_ACCESS
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    *Header             `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Payloads  map[string]*Payload `protobuf:"bytes,2,rep,name=payloads,proto3" json:"payloads,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Signature []byte              `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_token_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_model_token_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_model_token_token_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Token) GetPayloads() map[string]*Payload {
	if x != nil {
		return x.Payloads
	}
	return nil
}

func (x *Token) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_model_token_token_proto protoreflect.FileDescriptor

var file_model_token_token_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x6e, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x37, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x48, 0x61, 0x73,
	0x68, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x22, 0x47, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x59, 0x0a, 0x06, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x12, 0x1f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xd1, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x08, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x4b, 0x0a, 0x0d, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x49, 0x0a, 0x12, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4d, 0x41, 0x43,
	0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x53, 0x41, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x43, 0x44, 0x53, 0x41, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x44, 0x32, 0x35, 0x35, 0x31,
	0x39, 0x10, 0x04, 0x2a, 0x8f, 0x01, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x41, 0x33, 0x5f, 0x33, 0x38,
	0x34, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x48, 0x41, 0x33, 0x5f, 0x35, 0x31, 0x32, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x32, 0x42, 0x5f, 0x33, 0x38, 0x34,
	0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x32, 0x42, 0x5f, 0x35, 0x31,
	0x32, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x32, 0x53, 0x5f, 0x33,
	0x38, 0x34, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x32, 0x53, 0x5f,
	0x35, 0x31, 0x32, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x33, 0x5f,
	0x32, 0x35, 0x36, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x33, 0x5f,
	0x35, 0x31, 0x32, 0x10, 0x07, 0x2a, 0xdd, 0x03, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a,
	0x0a, 0x06, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45,
	0x46, 0x52, 0x45, 0x53, 0x48, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x45, 0x54,
	0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x03, 0x12, 0x10, 0x0a,
	0x0c, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x50, 0x48, 0x4f, 0x4e, 0x45, 0x10, 0x04, 0x12,
	0x13, 0x0a, 0x0f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x49,
	0x54, 0x59, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x41,
	0x44, 0x44, 0x52, 0x45, 0x53, 0x53, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x45, 0x52, 0x49,
	0x46, 0x59, 0x5f, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x07, 0x12, 0x0f, 0x0a,
	0x0b, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x41, 0x43, 0x45, 0x10, 0x08, 0x12, 0x16,
	0x0a, 0x12, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x46, 0x49, 0x4e, 0x47, 0x45, 0x52, 0x50,
	0x52, 0x49, 0x4e, 0x54, 0x10, 0x09, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59,
	0x5f, 0x56, 0x4f, 0x49, 0x43, 0x45, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x45, 0x52, 0x49,
	0x46, 0x59, 0x5f, 0x45, 0x59, 0x45, 0x10, 0x0b, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x45, 0x52, 0x49,
	0x46, 0x59, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x0c, 0x12, 0x13,
	0x0a, 0x0f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56, 0x49, 0x4f,
	0x52, 0x10, 0x0d, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x4c, 0x4f,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x45, 0x52, 0x49,
	0x46, 0x59, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x0f, 0x12, 0x14, 0x0a, 0x10, 0x56,
	0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x42, 0x49, 0x4f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x10,
	0x10, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x4f, 0x54, 0x50, 0x10,
	0x11, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x50, 0x49, 0x4e, 0x10,
	0x12, 0x12, 0x12, 0x0a, 0x0e, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x50, 0x41, 0x54, 0x54,
	0x45, 0x52, 0x4e, 0x10, 0x13, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f,
	0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x45,
	0x52, 0x49, 0x46, 0x59, 0x5f, 0x41, 0x4e, 0x53, 0x57, 0x45, 0x52, 0x10, 0x15, 0x12, 0x13, 0x0a,
	0x0f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x16, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x10, 0x17, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x10, 0x18, 0x2a, 0xc2, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x31,
	0x36, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x49, 0x4e,
	0x54, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49,
	0x4e, 0x54, 0x33, 0x32, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34,
	0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x33, 0x32, 0x10, 0x0a, 0x12,
	0x0b, 0x0a, 0x07, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45,
	0x53, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x0e, 0x12, 0x0b, 0x0a,
	0x07, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x0f, 0x42, 0x0d, 0x5a, 0x0b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_model_token_token_proto_rawDescOnce sync.Once
	file_model_token_token_proto_rawDescData = file_model_token_token_proto_rawDesc
)

func file_model_token_token_proto_rawDescGZIP() []byte {
	file_model_token_token_proto_rawDescOnce.Do(func() {
		file_model_token_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_token_token_proto_rawDescData)
	})
	return file_model_token_token_proto_rawDescData
}

var file_model_token_token_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_model_token_token_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_model_token_token_proto_goTypes = []interface{}{
	(SignatureAlgorithm)(0), // 0: token.SignatureAlgorithm
	(HashAlgorithm)(0),      // 1: token.HashAlgorithm
	(Type)(0),               // 2: token.Type
	(PayloadType)(0),        // 3: token.PayloadType
	(*Algorithm)(nil),       // 4: token.Algorithm
	(*Payload)(nil),         // 5: token.Payload
	(*Header)(nil),          // 6: token.Header
	(*Token)(nil),           // 7: token.Token
	nil,                     // 8: token.Token.PayloadsEntry
}
var file_model_token_token_proto_depIdxs = []int32{
	0, // 0: token.Algorithm.signature:type_name -> token.SignatureAlgorithm
	1, // 1: token.Algorithm.hash:type_name -> token.HashAlgorithm
	3, // 2: token.Payload.type:type_name -> token.PayloadType
	4, // 3: token.Header.algorithm:type_name -> token.Algorithm
	2, // 4: token.Header.type:type_name -> token.Type
	6, // 5: token.Token.header:type_name -> token.Header
	8, // 6: token.Token.payloads:type_name -> token.Token.PayloadsEntry
	5, // 7: token.Token.PayloadsEntry.value:type_name -> token.Payload
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_model_token_token_proto_init() }
func file_model_token_token_proto_init() {
	if File_model_token_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_token_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Algorithm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_token_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_token_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_model_token_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_model_token_token_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_token_token_proto_goTypes,
		DependencyIndexes: file_model_token_token_proto_depIdxs,
		EnumInfos:         file_model_token_token_proto_enumTypes,
		MessageInfos:      file_model_token_token_proto_msgTypes,
	}.Build()
	File_model_token_token_proto = out.File
	file_model_token_token_proto_rawDesc = nil
	file_model_token_token_proto_goTypes = nil
	file_model_token_token_proto_depIdxs = nil
}
