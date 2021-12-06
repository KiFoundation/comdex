// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/bandoracle/v1beta1/msg.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type MsgAddMarketRequest struct {
	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Symbol   string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	ScriptID uint64 `protobuf:"varint,3,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
	Id       uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgAddMarketRequest) Reset()         { *m = MsgAddMarketRequest{} }
func (m *MsgAddMarketRequest) String() string { return proto.CompactTextString(m) }
func (*MsgAddMarketRequest) ProtoMessage()    {}
func (*MsgAddMarketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3debb22a00d29a, []int{0}
}
func (m *MsgAddMarketRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddMarketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddMarketRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddMarketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddMarketRequest.Merge(m, src)
}
func (m *MsgAddMarketRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddMarketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddMarketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddMarketRequest proto.InternalMessageInfo

type MsgAddMarketResponse struct {
}

func (m *MsgAddMarketResponse) Reset()         { *m = MsgAddMarketResponse{} }
func (m *MsgAddMarketResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddMarketResponse) ProtoMessage()    {}
func (*MsgAddMarketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3debb22a00d29a, []int{1}
}
func (m *MsgAddMarketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddMarketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddMarketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddMarketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddMarketResponse.Merge(m, src)
}
func (m *MsgAddMarketResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddMarketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddMarketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddMarketResponse proto.InternalMessageInfo

type MsgUpdateMarketRequest struct {
	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Symbol   string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	ScriptID uint64 `protobuf:"varint,3,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
}

func (m *MsgUpdateMarketRequest) Reset()         { *m = MsgUpdateMarketRequest{} }
func (m *MsgUpdateMarketRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateMarketRequest) ProtoMessage()    {}
func (*MsgUpdateMarketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3debb22a00d29a, []int{2}
}
func (m *MsgUpdateMarketRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateMarketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateMarketRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateMarketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateMarketRequest.Merge(m, src)
}
func (m *MsgUpdateMarketRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateMarketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateMarketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateMarketRequest proto.InternalMessageInfo

type MsgUpdateMarketResponse struct {
}

func (m *MsgUpdateMarketResponse) Reset()         { *m = MsgUpdateMarketResponse{} }
func (m *MsgUpdateMarketResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateMarketResponse) ProtoMessage()    {}
func (*MsgUpdateMarketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3debb22a00d29a, []int{3}
}
func (m *MsgUpdateMarketResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateMarketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateMarketResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateMarketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateMarketResponse.Merge(m, src)
}
func (m *MsgUpdateMarketResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateMarketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateMarketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateMarketResponse proto.InternalMessageInfo

type MsgRemoveMarketForAssetRequest struct {
	From   string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Id     uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
}

func (m *MsgRemoveMarketForAssetRequest) Reset()         { *m = MsgRemoveMarketForAssetRequest{} }
func (m *MsgRemoveMarketForAssetRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveMarketForAssetRequest) ProtoMessage()    {}
func (*MsgRemoveMarketForAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3debb22a00d29a, []int{4}
}
func (m *MsgRemoveMarketForAssetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveMarketForAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveMarketForAssetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveMarketForAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveMarketForAssetRequest.Merge(m, src)
}
func (m *MsgRemoveMarketForAssetRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveMarketForAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveMarketForAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveMarketForAssetRequest proto.InternalMessageInfo

type MsgRemoveMarketForAssetResponse struct {
}

func (m *MsgRemoveMarketForAssetResponse) Reset()         { *m = MsgRemoveMarketForAssetResponse{} }
func (m *MsgRemoveMarketForAssetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveMarketForAssetResponse) ProtoMessage()    {}
func (*MsgRemoveMarketForAssetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b3debb22a00d29a, []int{5}
}
func (m *MsgRemoveMarketForAssetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRemoveMarketForAssetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveMarketForAssetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRemoveMarketForAssetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveMarketForAssetResponse.Merge(m, src)
}
func (m *MsgRemoveMarketForAssetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRemoveMarketForAssetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveMarketForAssetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveMarketForAssetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddMarketRequest)(nil), "comdex.bandoracle.v1beta1.MsgAddMarketRequest")
	proto.RegisterType((*MsgAddMarketResponse)(nil), "comdex.bandoracle.v1beta1.MsgAddMarketResponse")
	proto.RegisterType((*MsgUpdateMarketRequest)(nil), "comdex.bandoracle.v1beta1.MsgUpdateMarketRequest")
	proto.RegisterType((*MsgUpdateMarketResponse)(nil), "comdex.bandoracle.v1beta1.MsgUpdateMarketResponse")
	proto.RegisterType((*MsgRemoveMarketForAssetRequest)(nil), "comdex.bandoracle.v1beta1.MsgRemoveMarketForAssetRequest")
	proto.RegisterType((*MsgRemoveMarketForAssetResponse)(nil), "comdex.bandoracle.v1beta1.MsgRemoveMarketForAssetResponse")
}

func init() {
	proto.RegisterFile("comdex/bandoracle/v1beta1/msg.proto", fileDescriptor_9b3debb22a00d29a)
}

var fileDescriptor_9b3debb22a00d29a = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0xce, 0xcf, 0x4d,
	0x49, 0xad, 0xd0, 0x4f, 0x4a, 0xcc, 0x4b, 0xc9, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x33,
	0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x84, 0x28, 0xd2, 0x43, 0x28, 0xd2, 0x83, 0x2a, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0xab, 0xd2, 0x07, 0xb1, 0x20, 0x1a, 0x94, 0xb6, 0x30, 0x72, 0x09, 0xfb, 0x16, 0xa7, 0x3b, 0xa6,
	0xa4, 0xf8, 0x26, 0x16, 0x65, 0xa7, 0x96, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x29,
	0x73, 0xb1, 0xa4, 0x15, 0xe5, 0xe7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x3a, 0xf1, 0x7f, 0xba,
	0x27, 0xcf, 0x5d, 0x99, 0x98, 0x9b, 0x63, 0xa5, 0x04, 0x12, 0x55, 0x0a, 0x02, 0x4b, 0x0a, 0x69,
	0x72, 0xb1, 0x15, 0x57, 0xe6, 0x26, 0xe5, 0xe7, 0x48, 0x30, 0x81, 0x95, 0x09, 0x7e, 0xba, 0x27,
	0xcf, 0x0b, 0x51, 0x06, 0x11, 0x57, 0x0a, 0x82, 0x2a, 0x10, 0xb2, 0xe5, 0xe2, 0x2c, 0x4e, 0x2e,
	0xca, 0x2c, 0x28, 0x89, 0xcf, 0x4c, 0x91, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x71, 0x52, 0x78, 0x74,
	0x4f, 0x9e, 0x23, 0x18, 0x2c, 0xe8, 0xe9, 0xf2, 0xe9, 0x9e, 0xbc, 0x00, 0x54, 0x27, 0x4c, 0x99,
	0x52, 0x10, 0x07, 0x84, 0xed, 0x99, 0x22, 0xc4, 0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0x02, 0xd2,
	0x17, 0xc4, 0x94, 0x99, 0xa2, 0x24, 0xc6, 0x25, 0x82, 0xea, 0xea, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0xa5, 0xe5, 0x8c, 0x5c, 0x62, 0xbe, 0xc5, 0xe9, 0xa1, 0x05, 0x29, 0x89, 0x25, 0xa9, 0x83,
	0xd9, 0x47, 0x4a, 0x92, 0x5c, 0xe2, 0x18, 0x0e, 0x85, 0x7a, 0xa2, 0x81, 0x91, 0x4b, 0xce, 0xb7,
	0x38, 0x3d, 0x28, 0x35, 0x37, 0xbf, 0x0c, 0x2a, 0xe7, 0x96, 0x5f, 0xe4, 0x58, 0x5c, 0x4c, 0xa2,
	0x67, 0x20, 0x81, 0xc6, 0x04, 0x0b, 0x34, 0x24, 0xcf, 0x31, 0x13, 0xf0, 0x9c, 0x92, 0x22, 0x97,
	0x3c, 0x4e, 0x17, 0x40, 0x5c, 0xe9, 0x14, 0x76, 0xe2, 0xa1, 0x1c, 0xc3, 0x8a, 0x47, 0x72, 0x0c,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x92, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x49, 0x97, 0xba, 0xf9, 0x69, 0x69, 0x99, 0xc9, 0x99,
	0x89, 0x39, 0x50, 0xbe, 0x3e, 0x4a, 0x72, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x27,
	0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x6d, 0x70, 0x3b, 0xf0, 0x02, 0x00, 0x00,
}

