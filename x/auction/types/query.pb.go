// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type QueryAuctionRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryAuctionRequest) Reset()         { *m = QueryAuctionRequest{} }
func (m *QueryAuctionRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAuctionRequest) ProtoMessage()    {}
func (*QueryAuctionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3902e0355803c62, []int{2}
}
func (m *QueryAuctionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuctionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuctionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuctionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuctionRequest.Merge(m, src)
}
func (m *QueryAuctionRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuctionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuctionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuctionRequest proto.InternalMessageInfo

type QueryAuctionResponse struct {
	Auction CollateralAuction `protobuf:"bytes,1,opt,name=auction,proto3" json:"auction" yaml:"auction"`
}

func (m *QueryAuctionResponse) Reset()         { *m = QueryAuctionResponse{} }
func (m *QueryAuctionResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuctionResponse) ProtoMessage()    {}
func (*QueryAuctionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3902e0355803c62, []int{3}
}
func (m *QueryAuctionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuctionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuctionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuctionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuctionResponse.Merge(m, src)
}
func (m *QueryAuctionResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuctionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuctionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuctionResponse proto.InternalMessageInfo

type QueryAuctionsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty" yaml:"pagination"`
}

func (m *QueryAuctionsRequest) Reset()         { *m = QueryAuctionsRequest{} }
func (m *QueryAuctionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAuctionsRequest) ProtoMessage()    {}
func (*QueryAuctionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3902e0355803c62, []int{4}
}
func (m *QueryAuctionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuctionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuctionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuctionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuctionsRequest.Merge(m, src)
}
func (m *QueryAuctionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuctionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuctionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuctionsRequest proto.InternalMessageInfo

type QueryAuctionsResponse struct {
	Auctions   []CollateralAuction `protobuf:"bytes,1,rep,name=auctions,proto3" json:"auctions" yaml:"auctions"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty" yaml:"pagination"`
}

func (m *QueryAuctionsResponse) Reset()         { *m = QueryAuctionsResponse{} }
func (m *QueryAuctionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuctionsResponse) ProtoMessage()    {}
func (*QueryAuctionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3902e0355803c62, []int{5}
}
func (m *QueryAuctionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuctionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuctionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuctionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuctionsResponse.Merge(m, src)
}
func (m *QueryAuctionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuctionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuctionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuctionsResponse proto.InternalMessageInfo

type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3902e0355803c62, []int{6}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3902e0355803c62, []int{7}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryAuctionRequest)(nil), "comdex.auction.v1beta1.QueryAuctionRequest")
	proto.RegisterType((*QueryAuctionResponse)(nil), "comdex.auction.v1beta1.QueryAuctionResponse")
	proto.RegisterType((*QueryAuctionsRequest)(nil), "comdex.auction.v1beta1.QueryAuctionsRequest")
	proto.RegisterType((*QueryAuctionsResponse)(nil), "comdex.auction.v1beta1.QueryAuctionsResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "comdex.auction.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "comdex.auction.v1beta1.QueryParamsResponse")
}

func init() { proto.RegisterFile("comdex/auction/v1beta1/querier.proto", fileDescriptor_d3902e0355803c62) }

