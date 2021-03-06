// Code generated by protoc-gen-go. DO NOT EDIT.
// source: swagger.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetSpecRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSpecRequest) Reset()         { *m = GetSpecRequest{} }
func (m *GetSpecRequest) String() string { return proto.CompactTextString(m) }
func (*GetSpecRequest) ProtoMessage()    {}
func (*GetSpecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49635b75e059a131, []int{0}
}

func (m *GetSpecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpecRequest.Unmarshal(m, b)
}
func (m *GetSpecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpecRequest.Marshal(b, m, deterministic)
}
func (m *GetSpecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpecRequest.Merge(m, src)
}
func (m *GetSpecRequest) XXX_Size() int {
	return xxx_messageInfo_GetSpecRequest.Size(m)
}
func (m *GetSpecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpecRequest proto.InternalMessageInfo

type GetSpecResponse struct {
	SwaggerSpec          string   `protobuf:"bytes,1,opt,name=swagger_spec,json=swaggerSpec,proto3" json:"swagger_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSpecResponse) Reset()         { *m = GetSpecResponse{} }
func (m *GetSpecResponse) String() string { return proto.CompactTextString(m) }
func (*GetSpecResponse) ProtoMessage()    {}
func (*GetSpecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49635b75e059a131, []int{1}
}

func (m *GetSpecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSpecResponse.Unmarshal(m, b)
}
func (m *GetSpecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSpecResponse.Marshal(b, m, deterministic)
}
func (m *GetSpecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSpecResponse.Merge(m, src)
}
func (m *GetSpecResponse) XXX_Size() int {
	return xxx_messageInfo_GetSpecResponse.Size(m)
}
func (m *GetSpecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSpecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSpecResponse proto.InternalMessageInfo

func (m *GetSpecResponse) GetSwaggerSpec() string {
	if m != nil {
		return m.SwaggerSpec
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSpecRequest)(nil), "magma.orc8r.obsidian.swagger.GetSpecRequest")
	proto.RegisterType((*GetSpecResponse)(nil), "magma.orc8r.obsidian.swagger.GetSpecResponse")
}

func init() { proto.RegisterFile("swagger.proto", fileDescriptor_49635b75e059a131) }

var fileDescriptor_49635b75e059a131 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x4f, 0x4c,
	0x4f, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xc9, 0x4d, 0x4c, 0xcf, 0x4d,
	0xd4, 0xcb, 0x2f, 0x4a, 0xb6, 0x28, 0xd2, 0xcb, 0x4f, 0x2a, 0xce, 0x4c, 0xc9, 0x4c, 0xcc, 0xd3,
	0x83, 0xaa, 0x51, 0x12, 0xe0, 0xe2, 0x73, 0x4f, 0x2d, 0x09, 0x2e, 0x48, 0x4d, 0x0e, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x32, 0xe1, 0xe2, 0x87, 0x8b, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0x29, 0x72, 0xf1, 0x40, 0xd5, 0xc7, 0x17, 0x17, 0xa4, 0x26, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x71, 0x43, 0xc5, 0x40, 0x4a, 0x8d, 0xca, 0xb9, 0xb8, 0x83, 0x11, 0x5c, 0xa1, 0x0c,
	0x2e, 0x76, 0xa8, 0x21, 0x42, 0x3a, 0x7a, 0xf8, 0x1c, 0xa0, 0x87, 0x6a, 0xbb, 0x94, 0x2e, 0x91,
	0xaa, 0x21, 0x2e, 0x53, 0x62, 0x70, 0xe2, 0x88, 0x62, 0x03, 0xfb, 0xb3, 0x38, 0x09, 0x42, 0x1b,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x20, 0x9f, 0xa8, 0x56, 0x00, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SwaggerSpecClient is the client API for SwaggerSpec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SwaggerSpecClient interface {
	// GetSpec returns the Swagger specification of the service's REST API.
	GetSpec(ctx context.Context, in *GetSpecRequest, opts ...grpc.CallOption) (*GetSpecResponse, error)
}

type swaggerSpecClient struct {
	cc grpc.ClientConnInterface
}

func NewSwaggerSpecClient(cc grpc.ClientConnInterface) SwaggerSpecClient {
	return &swaggerSpecClient{cc}
}

func (c *swaggerSpecClient) GetSpec(ctx context.Context, in *GetSpecRequest, opts ...grpc.CallOption) (*GetSpecResponse, error) {
	out := new(GetSpecResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.obsidian.swagger.SwaggerSpec/GetSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwaggerSpecServer is the server API for SwaggerSpec service.
type SwaggerSpecServer interface {
	// GetSpec returns the Swagger specification of the service's REST API.
	GetSpec(context.Context, *GetSpecRequest) (*GetSpecResponse, error)
}

// UnimplementedSwaggerSpecServer can be embedded to have forward compatible implementations.
type UnimplementedSwaggerSpecServer struct {
}

func (*UnimplementedSwaggerSpecServer) GetSpec(ctx context.Context, req *GetSpecRequest) (*GetSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpec not implemented")
}

func RegisterSwaggerSpecServer(s *grpc.Server, srv SwaggerSpecServer) {
	s.RegisterService(&_SwaggerSpec_serviceDesc, srv)
}

func _SwaggerSpec_GetSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwaggerSpecServer).GetSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.obsidian.swagger.SwaggerSpec/GetSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwaggerSpecServer).GetSpec(ctx, req.(*GetSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwaggerSpec_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.obsidian.swagger.SwaggerSpec",
	HandlerType: (*SwaggerSpecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSpec",
			Handler:    _SwaggerSpec_GetSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swagger.proto",
}
