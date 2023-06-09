// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gauss/validator-dao/v1/dao.proto

package types

import (
	fmt "fmt"
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

type DaoBiz struct {
	Name string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fee  types.Coin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee"`
}

func (m *DaoBiz) Reset()         { *m = DaoBiz{} }
func (m *DaoBiz) String() string { return proto.CompactTextString(m) }
func (*DaoBiz) ProtoMessage()    {}
func (*DaoBiz) Descriptor() ([]byte, []int) {
	return fileDescriptor_af10164922e44483, []int{0}
}
func (m *DaoBiz) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DaoBiz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DaoBiz.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DaoBiz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DaoBiz.Merge(m, src)
}
func (m *DaoBiz) XXX_Size() int {
	return m.Size()
}
func (m *DaoBiz) XXX_DiscardUnknown() {
	xxx_messageInfo_DaoBiz.DiscardUnknown(m)
}

var xxx_messageInfo_DaoBiz proto.InternalMessageInfo

func (m *DaoBiz) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DaoBiz) GetFee() types.Coin {
	if m != nil {
		return m.Fee
	}
	return types.Coin{}
}

type AuthorizerBizs struct {
	Authorizer string   `protobuf:"bytes,1,opt,name=authorizer,proto3" json:"authorizer,omitempty"`
	Bizs       []DaoBiz `protobuf:"bytes,2,rep,name=bizs,proto3" json:"bizs"`
}

func (m *AuthorizerBizs) Reset()         { *m = AuthorizerBizs{} }
func (m *AuthorizerBizs) String() string { return proto.CompactTextString(m) }
func (*AuthorizerBizs) ProtoMessage()    {}
func (*AuthorizerBizs) Descriptor() ([]byte, []int) {
	return fileDescriptor_af10164922e44483, []int{1}
}
func (m *AuthorizerBizs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthorizerBizs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthorizerBizs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthorizerBizs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizerBizs.Merge(m, src)
}
func (m *AuthorizerBizs) XXX_Size() int {
	return m.Size()
}
func (m *AuthorizerBizs) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizerBizs.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizerBizs proto.InternalMessageInfo

func (m *AuthorizerBizs) GetAuthorizer() string {
	if m != nil {
		return m.Authorizer
	}
	return ""
}

func (m *AuthorizerBizs) GetBizs() []DaoBiz {
	if m != nil {
		return m.Bizs
	}
	return nil
}

type AcqBiz struct {
	From    string     `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	BizName string     `protobuf:"bytes,2,opt,name=biz_name,json=bizName,proto3" json:"biz_name,omitempty"`
	Amount  types.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	Price   types.Coin `protobuf:"bytes,4,opt,name=price,proto3" json:"price"`
}

func (m *AcqBiz) Reset()         { *m = AcqBiz{} }
func (m *AcqBiz) String() string { return proto.CompactTextString(m) }
func (*AcqBiz) ProtoMessage()    {}
func (*AcqBiz) Descriptor() ([]byte, []int) {
	return fileDescriptor_af10164922e44483, []int{2}
}
func (m *AcqBiz) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AcqBiz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AcqBiz.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AcqBiz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcqBiz.Merge(m, src)
}
func (m *AcqBiz) XXX_Size() int {
	return m.Size()
}
func (m *AcqBiz) XXX_DiscardUnknown() {
	xxx_messageInfo_AcqBiz.DiscardUnknown(m)
}

var xxx_messageInfo_AcqBiz proto.InternalMessageInfo

func (m *AcqBiz) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *AcqBiz) GetBizName() string {
	if m != nil {
		return m.BizName
	}
	return ""
}

func (m *AcqBiz) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func (m *AcqBiz) GetPrice() types.Coin {
	if m != nil {
		return m.Price
	}
	return types.Coin{}
}

type GranteeBizs struct {
	Grantee string   `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	Bizs    []AcqBiz `protobuf:"bytes,2,rep,name=bizs,proto3" json:"bizs"`
}

func (m *GranteeBizs) Reset()         { *m = GranteeBizs{} }
func (m *GranteeBizs) String() string { return proto.CompactTextString(m) }
func (*GranteeBizs) ProtoMessage()    {}
func (*GranteeBizs) Descriptor() ([]byte, []int) {
	return fileDescriptor_af10164922e44483, []int{3}
}
func (m *GranteeBizs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GranteeBizs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GranteeBizs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GranteeBizs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GranteeBizs.Merge(m, src)
}
func (m *GranteeBizs) XXX_Size() int {
	return m.Size()
}
func (m *GranteeBizs) XXX_DiscardUnknown() {
	xxx_messageInfo_GranteeBizs.DiscardUnknown(m)
}

