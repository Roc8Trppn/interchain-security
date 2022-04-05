// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/child/v1/child.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// CrossChainValidator defines the validators for CCV child module
type CrossChainValidator struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Power   int64  `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
}

func (m *CrossChainValidator) Reset()         { *m = CrossChainValidator{} }
func (m *CrossChainValidator) String() string { return proto.CompactTextString(m) }
func (*CrossChainValidator) ProtoMessage()    {}
func (*CrossChainValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00ed7c78fea69b9, []int{0}
}
func (m *CrossChainValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrossChainValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrossChainValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrossChainValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossChainValidator.Merge(m, src)
}
func (m *CrossChainValidator) XXX_Size() int {
	return m.Size()
}
func (m *CrossChainValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossChainValidator.DiscardUnknown(m)
}

var xxx_messageInfo_CrossChainValidator proto.InternalMessageInfo

func (m *CrossChainValidator) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *CrossChainValidator) GetPower() int64 {
	if m != nil {
		return m.Power
	}
	return 0
}

// Params defines the parameters for CCV child module
type Params struct {
	Enabled bool `protobuf:"varint,1,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	// distribution params
	BlocksPerDistributionTransmission int64  `protobuf:"varint,2,opt,name=BlocksPerDistributionTransmission,proto3" json:"BlocksPerDistributionTransmission,omitempty"`
	ProviderFeePoolAddrStr            string `protobuf:"bytes,3,opt,name=ProviderFeePoolAddrStr,proto3" json:"ProviderFeePoolAddrStr,omitempty"`
	DistributionTransmissionChannel   string `protobuf:"bytes,4,opt,name=DistributionTransmissionChannel,proto3" json:"DistributionTransmissionChannel,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00ed7c78fea69b9, []int{1}
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

func (m *Params) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Params) GetBlocksPerDistributionTransmission() int64 {
	if m != nil {
		return m.BlocksPerDistributionTransmission
	}
	return 0
}

func (m *Params) GetProviderFeePoolAddrStr() string {
	if m != nil {
		return m.ProviderFeePoolAddrStr
	}
	return ""
}

func (m *Params) GetDistributionTransmissionChannel() string {
	if m != nil {
		return m.DistributionTransmissionChannel
	}
	return ""
}

// LastTransmissionBlockHeight is the last time validator holding
// pools were transmitted to the provider chain
type LastTransmissionBlockHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=Height,proto3" json:"Height,omitempty"`
}

func (m *LastTransmissionBlockHeight) Reset()         { *m = LastTransmissionBlockHeight{} }
func (m *LastTransmissionBlockHeight) String() string { return proto.CompactTextString(m) }
func (*LastTransmissionBlockHeight) ProtoMessage()    {}
func (*LastTransmissionBlockHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_d00ed7c78fea69b9, []int{2}
}
func (m *LastTransmissionBlockHeight) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastTransmissionBlockHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastTransmissionBlockHeight.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastTransmissionBlockHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastTransmissionBlockHeight.Merge(m, src)
}
func (m *LastTransmissionBlockHeight) XXX_Size() int {
	return m.Size()
}
func (m *LastTransmissionBlockHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_LastTransmissionBlockHeight.DiscardUnknown(m)
}

var xxx_messageInfo_LastTransmissionBlockHeight proto.InternalMessageInfo

func (m *LastTransmissionBlockHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*CrossChainValidator)(nil), "interchain_security.ccv.child.v1.CrossChainValidator")
	proto.RegisterType((*Params)(nil), "interchain_security.ccv.child.v1.Params")
	proto.RegisterType((*LastTransmissionBlockHeight)(nil), "interchain_security.ccv.child.v1.LastTransmissionBlockHeight")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/child/v1/child.proto", fileDescriptor_d00ed7c78fea69b9)
}

