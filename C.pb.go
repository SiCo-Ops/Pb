// Code generated by protoc-gen-go. DO NOT EDIT.
// source: C.proto

package pb

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

type HookCreateCall struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hooktype             string   `protobuf:"bytes,2,opt,name=hooktype,proto3" json:"hooktype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookCreateCall) Reset()         { *m = HookCreateCall{} }
func (m *HookCreateCall) String() string { return proto.CompactTextString(m) }
func (*HookCreateCall) ProtoMessage()    {}
func (*HookCreateCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{0}
}

func (m *HookCreateCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookCreateCall.Unmarshal(m, b)
}
func (m *HookCreateCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookCreateCall.Marshal(b, m, deterministic)
}
func (m *HookCreateCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookCreateCall.Merge(m, src)
}
func (m *HookCreateCall) XXX_Size() int {
	return xxx_messageInfo_HookCreateCall.Size(m)
}
func (m *HookCreateCall) XXX_DiscardUnknown() {
	xxx_messageInfo_HookCreateCall.DiscardUnknown(m)
}

var xxx_messageInfo_HookCreateCall proto.InternalMessageInfo

func (m *HookCreateCall) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HookCreateCall) GetHooktype() string {
	if m != nil {
		return m.Hooktype
	}
	return ""
}

type HookCreateBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Hookname             string   `protobuf:"bytes,2,opt,name=hookname,proto3" json:"hookname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookCreateBack) Reset()         { *m = HookCreateBack{} }
func (m *HookCreateBack) String() string { return proto.CompactTextString(m) }
func (*HookCreateBack) ProtoMessage()    {}
func (*HookCreateBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{1}
}

func (m *HookCreateBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookCreateBack.Unmarshal(m, b)
}
func (m *HookCreateBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookCreateBack.Marshal(b, m, deterministic)
}
func (m *HookCreateBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookCreateBack.Merge(m, src)
}
func (m *HookCreateBack) XXX_Size() int {
	return xxx_messageInfo_HookCreateBack.Size(m)
}
func (m *HookCreateBack) XXX_DiscardUnknown() {
	xxx_messageInfo_HookCreateBack.DiscardUnknown(m)
}

var xxx_messageInfo_HookCreateBack proto.InternalMessageInfo

func (m *HookCreateBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HookCreateBack) GetHookname() string {
	if m != nil {
		return m.Hookname
	}
	return ""
}

