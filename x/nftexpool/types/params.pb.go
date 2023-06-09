// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gauss/nftexpool/v1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Params struct {
	// create fee
	PoolCreationFee types.Coin `protobuf:"bytes,1,opt,name=pool_creation_fee,json=poolCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pool_creation_fee" yaml:"pool_creation_fee"`
	// fee allot
	BurnPercent       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=burn_percent,json=burnPercent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"burn_percent" yaml:"burn_percent"`
	CommunityPercent  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=community_percent,json=communityPercent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_percent" yaml:"community_percent"`
	ValidatorsPercent github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=validators_percent,json=validatorsPercent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"validators_percent" yaml:"validators_percent"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbf4746d927302f7, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPoolCreationFee() types.Coin {
	if m != nil {
		return m.PoolCreationFee
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "gauss.pool.v1.Params")
}

func init() { proto.RegisterFile("gauss/nftexpool/v1/params.proto", fileDescriptor_bbf4746d927302f7) }

var fileDescriptor_bbf4746d927302f7 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x6f, 0xda, 0x40,
	0x14, 0x80, 0x7d, 0x2d, 0x42, 0xad, 0x69, 0xd5, 0xe2, 0x76, 0x00, 0x06, 0x1f, 0xf2, 0x50, 0x21,
	0x55, 0xbd, 0x93, 0x5b, 0xa9, 0x03, 0x23, 0x94, 0x0e, 0x6d, 0x07, 0xe4, 0xb1, 0x0b, 0x3a, 0x9b,
	0xc3, 0x58, 0xc5, 0x3e, 0xcb, 0x77, 0x76, 0xa1, 0x7f, 0xa0, 0x6b, 0xa4, 0x2c, 0x19, 0x99, 0xf3,
	0x4b, 0x18, 0x19, 0xa3, 0x0c, 0x4e, 0x04, 0x4b, 0x66, 0x7e, 0x41, 0xe4, 0x3b, 0x87, 0x20, 0xb1,
	0x84, 0xc5, 0x7e, 0xef, 0xf9, 0xbd, 0xef, 0x7b, 0xb2, 0x9e, 0x0e, 0x7d, 0x92, 0x72, 0x8e, 0xa3,
	0x89, 0xa0, 0xf3, 0x98, 0xb1, 0x19, 0xce, 0x6c, 0x1c, 0x93, 0x84, 0x84, 0x1c, 0xc5, 0x09, 0x13,
	0xcc, 0x78, 0x2d, 0x1b, 0x50, 0xf1, 0x0d, 0x65, 0x76, 0xcb, 0xf4, 0x18, 0x0f, 0x19, 0xc7, 0x2e,
	0xe1, 0x14, 0x67, 0xb6, 0x4b, 0x05, 0xb1, 0xb1, 0xc7, 0x82, 0x48, 0xb5, 0xb7, 0xde, 0xfb, 0xcc,
	0x67, 0x32, 0xc4, 0x45, 0xa4, 0xaa, 0xd6, 0xff, 0x8a, 0x5e, 0x1d, 0x4a, 0xaa, 0x71, 0x0e, 0xf4,
	0x7a, 0x01, 0x1b, 0x79, 0x09, 0x25, 0x22, 0x60, 0xd1, 0x68, 0x42, 0x69, 0x03, 0xb4, 0x41, 0xa7,
	0xf6, 0xb9, 0x89, 0x14, 0x1d, 0x15, 0x74, 0x54, 0xd2, 0x51, 0x9f, 0x05, 0x51, 0xef, 0xd7, 0x2a,
	0x87, 0xda, 0x2e, 0x87, 0x8d, 0x05, 0x09, 0x67, 0x5d, 0xeb, 0x88, 0x60, 0x5d, 0xde, 0xc0, 0x8e,
	0x1f, 0x88, 0x69, 0xea, 0x22, 0x8f, 0x85, 0xb8, 0x5c, 0x53, 0xbd, 0x3e, 0xf1, 0xf1, 0x1f, 0x2c,
	0x16, 0x31, 0xe5, 0x12, 0xc6, 0x9d, 0x37, 0xc5, 0x7c, 0xbf, 0x1c, 0xff, 0x4e, 0xa9, 0x31, 0xd5,
	0x5f, 0xb9, 0x69, 0x12, 0x8d, 0x62, 0x9a, 0x78, 0x34, 0x12, 0x8d, 0x67, 0x6d, 0xd0, 0x79, 0xd9,
	0x1b, 0x14, 0xd2, 0xeb, 0x1c, 0x7e, 0x78, 0x02, 0xf8, 0x1b, 0xf5, 0x76, 0x39, 0x7c, 0xa7, 0xd6,
	0x3b, 0x64, 0x59, 0x4e, 0xad, 0x48, 0x87, 0x2a, 0x33, 0xfe, 0xea, 0x75, 0x8f, 0x85, 0x61, 0x1a,
	0x05, 0x62, 0xb1, 0xd7, 0x3d, 0x97, 0xba, 0x1f, 0x27, 0xeb, 0xca, 0xbf, 0x71, 0x04, 0xb4, 0x9c,
	0xb7, 0xfb, 0xda, 0x83, 0xf8, 0x9f, 0x6e, 0x64, 0x64, 0x16, 0x8c, 0x89, 0x60, 0x09, 0xdf, 0x9b,
	0x2b, 0xd2, 0xfc, 0xf3, 0x64, 0x73, 0x53, 0x99, 0x8f, 0x89, 0x96, 0x53, 0x7f, 0x2c, 0x96, 0xee,
	0xee, 0x8b, 0x8b, 0x25, 0xd4, 0xee, 0x96, 0x10, 0xf4, 0x06, 0xab, 0x8d, 0x09, 0xd6, 0x1b, 0x13,
	0xdc, 0x6e, 0x4c, 0x70, 0xb6, 0x35, 0xb5, 0xf5, 0xd6, 0xd4, 0xae, 0xb6, 0xa6, 0xf6, 0xfb, 0xe3,
	0x81, 0x5b, 0x1d, 0xa5, 0x7a, 0x66, 0x5f, 0xf1, 0xfc, 0xe0, 0x3e, 0xe5, 0x12, 0x6e, 0x55, 0xde,
	0xd5, 0x97, 0xfb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xd3, 0x65, 0x16, 0xbf, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.PoolCreationFee.Equal(&that1.PoolCreationFee) {
		return false
	}
	if !this.BurnPercent.Equal(that1.BurnPercent) {
		return false
	}
	if !this.CommunityPercent.Equal(that1.CommunityPercent) {
		return false
	}
	if !this.ValidatorsPercent.Equal(that1.ValidatorsPercent) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ValidatorsPercent.Size()
		i -= size
		if _, err := m.ValidatorsPercent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.CommunityPercent.Size()
		i -= size
		if _, err := m.CommunityPercent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BurnPercent.Size()
		i -= size
		if _, err := m.BurnPercent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.PoolCreationFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PoolCreationFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.BurnPercent.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CommunityPercent.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.ValidatorsPercent.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolCreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnPercent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPercent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CommunityPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsPercent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ValidatorsPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
