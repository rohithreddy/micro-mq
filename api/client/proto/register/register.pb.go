// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/client/proto/register/register.proto

/*
Package api_client_register is a generated protocol buffer package.

It is generated from these files:
	api/client/proto/register/register.proto

It has these top-level messages:
	RegisterReq
	RegisterResp
*/
package api_client_register

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RegisterReq struct {
	ClientId string `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
}

func (m *RegisterReq) Reset()                    { *m = RegisterReq{} }
func (m *RegisterReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()               {}
func (*RegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type RegisterResp struct {
	ClientId string `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Pwd      string `protobuf:"bytes,3,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *RegisterResp) Reset()                    { *m = RegisterResp{} }
func (m *RegisterResp) String() string            { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()               {}
func (*RegisterResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterResp) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *RegisterResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterResp) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "api_client_register.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "api_client_register.RegisterResp")
}

func init() { proto.RegisterFile("api/client/proto/register/register.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x2c, 0xc8, 0xd4,
	0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4a, 0x4d,
	0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x82, 0x33, 0xf4, 0xc0, 0xe2, 0x42, 0xc2, 0x89, 0x05, 0x99, 0xf1,
	0x10, 0x95, 0xf1, 0x30, 0x29, 0x25, 0x4d, 0x2e, 0xee, 0x20, 0x28, 0x3b, 0x28, 0xb5, 0x50, 0x48,
	0x8a, 0x8b, 0x03, 0xa2, 0xc2, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xce, 0x57,
	0x0a, 0xe0, 0xe2, 0x41, 0x28, 0x2d, 0x2e, 0xc0, 0xa7, 0x56, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31,
	0x37, 0x55, 0x82, 0x09, 0x2c, 0x0e, 0x66, 0x0b, 0x09, 0x70, 0x31, 0x17, 0x94, 0xa7, 0x48, 0x30,
	0x83, 0x85, 0x40, 0xcc, 0x24, 0x36, 0xb0, 0xc3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd2,
	0xc6, 0xf2, 0x90, 0xc4, 0x00, 0x00, 0x00,
}
