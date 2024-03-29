// Code generated by protoc-gen-go. DO NOT EDIT.
// source: validator.proto

package validator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type Validator struct {
	Lt                   string   `protobuf:"bytes,1,opt,name=lt,proto3" json:"lt,omitempty"`
	Lte                  string   `protobuf:"bytes,2,opt,name=lte,proto3" json:"lte,omitempty"`
	Gt                   string   `protobuf:"bytes,3,opt,name=gt,proto3" json:"gt,omitempty"`
	Gte                  string   `protobuf:"bytes,4,opt,name=gte,proto3" json:"gte,omitempty"`
	Eq                   string   `protobuf:"bytes,7,opt,name=eq,proto3" json:"eq,omitempty"`
	Ne                   string   `protobuf:"bytes,8,opt,name=ne,proto3" json:"ne,omitempty"`
	In                   string   `protobuf:"bytes,5,opt,name=in,proto3" json:"in,omitempty"`
	Nin                  string   `protobuf:"bytes,6,opt,name=nin,proto3" json:"nin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf1c6ec7c0d80dd5, []int{0}
}

func (m *Validator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Validator.Unmarshal(m, b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return xxx_messageInfo_Validator.Size(m)
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetLt() string {
	if m != nil {
		return m.Lt
	}
	return ""
}

func (m *Validator) GetLte() string {
	if m != nil {
		return m.Lte
	}
	return ""
}

func (m *Validator) GetGt() string {
	if m != nil {
		return m.Gt
	}
	return ""
}

func (m *Validator) GetGte() string {
	if m != nil {
		return m.Gte
	}
	return ""
}

func (m *Validator) GetEq() string {
	if m != nil {
		return m.Eq
	}
	return ""
}

func (m *Validator) GetNe() string {
	if m != nil {
		return m.Ne
	}
	return ""
}

func (m *Validator) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Validator) GetNin() string {
	if m != nil {
		return m.Nin
	}
	return ""
}

var E_Validate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Validator)(nil),
	Field:         1000,
	Name:          "validator.validate",
	Tag:           "bytes,1000,opt,name=validate",
	Filename:      "validator.proto",
}

func init() {
	proto.RegisterType((*Validator)(nil), "validator.Validator")
	proto.RegisterExtension(E_Validate)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptor_bf1c6ec7c0d80dd5) }

var fileDescriptor_bf1c6ec7c0d80dd5 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xbf, 0x4a, 0x04, 0x31,
	0x10, 0xc6, 0xd9, 0x3d, 0x3d, 0x6f, 0x73, 0xa0, 0x12, 0x2c, 0x06, 0x41, 0x38, 0xac, 0xac, 0x72,
	0xa0, 0x9d, 0xa5, 0x85, 0xad, 0x72, 0x85, 0x85, 0xdd, 0xad, 0x3b, 0x86, 0x81, 0x30, 0xd9, 0x3f,
	0xa3, 0xef, 0xe1, 0x5b, 0xfa, 0x18, 0x92, 0x8c, 0x9b, 0xed, 0xf2, 0xcd, 0xfc, 0xf8, 0xc2, 0x6f,
	0xcc, 0xc5, 0xf7, 0x31, 0x50, 0x77, 0x94, 0x38, 0xba, 0x7e, 0x8c, 0x12, 0x6d, 0x53, 0x06, 0xd7,
	0x3b, 0x1f, 0xa3, 0x0f, 0xb8, 0xcf, 0x8b, 0xf6, 0xeb, 0x73, 0xdf, 0xe1, 0xf4, 0x31, 0x52, 0x5f,
	0xe0, 0xdb, 0x9f, 0xca, 0x34, 0x6f, 0x33, 0x6f, 0xcf, 0x4d, 0x1d, 0x04, 0xaa, 0x5d, 0x75, 0xd7,
	0x1c, 0xea, 0x20, 0xf6, 0xd2, 0xac, 0x82, 0x20, 0xd4, 0x79, 0x90, 0x9e, 0x89, 0xf0, 0x02, 0x2b,
	0x25, 0x7c, 0x26, 0xbc, 0x20, 0x9c, 0x28, 0xe1, 0x95, 0xc0, 0x01, 0xce, 0x94, 0xc0, 0x21, 0x65,
	0x46, 0xd8, 0x68, 0xe6, 0xbc, 0x27, 0x86, 0x53, 0xcd, 0xc4, 0xa9, 0x81, 0x89, 0x61, 0xad, 0x0d,
	0x4c, 0xfc, 0xf8, 0x6a, 0x36, 0xff, 0x0a, 0x68, 0x6f, 0x9c, 0x2a, 0xb8, 0x59, 0xc1, 0x3d, 0x13,
	0x86, 0xee, 0xa5, 0x17, 0x8a, 0x3c, 0xc1, 0x6f, 0xfa, 0x67, 0x7b, 0x7f, 0xe5, 0x96, 0x23, 0x14,
	0x9b, 0x43, 0x69, 0x79, 0xda, 0xbe, 0x2f, 0x47, 0x69, 0xd7, 0xb9, 0xea, 0xe1, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x52, 0x20, 0x0a, 0x5a, 0x39, 0x01, 0x00, 0x00,
}
