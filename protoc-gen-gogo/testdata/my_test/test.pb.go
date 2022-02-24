// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: my_test/test.proto

// This package holds interesting messages.

package test

import (
	fmt "fmt"
	proto "github.com/kubegames/protobuf/proto"
	_ "github.com/kubegames/protobuf/protoc-gen-gogo/testdata/multi"
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

type HatType int32

const (
	// deliberately skipping 0
	HatType_FEDORA HatType = 1
	HatType_FEZ    HatType = 2
)

var HatType_name = map[int32]string{
	1: "FEDORA",
	2: "FEZ",
}

var HatType_value = map[string]int32{
	"FEDORA": 1,
	"FEZ":    2,
}

func (x HatType) Enum() *HatType {
	p := new(HatType)
	*p = x
	return p
}

func (x HatType) String() string {
	return proto.EnumName(HatType_name, int32(x))
}

func (x *HatType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HatType_value, data, "HatType")
	if err != nil {
		return err
	}
	*x = HatType(value)
	return nil
}

func (HatType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0}
}

// This enum represents days of the week.
type Days int32

const (
	Days_MONDAY  Days = 1
	Days_TUESDAY Days = 2
	Days_LUNDI   Days = 1
)

var Days_name = map[int32]string{
	1: "MONDAY",
	2: "TUESDAY",
	// Duplicate value: 1: "LUNDI",
}

var Days_value = map[string]int32{
	"MONDAY":  1,
	"TUESDAY": 2,
	"LUNDI":   1,
}

func (x Days) Enum() *Days {
	p := new(Days)
	*p = x
	return p
}

func (x Days) String() string {
	return proto.EnumName(Days_name, int32(x))
}

func (x *Days) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Days_value, data, "Days")
	if err != nil {
		return err
	}
	*x = Days(value)
	return nil
}

func (Days) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1}
}

type Request_Color int32

const (
	Request_RED   Request_Color = 0
	Request_GREEN Request_Color = 1
	Request_BLUE  Request_Color = 2
)

var Request_Color_name = map[int32]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}

var Request_Color_value = map[string]int32{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Request_Color) Enum() *Request_Color {
	p := new(Request_Color)
	*p = x
	return p
}

func (x Request_Color) String() string {
	return proto.EnumName(Request_Color_name, int32(x))
}

func (x *Request_Color) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Color_value, data, "Request_Color")
	if err != nil {
		return err
	}
	*x = Request_Color(value)
	return nil
}

func (Request_Color) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0, 0}
}

type Reply_Entry_Game int32

const (
	Reply_Entry_FOOTBALL Reply_Entry_Game = 1
	Reply_Entry_TENNIS   Reply_Entry_Game = 2
)

var Reply_Entry_Game_name = map[int32]string{
	1: "FOOTBALL",
	2: "TENNIS",
}

var Reply_Entry_Game_value = map[string]int32{
	"FOOTBALL": 1,
	"TENNIS":   2,
}

func (x Reply_Entry_Game) Enum() *Reply_Entry_Game {
	p := new(Reply_Entry_Game)
	*p = x
	return p
}

func (x Reply_Entry_Game) String() string {
	return proto.EnumName(Reply_Entry_Game_name, int32(x))
}

func (x *Reply_Entry_Game) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Reply_Entry_Game_value, data, "Reply_Entry_Game")
	if err != nil {
		return err
	}
	*x = Reply_Entry_Game(value)
	return nil
}

func (Reply_Entry_Game) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1, 0, 0}
}

