// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/commitment/params.proto

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

// Params defines the parameters for the module.
type Params struct {
	VestingInfos []*VestingInfo `protobuf:"bytes,1,rep,name=vesting_infos,json=vestingInfos,proto3" json:"vesting_infos,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e317feaf73ff7e, []int{0}
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

func (m *Params) GetVestingInfos() []*VestingInfo {
	if m != nil {
		return m.VestingInfos
	}
	return nil
}

type VestingInfo struct {
	BaseDenom       string                                 `protobuf:"bytes,1,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	VestingDenom    string                                 `protobuf:"bytes,2,opt,name=vesting_denom,json=vestingDenom,proto3" json:"vesting_denom,omitempty"`
	EpochIdentifier string                                 `protobuf:"bytes,3,opt,name=epoch_identifier,json=epochIdentifier,proto3" json:"epoch_identifier,omitempty"`
	NumEpochs       int64                                  `protobuf:"varint,4,opt,name=num_epochs,json=numEpochs,proto3" json:"num_epochs,omitempty"`
	VestNowFactor   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=vest_now_factor,json=vestNowFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"vest_now_factor" yaml:"amount"`
}

func (m *VestingInfo) Reset()         { *m = VestingInfo{} }
func (m *VestingInfo) String() string { return proto.CompactTextString(m) }
func (*VestingInfo) ProtoMessage()    {}
func (*VestingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_92e317feaf73ff7e, []int{1}
}
func (m *VestingInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VestingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VestingInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VestingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VestingInfo.Merge(m, src)
}
func (m *VestingInfo) XXX_Size() int {
	return m.Size()
}
func (m *VestingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VestingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VestingInfo proto.InternalMessageInfo

func (m *VestingInfo) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *VestingInfo) GetVestingDenom() string {
	if m != nil {
		return m.VestingDenom
	}
	return ""
}

func (m *VestingInfo) GetEpochIdentifier() string {
	if m != nil {
		return m.EpochIdentifier
	}
	return ""
}

func (m *VestingInfo) GetNumEpochs() int64 {
	if m != nil {
		return m.NumEpochs
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "elysnetwork.elys.commitment.Params")
	proto.RegisterType((*VestingInfo)(nil), "elysnetwork.elys.commitment.VestingInfo")
}

func init() { proto.RegisterFile("elys/commitment/params.proto", fileDescriptor_92e317feaf73ff7e) }

var fileDescriptor_92e317feaf73ff7e = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x93, 0xb6, 0xb7, 0xd0, 0xe9, 0x2d, 0xbd, 0x84, 0xbb, 0x08, 0xf7, 0x4f, 0x52, 0x22,
	0x48, 0x5c, 0x34, 0x01, 0xdd, 0x75, 0x23, 0x14, 0x15, 0xb2, 0x50, 0x24, 0x0b, 0x17, 0x82, 0x84,
	0x34, 0x9d, 0xa4, 0xa1, 0x9d, 0x39, 0x21, 0x33, 0x69, 0xed, 0xde, 0x07, 0x70, 0xe9, 0xd2, 0xc7,
	0xe9, 0xb2, 0x4b, 0x71, 0x51, 0xa4, 0x7d, 0x03, 0x9f, 0x40, 0x66, 0x52, 0x6d, 0x57, 0xae, 0x72,
	0xce, 0x97, 0xdf, 0xf9, 0xce, 0x9c, 0x73, 0xd0, 0x3f, 0x3c, 0x99, 0x33, 0x37, 0x02, 0x42, 0x52,
	0x4e, 0x30, 0xe5, 0x6e, 0x16, 0xe6, 0x21, 0x61, 0x4e, 0x96, 0x03, 0x07, 0xed, 0xaf, 0xf8, 0x4b,
	0x31, 0x9f, 0x41, 0x3e, 0x76, 0x44, 0xec, 0xec, 0xc8, 0x3f, 0xbf, 0x13, 0x48, 0x40, 0x72, 0xae,
	0x88, 0xca, 0x12, 0xeb, 0x0e, 0xd5, 0xaf, 0xa5, 0x85, 0x76, 0x89, 0x5a, 0x53, 0xcc, 0x78, 0x4a,
	0x93, 0x20, 0xa5, 0x31, 0x30, 0x5d, 0xed, 0x54, 0xed, 0xe6, 0xb1, 0xed, 0x7c, 0x63, 0xea, 0xdc,
	0x94, 0x15, 0x1e, 0x8d, 0xc1, 0xff, 0x39, 0xdd, 0x25, 0xac, 0x57, 0x7b, 0x7a, 0x36, 0x15, 0xeb,
	0xa1, 0x82, 0x9a, 0x7b, 0x8c, 0xf6, 0x1f, 0xa1, 0x41, 0xc8, 0x70, 0x30, 0xc4, 0x14, 0x88, 0xae,
	0x76, 0x54, 0xbb, 0xe1, 0x37, 0x84, 0x72, 0x26, 0x04, 0xed, 0x60, 0xf7, 0x86, 0x92, 0xa8, 0x48,
	0xe2, 0xd3, 0xb9, 0x84, 0x8e, 0xd0, 0x2f, 0x9c, 0x41, 0x34, 0x0a, 0xd2, 0x21, 0xa6, 0x3c, 0x8d,
	0x53, 0x9c, 0xeb, 0x55, 0xc9, 0xb5, 0xa5, 0xee, 0x7d, 0xc9, 0xa2, 0x1d, 0x2d, 0x48, 0x20, 0x65,
	0xa6, 0xd7, 0x3a, 0xaa, 0x5d, 0xf5, 0x1b, 0xb4, 0x20, 0xe7, 0x52, 0xd0, 0x12, 0xd4, 0x16, 0xce,
	0x01, 0x85, 0x59, 0x10, 0x87, 0x11, 0x87, 0x5c, 0xff, 0x21, 0x8c, 0xfa, 0xa7, 0x8b, 0x95, 0xa9,
	0xbc, 0xae, 0xcc, 0xc3, 0x24, 0xe5, 0xa3, 0x62, 0x20, 0x26, 0x76, 0x23, 0x60, 0x04, 0xd8, 0xf6,
	0xd3, 0x65, 0xc3, 0xb1, 0xcb, 0xe7, 0x19, 0x66, 0x8e, 0x47, 0xf9, 0xfb, 0xca, 0x6c, 0xcd, 0x43,
	0x32, 0xe9, 0x59, 0x21, 0x81, 0x82, 0x72, 0xcb, 0x97, 0x63, 0x5c, 0xc1, 0xec, 0x42, 0xba, 0xf6,
	0xbd, 0xc5, 0xda, 0x50, 0x97, 0x6b, 0x43, 0x7d, 0x5b, 0x1b, 0xea, 0xe3, 0xc6, 0x50, 0x96, 0x1b,
	0x43, 0x79, 0xd9, 0x18, 0xca, 0xad, 0xbb, 0xd7, 0x41, 0x2c, 0xb7, 0xbb, 0xdd, 0xb4, 0x4c, 0xdc,
	0xfb, 0xfd, 0x53, 0xcb, 0x76, 0x83, 0xba, 0xbc, 0xdb, 0xc9, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x52, 0xb3, 0x74, 0x3b, 0x0a, 0x02, 0x00, 0x00,
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
	if len(m.VestingInfos) > 0 {
		for iNdEx := len(m.VestingInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestingInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VestingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VestingInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VestingInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.VestNowFactor.Size()
		i -= size
		if _, err := m.VestNowFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.NumEpochs != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.NumEpochs))
		i--
		dAtA[i] = 0x20
	}
	if len(m.EpochIdentifier) > 0 {
		i -= len(m.EpochIdentifier)
		copy(dAtA[i:], m.EpochIdentifier)
		i = encodeVarintParams(dAtA, i, uint64(len(m.EpochIdentifier)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VestingDenom) > 0 {
		i -= len(m.VestingDenom)
		copy(dAtA[i:], m.VestingDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.VestingDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0xa
	}
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
	if len(m.VestingInfos) > 0 {
		for _, e := range m.VestingInfos {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *VestingInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.VestingDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.EpochIdentifier)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.NumEpochs != 0 {
		n += 1 + sovParams(uint64(m.NumEpochs))
	}
	l = m.VestNowFactor.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field VestingInfos", wireType)
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
			m.VestingInfos = append(m.VestingInfos, &VestingInfo{})
			if err := m.VestingInfos[len(m.VestingInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *VestingInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: VestingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VestingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
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
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingDenom", wireType)
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
			m.VestingDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochIdentifier", wireType)
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
			m.EpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochs", wireType)
			}
			m.NumEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestNowFactor", wireType)
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
			if err := m.VestNowFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
