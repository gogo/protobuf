// Code generated by protoc-gen-gogo. DO NOT EDIT.
// deprecated/deprecated.proto is a deprecated file.

// package deprecated contains only deprecated messages and services.

package deprecated

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DeprecatedEnum contains deprecated values.
type DeprecatedEnum int32 // Deprecated: Do not use.
const (
	// DEPRECATED is the iota value of this enum.
	DeprecatedEnum_DEPRECATED DeprecatedEnum = 0 // Deprecated: Do not use.
)

var DeprecatedEnum_name = map[int32]string{
	0: "DEPRECATED",
}

var DeprecatedEnum_value = map[string]int32{
	"DEPRECATED": 0,
}

func (x DeprecatedEnum) String() string {
	return proto.EnumName(DeprecatedEnum_name, int32(x))
}

func (DeprecatedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{0}
}

// DeprecatedRequest is a request to DeprecatedCall.
//
// Deprecated: Do not use.
type DeprecatedRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeprecatedRequest) Reset()         { *m = DeprecatedRequest{} }
func (m *DeprecatedRequest) String() string { return proto.CompactTextString(m) }
func (*DeprecatedRequest) ProtoMessage()    {}
func (*DeprecatedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{0}
}
func (m *DeprecatedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedRequest.Unmarshal(m, b)
}
func (m *DeprecatedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedRequest.Marshal(b, m, deterministic)
}
func (m *DeprecatedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedRequest.Merge(m, src)
}
func (m *DeprecatedRequest) XXX_Size() int {
	return xxx_messageInfo_DeprecatedRequest.Size(m)
}
func (m *DeprecatedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedRequest proto.InternalMessageInfo

// Deprecated: Do not use.
type DeprecatedResponse struct {
	// DeprecatedField contains a DeprecatedEnum.
	DeprecatedField DeprecatedEnum `protobuf:"varint,1,opt,name=deprecated_field,json=deprecatedField,proto3,enum=deprecated.DeprecatedEnum" json:"deprecated_field,omitempty"` // Deprecated: Do not use.
	// DeprecatedOneof contains a deprecated field.
	//
	// Types that are valid to be assigned to DeprecatedOneof:
	//	*DeprecatedResponse_DeprecatedOneofField
	DeprecatedOneof      isDeprecatedResponse_DeprecatedOneof `protobuf_oneof:"deprecated_oneof"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *DeprecatedResponse) Reset()         { *m = DeprecatedResponse{} }
func (m *DeprecatedResponse) String() string { return proto.CompactTextString(m) }
func (*DeprecatedResponse) ProtoMessage()    {}
func (*DeprecatedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f64ba265cd7eae3f, []int{1}
}
func (m *DeprecatedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeprecatedResponse.Unmarshal(m, b)
}
func (m *DeprecatedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeprecatedResponse.Marshal(b, m, deterministic)
}
func (m *DeprecatedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeprecatedResponse.Merge(m, src)
}
func (m *DeprecatedResponse) XXX_Size() int {
	return xxx_messageInfo_DeprecatedResponse.Size(m)
}
func (m *DeprecatedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeprecatedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeprecatedResponse proto.InternalMessageInfo

type isDeprecatedResponse_DeprecatedOneof interface {
	isDeprecatedResponse_DeprecatedOneof()
}

type DeprecatedResponse_DeprecatedOneofField struct {
	DeprecatedOneofField string `protobuf:"bytes,2,opt,name=deprecated_oneof_field,json=deprecatedOneofField,proto3,oneof" json:"deprecated_oneof_field,omitempty"`
}

func (*DeprecatedResponse_DeprecatedOneofField) isDeprecatedResponse_DeprecatedOneof() {}

func (m *DeprecatedResponse) GetDeprecatedOneof() isDeprecatedResponse_DeprecatedOneof {
	if m != nil {
		return m.DeprecatedOneof
	}
	return nil
}

// Deprecated: Do not use.
func (m *DeprecatedResponse) GetDeprecatedField() DeprecatedEnum {
	if m != nil {
		return m.DeprecatedField
	}
	return DeprecatedEnum_DEPRECATED
}

// Deprecated: Do not use.
func (m *DeprecatedResponse) GetDeprecatedOneofField() string {
	if x, ok := m.GetDeprecatedOneof().(*DeprecatedResponse_DeprecatedOneofField); ok {
		return x.DeprecatedOneofField
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DeprecatedResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DeprecatedResponse_DeprecatedOneofField)(nil),
	}
}

func init() {
	proto.RegisterEnum("deprecated.DeprecatedEnum", DeprecatedEnum_name, DeprecatedEnum_value)
	proto.RegisterType((*DeprecatedRequest)(nil), "deprecated.DeprecatedRequest")
	proto.RegisterType((*DeprecatedResponse)(nil), "deprecated.DeprecatedResponse")
}

func init() { proto.RegisterFile("deprecated/deprecated.proto", fileDescriptor_f64ba265cd7eae3f) }

var fileDescriptor_f64ba265cd7eae3f = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x49, 0x2d, 0x28,
	0x4a, 0x4d, 0x4e, 0x2c, 0x49, 0x4d, 0xd1, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x4a, 0xe2, 0x5c, 0x82, 0x2e, 0x70, 0x5e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
	0x89, 0x15, 0x93, 0x04, 0xa3, 0xd2, 0x66, 0x46, 0x2e, 0x21, 0x64, 0x99, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0x21, 0x4f, 0x2e, 0x01, 0x84, 0xee, 0xf8, 0xb4, 0xcc, 0xd4, 0x9c, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0x3e, 0x23, 0x29, 0x3d, 0x24, 0x8b, 0x10, 0x3a, 0x5d, 0xf3, 0x4a, 0x73, 0x9d,
	0x98, 0x24, 0x18, 0x83, 0xf8, 0x11, 0xd2, 0x6e, 0x20, 0x6d, 0x42, 0x56, 0x5c, 0x62, 0x48, 0x46,
	0xe5, 0xe7, 0xa5, 0xe6, 0xa7, 0x41, 0x0d, 0x64, 0x52, 0x60, 0xd4, 0xe0, 0x04, 0x69, 0xf2, 0x60,
	0x08, 0x12, 0x41, 0xa8, 0xf1, 0x07, 0x29, 0x01, 0xeb, 0x05, 0xb9, 0xd0, 0x49, 0x08, 0xc5, 0x29,
	0x60, 0xfd, 0x5a, 0x1a, 0x5c, 0x7c, 0xa8, 0x56, 0x0b, 0x09, 0x71, 0x71, 0xb9, 0xb8, 0x06, 0x04,
	0xb9, 0x3a, 0x3b, 0x86, 0xb8, 0xba, 0x08, 0x30, 0x48, 0x31, 0x71, 0x30, 0x4a, 0x31, 0x49, 0x30,
	0x1a, 0xe5, 0x21, 0x7b, 0x3c, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x04, 0x59, 0xbb,
	0x73, 0x62, 0x4e, 0x8e, 0x90, 0x2c, 0x76, 0x5f, 0x41, 0x43, 0x4a, 0x4a, 0x0e, 0x97, 0x34, 0x24,
	0xb8, 0x94, 0x98, 0x3b, 0x98, 0x18, 0xa5, 0x40, 0x84, 0x93, 0x63, 0x94, 0x4d, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x7a, 0x7e, 0x7a, 0xbe, 0x3e, 0x38, 0x36, 0x92,
	0x4a, 0xd3, 0x20, 0x8c, 0x64, 0xdd, 0xf4, 0xd4, 0x3c, 0x5d, 0xb0, 0x44, 0x49, 0x6a, 0x71, 0x49,
	0x4a, 0x62, 0x49, 0x22, 0x52, 0xd4, 0xed, 0x60, 0x64, 0x4c, 0x62, 0x03, 0xab, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0x08, 0x22, 0x8d, 0xdd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeprecatedServiceClient is the client API for DeprecatedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Do not use.
type DeprecatedServiceClient interface {
	// DeprecatedCall takes a DeprecatedRequest and returns a DeprecatedResponse.
	//
	// Deprecated: Do not use.
	DeprecatedCall(ctx context.Context, in *DeprecatedRequest, opts ...grpc.CallOption) (*DeprecatedResponse, error)
}

type deprecatedServiceClient struct {
	cc grpc.ClientConnInterface
}

// Deprecated: Do not use.
func NewDeprecatedServiceClient(cc grpc.ClientConnInterface) DeprecatedServiceClient {
	return &deprecatedServiceClient{cc}
}

// Deprecated: Do not use.
func (c *deprecatedServiceClient) DeprecatedCall(ctx context.Context, in *DeprecatedRequest, opts ...grpc.CallOption) (*DeprecatedResponse, error) {
	out := new(DeprecatedResponse)
	err := c.cc.Invoke(ctx, "/deprecated.DeprecatedService/DeprecatedCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeprecatedServiceServer is the server API for DeprecatedService service.
//
// Deprecated: Do not use.
type DeprecatedServiceServer interface {
	// DeprecatedCall takes a DeprecatedRequest and returns a DeprecatedResponse.
	//
	// Deprecated: Do not use.
	DeprecatedCall(context.Context, *DeprecatedRequest) (*DeprecatedResponse, error)
}

// Deprecated: Do not use.
// UnimplementedDeprecatedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeprecatedServiceServer struct {
}

func (*UnimplementedDeprecatedServiceServer) DeprecatedCall(ctx context.Context, req *DeprecatedRequest) (*DeprecatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeprecatedCall not implemented")
}

// Deprecated: Do not use.
func RegisterDeprecatedServiceServer(s *grpc.Server, srv DeprecatedServiceServer) {
	s.RegisterService(&_DeprecatedService_serviceDesc, srv)
}

func _DeprecatedService_DeprecatedCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeprecatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeprecatedServiceServer).DeprecatedCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deprecated.DeprecatedService/DeprecatedCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeprecatedServiceServer).DeprecatedCall(ctx, req.(*DeprecatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeprecatedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "deprecated.DeprecatedService",
	HandlerType: (*DeprecatedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeprecatedCall",
			Handler:    _DeprecatedService_DeprecatedCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deprecated/deprecated.proto",
}