var fileDescriptor_d00ed7c78fea69b9 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbd, 0xae, 0xd3, 0x30,
	0x18, 0xad, 0x29, 0x14, 0xb0, 0x98, 0x42, 0x55, 0x45, 0x45, 0x0a, 0xa1, 0x53, 0x07, 0x88, 0x15,
	0x21, 0x60, 0xa6, 0xa5, 0xa8, 0x43, 0x87, 0x10, 0x10, 0x03, 0x0b, 0x72, 0x6c, 0x2b, 0xb1, 0x48,
	0xfc, 0x55, 0xfe, 0xdc, 0x40, 0xdf, 0x82, 0xc7, 0x62, 0xec, 0xc8, 0x78, 0xd5, 0xbe, 0xc1, 0x7d,
	0x82, 0xab, 0x26, 0xed, 0xbd, 0x1d, 0x6e, 0xd5, 0xed, 0x1c, 0x9d, 0x1f, 0x59, 0x3e, 0x1f, 0x7d,
	0xad, 0x8d, 0x53, 0x56, 0x14, 0x5c, 0x9b, 0x9f, 0xa8, 0xc4, 0xca, 0x6a, 0xb7, 0x66, 0x42, 0xd4,
	0x4c, 0x14, 0xba, 0x94, 0xac, 0x8e, 0x5b, 0x10, 0x2d, 0x2d, 0x38, 0xf0, 0xc2, 0x7b, 0xdc, 0x91,
	0x10, 0x75, 0xd4, 0x9a, 0xea, 0x78, 0xd8, 0xcf, 0x21, 0x87, 0xc6, 0xcc, 0xf6, 0xa8, 0xcd, 0x0d,
	0x03, 0x01, 0x58, 0x01, 0xb2, 0x8c, 0xa3, 0x62, 0x75, 0x9c, 0x29, 0xc7, 0x63, 0x26, 0x40, 0x9b,
	0x56, 0x1f, 0xcd, 0xe8, 0xf3, 0xa9, 0x05, 0xc4, 0xe9, 0xbe, 0xf9, 0x3b, 0x2f, 0xb5, 0xe4, 0x0e,
	0xac, 0xe7, 0xd3, 0xc7, 0x5c, 0x4a, 0xab, 0x10, 0x7d, 0x12, 0x92, 0xf1, 0xb3, 0xf4, 0x48, 0xbd,
	0x3e, 0x7d, 0xb4, 0x84, 0xdf, 0xca, 0xfa, 0x0f, 0x42, 0x32, 0xee, 0xa6, 0x2d, 0x19, 0x5d, 0x13,
	0xda, 0x4b, 0xb8, 0xe5, 0x15, 0xee, 0xa3, 0x33, 0xc3, 0xb3, 0x52, 0xc9, 0x26, 0xfa, 0x24, 0x3d,
	0x52, 0x6f, 0x41, 0x5f, 0x4d, 0x4a, 0x10, 0xbf, 0x30, 0x51, 0xf6, 0x93, 0x46, 0x67, 0x75, 0xb6,
	0x72, 0x1a, 0xcc, 0x37, 0xcb, 0x0d, 0x56, 0x1a, 0x51, 0x83, 0x39, 0xd4, 0x5e, 0x36, 0x7a, 0xef,
	0xe9, 0x20, 0xb1, 0x50, 0x6b, 0xa9, 0xec, 0x67, 0xa5, 0x12, 0x80, 0xf2, 0xa3, 0x94, 0xf6, 0xab,
	0xb3, 0x7e, 0x37, 0x24, 0xe3, 0xa7, 0xe9, 0x19, 0xd5, 0x9b, 0xd3, 0x97, 0xe7, 0x3a, 0xa7, 0x05,
	0x37, 0x46, 0x95, 0xfe, 0xc3, 0xa6, 0xe0, 0x92, 0x6d, 0xf4, 0x8e, 0xbe, 0x58, 0x70, 0x74, 0xa7,
	0x52, 0xf3, 0xec, 0xb9, 0xd2, 0x79, 0xe1, 0xbc, 0x01, 0xed, 0xb5, 0xa8, 0xf9, 0x87, 0x6e, 0x7a,
	0x60, 0x93, 0x2f, 0xff, 0xb6, 0x01, 0xd9, 0x6c, 0x03, 0x72, 0xb5, 0x0d, 0xc8, 0xdf, 0x5d, 0xd0,
	0xd9, 0xec, 0x82, 0xce, 0xff, 0x5d, 0xd0, 0xf9, 0xf1, 0x21, 0xd7, 0xae, 0x58, 0x65, 0x91, 0x80,
	0x8a, 0x1d, 0x76, 0xbb, 0x9b, 0xfd, 0xcd, 0xed, 0x91, 0xfc, 0x39, 0x39, 0x13, 0xb7, 0x5e, 0x2a,
	0xcc, 0x7a, 0xcd, 0x98, 0x6f, 0x6f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x3c, 0x05, 0x2a, 0x54,
	0x02, 0x00, 0x00,
}