var xxx_messageInfo_GranteeBizs proto.InternalMessageInfo

func (m *GranteeBizs) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *GranteeBizs) GetBizs() []AcqBiz {
	if m != nil {
		return m.Bizs
	}
	return nil
}

func init() {
	proto.RegisterType((*DaoBiz)(nil), "gauss.validatordao.v1.DaoBiz")
	proto.RegisterType((*AuthorizerBizs)(nil), "gauss.validatordao.v1.AuthorizerBizs")
	proto.RegisterType((*AcqBiz)(nil), "gauss.validatordao.v1.AcqBiz")
	proto.RegisterType((*GranteeBizs)(nil), "gauss.validatordao.v1.GranteeBizs")
}

func init() { proto.RegisterFile("gauss/validator-dao/v1/dao.proto", fileDescriptor_af10164922e44483) }

var fileDescriptor_af10164922e44483 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x4d, 0xda, 0x90, 0x82, 0x2b, 0x31, 0x58, 0x20, 0xa5, 0x95, 0x30, 0x51, 0xa6, 0x2e, 0xd8,
	0x4a, 0x11, 0x74, 0x6e, 0x40, 0x42, 0x2c, 0x20, 0x75, 0x64, 0x01, 0x27, 0x75, 0x53, 0x4b, 0x24,
	0xb7, 0xc4, 0x4e, 0x04, 0xf9, 0x0a, 0xfe, 0x82, 0x5f, 0xe9, 0xd8, 0x91, 0x09, 0xa1, 0xf6, 0x47,
	0x9e, 0x62, 0xa7, 0x55, 0xf5, 0xf4, 0x86, 0x6e, 0xd7, 0xf7, 0x9e, 0x7b, 0x7c, 0xce, 0xd5, 0x41,
	0x61, 0xce, 0x6b, 0xa5, 0x58, 0xc3, 0xbf, 0xcb, 0x35, 0xd7, 0x50, 0xbd, 0x5a, 0x73, 0x60, 0x4d,
	0xcc, 0xd6, 0x1c, 0xe8, 0xae, 0x02, 0x0d, 0xf8, 0xb9, 0x41, 0xd0, 0x0b, 0xa2, 0x9b, 0x34, 0xf1,
	0xf4, 0x59, 0x0e, 0x39, 0x18, 0x04, 0xeb, 0x2a, 0x0b, 0x9e, 0x92, 0x0c, 0x54, 0x01, 0x8a, 0xa5,
	0x5c, 0x09, 0xd6, 0xc4, 0xa9, 0xd0, 0x3c, 0x66, 0x19, 0xc8, 0xd2, 0xce, 0xa3, 0xcf, 0xc8, 0x7f,
	0xcf, 0x21, 0x91, 0x2d, 0xc6, 0xc8, 0x2b, 0x79, 0x21, 0x02, 0x37, 0x74, 0x67, 0x4f, 0x56, 0xa6,
	0xc6, 0x31, 0x1a, 0x6e, 0x84, 0x08, 0x06, 0xa1, 0x3b, 0x1b, 0xcf, 0x27, 0xd4, 0x72, 0xd1, 0x8e,
	0x8b, 0xf6, 0x5c, 0xf4, 0x1d, 0xc8, 0x32, 0xf1, 0xf6, 0xff, 0x5e, 0x3a, 0xab, 0x0e, 0x1b, 0x49,
	0xf4, 0x74, 0x59, 0xeb, 0x2d, 0x54, 0xb2, 0x15, 0x55, 0x22, 0x5b, 0x85, 0x09, 0x42, 0xfc, 0xd2,
	0xe9, 0xe9, 0xaf, 0x3a, 0x78, 0x81, 0xbc, 0x54, 0xb6, 0x2a, 0x18, 0x84, 0xc3, 0xd9, 0x78, 0xfe,
	0x82, 0x3e, 0x68, 0x8f, 0x5a, 0x95, 0xfd, 0x4f, 0x66, 0x21, 0xfa, 0xe3, 0x22, 0x7f, 0x99, 0xfd,
	0xe8, 0xc5, 0x6f, 0x2a, 0x28, 0xce, 0xe2, 0xbb, 0x1a, 0x4f, 0xd0, 0xe3, 0x54, 0xb6, 0x5f, 0x8d,
	0xa9, 0x81, 0xe9, 0x8f, 0x52, 0xd9, 0x7e, 0xea, 0x7c, 0x2d, 0x90, 0xcf, 0x0b, 0xa8, 0x4b, 0x1d,
	0x0c, 0x6f, 0xb3, 0xd6, 0xc3, 0xf1, 0x1b, 0xf4, 0x68, 0x57, 0xc9, 0x4c, 0x04, 0xde, 0x6d, 0x7b,
	0x16, 0x1d, 0x7d, 0x43, 0xe3, 0x0f, 0x15, 0x2f, 0xb5, 0x10, 0xe6, 0x22, 0x01, 0x1a, 0xe5, 0xf6,
	0xd9, 0x0b, 0x3e, 0x3f, 0x6f, 0xbc, 0x85, 0x35, 0x7d, 0x7d, 0x8b, 0xe4, 0xe3, 0xfe, 0x48, 0xdc,
	0xc3, 0x91, 0xb8, 0xff, 0x8f, 0xc4, 0xfd, 0x7d, 0x22, 0xce, 0xe1, 0x44, 0x9c, 0xbf, 0x27, 0xe2,
	0x7c, 0x61, 0xb9, 0xd4, 0xdb, 0x3a, 0xa5, 0x19, 0x14, 0xcc, 0x66, 0xab, 0x4f, 0xd8, 0x5b, 0xf6,
	0xf3, 0x5e, 0xcc, 0xf4, 0xaf, 0x9d, 0x50, 0xa9, 0x6f, 0x92, 0xf1, 0xfa, 0x2e, 0x00, 0x00, 0xff,
	0xff, 0x37, 0xd0, 0xb7, 0xd0, 0x8a, 0x02, 0x00, 0x00,
}