type HookQueryCall struct {
	Hookname             string   `protobuf:"bytes,1,opt,name=hookname,proto3" json:"hookname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookQueryCall) Reset()         { *m = HookQueryCall{} }
func (m *HookQueryCall) String() string { return proto.CompactTextString(m) }
func (*HookQueryCall) ProtoMessage()    {}
func (*HookQueryCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{2}
}

func (m *HookQueryCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookQueryCall.Unmarshal(m, b)
}
func (m *HookQueryCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookQueryCall.Marshal(b, m, deterministic)
}
func (m *HookQueryCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookQueryCall.Merge(m, src)
}
func (m *HookQueryCall) XXX_Size() int {
	return xxx_messageInfo_HookQueryCall.Size(m)
}
func (m *HookQueryCall) XXX_DiscardUnknown() {
	xxx_messageInfo_HookQueryCall.DiscardUnknown(m)
}

var xxx_messageInfo_HookQueryCall proto.InternalMessageInfo

func (m *HookQueryCall) GetHookname() string {
	if m != nil {
		return m.Hookname
	}
	return ""
}

type HookQueryBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Hookid               string   `protobuf:"bytes,2,opt,name=hookid,proto3" json:"hookid,omitempty"`
	Hooktype             string   `protobuf:"bytes,3,opt,name=hooktype,proto3" json:"hooktype,omitempty"`
	Belong               string   `protobuf:"bytes,4,opt,name=belong,proto3" json:"belong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookQueryBack) Reset()         { *m = HookQueryBack{} }
func (m *HookQueryBack) String() string { return proto.CompactTextString(m) }
func (*HookQueryBack) ProtoMessage()    {}
func (*HookQueryBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{3}
}

func (m *HookQueryBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookQueryBack.Unmarshal(m, b)
}
func (m *HookQueryBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookQueryBack.Marshal(b, m, deterministic)
}
func (m *HookQueryBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookQueryBack.Merge(m, src)
}
func (m *HookQueryBack) XXX_Size() int {
	return xxx_messageInfo_HookQueryBack.Size(m)
}
func (m *HookQueryBack) XXX_DiscardUnknown() {
	xxx_messageInfo_HookQueryBack.DiscardUnknown(m)
}

var xxx_messageInfo_HookQueryBack proto.InternalMessageInfo

func (m *HookQueryBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HookQueryBack) GetHookid() string {
	if m != nil {
		return m.Hookid
	}
	return ""
}

func (m *HookQueryBack) GetHooktype() string {
	if m != nil {
		return m.Hooktype
	}
	return ""
}

func (m *HookQueryBack) GetBelong() string {
	if m != nil {
		return m.Belong
	}
	return ""
}

type HookAuthCall struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Hookname             string   `protobuf:"bytes,2,opt,name=hookname,proto3" json:"hookname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookAuthCall) Reset()         { *m = HookAuthCall{} }
func (m *HookAuthCall) String() string { return proto.CompactTextString(m) }
func (*HookAuthCall) ProtoMessage()    {}
func (*HookAuthCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{4}
}

func (m *HookAuthCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookAuthCall.Unmarshal(m, b)
}
func (m *HookAuthCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookAuthCall.Marshal(b, m, deterministic)
}
func (m *HookAuthCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookAuthCall.Merge(m, src)
}
func (m *HookAuthCall) XXX_Size() int {
	return xxx_messageInfo_HookAuthCall.Size(m)
}
func (m *HookAuthCall) XXX_DiscardUnknown() {
	xxx_messageInfo_HookAuthCall.DiscardUnknown(m)
}

var xxx_messageInfo_HookAuthCall proto.InternalMessageInfo

func (m *HookAuthCall) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HookAuthCall) GetHookname() string {
	if m != nil {
		return m.Hookname
	}
	return ""
}

type HookAuthBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Hookid               string   `protobuf:"bytes,2,opt,name=hookid,proto3" json:"hookid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookAuthBack) Reset()         { *m = HookAuthBack{} }
func (m *HookAuthBack) String() string { return proto.CompactTextString(m) }
func (*HookAuthBack) ProtoMessage()    {}
func (*HookAuthBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{5}
}

func (m *HookAuthBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookAuthBack.Unmarshal(m, b)
}
func (m *HookAuthBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookAuthBack.Marshal(b, m, deterministic)
}
func (m *HookAuthBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookAuthBack.Merge(m, src)
}
func (m *HookAuthBack) XXX_Size() int {
	return xxx_messageInfo_HookAuthBack.Size(m)
}
func (m *HookAuthBack) XXX_DiscardUnknown() {
	xxx_messageInfo_HookAuthBack.DiscardUnknown(m)
}

var xxx_messageInfo_HookAuthBack proto.InternalMessageInfo

func (m *HookAuthBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HookAuthBack) GetHookid() string {
	if m != nil {
		return m.Hookid
	}
	return ""
}

type HookUpdateNameCall struct {
	Hookid               string   `protobuf:"bytes,1,opt,name=hookid,proto3" json:"hookid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookUpdateNameCall) Reset()         { *m = HookUpdateNameCall{} }
func (m *HookUpdateNameCall) String() string { return proto.CompactTextString(m) }
func (*HookUpdateNameCall) ProtoMessage()    {}
func (*HookUpdateNameCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{6}
}

func (m *HookUpdateNameCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookUpdateNameCall.Unmarshal(m, b)
}
func (m *HookUpdateNameCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookUpdateNameCall.Marshal(b, m, deterministic)
}
func (m *HookUpdateNameCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookUpdateNameCall.Merge(m, src)
}
func (m *HookUpdateNameCall) XXX_Size() int {
	return xxx_messageInfo_HookUpdateNameCall.Size(m)
}
func (m *HookUpdateNameCall) XXX_DiscardUnknown() {
	xxx_messageInfo_HookUpdateNameCall.DiscardUnknown(m)
}

var xxx_messageInfo_HookUpdateNameCall proto.InternalMessageInfo

func (m *HookUpdateNameCall) GetHookid() string {
	if m != nil {
		return m.Hookid
	}
	return ""
}

type HookUpdateNameBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Hookname             string   `protobuf:"bytes,2,opt,name=hookname,proto3" json:"hookname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookUpdateNameBack) Reset()         { *m = HookUpdateNameBack{} }
func (m *HookUpdateNameBack) String() string { return proto.CompactTextString(m) }
func (*HookUpdateNameBack) ProtoMessage()    {}
func (*HookUpdateNameBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{7}
}

func (m *HookUpdateNameBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookUpdateNameBack.Unmarshal(m, b)
}
func (m *HookUpdateNameBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookUpdateNameBack.Marshal(b, m, deterministic)
}
func (m *HookUpdateNameBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookUpdateNameBack.Merge(m, src)
}
func (m *HookUpdateNameBack) XXX_Size() int {
	return xxx_messageInfo_HookUpdateNameBack.Size(m)
}
func (m *HookUpdateNameBack) XXX_DiscardUnknown() {
	xxx_messageInfo_HookUpdateNameBack.DiscardUnknown(m)
}

var xxx_messageInfo_HookUpdateNameBack proto.InternalMessageInfo

func (m *HookUpdateNameBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HookUpdateNameBack) GetHookname() string {
	if m != nil {
		return m.Hookname
	}
	return ""
}

type HookReceiveCall struct {
	Hookname             string   `protobuf:"bytes,1,opt,name=hookname,proto3" json:"hookname,omitempty"`
	Hooktype             string   `protobuf:"bytes,2,opt,name=hooktype,proto3" json:"hooktype,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookReceiveCall) Reset()         { *m = HookReceiveCall{} }
func (m *HookReceiveCall) String() string { return proto.CompactTextString(m) }
func (*HookReceiveCall) ProtoMessage()    {}
func (*HookReceiveCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{8}
}

func (m *HookReceiveCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookReceiveCall.Unmarshal(m, b)
}
func (m *HookReceiveCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookReceiveCall.Marshal(b, m, deterministic)
}
func (m *HookReceiveCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookReceiveCall.Merge(m, src)
}
func (m *HookReceiveCall) XXX_Size() int {
	return xxx_messageInfo_HookReceiveCall.Size(m)
}
func (m *HookReceiveCall) XXX_DiscardUnknown() {
	xxx_messageInfo_HookReceiveCall.DiscardUnknown(m)
}

var xxx_messageInfo_HookReceiveCall proto.InternalMessageInfo

func (m *HookReceiveCall) GetHookname() string {
	if m != nil {
		return m.Hookname
	}
	return ""
}

func (m *HookReceiveCall) GetHooktype() string {
	if m != nil {
		return m.Hooktype
	}
	return ""
}

func (m *HookReceiveCall) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type HookReceiveBack struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Hookid               string   `protobuf:"bytes,2,opt,name=hookid,proto3" json:"hookid,omitempty"`
	Params               []byte   `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HookReceiveBack) Reset()         { *m = HookReceiveBack{} }
func (m *HookReceiveBack) String() string { return proto.CompactTextString(m) }
func (*HookReceiveBack) ProtoMessage()    {}
func (*HookReceiveBack) Descriptor() ([]byte, []int) {
	return fileDescriptor_598472221a30ffea, []int{9}
}

func (m *HookReceiveBack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HookReceiveBack.Unmarshal(m, b)
}
func (m *HookReceiveBack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HookReceiveBack.Marshal(b, m, deterministic)
}
func (m *HookReceiveBack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HookReceiveBack.Merge(m, src)
}
func (m *HookReceiveBack) XXX_Size() int {
	return xxx_messageInfo_HookReceiveBack.Size(m)
}
func (m *HookReceiveBack) XXX_DiscardUnknown() {
	xxx_messageInfo_HookReceiveBack.DiscardUnknown(m)
}

var xxx_messageInfo_HookReceiveBack proto.InternalMessageInfo

func (m *HookReceiveBack) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HookReceiveBack) GetHookid() string {
	if m != nil {
		return m.Hookid
	}
	return ""
}

func (m *HookReceiveBack) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*HookCreateCall)(nil), "pb.HookCreateCall")
	proto.RegisterType((*HookCreateBack)(nil), "pb.HookCreateBack")
	proto.RegisterType((*HookQueryCall)(nil), "pb.HookQueryCall")
	proto.RegisterType((*HookQueryBack)(nil), "pb.HookQueryBack")
	proto.RegisterType((*HookAuthCall)(nil), "pb.HookAuthCall")
	proto.RegisterType((*HookAuthBack)(nil), "pb.HookAuthBack")
	proto.RegisterType((*HookUpdateNameCall)(nil), "pb.HookUpdateNameCall")
	proto.RegisterType((*HookUpdateNameBack)(nil), "pb.HookUpdateNameBack")
	proto.RegisterType((*HookReceiveCall)(nil), "pb.HookReceiveCall")
	proto.RegisterType((*HookReceiveBack)(nil), "pb.HookReceiveBack")
}

