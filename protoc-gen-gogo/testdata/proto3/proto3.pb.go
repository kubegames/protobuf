// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3/proto3.proto

package proto3

import (
	fmt "fmt"
	proto "github.com/kubegames/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Request_Flavour int32

const (
	Request_SWEET         Request_Flavour = 0
	Request_SOUR          Request_Flavour = 1
	Request_UMAMI         Request_Flavour = 2
	Request_GOPHERLICIOUS Request_Flavour = 3
)

var Request_Flavour_name = map[int32]string{
	0: "SWEET",
	1: "SOUR",
	2: "UMAMI",
	3: "GOPHERLICIOUS",
}

var Request_Flavour_value = map[string]int32{
	"SWEET":         0,
	"SOUR":          1,
	"UMAMI":         2,
	"GOPHERLICIOUS": 3,
}

func (x Request_Flavour) String() string {
	return proto.EnumName(Request_Flavour_name, int32(x))
}

func (Request_Flavour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0, 0}
}

type Request struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key                  []int64         `protobuf:"varint,2,rep,packed,name=key,proto3" json:"key,omitempty"`
	Taste                Request_Flavour `protobuf:"varint,3,opt,name=taste,proto3,enum=proto3.Request_Flavour" json:"taste,omitempty"`
	Book                 *Book           `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty"`
	Unpacked             []int64         `protobuf:"varint,5,rep,name=unpacked,proto3" json:"unpacked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetTaste() Request_Flavour {
	if m != nil {
		return m.Taste
	}
	return Request_SWEET
}

func (m *Request) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *Request) GetUnpacked() []int64 {
	if m != nil {
		return m.Unpacked
	}
	return nil
}

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{1}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto3.Request_Flavour", Request_Flavour_name, Request_Flavour_value)
	proto.RegisterType((*Request)(nil), "proto3.Request")
	proto.RegisterType((*Book)(nil), "proto3.Book")
}

func init() { proto.RegisterFile("proto3/proto3.proto", fileDescriptor_ab04eb4084a521db) }

var fileDescriptor_ab04eb4084a521db = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x4d, 0x4f, 0xf2, 0x40,
	0x14, 0x85, 0x99, 0x7e, 0xbc, 0xc0, 0x7d, 0xd1, 0x8c, 0x57, 0x13, 0xc7, 0x8d, 0x99, 0xb0, 0xea,
	0x86, 0x92, 0xe0, 0x42, 0x63, 0xdc, 0x88, 0xa2, 0x92, 0x48, 0x30, 0x83, 0xc4, 0xc4, 0x8d, 0x99,
	0xc2, 0x58, 0x49, 0x81, 0xc1, 0x76, 0x2a, 0xf1, 0xcf, 0xfa, 0x5b, 0x4c, 0x3b, 0xc5, 0xd5, 0x3d,
	0xf7, 0x23, 0xcf, 0xc9, 0x3d, 0x70, 0xb8, 0x49, 0xb5, 0xd1, 0x67, 0x5d, 0x5b, 0xc2, 0xb2, 0xe0,
	0x3f, 0xdb, 0xb5, 0x7f, 0x08, 0xd4, 0x85, 0xfa, 0xcc, 0x55, 0x66, 0x10, 0xc1, 0x5b, 0xcb, 0x95,
	0x62, 0x84, 0x93, 0xa0, 0x29, 0x4a, 0x8d, 0x14, 0xdc, 0x44, 0x7d, 0x33, 0x87, 0xbb, 0x81, 0x2b,
	0x0a, 0x89, 0x1d, 0xf0, 0x8d, 0xcc, 0x8c, 0x62, 0x2e, 0x27, 0xc1, 0x7e, 0xef, 0x38, 0xac, 0xb8,
	0x15, 0x25, 0xbc, 0x5b, 0xca, 0x2f, 0x9d, 0xa7, 0xc2, 0x5e, 0x21, 0x07, 0x2f, 0xd2, 0x3a, 0x61,
	0x1e, 0x27, 0xc1, 0xff, 0x5e, 0x6b, 0x77, 0xdd, 0xd7, 0x3a, 0x11, 0xe5, 0x06, 0x4f, 0xa1, 0x91,
	0xaf, 0x37, 0x72, 0x96, 0xa8, 0x39, 0xf3, 0x0b, 0x9f, 0xbe, 0x43, 0x6b, 0xe2, 0x6f, 0xd6, 0xbe,
	0x82, 0x7a, 0xc5, 0xc4, 0x26, 0xf8, 0x93, 0x97, 0xc1, 0xe0, 0x99, 0xd6, 0xb0, 0x01, 0xde, 0x64,
	0x3c, 0x15, 0x94, 0x14, 0xc3, 0xe9, 0xe8, 0x7a, 0x34, 0xa4, 0x0e, 0x1e, 0xc0, 0xde, 0xfd, 0xf8,
	0xe9, 0x61, 0x20, 0x1e, 0x87, 0x37, 0xc3, 0xf1, 0x74, 0x42, 0xdd, 0xf6, 0x39, 0x78, 0x85, 0x17,
	0x1e, 0x81, 0x6f, 0x16, 0x66, 0xb9, 0xfb, 0xce, 0x36, 0x78, 0x02, 0x8d, 0x54, 0x6e, 0xdf, 0xe6,
	0xd2, 0x48, 0xe6, 0x70, 0x12, 0xb4, 0x44, 0x3d, 0x95, 0xdb, 0x5b, 0x69, 0x64, 0xff, 0xf2, 0xf5,
	0x22, 0x5e, 0x98, 0x8f, 0x3c, 0x0a, 0x67, 0x7a, 0xd5, 0x8d, 0x75, 0xac, 0x6d, 0x82, 0x51, 0xfe,
	0x6e, 0xc5, 0xac, 0x13, 0xab, 0x75, 0xa7, 0x5c, 0x18, 0x95, 0x99, 0x82, 0x51, 0x65, 0x1c, 0x55,
	0xe9, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xe1, 0xfa, 0x46, 0x7b, 0x01, 0x00, 0x00,
}