func (m *CrossChainValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossChainValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrossChainValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Power != 0 {
		i = encodeVarintChild(dAtA, i, uint64(m.Power))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintChild(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	if len(m.DistributionTransmissionChannel) > 0 {
		i -= len(m.DistributionTransmissionChannel)
		copy(dAtA[i:], m.DistributionTransmissionChannel)
		i = encodeVarintChild(dAtA, i, uint64(len(m.DistributionTransmissionChannel)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProviderFeePoolAddrStr) > 0 {
		i -= len(m.ProviderFeePoolAddrStr)
		copy(dAtA[i:], m.ProviderFeePoolAddrStr)
		i = encodeVarintChild(dAtA, i, uint64(len(m.ProviderFeePoolAddrStr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlocksPerDistributionTransmission != 0 {
		i = encodeVarintChild(dAtA, i, uint64(m.BlocksPerDistributionTransmission))
		i--
		dAtA[i] = 0x10
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LastTransmissionBlockHeight) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastTransmissionBlockHeight) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastTransmissionBlockHeight) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintChild(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintChild(dAtA []byte, offset int, v uint64) int {
	offset -= sovChild(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CrossChainValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovChild(uint64(l))
	}
	if m.Power != 0 {
		n += 1 + sovChild(uint64(m.Power))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Enabled {
		n += 2
	}
	if m.BlocksPerDistributionTransmission != 0 {
		n += 1 + sovChild(uint64(m.BlocksPerDistributionTransmission))
	}
	l = len(m.ProviderFeePoolAddrStr)
	if l > 0 {
		n += 1 + l + sovChild(uint64(l))
	}
	l = len(m.DistributionTransmissionChannel)
	if l > 0 {
		n += 1 + l + sovChild(uint64(l))
	}
	return n
}

func (m *LastTransmissionBlockHeight) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovChild(uint64(m.Height))
	}
	return n
}

func sovChild(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozChild(x uint64) (n int) {
	return sovChild(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CrossChainValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChild
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
			return fmt.Errorf("proto: CrossChainValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossChainValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthChild
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChild
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Power", wireType)
			}
			m.Power = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Power |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChild(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChild
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChild
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerDistributionTransmission", wireType)
			}
			m.BlocksPerDistributionTransmission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerDistributionTransmission |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderFeePoolAddrStr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
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
				return ErrInvalidLengthChild
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChild
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderFeePoolAddrStr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionTransmissionChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
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
				return ErrInvalidLengthChild
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthChild
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributionTransmissionChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChild(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChild
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
func (m *LastTransmissionBlockHeight) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChild
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
			return fmt.Errorf("proto: LastTransmissionBlockHeight: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastTransmissionBlockHeight: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChild
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipChild(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthChild
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
func skipChild(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChild
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
					return 0, ErrIntOverflowChild
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
					return 0, ErrIntOverflowChild
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
				return 0, ErrInvalidLengthChild
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupChild
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthChild
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthChild        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChild          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupChild = fmt.Errorf("proto: unexpected end of group")
)
