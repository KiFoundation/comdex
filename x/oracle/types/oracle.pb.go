// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/oracle/v1beta1/oracle.proto

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

type Market struct {
	Symbol   string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty" yaml:"symbol"`
	ScriptID uint64 `protobuf:"varint,2,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ea76e22e2125a9, []int{0}
}
func (m *Market) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Market.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return m.Size()
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

type Calldata struct {
	Multiplier uint64 `protobuf:"varint,1,opt,name=multiplier,proto3" json:"multiplier,omitempty" yaml:"multiplier"`
}

func (m *Calldata) Reset()         { *m = Calldata{} }
func (m *Calldata) String() string { return proto.CompactTextString(m) }
func (*Calldata) ProtoMessage()    {}
func (*Calldata) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ea76e22e2125a9, []int{1}
}
func (m *Calldata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Calldata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Calldata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Calldata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calldata.Merge(m, src)
}
func (m *Calldata) XXX_Size() int {
	return m.Size()
}
func (m *Calldata) XXX_DiscardUnknown() {
	xxx_messageInfo_Calldata.DiscardUnknown(m)
}

var xxx_messageInfo_Calldata proto.InternalMessageInfo

type Result struct {
	Rates uint64 `protobuf:"varint,1,opt,name=rates,proto3" json:"rates,omitempty" yaml:"rates"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ea76e22e2125a9, []int{2}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Result.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return m.Size()
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Market)(nil), "comdex.oracle.v1beta1.Market")
	proto.RegisterType((*Calldata)(nil), "comdex.oracle.v1beta1.Calldata")
	proto.RegisterType((*Result)(nil), "comdex.oracle.v1beta1.Result")
}

func init() {
	proto.RegisterFile("comdex/oracle/v1beta1/oracle.proto", fileDescriptor_52ea76e22e2125a9)
}

var fileDescriptor_52ea76e22e2125a9 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xe3, 0x4f, 0xfd, 0xa2, 0xd6, 0x02, 0xa9, 0x8d, 0xa8, 0x54, 0x31, 0x38, 0x95, 0x07,
	0x54, 0x06, 0x9a, 0x56, 0x88, 0x05, 0x89, 0x81, 0xc0, 0xd2, 0x81, 0xc5, 0x6c, 0x2c, 0xc8, 0x49,
	0xdc, 0x62, 0xe1, 0xc8, 0x91, 0xe3, 0x20, 0x72, 0x17, 0x5c, 0x06, 0x97, 0xd2, 0xb1, 0x23, 0x53,
	0x04, 0xc9, 0x1d, 0xe4, 0x0a, 0x10, 0xb1, 0xf9, 0xd9, 0xce, 0x79, 0xcf, 0xf3, 0xe8, 0x48, 0x2f,
	0xc4, 0xb1, 0x4c, 0x13, 0xf6, 0x1c, 0x48, 0x45, 0x63, 0xc1, 0x82, 0xa7, 0x65, 0xc4, 0x34, 0x5d,
	0xda, 0x75, 0x9e, 0x29, 0xa9, 0xa5, 0x37, 0x36, 0xcc, 0xdc, 0x86, 0x96, 0x39, 0x3c, 0xd8, 0xc8,
	0x8d, 0xec, 0x88, 0xe0, 0x6b, 0x32, 0x30, 0x56, 0xd0, 0xbd, 0xa1, 0xea, 0x91, 0x69, 0xef, 0x18,
	0xba, 0x79, 0x99, 0x46, 0x52, 0x4c, 0xc0, 0x14, 0xcc, 0x06, 0xe1, 0xa8, 0xad, 0xfc, 0xfd, 0x92,
	0xa6, 0xe2, 0x1c, 0x9b, 0x1c, 0x13, 0x0b, 0x78, 0x17, 0x70, 0x90, 0xc7, 0x8a, 0x67, 0xfa, 0x9e,
	0x27, 0x93, 0x7f, 0x53, 0x30, 0xeb, 0x85, 0xd3, 0xba, 0xf2, 0xfb, 0xb7, 0x5d, 0xb8, 0xba, 0x6e,
	0x2b, 0x7f, 0x68, 0xcd, 0x6f, 0x0c, 0x93, 0xbe, 0x99, 0x57, 0x09, 0xbe, 0x84, 0xfd, 0x2b, 0x2a,
	0x44, 0x42, 0x35, 0xf5, 0xce, 0x20, 0x4c, 0x0b, 0xa1, 0x79, 0x26, 0x38, 0x53, 0xdd, 0xe7, 0x5e,
	0x38, 0x6e, 0x2b, 0x7f, 0x64, 0xfc, 0xdf, 0x1b, 0x26, 0x7f, 0x40, 0xbc, 0x80, 0x2e, 0x61, 0x79,
	0x21, 0xb4, 0x77, 0x04, 0xff, 0x2b, 0xaa, 0x59, 0x6e, 0xdd, 0x61, 0x5b, 0xf9, 0x7b, 0xc6, 0xed,
	0x62, 0x4c, 0xcc, 0x39, 0x24, 0xdb, 0x0f, 0xe4, 0xbc, 0xd6, 0xc8, 0xd9, 0xd6, 0x08, 0xec, 0x6a,
	0x04, 0xde, 0x6b, 0x04, 0x5e, 0x1a, 0xe4, 0xec, 0x1a, 0xe4, 0xbc, 0x35, 0xc8, 0xb9, 0x5b, 0x6c,
	0xb8, 0x7e, 0x28, 0xa2, 0x79, 0x2c, 0xd3, 0xc0, 0x54, 0x78, 0x22, 0xd7, 0x6b, 0x1e, 0x73, 0x2a,
	0xec, 0x1e, 0xfc, 0x14, 0xaf, 0xcb, 0x8c, 0xe5, 0x91, 0xdb, 0x75, 0x78, 0xfa, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x83, 0x1e, 0x27, 0x51, 0x96, 0x01, 0x00, 0x00,
}

func (m *Market) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Market) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Market) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ScriptID != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.ScriptID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Calldata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Calldata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Calldata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Result) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Result) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Result) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rates != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.Rates))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Market) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if m.ScriptID != 0 {
		n += 1 + sovOracle(uint64(m.ScriptID))
	}
	return n
}

func (m *Calldata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Multiplier != 0 {
		n += 1 + sovOracle(uint64(m.Multiplier))
	}
	return n
}

func (m *Result) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Rates != 0 {
		n += 1 + sovOracle(uint64(m.Rates))
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Market) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Market: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Market: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptID", wireType)
			}
			m.ScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *Calldata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Calldata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Calldata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *Result) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Result: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Result: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
			m.Rates = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rates |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