// This is a message that might be sent somewhere.
type Request struct {
	Key []int64 `protobuf:"varint,1,rep,name=key" json:"key,omitempty"`
	//  optional imp.ImportedMessage imported_message = 2;
	Hue *Request_Color `protobuf:"varint,3,opt,name=hue,enum=my.test.Request_Color" json:"hue,omitempty"`
	Hat *HatType       `protobuf:"varint,4,opt,name=hat,enum=my.test.HatType,def=1" json:"hat,omitempty"`
	//  optional imp.ImportedMessage.Owner owner = 6;
	Deadline  *float32           `protobuf:"fixed32,7,opt,name=deadline,def=inf" json:"deadline,omitempty"`
	Somegroup *Request_SomeGroup `protobuf:"group,8,opt,name=SomeGroup,json=somegroup" json:"somegroup,omitempty"`
	// This is a map field. It will generate map[int32]string.
	NameMapping map[int32]string `protobuf:"bytes,14,rep,name=name_mapping,json=nameMapping" json:"name_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// This is a map field whose value type is a message.
	MsgMapping map[int64]*Reply `protobuf:"bytes,15,rep,name=msg_mapping,json=msgMapping" json:"msg_mapping,omitempty" protobuf_key:"zigzag64,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Reset_     *int32           `protobuf:"varint,12,opt,name=reset" json:"reset,omitempty"`
	// This field should not conflict with any getters.
	GetKey_              *string  `protobuf:"bytes,16,opt,name=get_key,json=getKey" json:"get_key,omitempty"`
	FloatNinf            *float32 `protobuf:"fixed32,20,opt,name=float_ninf,json=floatNinf,def=-inf" json:"float_ninf,omitempty"`
	FloatPinf            *float32 `protobuf:"fixed32,21,opt,name=float_pinf,json=floatPinf,def=inf" json:"float_pinf,omitempty"`
	FloatExp             *float32 `protobuf:"fixed32,22,opt,name=float_exp,json=floatExp,def=1e+09" json:"float_exp,omitempty"`
	DoubleNinf           *float64 `protobuf:"fixed64,23,opt,name=double_ninf,json=doubleNinf,def=-inf" json:"double_ninf,omitempty"`
	DoublePinf           *float64 `protobuf:"fixed64,24,opt,name=double_pinf,json=doublePinf,def=inf" json:"double_pinf,omitempty"`
	DoubleExp            *float64 `protobuf:"fixed64,25,opt,name=double_exp,json=doubleExp,def=1e+09" json:"double_exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0}
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

const Default_Request_Hat HatType = HatType_FEDORA

var Default_Request_Deadline float32 = float32(math.Inf(1))
var Default_Request_FloatNinf float32 = float32(math.Inf(-1))
var Default_Request_FloatPinf float32 = float32(math.Inf(1))

const Default_Request_FloatExp float32 = 1e+09

var Default_Request_DoubleNinf float64 = math.Inf(-1)
var Default_Request_DoublePinf float64 = math.Inf(1)

const Default_Request_DoubleExp float64 = 1e+09

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetHue() Request_Color {
	if m != nil && m.Hue != nil {
		return *m.Hue
	}
	return Request_RED
}

func (m *Request) GetHat() HatType {
	if m != nil && m.Hat != nil {
		return *m.Hat
	}
	return Default_Request_Hat
}

func (m *Request) GetDeadline() float32 {
	if m != nil && m.Deadline != nil {
		return *m.Deadline
	}
	return Default_Request_Deadline
}

func (m *Request) GetSomegroup() *Request_SomeGroup {
	if m != nil {
		return m.Somegroup
	}
	return nil
}

func (m *Request) GetNameMapping() map[int32]string {
	if m != nil {
		return m.NameMapping
	}
	return nil
}

func (m *Request) GetMsgMapping() map[int64]*Reply {
	if m != nil {
		return m.MsgMapping
	}
	return nil
}

func (m *Request) GetReset_() int32 {
	if m != nil && m.Reset_ != nil {
		return *m.Reset_
	}
	return 0
}

func (m *Request) GetGetKey_() string {
	if m != nil && m.GetKey_ != nil {
		return *m.GetKey_
	}
	return ""
}

func (m *Request) GetFloatNinf() float32 {
	if m != nil && m.FloatNinf != nil {
		return *m.FloatNinf
	}
	return Default_Request_FloatNinf
}

func (m *Request) GetFloatPinf() float32 {
	if m != nil && m.FloatPinf != nil {
		return *m.FloatPinf
	}
	return Default_Request_FloatPinf
}

func (m *Request) GetFloatExp() float32 {
	if m != nil && m.FloatExp != nil {
		return *m.FloatExp
	}
	return Default_Request_FloatExp
}

func (m *Request) GetDoubleNinf() float64 {
	if m != nil && m.DoubleNinf != nil {
		return *m.DoubleNinf
	}
	return Default_Request_DoubleNinf
}

func (m *Request) GetDoublePinf() float64 {
	if m != nil && m.DoublePinf != nil {
		return *m.DoublePinf
	}
	return Default_Request_DoublePinf
}

func (m *Request) GetDoubleExp() float64 {
	if m != nil && m.DoubleExp != nil {
		return *m.DoubleExp
	}
	return Default_Request_DoubleExp
}

type Request_SomeGroup struct {
	GroupField           *int32   `protobuf:"varint,9,opt,name=group_field,json=groupField" json:"group_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_SomeGroup) Reset()         { *m = Request_SomeGroup{} }
func (m *Request_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*Request_SomeGroup) ProtoMessage()    {}
func (*Request_SomeGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{0, 0}
}
func (m *Request_SomeGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_SomeGroup.Unmarshal(m, b)
}
func (m *Request_SomeGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_SomeGroup.Marshal(b, m, deterministic)
}
func (m *Request_SomeGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_SomeGroup.Merge(m, src)
}
func (m *Request_SomeGroup) XXX_Size() int {
	return xxx_messageInfo_Request_SomeGroup.Size(m)
}
func (m *Request_SomeGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_SomeGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Request_SomeGroup proto.InternalMessageInfo

func (m *Request_SomeGroup) GetGroupField() int32 {
	if m != nil && m.GroupField != nil {
		return *m.GroupField
	}
	return 0
}

type Reply struct {
	Found                        []*Reply_Entry `protobuf:"bytes,1,rep,name=found" json:"found,omitempty"`
	CompactKeys                  []int32        `protobuf:"varint,2,rep,packed,name=compact_keys,json=compactKeys" json:"compact_keys,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1}
}

var extRange_Reply = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*Reply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Reply
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetFound() []*Reply_Entry {
	if m != nil {
		return m.Found
	}
	return nil
}

func (m *Reply) GetCompactKeys() []int32 {
	if m != nil {
		return m.CompactKeys
	}
	return nil
}

type Reply_Entry struct {
	KeyThatNeeds_1234Camel_CasIng *int64   `protobuf:"varint,1,req,name=key_that_needs_1234camel_CasIng,json=keyThatNeeds1234camelCasIng" json:"key_that_needs_1234camel_CasIng,omitempty"`
	Value                         *int64   `protobuf:"varint,2,opt,name=value,def=7" json:"value,omitempty"`
	XMyFieldName_2                *int64   `protobuf:"varint,3,opt,name=_my_field_name_2,json=MyFieldName2" json:"_my_field_name_2,omitempty"`
	XXX_NoUnkeyedLiteral          struct{} `json:"-"`
	XXX_unrecognized              []byte   `json:"-"`
	XXX_sizecache                 int32    `json:"-"`
}

func (m *Reply_Entry) Reset()         { *m = Reply_Entry{} }
func (m *Reply_Entry) String() string { return proto.CompactTextString(m) }
func (*Reply_Entry) ProtoMessage()    {}
func (*Reply_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{1, 0}
}
func (m *Reply_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply_Entry.Unmarshal(m, b)
}
func (m *Reply_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply_Entry.Marshal(b, m, deterministic)
}
func (m *Reply_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply_Entry.Merge(m, src)
}
func (m *Reply_Entry) XXX_Size() int {
	return xxx_messageInfo_Reply_Entry.Size(m)
}
func (m *Reply_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Reply_Entry proto.InternalMessageInfo

const Default_Reply_Entry_Value int64 = 7

func (m *Reply_Entry) GetKeyThatNeeds_1234Camel_CasIng() int64 {
	if m != nil && m.KeyThatNeeds_1234Camel_CasIng != nil {
		return *m.KeyThatNeeds_1234Camel_CasIng
	}
	return 0
}

func (m *Reply_Entry) GetValue() int64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return Default_Reply_Entry_Value
}

func (m *Reply_Entry) GetXMyFieldName_2() int64 {
	if m != nil && m.XMyFieldName_2 != nil {
		return *m.XMyFieldName_2
	}
	return 0
}

type OtherBase struct {
	Name                         *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *OtherBase) Reset()         { *m = OtherBase{} }
func (m *OtherBase) String() string { return proto.CompactTextString(m) }
func (*OtherBase) ProtoMessage()    {}
func (*OtherBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{2}
}

var extRange_OtherBase = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*OtherBase) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OtherBase
}

func (m *OtherBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherBase.Unmarshal(m, b)
}
func (m *OtherBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherBase.Marshal(b, m, deterministic)
}
func (m *OtherBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherBase.Merge(m, src)
}
func (m *OtherBase) XXX_Size() int {
	return xxx_messageInfo_OtherBase.Size(m)
}
func (m *OtherBase) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherBase.DiscardUnknown(m)
}

var xxx_messageInfo_OtherBase proto.InternalMessageInfo

func (m *OtherBase) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ReplyExtensions struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyExtensions) Reset()         { *m = ReplyExtensions{} }
func (m *ReplyExtensions) String() string { return proto.CompactTextString(m) }
func (*ReplyExtensions) ProtoMessage()    {}
func (*ReplyExtensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{3}
}
func (m *ReplyExtensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyExtensions.Unmarshal(m, b)
}
func (m *ReplyExtensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyExtensions.Marshal(b, m, deterministic)
}
func (m *ReplyExtensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyExtensions.Merge(m, src)
}
func (m *ReplyExtensions) XXX_Size() int {
	return xxx_messageInfo_ReplyExtensions.Size(m)
}
func (m *ReplyExtensions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyExtensions.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyExtensions proto.InternalMessageInfo

var E_ReplyExtensions_Time = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*float64)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.time",
	Tag:           "fixed64,101,opt,name=time",
	Filename:      "my_test/test.proto",
}

var E_ReplyExtensions_Carrot = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         105,
	Name:          "my.test.ReplyExtensions.carrot",
	Tag:           "bytes,105,opt,name=carrot",
	Filename:      "my_test/test.proto",
}

var E_ReplyExtensions_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*OtherBase)(nil),
	ExtensionType: (*ReplyExtensions)(nil),
	Field:         101,
	Name:          "my.test.ReplyExtensions.donut",
	Tag:           "bytes,101,opt,name=donut",
	Filename:      "my_test/test.proto",
}

type OtherReplyExtensions struct {
	Key                  *int32   `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OtherReplyExtensions) Reset()         { *m = OtherReplyExtensions{} }
func (m *OtherReplyExtensions) String() string { return proto.CompactTextString(m) }
func (*OtherReplyExtensions) ProtoMessage()    {}
func (*OtherReplyExtensions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{4}
}
func (m *OtherReplyExtensions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OtherReplyExtensions.Unmarshal(m, b)
}
func (m *OtherReplyExtensions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OtherReplyExtensions.Marshal(b, m, deterministic)
}
func (m *OtherReplyExtensions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OtherReplyExtensions.Merge(m, src)
}
func (m *OtherReplyExtensions) XXX_Size() int {
	return xxx_messageInfo_OtherReplyExtensions.Size(m)
}
func (m *OtherReplyExtensions) XXX_DiscardUnknown() {
	xxx_messageInfo_OtherReplyExtensions.DiscardUnknown(m)
}

var xxx_messageInfo_OtherReplyExtensions proto.InternalMessageInfo

func (m *OtherReplyExtensions) GetKey() int32 {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return 0
}

type OldReply struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `protobuf_messageset:"1" json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *OldReply) Reset()         { *m = OldReply{} }
func (m *OldReply) String() string { return proto.CompactTextString(m) }
func (*OldReply) ProtoMessage()    {}
func (*OldReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{5}
}

var extRange_OldReply = []proto.ExtensionRange{
	{Start: 100, End: 2147483646},
}

func (*OldReply) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_OldReply
}

func (m *OldReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OldReply.Unmarshal(m, b)
}
func (m *OldReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OldReply.Marshal(b, m, deterministic)
}
func (m *OldReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldReply.Merge(m, src)
}
func (m *OldReply) XXX_Size() int {
	return xxx_messageInfo_OldReply.Size(m)
}
func (m *OldReply) XXX_DiscardUnknown() {
	xxx_messageInfo_OldReply.DiscardUnknown(m)
}

var xxx_messageInfo_OldReply proto.InternalMessageInfo

type Communique struct {
	MakeMeCry *bool `protobuf:"varint,1,opt,name=make_me_cry,json=makeMeCry" json:"make_me_cry,omitempty"`
	// This is a oneof, called "union".
	//
	// Types that are valid to be assigned to Union:
	//	*Communique_Number
	//	*Communique_Name
	//	*Communique_Data
	//	*Communique_TempC
	//	*Communique_Height
	//	*Communique_Today
	//	*Communique_Maybe
	//	*Communique_Delta_
	//	*Communique_Msg
	//	*Communique_Somegroup
	Union                isCommunique_Union `protobuf_oneof:"union"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Communique) Reset()         { *m = Communique{} }
func (m *Communique) String() string { return proto.CompactTextString(m) }
func (*Communique) ProtoMessage()    {}
func (*Communique) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{6}
}
func (m *Communique) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Communique.Unmarshal(m, b)
}
func (m *Communique) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Communique.Marshal(b, m, deterministic)
}
func (m *Communique) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Communique.Merge(m, src)
}
func (m *Communique) XXX_Size() int {
	return xxx_messageInfo_Communique.Size(m)
}
func (m *Communique) XXX_DiscardUnknown() {
	xxx_messageInfo_Communique.DiscardUnknown(m)
}