func init() { proto.RegisterFile("C.proto", fileDescriptor_598472221a30ffea) }

var fileDescriptor_598472221a30ffea = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xdf, 0x76, 0x63, 0x7f, 0xce, 0xbb, 0x4d, 0x8d, 0x30, 0x4a, 0xaf, 0x24, 0x57, 0x82,
	0x32, 0x61, 0x22, 0x88, 0x78, 0xe1, 0xac, 0x17, 0x5e, 0x89, 0x56, 0xf6, 0x01, 0xd2, 0x36, 0xb8,
	0xb2, 0x6e, 0x09, 0xb5, 0x1b, 0xec, 0x0b, 0xf9, 0x39, 0xe5, 0xa4, 0xe9, 0xda, 0x94, 0x39, 0x99,
	0x77, 0x3d, 0x87, 0xe7, 0xf7, 0x9c, 0x26, 0xcf, 0x09, 0xb4, 0xbd, 0x91, 0x4c, 0x45, 0x26, 0x88,
	0x2d, 0x03, 0x7a, 0x0f, 0x83, 0x67, 0x21, 0xe6, 0x5e, 0xca, 0x59, 0xc6, 0x3d, 0x96, 0x24, 0x64,
	0x00, 0x76, 0x1c, 0x39, 0xd6, 0x99, 0x75, 0xde, 0xf5, 0xed, 0x38, 0x22, 0x2e, 0x74, 0x66, 0x42,
	0xcc, 0xb3, 0x8d, 0xe4, 0x8e, 0xad, 0xba, 0xdb, 0x9a, 0x3e, 0x54, 0xe9, 0x47, 0x16, 0xce, 0x09,
	0x81, 0x66, 0x28, 0x22, 0xae, 0xf8, 0x86, 0xaf, 0xbe, 0x0b, 0x87, 0x25, 0x5b, 0x18, 0x0e, 0x58,
	0xd3, 0x0b, 0xe8, 0xa3, 0xc3, 0xdb, 0x8a, 0xa7, 0x1b, 0x35, 0xbe, 0x2a, 0xb6, 0x6a, 0x62, 0x51,
	0x11, 0xff, 0x38, 0x6d, 0x08, 0x2d, 0x04, 0xe2, 0x48, 0xcf, 0xd2, 0x95, 0x71, 0x8e, 0x86, 0x79,
	0x0e, 0x64, 0x02, 0x9e, 0x88, 0xe5, 0x87, 0xd3, 0xcc, 0x99, 0xbc, 0xa2, 0x77, 0xd0, 0xc3, 0x81,
	0x93, 0x55, 0x36, 0xdb, 0x77, 0x37, 0x3b, 0x4f, 0x56, 0x61, 0x0f, 0xfd, 0x57, 0x7a, 0x09, 0x04,
	0xd9, 0xa9, 0x8c, 0x58, 0xc6, 0x5f, 0xd8, 0x22, 0x4f, 0xa6, 0x54, 0x5b, 0x86, 0xfa, 0xa9, 0xae,
	0xfe, 0x53, 0x12, 0x21, 0x1c, 0xa1, 0x8b, 0xcf, 0x43, 0x1e, 0xaf, 0xf9, 0x6f, 0x59, 0xec, 0x5b,
	0x0b, 0xe2, 0x40, 0x5b, 0xb2, 0x4d, 0x22, 0x58, 0xa4, 0x6e, 0xba, 0xe7, 0x17, 0x25, 0x9d, 0x1a,
	0x43, 0x0e, 0xce, 0x70, 0x08, 0x2d, 0xc9, 0x52, 0xb6, 0xf8, 0xd4, 0xbe, 0xba, 0x1a, 0x7f, 0xd9,
	0xf0, 0x1f, 0x7d, 0xdf, 0x79, 0xba, 0x8e, 0x43, 0x4e, 0xae, 0xa0, 0x8d, 0xf7, 0xee, 0xbf, 0x7a,
	0xe4, 0x78, 0x24, 0x83, 0x51, 0x35, 0x44, 0xd7, 0xe8, 0xe0, 0x2f, 0xd0, 0x7f, 0xe4, 0x06, 0xba,
	0xf9, 0x12, 0x23, 0x42, 0x0a, 0x41, 0xf9, 0x2a, 0xdc, 0x5a, 0x4f, 0x63, 0x63, 0xe8, 0xa8, 0x65,
	0x44, 0xea, 0xa4, 0x50, 0x6c, 0x77, 0xd9, 0x35, 0x5b, 0x9a, 0x99, 0x40, 0xbf, 0x4c, 0x0a, 0xc1,
	0x61, 0xa1, 0x32, 0xe3, 0x76, 0x77, 0xf4, 0xb5, 0xc5, 0x2d, 0x80, 0xbe, 0x41, 0xe4, 0x4f, 0x0b,
	0x5d, 0x25, 0x3a, 0xb7, 0xde, 0xcc, 0xc9, 0xa0, 0xa5, 0x5e, 0xfe, 0xf5, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x84, 0xb0, 0x2d, 0xe2, 0x04, 0x04, 0x00, 0x00,
}
