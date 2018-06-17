// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/define.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	api/proto/define.proto

It has these top-level messages:
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type BASE_PATH int32

const (
	BASE_PATH_api BASE_PATH = 0
)

var BASE_PATH_name = map[int32]string{
	0: "api",
}
var BASE_PATH_value = map[string]int32{
	"api": 0,
}

func (x BASE_PATH) String() string {
	return proto1.EnumName(BASE_PATH_name, int32(x))
}
func (BASE_PATH) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto1.RegisterEnum("proto.BASE_PATH", BASE_PATH_name, BASE_PATH_value)
}

func init() { proto1.RegisterFile("api/proto/define.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 75 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x49, 0x4d, 0xcb, 0xcc, 0x4b, 0xd5, 0x03, 0x73, 0x84,
	0x58, 0xc1, 0x94, 0x96, 0x08, 0x17, 0xa7, 0x93, 0x63, 0xb0, 0x6b, 0x7c, 0x80, 0x63, 0x88, 0x87,
	0x10, 0x3b, 0x17, 0x73, 0x62, 0x41, 0xa6, 0x00, 0x43, 0x12, 0x1b, 0x58, 0xd2, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x59, 0xdd, 0xae, 0x3d, 0x00, 0x00, 0x00,
}
