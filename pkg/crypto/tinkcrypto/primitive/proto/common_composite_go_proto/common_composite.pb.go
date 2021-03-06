// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common_composite.proto

package common_composite_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type KeyType int32

const (
	KeyType_UNKNOWN_KEY_TYPE KeyType = 0
	KeyType_EC               KeyType = 1
	KeyType_OKP              KeyType = 2
)

var KeyType_name = map[int32]string{
	0: "UNKNOWN_KEY_TYPE",
	1: "EC",
	2: "OKP",
}

var KeyType_value = map[string]int32{
	"UNKNOWN_KEY_TYPE": 0,
	"EC":               1,
	"OKP":              2,
}

func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}

func (KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c0933941e418c01, []int{0}
}

func init() {
	proto.RegisterEnum("google.crypto.tink.KeyType", KeyType_name, KeyType_value)
}

func init() { proto.RegisterFile("proto/common_composite.proto", fileDescriptor_3c0933941e418c01) }

var fileDescriptor_3c0933941e418c01 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x86, 0x9d, 0x11, 0xa6, 0x90, 0x55, 0x08, 0x2e, 0x7b, 0x02, 0xa1, 0x89, 0xe0, 0x0d, 0x2a,
	0xb3, 0x90, 0x40, 0x9a, 0xc5, 0x88, 0xd4, 0x4d, 0xe8, 0xc4, 0x67, 0x1a, 0xda, 0xf4, 0x85, 0xd7,
	0xa8, 0xe4, 0x10, 0x5e, 0xc2, 0x93, 0x8a, 0x33, 0xee, 0xec, 0xee, 0xbd, 0x0f, 0xbe, 0xff, 0xe7,
	0x67, 0xcb, 0x4c, 0x58, 0x50, 0x79, 0x4c, 0x09, 0x4f, 0xce, 0x63, 0xca, 0x78, 0x8e, 0x05, 0xe4,
	0x84, 0x85, 0x08, 0x88, 0xe1, 0x08, 0xd2, 0x53, 0xcd, 0x05, 0x65, 0x89, 0xa7, 0xc3, 0xed, 0x1d,
	0x5b, 0x68, 0xa8, 0x43, 0xcd, 0x20, 0x6e, 0x18, 0x7f, 0x32, 0xda, 0x6c, 0x9e, 0x8d, 0xd3, 0xfd,
	0xd6, 0x0d, 0x5b, 0xdb, 0xf3, 0x2b, 0xd1, 0xb1, 0xb6, 0x7f, 0xe0, 0x8d, 0x58, 0xb0, 0xeb, 0x8d,
	0xb6, 0xbc, 0x5d, 0x7f, 0x35, 0x6c, 0xe9, 0x31, 0xc9, 0xff, 0x61, 0x73, 0x8d, 0x6d, 0x5e, 0xc6,
	0x10, 0xcb, 0xfe, 0x7d, 0x94, 0x1e, 0x93, 0xda, 0xd7, 0x0c, 0x74, 0x84, 0xd7, 0x00, 0xa4, 0x76,
	0x14, 0xe1, 0xbc, 0x7a, 0xa3, 0x5d, 0x82, 0x4f, 0xa4, 0xc3, 0x2a, 0xa0, 0x9a, 0x75, 0xf5, 0xab,
	0xff, 0x9d, 0x99, 0x62, 0x8a, 0x25, 0x7e, 0x80, 0xba, 0x3c, 0xc6, 0x4d, 0xf8, 0xbb, 0xed, 0x86,
	0x47, 0xa3, 0xed, 0x7a, 0xec, 0xa6, 0xff, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xe5, 0x85,
	0xe6, 0xfc, 0x00, 0x00, 0x00,
}
