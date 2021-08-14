// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/gov.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type AddPoolProposal struct {
	Title            string                                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	Description      string                                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	DenomIn          string                                 `protobuf:"bytes,3,opt,name=denom_in,json=denomIn,proto3" json:"denom_in,omitempty" yaml:"denom_in"`
	DenomOut         string                                 `protobuf:"bytes,4,opt,name=denom_out,json=denomOut,proto3" json:"denom_out,omitempty" yaml:"denom_out"`
	LiquidationRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=liquidation_ratio,json=liquidationRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_ratio" yaml:"liquidation_ratio"`
}

func (m *AddPoolProposal) Reset()         { *m = AddPoolProposal{} }
func (m *AddPoolProposal) String() string { return proto.CompactTextString(m) }
func (*AddPoolProposal) ProtoMessage()    {}
func (*AddPoolProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_31c5aab0360b917f, []int{0}
}
func (m *AddPoolProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddPoolProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddPoolProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddPoolProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPoolProposal.Merge(m, src)
}
func (m *AddPoolProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddPoolProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPoolProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddPoolProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddPoolProposal)(nil), "comdex.asset.v1beta1.AddPoolProposal")
}

func init() { proto.RegisterFile("comdex/asset/v1beta1/gov.proto", fileDescriptor_31c5aab0360b917f) }

var fileDescriptor_31c5aab0360b917f = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0xa1, 0x5a, 0xb5, 0x68, 0xd2, 0x8a, 0x8d, 0x21, 0x1e, 0x16, 0xc3, 0xa1, 0xf1, 0x52,
	0x36, 0x8d, 0x17, 0xe3, 0xcd, 0xc6, 0x8b, 0x5e, 0xac, 0x1c, 0xbd, 0x34, 0x14, 0xb6, 0xb8, 0x11,
	0x18, 0x64, 0x97, 0x6a, 0xdf, 0xc2, 0xc7, 0xf0, 0x51, 0x7a, 0xb3, 0x47, 0xe3, 0x81, 0x28, 0x7d,
	0x03, 0x9e, 0xc0, 0xb0, 0xdb, 0x26, 0x18, 0x2f, 0xcc, 0x30, 0xdf, 0xff, 0xcf, 0x6c, 0xf2, 0x6b,
	0xc8, 0x83, 0xc8, 0x27, 0xaf, 0xd8, 0x65, 0x8c, 0x70, 0x3c, 0x1b, 0x4c, 0x08, 0x77, 0x07, 0x38,
	0x80, 0x99, 0x9d, 0xa4, 0xc0, 0x41, 0xef, 0x4a, 0x6e, 0x0b, 0x6e, 0xaf, 0xf9, 0x49, 0x37, 0x80,
	0x00, 0x84, 0x00, 0x57, 0x9d, 0xd4, 0x5a, 0x1f, 0x0d, 0xad, 0x7d, 0xe5, 0xfb, 0x23, 0x80, 0x70,
	0x94, 0x42, 0x02, 0xcc, 0x0d, 0xf5, 0x9e, 0xd6, 0xe4, 0x94, 0x87, 0xc4, 0x50, 0x4f, 0xd5, 0xb3,
	0xd6, 0xb0, 0x53, 0xe6, 0xe6, 0xc1, 0xdc, 0x8d, 0xc2, 0x4b, 0x4b, 0x8c, 0x2d, 0x47, 0x62, 0xfd,
	0x42, 0xdb, 0xf7, 0x09, 0xf3, 0x52, 0x9a, 0x70, 0x0a, 0xb1, 0xd1, 0x10, 0xea, 0xe3, 0x32, 0x37,
	0x75, 0xa9, 0xae, 0x41, 0xcb, 0xa9, 0x4b, 0x75, 0x5b, 0xdb, 0xf3, 0x49, 0x0c, 0xd1, 0x98, 0xc6,
	0xc6, 0x96, 0xb0, 0x1d, 0x95, 0xb9, 0xd9, 0xde, 0xd8, 0x24, 0xb1, 0x9c, 0x5d, 0xd1, 0xde, 0xc4,
	0xfa, 0x40, 0x6b, 0xc9, 0x29, 0x64, 0xdc, 0xd8, 0x16, 0x86, 0x6e, 0x99, 0x9b, 0x9d, 0xba, 0x01,
	0x32, 0x6e, 0x39, 0x72, 0xed, 0x5d, 0xc6, 0xf5, 0x17, 0xed, 0x30, 0xa4, 0xcf, 0x19, 0xf5, 0xdd,
	0xea, 0xe2, 0x38, 0xad, 0x8a, 0xd1, 0x14, 0xd6, 0xdb, 0x45, 0x6e, 0x2a, 0x5f, 0xb9, 0xd9, 0x0b,
	0x28, 0x7f, 0xcc, 0x26, 0xb6, 0x07, 0x11, 0xf6, 0x80, 0x45, 0xc0, 0xd6, 0xa5, 0xcf, 0xfc, 0x27,
	0xcc, 0xe7, 0x09, 0x61, 0xf6, 0x35, 0xf1, 0xca, 0xdc, 0x34, 0xe4, 0xa1, 0x7f, 0x0b, 0x2d, 0xa7,
	0x53, 0x9b, 0x39, 0xd5, 0x77, 0x78, 0xbf, 0xf8, 0x41, 0xca, 0x7b, 0x81, 0x94, 0x45, 0x81, 0xd4,
	0x65, 0x81, 0xd4, 0xef, 0x02, 0xa9, 0x6f, 0x2b, 0xa4, 0x2c, 0x57, 0x48, 0xf9, 0x5c, 0x21, 0xe5,
	0x01, 0xff, 0xb9, 0x5b, 0x45, 0xd5, 0x87, 0xe9, 0x94, 0x7a, 0xd4, 0x0d, 0xd7, 0xff, 0x78, 0x13,
	0xae, 0x78, 0xc4, 0x64, 0x47, 0x64, 0x75, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x87, 0xcd,
	0xf0, 0xf9, 0x01, 0x00, 0x00,
}

func (m *AddPoolProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddPoolProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddPoolProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LiquidationRatio.Size()
		i -= size
		if _, err := m.LiquidationRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGov(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DenomOut) > 0 {
		i -= len(m.DenomOut)
		copy(dAtA[i:], m.DenomOut)
		i = encodeVarintGov(dAtA, i, uint64(len(m.DenomOut)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DenomIn) > 0 {
		i -= len(m.DenomIn)
		copy(dAtA[i:], m.DenomIn)
		i = encodeVarintGov(dAtA, i, uint64(len(m.DenomIn)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddPoolProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.DenomIn)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.DenomOut)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = m.LiquidationRatio.Size()
	n += 1 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddPoolProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
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
			return fmt.Errorf("proto: AddPoolProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddPoolProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
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
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
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
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
					return 0, ErrIntOverflowGov
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
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
