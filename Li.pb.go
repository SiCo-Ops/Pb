// Code generated by protoc-gen-go.
// source: Li.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CloudResponse struct {
	Code int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *CloudResponse) Reset()                    { *m = CloudResponse{} }
func (m *CloudResponse) String() string            { return proto.CompactTextString(m) }
func (*CloudResponse) ProtoMessage()               {}
func (*CloudResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CloudResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CloudResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CloudResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudRequest struct {
	Bsns     string            `protobuf:"bytes,1,opt,name=bsns" json:"bsns,omitempty"`
	Action   string            `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Region   string            `protobuf:"bytes,3,opt,name=region" json:"region,omitempty"`
	CloudId  string            `protobuf:"bytes,4,opt,name=cloudId" json:"cloudId,omitempty"`
	CloudKey string            `protobuf:"bytes,5,opt,name=cloudKey" json:"cloudKey,omitempty"`
	Id       string            `protobuf:"bytes,6,opt,name=id" json:"id,omitempty"`
	Params   map[string]string `protobuf:"bytes,7,rep,name=params" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudRequest) Reset()                    { *m = CloudRequest{} }
func (m *CloudRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudRequest) ProtoMessage()               {}
func (*CloudRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CloudRequest) GetBsns() string {
	if m != nil {
		return m.Bsns
	}
	return ""
}

func (m *CloudRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *CloudRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *CloudRequest) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

func (m *CloudRequest) GetCloudKey() string {
	if m != nil {
		return m.CloudKey
	}
	return ""
}

func (m *CloudRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CloudRequest) GetParams() map[string]string {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*CloudResponse)(nil), "pb.Cloud_response")
	proto.RegisterType((*CloudRequest)(nil), "pb.Cloud_request")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cloud_API service

type Cloud_APIClient interface {
	Qcloud(ctx context.Context, in *CloudRequest, opts ...grpc.CallOption) (*CloudResponse, error)
}

type cloud_APIClient struct {
	cc *grpc.ClientConn
}

func NewCloud_APIClient(cc *grpc.ClientConn) Cloud_APIClient {
	return &cloud_APIClient{cc}
}

func (c *cloud_APIClient) Qcloud(ctx context.Context, in *CloudRequest, opts ...grpc.CallOption) (*CloudResponse, error) {
	out := new(CloudResponse)
	err := grpc.Invoke(ctx, "/pb.Cloud_API/qcloud", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cloud_API service

type Cloud_APIServer interface {
	Qcloud(context.Context, *CloudRequest) (*CloudResponse, error)
}

func RegisterCloud_APIServer(s *grpc.Server, srv Cloud_APIServer) {
	s.RegisterService(&_Cloud_API_serviceDesc, srv)
}

func _Cloud_API_Qcloud_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Cloud_APIServer).Qcloud(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Cloud_API/Qcloud",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Cloud_APIServer).Qcloud(ctx, req.(*CloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cloud_API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Cloud_API",
	HandlerType: (*Cloud_APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "qcloud",
			Handler:    _Cloud_API_Qcloud_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Li.proto",
}

func init() { proto.RegisterFile("Li.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xff, 0xd9, 0x34, 0x69, 0x33, 0xfd, 0x5b, 0x74, 0x10, 0x59, 0x02, 0x42, 0xc8, 0x29,
	0xa7, 0x80, 0x15, 0x41, 0x3d, 0x08, 0x22, 0x1e, 0xaa, 0x1e, 0x4a, 0x5e, 0x40, 0x36, 0xd9, 0xa5,
	0x04, 0xdb, 0x6c, 0x9a, 0x4d, 0x84, 0x3c, 0x83, 0x2f, 0x2d, 0x3b, 0x49, 0xa5, 0xe2, 0xed, 0xfb,
	0x7e, 0x33, 0x7c, 0x3b, 0xb3, 0x03, 0xb3, 0xb7, 0x32, 0xad, 0x1b, 0xdd, 0x6a, 0x64, 0x75, 0x1e,
	0xbf, 0xc0, 0xe2, 0x69, 0xab, 0x3b, 0xf9, 0xde, 0x28, 0x53, 0xeb, 0xca, 0x28, 0x44, 0x98, 0x14,
	0x5a, 0x2a, 0xee, 0x44, 0x4e, 0xe2, 0x65, 0xa4, 0xf1, 0x14, 0xdc, 0x9d, 0xd9, 0x70, 0x16, 0x39,
	0x49, 0x90, 0x59, 0x69, 0xbb, 0xa4, 0x68, 0x05, 0x77, 0x23, 0x27, 0xf9, 0x9f, 0x91, 0x8e, 0xbf,
	0x18, 0x9c, 0x1c, 0xc2, 0xf6, 0x9d, 0x32, 0xad, 0xed, 0xca, 0x4d, 0x65, 0x28, 0x2b, 0xc8, 0x48,
	0xe3, 0x05, 0xf8, 0xa2, 0x68, 0x4b, 0x5d, 0x8d, 0x71, 0xa3, 0xb3, 0xbc, 0x51, 0x1b, 0xcb, 0xdd,
	0x81, 0x0f, 0x0e, 0x39, 0x4c, 0x0b, 0x1b, 0xba, 0x92, 0x7c, 0x42, 0x85, 0x83, 0xc5, 0x10, 0x66,
	0x24, 0x5f, 0x55, 0xcf, 0x3d, 0x2a, 0xfd, 0x78, 0x5c, 0x00, 0x2b, 0x25, 0xf7, 0x89, 0xb2, 0x52,
	0xe2, 0x0d, 0xf8, 0xb5, 0x68, 0xc4, 0xce, 0xf0, 0x69, 0xe4, 0x26, 0xf3, 0xe5, 0x65, 0x5a, 0xe7,
	0xe9, 0xaf, 0x61, 0xd3, 0x35, 0xd5, 0x9f, 0xab, 0xb6, 0xe9, 0xb3, 0xb1, 0x39, 0xbc, 0x83, 0xf9,
	0x11, 0xb6, 0xff, 0xf0, 0xa1, 0xfa, 0x71, 0x1d, 0x2b, 0xf1, 0x1c, 0xbc, 0x4f, 0xb1, 0xed, 0xd4,
	0xb8, 0xcc, 0x60, 0xee, 0xd9, 0xad, 0xb3, 0x7c, 0x80, 0x60, 0xc8, 0x7f, 0x5c, 0xaf, 0xf0, 0x0a,
	0xfc, 0x3d, 0xcd, 0x86, 0x67, 0x7f, 0x1e, 0x0e, 0xf1, 0x18, 0x0d, 0x57, 0x88, 0xff, 0xe5, 0x3e,
	0x1d, 0xe9, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x6f, 0x9b, 0x26, 0xb0, 0x01, 0x00, 0x00,
}