func (m *MsgAddMarketRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddMarketRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddMarketRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x20
	}
	if m.ScriptID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ScriptID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddMarketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddMarketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddMarketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateMarketRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateMarketRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateMarketRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ScriptID != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.ScriptID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateMarketResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateMarketResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateMarketResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRemoveMarketForAssetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveMarketForAssetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveMarketForAssetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintMsg(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRemoveMarketForAssetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveMarketForAssetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveMarketForAssetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddMarketRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ScriptID != 0 {
		n += 1 + sovMsg(uint64(m.ScriptID))
	}
	if m.Id != 0 {
		n += 1 + sovMsg(uint64(m.Id))
	}
	return n
}

func (m *MsgAddMarketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateMarketRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.ScriptID != 0 {
		n += 1 + sovMsg(uint64(m.ScriptID))
	}
	return n
}

func (m *MsgUpdateMarketResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRemoveMarketForAssetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovMsg(uint64(m.Id))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MsgRemoveMarketForAssetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddMarketRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgAddMarketRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddMarketRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptID", wireType)
			}
			m.ScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgAddMarketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgAddMarketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddMarketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUpdateMarketRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdateMarketRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateMarketRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptID", wireType)
			}
			m.ScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgUpdateMarketResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgUpdateMarketResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateMarketResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgRemoveMarketForAssetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgRemoveMarketForAssetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveMarketForAssetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MsgRemoveMarketForAssetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MsgRemoveMarketForAssetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveMarketForAssetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
