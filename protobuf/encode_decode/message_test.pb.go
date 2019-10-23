// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_test.proto

package encode_decode

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

type VarintsA struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VarintsA) Reset()         { *m = VarintsA{} }
func (m *VarintsA) String() string { return proto.CompactTextString(m) }
func (*VarintsA) ProtoMessage()    {}
func (*VarintsA) Descriptor() ([]byte, []int) {
	return fileDescriptor_1df2290959daf7b8, []int{0}
}

func (m *VarintsA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VarintsA.Unmarshal(m, b)
}
func (m *VarintsA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VarintsA.Marshal(b, m, deterministic)
}
func (m *VarintsA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VarintsA.Merge(m, src)
}
func (m *VarintsA) XXX_Size() int {
	return xxx_messageInfo_VarintsA.Size(m)
}
func (m *VarintsA) XXX_DiscardUnknown() {
	xxx_messageInfo_VarintsA.DiscardUnknown(m)
}

var xxx_messageInfo_VarintsA proto.InternalMessageInfo

func (m *VarintsA) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

type ZigZagA struct {
	B                    int32    `protobuf:"zigzag32,1,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZigZagA) Reset()         { *m = ZigZagA{} }
func (m *ZigZagA) String() string { return proto.CompactTextString(m) }
func (*ZigZagA) ProtoMessage()    {}
func (*ZigZagA) Descriptor() ([]byte, []int) {
	return fileDescriptor_1df2290959daf7b8, []int{1}
}

func (m *ZigZagA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZigZagA.Unmarshal(m, b)
}
func (m *ZigZagA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZigZagA.Marshal(b, m, deterministic)
}
func (m *ZigZagA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZigZagA.Merge(m, src)
}
func (m *ZigZagA) XXX_Size() int {
	return xxx_messageInfo_ZigZagA.Size(m)
}
func (m *ZigZagA) XXX_DiscardUnknown() {
	xxx_messageInfo_ZigZagA.DiscardUnknown(m)
}

var xxx_messageInfo_ZigZagA proto.InternalMessageInfo

func (m *ZigZagA) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type StringA struct {
	A                    string   `protobuf:"bytes,7,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringA) Reset()         { *m = StringA{} }
func (m *StringA) String() string { return proto.CompactTextString(m) }
func (*StringA) ProtoMessage()    {}
func (*StringA) Descriptor() ([]byte, []int) {
	return fileDescriptor_1df2290959daf7b8, []int{2}
}

func (m *StringA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringA.Unmarshal(m, b)
}
func (m *StringA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringA.Marshal(b, m, deterministic)
}
func (m *StringA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringA.Merge(m, src)
}
func (m *StringA) XXX_Size() int {
	return xxx_messageInfo_StringA.Size(m)
}
func (m *StringA) XXX_DiscardUnknown() {
	xxx_messageInfo_StringA.DiscardUnknown(m)
}

var xxx_messageInfo_StringA proto.InternalMessageInfo

func (m *StringA) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

type ReapeatedA struct {
	A                    []int32  `protobuf:"varint,3,rep,packed,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReapeatedA) Reset()         { *m = ReapeatedA{} }
func (m *ReapeatedA) String() string { return proto.CompactTextString(m) }
func (*ReapeatedA) ProtoMessage()    {}
func (*ReapeatedA) Descriptor() ([]byte, []int) {
	return fileDescriptor_1df2290959daf7b8, []int{3}
}

func (m *ReapeatedA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReapeatedA.Unmarshal(m, b)
}
func (m *ReapeatedA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReapeatedA.Marshal(b, m, deterministic)
}
func (m *ReapeatedA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReapeatedA.Merge(m, src)
}
func (m *ReapeatedA) XXX_Size() int {
	return xxx_messageInfo_ReapeatedA.Size(m)
}
func (m *ReapeatedA) XXX_DiscardUnknown() {
	xxx_messageInfo_ReapeatedA.DiscardUnknown(m)
}

var xxx_messageInfo_ReapeatedA proto.InternalMessageInfo

func (m *ReapeatedA) GetA() []int32 {
	if m != nil {
		return m.A
	}
	return nil
}

type MultiFieldA struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    string   `protobuf:"bytes,7,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiFieldA) Reset()         { *m = MultiFieldA{} }
func (m *MultiFieldA) String() string { return proto.CompactTextString(m) }
func (*MultiFieldA) ProtoMessage()    {}
func (*MultiFieldA) Descriptor() ([]byte, []int) {
	return fileDescriptor_1df2290959daf7b8, []int{4}
}

func (m *MultiFieldA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiFieldA.Unmarshal(m, b)
}
func (m *MultiFieldA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiFieldA.Marshal(b, m, deterministic)
}
func (m *MultiFieldA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiFieldA.Merge(m, src)
}
func (m *MultiFieldA) XXX_Size() int {
	return xxx_messageInfo_MultiFieldA.Size(m)
}
func (m *MultiFieldA) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiFieldA.DiscardUnknown(m)
}

var xxx_messageInfo_MultiFieldA proto.InternalMessageInfo

func (m *MultiFieldA) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *MultiFieldA) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

func init() {
	proto.RegisterType((*VarintsA)(nil), "encode_decode.VarintsA")
	proto.RegisterType((*ZigZagA)(nil), "encode_decode.ZigZagA")
	proto.RegisterType((*StringA)(nil), "encode_decode.StringA")
	proto.RegisterType((*ReapeatedA)(nil), "encode_decode.ReapeatedA")
	proto.RegisterType((*MultiFieldA)(nil), "encode_decode.MultiFieldA")
}

func init() { proto.RegisterFile("message_test.proto", fileDescriptor_1df2290959daf7b8) }

var fileDescriptor_1df2290959daf7b8 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xca, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x8d, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2,
	0x4d, 0xcd, 0x4b, 0xce, 0x4f, 0x49, 0x8d, 0x4f, 0x49, 0x05, 0x51, 0x4a, 0x12, 0x5c, 0x1c, 0x61,
	0x89, 0x45, 0x99, 0x79, 0x25, 0xc5, 0x8e, 0x42, 0x3c, 0x5c, 0x8c, 0x89, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xac, 0x41, 0x8c, 0x89, 0x4a, 0xe2, 0x5c, 0xec, 0x51, 0x99, 0xe9, 0x51, 0x89, 0xe9, 0x60,
	0x89, 0x24, 0xb0, 0x84, 0x60, 0x10, 0x63, 0x12, 0x48, 0x22, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x1d,
	0xaa, 0x83, 0x5d, 0x81, 0x51, 0x83, 0x13, 0xa4, 0x43, 0x8a, 0x8b, 0x2b, 0x28, 0x35, 0xb1, 0x20,
	0x35, 0xb1, 0x24, 0x35, 0x05, 0x2a, 0xc7, 0xac, 0xc0, 0x0c, 0x31, 0x4d, 0x93, 0x8b, 0xdb, 0xb7,
	0x34, 0xa7, 0x24, 0xd3, 0x2d, 0x33, 0x35, 0x27, 0x05, 0xcd, 0x2a, 0x88, 0xf9, 0x50, 0x63, 0x92,
	0x92, 0xd8, 0xc0, 0x0e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x41, 0x2b, 0x4b, 0xbe,
	0x00, 0x00, 0x00,
}