var xxx_messageInfo_Communique proto.InternalMessageInfo

type isCommunique_Union interface {
	isCommunique_Union()
}

type Communique_Number struct {
	Number int32 `protobuf:"varint,5,opt,name=number,oneof" json:"number,omitempty"`
}
type Communique_Name struct {
	Name string `protobuf:"bytes,6,opt,name=name,oneof" json:"name,omitempty"`
}
type Communique_Data struct {
	Data []byte `protobuf:"bytes,7,opt,name=data,oneof" json:"data,omitempty"`
}
type Communique_TempC struct {
	TempC float64 `protobuf:"fixed64,8,opt,name=temp_c,json=tempC,oneof" json:"temp_c,omitempty"`
}
type Communique_Height struct {
	Height float32 `protobuf:"fixed32,9,opt,name=height,oneof" json:"height,omitempty"`
}
type Communique_Today struct {
	Today Days `protobuf:"varint,10,opt,name=today,enum=my.test.Days,oneof" json:"today,omitempty"`
}
type Communique_Maybe struct {
	Maybe bool `protobuf:"varint,11,opt,name=maybe,oneof" json:"maybe,omitempty"`
}
type Communique_Delta_ struct {
	Delta int32 `protobuf:"zigzag32,12,opt,name=delta,oneof" json:"delta,omitempty"`
}
type Communique_Msg struct {
	Msg *Reply `protobuf:"bytes,16,opt,name=msg,oneof" json:"msg,omitempty"`
}
type Communique_Somegroup struct {
	Somegroup *Communique_SomeGroup `protobuf:"group,14,opt,name=SomeGroup,json=somegroup,oneof" json:"somegroup,omitempty"`
}