func (m *DaoBiz) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DaoBiz) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DaoBiz) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDao(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthorizerBizs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizerBizs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthorizerBizs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bizs) > 0 {
		for iNdEx := len(m.Bizs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bizs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Authorizer) > 0 {
		i -= len(m.Authorizer)
		copy(dAtA[i:], m.Authorizer)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Authorizer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AcqBiz) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AcqBiz) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AcqBiz) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDao(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintDao(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.BizName) > 0 {
		i -= len(m.BizName)
		copy(dAtA[i:], m.BizName)
		i = encodeVarintDao(dAtA, i, uint64(len(m.BizName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintDao(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GranteeBizs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GranteeBizs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GranteeBizs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bizs) > 0 {
		for iNdEx := len(m.Bizs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Bizs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDao(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintDao(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDao(dAtA []byte, offset int, v uint64) int {
	offset -= sovDao(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DaoBiz) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = m.Fee.Size()
	n += 1 + l + sovDao(uint64(l))
	return n
}

func (m *AuthorizerBizs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authorizer)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	if len(m.Bizs) > 0 {
		for _, e := range m.Bizs {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	return n
}

func (m *AcqBiz) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = len(m.BizName)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovDao(uint64(l))
	l = m.Price.Size()
	n += 1 + l + sovDao(uint64(l))
	return n
}

func (m *GranteeBizs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovDao(uint64(l))
	}
	if len(m.Bizs) > 0 {
		for _, e := range m.Bizs {
			l = e.Size()
			n += 1 + l + sovDao(uint64(l))
		}
	}
	return n
}

func sovDao(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDao(x uint64) (n int) {
	return sovDao(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DaoBiz) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: DaoBiz: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DaoBiz: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *AuthorizerBizs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: AuthorizerBizs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizerBizs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorizer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authorizer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bizs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bizs = append(m.Bizs, DaoBiz{})
			if err := m.Bizs[len(m.Bizs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *AcqBiz) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: AcqBiz: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AcqBiz: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BizName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BizName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func (m *GranteeBizs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDao
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
			return fmt.Errorf("proto: GranteeBizs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GranteeBizs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bizs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDao
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
				return ErrInvalidLengthDao
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDao
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bizs = append(m.Bizs, AcqBiz{})
			if err := m.Bizs[len(m.Bizs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDao(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDao
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
func skipDao(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDao
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
					return 0, ErrIntOverflowDao
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
					return 0, ErrIntOverflowDao
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
				return 0, ErrInvalidLengthDao
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDao
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDao
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDao        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDao          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDao = fmt.Errorf("proto: unexpected end of group")
)