var fileDescriptor_d3902e0355803c62 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6a, 0x13, 0x41,
	0x18, 0xdf, 0x49, 0x25, 0x96, 0xa9, 0xad, 0x64, 0x9a, 0xd4, 0x12, 0x64, 0x12, 0x06, 0x8d, 0xb1,
	0xad, 0x3b, 0xb4, 0xde, 0xbc, 0x19, 0x6f, 0x82, 0x50, 0xd7, 0x9b, 0x62, 0x61, 0x92, 0x0c, 0xdb,
	0x85, 0xcd, 0xce, 0x36, 0xb3, 0x5b, 0x1a, 0xc4, 0x8b, 0x0f, 0x20, 0x82, 0xa0, 0xaf, 0xe0, 0xa3,
	0xe4, 0x58, 0xf4, 0xa0, 0xa7, 0xa0, 0x89, 0x4f, 0xd0, 0x27, 0x90, 0xcc, 0x7c, 0xdb, 0x76, 0xb5,
	0xb1, 0xe9, 0x29, 0x9b, 0xef, 0xdf, 0xef, 0xcf, 0xf7, 0xed, 0xe2, 0xd2, 0xbe, 0x0c, 0x43, 0xc5,
	0x0f, 0x52, 0xd9, 0x1f, 0xb8, 0x71, 0x5f, 0x25, 0x8a, 0xac, 0x75, 0x54, 0xaf, 0x2b, 0x8f, 0x5c,
	0x91, 0x76, 0x92, 0x40, 0x45, 0xee, 0xe1, 0x76, 0x5b, 0x26, 0x62, 0xbb, 0xba, 0xd1, 0x51, 0xba,
	0xa7, 0x34, 0x6f, 0x0b, 0x2d, 0x6d, 0x03, 0x87, 0x14, 0x8f, 0x85, 0x1f, 0x44, 0xc2, 0x54, 0x9b,
	0x19, 0xd5, 0xdb, 0xbe, 0x52, 0x7e, 0x28, 0xb9, 0x88, 0x03, 0x2e, 0xa2, 0x48, 0x25, 0x26, 0xa9,
	0x21, 0x5b, 0xf6, 0x95, 0xaf, 0xcc, 0x23, 0x9f, 0x3e, 0xd9, 0x28, 0x5b, 0xc5, 0xa5, 0x27, 0x2a,
	0x0c, 0x45, 0x22, 0xfb, 0x22, 0x7c, 0x6c, 0xc1, 0xd9, 0x22, 0x2e, 0xee, 0x8a, 0xbe, 0xe8, 0x69,
	0x76, 0x17, 0xaf, 0x3e, 0x9f, 0x82, 0x42, 0xc6, 0x93, 0x07, 0xa9, 0xd4, 0x09, 0x59, 0xc1, 0x85,
	0xa0, 0xbb, 0x8e, 0xea, 0xa8, 0x79, 0xcd, 0x2b, 0x04, 0x5d, 0xa6, 0x71, 0x39, 0x5f, 0xa6, 0x63,
	0x15, 0x69, 0x49, 0x5e, 0xe1, 0xeb, 0x20, 0xc8, 0x14, 0x2f, 0xed, 0xdc, 0x77, 0x2f, 0xd6, 0xe9,
	0xfe, 0x43, 0xa2, 0xb5, 0x36, 0x1c, 0xd5, 0x9c, 0x93, 0x51, 0x6d, 0x65, 0x20, 0x7a, 0xe1, 0x23,
	0x06, 0x0d, 0xcc, 0xcb, 0x26, 0xb2, 0x34, 0x0f, 0xaa, 0x33, 0x72, 0xaf, 0x31, 0x3e, 0xb3, 0x06,
	0x70, 0x1b, 0xae, 0xf5, 0xd1, 0x9d, 0xfa, 0xe8, 0x5a, 0xe3, 0x33, 0xe8, 0x5d, 0xe1, 0x4b, 0xe8,
	0x6d, 0x55, 0x4e, 0x46, 0xb5, 0x92, 0x05, 0x3c, 0x9b, 0xc1, 0xbc, 0x73, 0x03, 0xd9, 0x77, 0x84,
	0x2b, 0x7f, 0xe1, 0x82, 0xda, 0x3d, 0xbc, 0x08, 0xdc, 0xf4, 0x3a, 0xaa, 0x2f, 0x5c, 0x4d, 0xee,
	0x2d, 0x90, 0x7b, 0x33, 0x27, 0x57, 0x33, 0xef, 0x74, 0x26, 0xd9, 0xcb, 0x09, 0x2b, 0x18, 0x61,
	0xf7, 0x2e, 0x15, 0x66, 0xc9, 0xcd, 0xa3, 0xac, 0x8c, 0x89, 0x11, 0x66, 0x77, 0x0f, 0x96, 0xb0,
	0x2e, 0x9c, 0x40, 0x16, 0x05, 0xb1, 0xcf, 0x70, 0x31, 0x36, 0x11, 0x70, 0x98, 0xce, 0x92, 0x6a,
	0xfb, 0x5a, 0x15, 0xd0, 0xb7, 0x9c, 0x71, 0x30, 0xf7, 0xe5, 0xc1, 0x90, 0x9d, 0xaf, 0x0b, 0xf8,
	0x86, 0x81, 0x79, 0x21, 0xfb, 0x87, 0x41, 0x47, 0x92, 0xcf, 0x08, 0x02, 0x60, 0x10, 0xd9, 0x9c,
	0x05, 0x70, 0xc1, 0x81, 0x56, 0xb7, 0xe6, 0x2b, 0xb6, 0x5a, 0xd8, 0xd6, 0xbb, 0x6f, 0xbf, 0x3f,
	0x16, 0x1a, 0xe4, 0x0e, 0xb7, 0x5d, 0x1c, 0xba, 0x4e, 0x5f, 0xb5, 0xec, 0xff, 0x9b, 0xa0, 0xfb,
	0x96, 0x7c, 0x42, 0x78, 0x39, 0x77, 0x00, 0x64, 0x2e, 0xb4, 0xcc, 0xd0, 0xea, 0x83, 0x39, 0xab,
	0x81, 0x5c, 0xd3, 0x90, 0x63, 0xa4, 0x7e, 0x09, 0x39, 0x4d, 0xde, 0x23, 0xbc, 0x74, 0x6e, 0x55,
	0x64, 0xe3, 0xbf, 0x40, 0xb9, 0x2d, 0x57, 0x37, 0xe7, 0xaa, 0x05, 0x4a, 0x0d, 0x43, 0xa9, 0x4e,
	0xe8, 0x2c, 0x4a, 0x76, 0xa9, 0xad, 0xa7, 0xc3, 0x5f, 0xd4, 0xf9, 0x32, 0xa6, 0xce, 0x70, 0x4c,
	0xd1, 0xf1, 0x98, 0xa2, 0x9f, 0x63, 0x8a, 0x3e, 0x4c, 0xa8, 0x73, 0x3c, 0xa1, 0xce, 0x8f, 0x09,
	0x75, 0x5e, 0x36, 0xfd, 0x20, 0xd9, 0x4f, 0xdb, 0x53, 0x70, 0x6e, 0x0e, 0x39, 0x12, 0x69, 0xc2,
	0xed, 0xe7, 0xf1, 0x08, 0x7e, 0x93, 0x41, 0x2c, 0x75, 0xbb, 0x68, 0xbe, 0x57, 0x0f, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x65, 0x7a, 0x8f, 0x91, 0x3c, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	QueryAuction(ctx context.Context, in *QueryAuctionRequest, opts ...grpc.CallOption) (*QueryAuctionResponse, error)
	QueryAuctions(ctx context.Context, in *QueryAuctionsRequest, opts ...grpc.CallOption) (*QueryAuctionsResponse, error)
	QueryParams(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) QueryAuction(ctx context.Context, in *QueryAuctionRequest, opts ...grpc.CallOption) (*QueryAuctionResponse, error) {
	out := new(QueryAuctionResponse)
	err := c.cc.Invoke(ctx, "/comdex.auction.v1beta1.QueryService/QueryAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryAuctions(ctx context.Context, in *QueryAuctionsRequest, opts ...grpc.CallOption) (*QueryAuctionsResponse, error) {
	out := new(QueryAuctionsResponse)
	err := c.cc.Invoke(ctx, "/comdex.auction.v1beta1.QueryService/QueryAuctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) QueryParams(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/comdex.auction.v1beta1.QueryService/QueryParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	QueryAuction(context.Context, *QueryAuctionRequest) (*QueryAuctionResponse, error)
	QueryAuctions(context.Context, *QueryAuctionsRequest) (*QueryAuctionsResponse, error)
	QueryParams(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) QueryAuction(ctx context.Context, req *QueryAuctionRequest) (*QueryAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAuction not implemented")
}
func (*UnimplementedQueryServiceServer) QueryAuctions(ctx context.Context, req *QueryAuctionsRequest) (*QueryAuctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAuctions not implemented")
}
func (*UnimplementedQueryServiceServer) QueryParams(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryParams not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_QueryAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.auction.v1beta1.QueryService/QueryAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryAuction(ctx, req.(*QueryAuctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryAuctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryAuctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.auction.v1beta1.QueryService/QueryAuctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryAuctions(ctx, req.(*QueryAuctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_QueryParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).QueryParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comdex.auction.v1beta1.QueryService/QueryParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).QueryParams(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdex.auction.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryAuction",
			Handler:    _QueryService_QueryAuction_Handler,
		},
		{
			MethodName: "QueryAuctions",
			Handler:    _QueryService_QueryAuctions_Handler,
		},
		{
			MethodName: "QueryParams",
			Handler:    _QueryService_QueryParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comdex/auction/v1beta1/querier.proto",
}

func (m *QueryAuctionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuctionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuctionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryAuctionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuctionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuctionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Auction.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAuctionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuctionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuctionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAuctionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuctionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuctionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Auctions) > 0 {
		for iNdEx := len(m.Auctions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Auctions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *QueryAuctionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryAuctionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Auction.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAuctionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAuctionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Auctions) > 0 {
		for _, e := range m.Auctions {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *QueryAuctionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuctionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuctionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAuctionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuctionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuctionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Auction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAuctionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuctionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuctionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAuctionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuctionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuctionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctions = append(m.Auctions, CollateralAuction{})
			if err := m.Auctions[len(m.Auctions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