func (*Communique_Number) isCommunique_Union()    {}
func (*Communique_Name) isCommunique_Union()      {}
func (*Communique_Data) isCommunique_Union()      {}
func (*Communique_TempC) isCommunique_Union()     {}
func (*Communique_Height) isCommunique_Union()    {}
func (*Communique_Today) isCommunique_Union()     {}
func (*Communique_Maybe) isCommunique_Union()     {}
func (*Communique_Delta_) isCommunique_Union()    {}
func (*Communique_Msg) isCommunique_Union()       {}
func (*Communique_Somegroup) isCommunique_Union() {}

func (m *Communique) GetUnion() isCommunique_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (m *Communique) GetMakeMeCry() bool {
	if m != nil && m.MakeMeCry != nil {
		return *m.MakeMeCry
	}
	return false
}

func (m *Communique) GetNumber() int32 {
	if x, ok := m.GetUnion().(*Communique_Number); ok {
		return x.Number
	}
	return 0
}

func (m *Communique) GetName() string {
	if x, ok := m.GetUnion().(*Communique_Name); ok {
		return x.Name
	}
	return ""
}

func (m *Communique) GetData() []byte {
	if x, ok := m.GetUnion().(*Communique_Data); ok {
		return x.Data
	}
	return nil
}

func (m *Communique) GetTempC() float64 {
	if x, ok := m.GetUnion().(*Communique_TempC); ok {
		return x.TempC
	}
	return 0
}

func (m *Communique) GetHeight() float32 {
	if x, ok := m.GetUnion().(*Communique_Height); ok {
		return x.Height
	}
	return 0
}

func (m *Communique) GetToday() Days {
	if x, ok := m.GetUnion().(*Communique_Today); ok {
		return x.Today
	}
	return Days_MONDAY
}

func (m *Communique) GetMaybe() bool {
	if x, ok := m.GetUnion().(*Communique_Maybe); ok {
		return x.Maybe
	}
	return false
}

func (m *Communique) GetDelta() int32 {
	if x, ok := m.GetUnion().(*Communique_Delta_); ok {
		return x.Delta
	}
	return 0
}

func (m *Communique) GetMsg() *Reply {
	if x, ok := m.GetUnion().(*Communique_Msg); ok {
		return x.Msg
	}
	return nil
}

func (m *Communique) GetSomegroup() *Communique_SomeGroup {
	if x, ok := m.GetUnion().(*Communique_Somegroup); ok {
		return x.Somegroup
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Communique) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Communique_Number)(nil),
		(*Communique_Name)(nil),
		(*Communique_Data)(nil),
		(*Communique_TempC)(nil),
		(*Communique_Height)(nil),
		(*Communique_Today)(nil),
		(*Communique_Maybe)(nil),
		(*Communique_Delta_)(nil),
		(*Communique_Msg)(nil),
		(*Communique_Somegroup)(nil),
	}
}

type Communique_SomeGroup struct {
	Member               *string  `protobuf:"bytes,15,opt,name=member" json:"member,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Communique_SomeGroup) Reset()         { *m = Communique_SomeGroup{} }
func (m *Communique_SomeGroup) String() string { return proto.CompactTextString(m) }
func (*Communique_SomeGroup) ProtoMessage()    {}
func (*Communique_SomeGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{6, 0}
}
func (m *Communique_SomeGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Communique_SomeGroup.Unmarshal(m, b)
}
func (m *Communique_SomeGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Communique_SomeGroup.Marshal(b, m, deterministic)
}
func (m *Communique_SomeGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Communique_SomeGroup.Merge(m, src)
}
func (m *Communique_SomeGroup) XXX_Size() int {
	return xxx_messageInfo_Communique_SomeGroup.Size(m)
}
func (m *Communique_SomeGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_Communique_SomeGroup.DiscardUnknown(m)
}

var xxx_messageInfo_Communique_SomeGroup proto.InternalMessageInfo

func (m *Communique_SomeGroup) GetMember() string {
	if m != nil && m.Member != nil {
		return *m.Member
	}
	return ""
}

type Communique_Delta struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Communique_Delta) Reset()         { *m = Communique_Delta{} }
func (m *Communique_Delta) String() string { return proto.CompactTextString(m) }
func (*Communique_Delta) ProtoMessage()    {}
func (*Communique_Delta) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c9b60a40d5131b9, []int{6, 1}
}
func (m *Communique_Delta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Communique_Delta.Unmarshal(m, b)
}
func (m *Communique_Delta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Communique_Delta.Marshal(b, m, deterministic)
}
func (m *Communique_Delta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Communique_Delta.Merge(m, src)
}
func (m *Communique_Delta) XXX_Size() int {
	return xxx_messageInfo_Communique_Delta.Size(m)
}
func (m *Communique_Delta) XXX_DiscardUnknown() {
	xxx_messageInfo_Communique_Delta.DiscardUnknown(m)
}

var xxx_messageInfo_Communique_Delta proto.InternalMessageInfo

var E_Tag = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*string)(nil),
	Field:         103,
	Name:          "my.test.tag",
	Tag:           "bytes,103,opt,name=tag",
	Filename:      "my_test/test.proto",
}

var E_Donut = &proto.ExtensionDesc{
	ExtendedType:  (*Reply)(nil),
	ExtensionType: (*OtherReplyExtensions)(nil),
	Field:         106,
	Name:          "my.test.donut",
	Tag:           "bytes,106,opt,name=donut",
	Filename:      "my_test/test.proto",
}

func init() {
	proto.RegisterEnum("my.test.HatType", HatType_name, HatType_value)
	proto.RegisterEnum("my.test.Days", Days_name, Days_value)
	proto.RegisterEnum("my.test.Request_Color", Request_Color_name, Request_Color_value)
	proto.RegisterEnum("my.test.Reply_Entry_Game", Reply_Entry_Game_name, Reply_Entry_Game_value)
	proto.RegisterType((*Request)(nil), "my.test.Request")
	proto.RegisterMapType((map[int64]*Reply)(nil), "my.test.Request.MsgMappingEntry")
	proto.RegisterMapType((map[int32]string)(nil), "my.test.Request.NameMappingEntry")
	proto.RegisterType((*Request_SomeGroup)(nil), "my.test.Request.SomeGroup")
	proto.RegisterType((*Reply)(nil), "my.test.Reply")
	proto.RegisterType((*Reply_Entry)(nil), "my.test.Reply.Entry")
	proto.RegisterType((*OtherBase)(nil), "my.test.OtherBase")
	proto.RegisterExtension(E_ReplyExtensions_Time)
	proto.RegisterExtension(E_ReplyExtensions_Carrot)
	proto.RegisterExtension(E_ReplyExtensions_Donut)
	proto.RegisterType((*ReplyExtensions)(nil), "my.test.ReplyExtensions")
	proto.RegisterType((*OtherReplyExtensions)(nil), "my.test.OtherReplyExtensions")
	proto.RegisterType((*OldReply)(nil), "my.test.OldReply")
	proto.RegisterType((*Communique)(nil), "my.test.Communique")
	proto.RegisterType((*Communique_SomeGroup)(nil), "my.test.Communique.SomeGroup")
	proto.RegisterType((*Communique_Delta)(nil), "my.test.Communique.Delta")
	proto.RegisterExtension(E_Tag)
	proto.RegisterExtension(E_Donut)
}

func init() { proto.RegisterFile("my_test/test.proto", fileDescriptor_2c9b60a40d5131b9) }

var fileDescriptor_2c9b60a40d5131b9 = []byte{
	// 1145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xdd, 0x6e, 0xdb, 0x46,
	0x13, 0x35, 0x49, 0x51, 0x3f, 0x23, 0x7f, 0x36, 0xbf, 0x85, 0x6b, 0xb3, 0x2a, 0x92, 0xb0, 0x4a,
	0x5c, 0x28, 0x6e, 0x23, 0xc7, 0x6a, 0x81, 0xa6, 0x2a, 0x1a, 0xc4, 0xb2, 0xe8, 0x38, 0x88, 0x2d,
	0x17, 0x1b, 0xe7, 0xa2, 0xb9, 0x21, 0x68, 0x71, 0x45, 0xb1, 0xd6, 0x92, 0x8c, 0xb8, 0x2c, 0xcc,
	0x3b, 0x3f, 0x45, 0xfb, 0x1a, 0xbd, 0xef, 0x33, 0xf4, 0x99, 0x5c, 0xec, 0xac, 0x22, 0xc9, 0x56,
	0x51, 0x5d, 0x10, 0x9c, 0x33, 0x67, 0xce, 0x72, 0x67, 0x66, 0x67, 0x05, 0x84, 0x17, 0x9e, 0x60,
	0x99, 0xd8, 0x97, 0x8f, 0x76, 0x3a, 0x4d, 0x44, 0x42, 0x2a, 0xbc, 0x68, 0x4b, 0xb3, 0x41, 0x78,
	0x3e, 0x11, 0xd1, 0x3e, 0x3e, 0x0f, 0x94, 0xb3, 0xf9, 0x77, 0x19, 0x2a, 0x94, 0x7d, 0xcc, 0x59,
	0x26, 0x88, 0x05, 0xc6, 0x15, 0x2b, 0x6c, 0xcd, 0x31, 0x5a, 0x06, 0x95, 0xaf, 0xa4, 0x05, 0xc6,
	0x38, 0x67, 0xb6, 0xe1, 0x68, 0xad, 0x8d, 0xce, 0x76, 0x7b, 0x26, 0xd4, 0x9e, 0x05, 0xb4, 0x8f,
	0x92, 0x49, 0x32, 0xa5, 0x92, 0x42, 0xf6, 0xc0, 0x18, 0xfb, 0xc2, 0x2e, 0x21, 0xd3, 0x9a, 0x33,
	0x4f, 0x7c, 0x71, 0x51, 0xa4, 0xac, 0x5b, 0x3e, 0x76, 0xfb, 0xe7, 0xf4, 0x90, 0x4a, 0x12, 0x79,
	0x04, 0xd5, 0x80, 0xf9, 0xc1, 0x24, 0x8a, 0x99, 0x5d, 0x71, 0xb4, 0x96, 0xde, 0x35, 0xa2, 0x78,
	0x44, 0xe7, 0x20, 0x79, 0x01, 0xb5, 0x2c, 0xe1, 0x2c, 0x9c, 0x26, 0x79, 0x6a, 0x57, 0x1d, 0xad,
	0x05, 0x9d, 0xc6, 0xca, 0xe2, 0xef, 0x12, 0xce, 0x5e, 0x4b, 0x06, 0x5d, 0x90, 0x49, 0x1f, 0xd6,
	0x63, 0x9f, 0x33, 0x8f, 0xfb, 0x69, 0x1a, 0xc5, 0xa1, 0xbd, 0xe1, 0x18, 0xad, 0x7a, 0xe7, 0xcb,
	0x95, 0xe0, 0x81, 0xcf, 0xd9, 0x99, 0xe2, 0xb8, 0xb1, 0x98, 0x16, 0xb4, 0x1e, 0x2f, 0x10, 0x72,
	0x08, 0x75, 0x9e, 0x85, 0x73, 0x91, 0x4d, 0x14, 0x71, 0x56, 0x44, 0xce, 0xb2, 0xf0, 0x8e, 0x06,
	0xf0, 0x39, 0x40, 0xb6, 0xc0, 0x9c, 0xb2, 0x8c, 0x09, 0x7b, 0xdd, 0xd1, 0x5a, 0x26, 0x55, 0x06,
	0xd9, 0x81, 0x4a, 0xc8, 0x84, 0x27, 0xb3, 0x6c, 0x39, 0x5a, 0xab, 0x46, 0xcb, 0x21, 0x13, 0x6f,
	0x59, 0x41, 0x1e, 0x03, 0x8c, 0x26, 0x89, 0x2f, 0xbc, 0x38, 0x8a, 0x47, 0xf6, 0x16, 0x26, 0xa5,
	0xf4, 0x4c, 0x66, 0xa5, 0x86, 0xf8, 0x20, 0x8a, 0x47, 0xa4, 0xf9, 0x89, 0x94, 0x4a, 0xd2, 0x67,
	0x8b, 0xcc, 0x29, 0xce, 0xcf, 0x8a, 0xa3, 0x0c, 0x8f, 0x5d, 0xa7, 0xf6, 0x36, 0x52, 0xcc, 0x03,
	0xf6, 0xf5, 0xf3, 0x1f, 0x68, 0x15, 0x71, 0xf7, 0x3a, 0x25, 0xbb, 0x50, 0x0f, 0x92, 0xfc, 0x72,
	0xc2, 0xd4, 0x6a, 0x3b, 0x8e, 0xd6, 0xd2, 0x66, 0xab, 0x81, 0x72, 0xe0, 0x72, 0x4f, 0xe6, 0x34,
	0x5c, 0xcf, 0x46, 0x9a, 0xb1, 0xc4, 0xc2, 0x05, 0x9f, 0xc2, 0xcc, 0xc2, 0x15, 0x3f, 0x47, 0x12,
	0x1c, 0x3c, 0xff, 0xf4, 0xa3, 0x35, 0xe5, 0x75, 0xaf, 0xd3, 0xc6, 0x37, 0x50, 0x9b, 0x17, 0x8d,
	0x3c, 0x82, 0x3a, 0x96, 0xcc, 0x1b, 0x45, 0x6c, 0x12, 0xd8, 0x35, 0x4c, 0x13, 0x20, 0x74, 0x2c,
	0x91, 0xc6, 0x4b, 0xb0, 0xee, 0x57, 0x69, 0xd1, 0xa1, 0x92, 0x8c, 0x1d, 0xba, 0x05, 0xe6, 0x6f,
	0xfe, 0x24, 0x67, 0xb6, 0x8e, 0xf9, 0x54, 0x46, 0x57, 0x7f, 0xa1, 0x35, 0xce, 0x60, 0xf3, 0x5e,
	0x81, 0x96, 0xc3, 0x89, 0x0a, 0x7f, 0xb2, 0x1c, 0x5e, 0xef, 0x6c, 0x2c, 0xd5, 0x38, 0x9d, 0x14,
	0x4b, 0x72, 0xcd, 0x5d, 0x30, 0xb1, 0xdd, 0x49, 0x05, 0x0c, 0xea, 0xf6, 0xad, 0x35, 0x52, 0x03,
	0xf3, 0x35, 0x75, 0xdd, 0x81, 0xa5, 0x91, 0x2a, 0x94, 0x7a, 0xa7, 0xef, 0x5d, 0x4b, 0x6f, 0xfe,
	0xa1, 0x83, 0x89, 0xb1, 0x64, 0x0f, 0xcc, 0x51, 0x92, 0xc7, 0x01, 0x9e, 0xa7, 0x7a, 0x67, 0xeb,
	0xae, 0x74, 0x5b, 0xb5, 0x8c, 0xa2, 0x90, 0x5d, 0x58, 0x1f, 0x26, 0x3c, 0xf5, 0x87, 0xd8, 0x1b,
	0x99, 0xad, 0x3b, 0x46, 0xcb, 0xec, 0xe9, 0x96, 0x46, 0xeb, 0x33, 0xfc, 0x2d, 0x2b, 0xb2, 0xc6,
	0x9f, 0x1a, 0x98, 0x6a, 0x27, 0x7d, 0x78, 0x74, 0xc5, 0x0a, 0x4f, 0x8c, 0x65, 0xcb, 0x30, 0x16,
	0x64, 0xde, 0x41, 0xe7, 0xdb, 0xef, 0x86, 0x3e, 0x67, 0x13, 0xef, 0xc8, 0xcf, 0xde, 0xc4, 0xa1,
	0xad, 0x39, 0x7a, 0xcb, 0xa0, 0x5f, 0x5c, 0xb1, 0xe2, 0x62, 0xec, 0x8b, 0x81, 0x24, 0xcd, 0x39,
	0x8a, 0x42, 0x76, 0x96, 0x77, 0x6f, 0x74, 0xb5, 0xef, 0x67, 0x1b, 0x26, 0x5f, 0x81, 0xe5, 0xf1,
	0x42, 0x95, 0xc6, 0xc3, 0x03, 0xd5, 0xc1, 0x21, 0x60, 0xd0, 0xf5, 0xb3, 0x02, 0xcb, 0x23, 0x4b,
	0xd3, 0x69, 0x3a, 0x50, 0x7a, 0xed, 0x73, 0x46, 0xd6, 0xa1, 0x7a, 0x7c, 0x7e, 0x7e, 0xd1, 0x3b,
	0x3c, 0x3d, 0xb5, 0x34, 0x02, 0x50, 0xbe, 0x70, 0x07, 0x83, 0x37, 0xef, 0x2c, 0x7d, 0xaf, 0x5a,
	0x0d, 0xac, 0x9b, 0x9b, 0x9b, 0x1b, 0xbd, 0xf9, 0x14, 0x6a, 0xe7, 0x62, 0xcc, 0xa6, 0x3d, 0x3f,
	0x63, 0x84, 0x40, 0x49, 0xca, 0x62, 0x29, 0x6a, 0x14, 0xdf, 0x97, 0xa8, 0x7f, 0x69, 0xb0, 0x89,
	0x59, 0x72, 0xaf, 0x05, 0x8b, 0xb3, 0x28, 0x89, 0xb3, 0x4e, 0x13, 0x4a, 0x22, 0xe2, 0x8c, 0xdc,
	0x2b, 0x91, 0xcd, 0x64, 0xc7, 0x51, 0xf4, 0x75, 0x5e, 0x41, 0x79, 0xe8, 0x4f, 0xa7, 0x89, 0x58,
	0x61, 0x45, 0x58, 0x5e, 0xfb, 0x2e, 0xba, 0x50, 0xa7, 0xb3, 0xb8, 0x4e, 0x0f, 0xcc, 0x20, 0x89,
	0x73, 0x41, 0xc8, 0x9c, 0x3a, 0xff, 0x68, 0x5c, 0xea, 0xbf, 0x44, 0x54, 0x68, 0xb3, 0x05, 0x5b,
	0x18, 0x73, 0xcf, 0xbd, 0xda, 0xbc, 0x4d, 0x1b, 0xaa, 0xe7, 0x93, 0x00, 0x79, 0xb8, 0xfb, 0xdb,
	0xdb, 0xdb, 0xdb, 0x4a, 0x57, 0xaf, 0x6a, 0xcd, 0xdf, 0x0d, 0x80, 0xa3, 0x84, 0xf3, 0x3c, 0x8e,
	0x3e, 0xe6, 0x8c, 0x3c, 0x84, 0x3a, 0xf7, 0xaf, 0x98, 0xc7, 0x99, 0x37, 0x9c, 0x2a, 0x89, 0x2a,
	0xad, 0x49, 0xe8, 0x8c, 0x1d, 0x4d, 0x0b, 0x62, 0x43, 0x39, 0xce, 0xf9, 0x25, 0x9b, 0xda, 0xa6,
	0x54, 0x3f, 0x59, 0xa3, 0x33, 0x9b, 0x6c, 0xcd, 0x12, 0x5d, 0x96, 0x89, 0x3e, 0x59, 0x53, 0xa9,
	0x96, 0x68, 0xe0, 0x0b, 0x1f, 0xa7, 0xef, 0xba, 0x44, 0xa5, 0x45, 0x76, 0xa0, 0x2c, 0x18, 0x4f,
	0xbd, 0x21, 0xce, 0x5c, 0xed, 0x64, 0x8d, 0x9a, 0xd2, 0x3e, 0x92, 0xf2, 0x63, 0x16, 0x85, 0x63,
	0x81, 0xc7, 0x54, 0x97, 0xf2, 0xca, 0x26, 0xbb, 0x60, 0x8a, 0x24, 0xf0, 0x0b, 0x1b, 0x70, 0xf0,
	0xff, 0x6f, 0x9e, 0x9b, 0xbe, 0x5f, 0x64, 0x28, 0x20, 0xbd, 0x64, 0x1b, 0x4c, 0xee, 0x17, 0x97,
	0xcc, 0xae, 0xcb, 0x2f, 0x97, 0x38, 0x9a, 0x12, 0x0f, 0xd8, 0x44, 0xf8, 0x38, 0x25, 0xff, 0x2f,
	0x71, 0x34, 0x49, 0x13, 0x0c, 0x9e, 0x85, 0x38, 0x23, 0x57, 0x0e, 0xe5, 0xc9, 0x1a, 0x95, 0x4e,
	0xf2, 0xd3, 0xf2, 0x25, 0xb1, 0x81, 0x97, 0xc4, 0x83, 0x39, 0x73, 0x91, 0xbb, 0xc5, 0x3d, 0x71,
	0xb2, 0xb6, 0x74, 0x53, 0x34, 0x1e, 0x2f, 0x0f, 0xa3, 0x6d, 0x28, 0x73, 0x86, 0xf9, 0xdb, 0x54,
	0x63, 0x59, 0x59, 0x8d, 0x0a, 0x98, 0x7d, 0xf9, 0x41, 0xbd, 0x0a, 0x98, 0x79, 0x1c, 0x25, 0xf1,
	0xde, 0x43, 0xa8, 0xcc, 0xee, 0x34, 0xd9, 0xe6, 0xea, 0x56, 0xb3, 0x34, 0x39, 0x14, 0x8e, 0xdd,
	0x0f, 0x96, 0xbe, 0xd7, 0x86, 0x92, 0xdc, 0xba, 0x74, 0x9e, 0x9d, 0x0f, 0xfa, 0x87, 0xbf, 0x58,
	0x1a, 0xa9, 0x43, 0xe5, 0xe2, 0xbd, 0xfb, 0x4e, 0x1a, 0xba, 0x9c, 0x1a, 0xa7, 0xef, 0x07, 0xfd,
	0x37, 0x96, 0xd6, 0xd0, 0x2d, 0xad, 0xeb, 0x80, 0x21, 0xfc, 0x70, 0xa5, 0x5f, 0x43, 0xfc, 0x0c,
	0xe9, 0xea, 0x1e, 0x7d, 0x6a, 0xc9, 0xfb, 0x9c, 0x5f, 0x31, 0x3b, 0x0f, 0xee, 0x36, 0xea, 0xbf,
	0xf7, 0x64, 0xef, 0xd5, 0x87, 0x97, 0x61, 0x24, 0xc6, 0xf9, 0x65, 0x7b, 0x98, 0xf0, 0xfd, 0x30,
	0x09, 0x93, 0x7d, 0xbc, 0xff, 0x2f, 0xf3, 0x91, 0x7a, 0x19, 0x3e, 0x0b, 0x59, 0xfc, 0x0c, 0x1d,
	0x52, 0x4c, 0xf6, 0xc3, 0xfe, 0xec, 0xbf, 0xc4, 0x8f, 0xf2, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0b, 0xc8, 0x22, 0xdb, 0x5a, 0x08, 0x00, 0x00,
}
