// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/tradeshield/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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
	MarketOrderEnabled   bool   `protobuf:"varint,1,opt,name=market_order_enabled,json=marketOrderEnabled,proto3" json:"market_order_enabled,omitempty"`
	StakeEnabled         bool   `protobuf:"varint,2,opt,name=stake_enabled,json=stakeEnabled,proto3" json:"stake_enabled,omitempty"`
	ProcessOrdersEnabled bool   `protobuf:"varint,3,opt,name=process_orders_enabled,json=processOrdersEnabled,proto3" json:"process_orders_enabled,omitempty"`
	SwapEnabled          bool   `protobuf:"varint,4,opt,name=swap_enabled,json=swapEnabled,proto3" json:"swap_enabled,omitempty"`
	PerpetualEnabled     bool   `protobuf:"varint,5,opt,name=perpetual_enabled,json=perpetualEnabled,proto3" json:"perpetual_enabled,omitempty"`
	RewardEnabled        bool   `protobuf:"varint,6,opt,name=reward_enabled,json=rewardEnabled,proto3" json:"reward_enabled,omitempty"`
	LeverageEnabled      bool   `protobuf:"varint,7,opt,name=leverage_enabled,json=leverageEnabled,proto3" json:"leverage_enabled,omitempty"`
	LimitProcessOrder    uint64 `protobuf:"varint,8,opt,name=limit_process_order,json=limitProcessOrder,proto3" json:"limit_process_order,omitempty"`
	// For incentive system v2
	RewardPercentage cosmossdk_io_math.LegacyDec `protobuf:"bytes,9,opt,name=reward_percentage,json=rewardPercentage,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"reward_percentage"`
	MarginError      cosmossdk_io_math.LegacyDec `protobuf:"bytes,10,opt,name=margin_error,json=marginError,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"margin_error"`
	MinimumDeposit   cosmossdk_io_math.Int       `protobuf:"bytes,11,opt,name=minimum_deposit,json=minimumDeposit,proto3,customtype=cosmossdk.io/math.Int" json:"minimum_deposit"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_01646d3b2ec43fec, []int{0}
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

func (m *Params) GetMarketOrderEnabled() bool {
	if m != nil {
		return m.MarketOrderEnabled
	}
	return false
}

func (m *Params) GetStakeEnabled() bool {
	if m != nil {
		return m.StakeEnabled
	}
	return false
}

func (m *Params) GetProcessOrdersEnabled() bool {
	if m != nil {
		return m.ProcessOrdersEnabled
	}
	return false
}

func (m *Params) GetSwapEnabled() bool {
	if m != nil {
		return m.SwapEnabled
	}
	return false
}

func (m *Params) GetPerpetualEnabled() bool {
	if m != nil {
		return m.PerpetualEnabled
	}
	return false
}

func (m *Params) GetRewardEnabled() bool {
	if m != nil {
		return m.RewardEnabled
	}
	return false
}

func (m *Params) GetLeverageEnabled() bool {
	if m != nil {
		return m.LeverageEnabled
	}
	return false
}

func (m *Params) GetLimitProcessOrder() uint64 {
	if m != nil {
		return m.LimitProcessOrder
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "elys.tradeshield.Params")
}

func init() { proto.RegisterFile("elys/tradeshield/params.proto", fileDescriptor_01646d3b2ec43fec) }

var fileDescriptor_01646d3b2ec43fec = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0x80, 0x1b, 0x28, 0x65, 0x73, 0xbb, 0xad, 0x0d, 0x05, 0x85, 0x21, 0xb2, 0x02, 0x42, 0x2a,
	0x9a, 0x96, 0x0c, 0xc1, 0x13, 0x4c, 0xdd, 0x61, 0x08, 0x89, 0xaa, 0xda, 0x89, 0x03, 0x91, 0x9b,
	0xfc, 0x4a, 0xad, 0xc6, 0xb1, 0xf5, 0xdb, 0xa5, 0xf4, 0x2d, 0x78, 0x18, 0x1e, 0x62, 0xc7, 0x89,
	0x13, 0xe2, 0x30, 0xa1, 0xf6, 0x11, 0x78, 0x01, 0x14, 0x3b, 0x09, 0x9d, 0xb8, 0x71, 0x8b, 0xff,
	0xef, 0xcb, 0xa7, 0x58, 0x76, 0xc8, 0x53, 0xc8, 0x56, 0x2a, 0xd4, 0x48, 0x13, 0x50, 0x33, 0x06,
	0x59, 0x12, 0x4a, 0x8a, 0x94, 0xab, 0x40, 0xa2, 0xd0, 0xc2, 0xed, 0x16, 0x38, 0xd8, 0xc2, 0x87,
	0xfd, 0x54, 0xa4, 0xc2, 0xc0, 0xb0, 0x78, 0xb2, 0xde, 0xe1, 0xe3, 0x58, 0x28, 0x2e, 0x54, 0x64,
	0x81, 0x5d, 0x58, 0xf4, 0xfc, 0x77, 0x93, 0xb4, 0xc6, 0xa6, 0xe9, 0x9e, 0x92, 0x3e, 0xa7, 0x38,
	0x07, 0x1d, 0x09, 0x4c, 0x00, 0x23, 0xc8, 0xe9, 0x34, 0x83, 0xc4, 0x73, 0x06, 0xce, 0x70, 0x67,
	0xe2, 0x5a, 0xf6, 0xa1, 0x40, 0xe7, 0x96, 0xb8, 0x2f, 0xc8, 0x9e, 0xd2, 0x74, 0x0e, 0xb5, 0x7a,
	0xc7, 0xa8, 0x1d, 0x33, 0xac, 0xa4, 0xb7, 0xe4, 0x91, 0x44, 0x11, 0x83, 0x52, 0xb6, 0xab, 0x6a,
	0xfb, 0xae, 0xb1, 0xfb, 0x25, 0x35, 0x65, 0x55, 0xbd, 0xf5, 0x8c, 0x74, 0xd4, 0x92, 0xca, 0xda,
	0x6d, 0x1a, 0xb7, 0x5d, 0xcc, 0x2a, 0xe5, 0x98, 0xf4, 0x24, 0xa0, 0x04, 0xbd, 0xa0, 0x59, 0xed,
	0xdd, 0x33, 0x5e, 0xb7, 0x06, 0x95, 0xfc, 0x92, 0xec, 0x23, 0x2c, 0x29, 0x26, 0xb5, 0xd9, 0x32,
	0xe6, 0x9e, 0x9d, 0x56, 0xda, 0x2b, 0xd2, 0xcd, 0xe0, 0x33, 0x20, 0x4d, 0xff, 0x6e, 0xea, 0xbe,
	0x11, 0x0f, 0xaa, 0x79, 0xa5, 0x06, 0xe4, 0x41, 0xc6, 0x38, 0xd3, 0xd1, 0xad, 0xdd, 0x79, 0x3b,
	0x03, 0x67, 0xd8, 0x9c, 0xf4, 0x0c, 0x1a, 0x6f, 0xed, 0xcc, 0xfd, 0x44, 0x7a, 0xe5, 0x17, 0x48,
	0xc0, 0x18, 0x72, 0x4d, 0x53, 0xf0, 0x76, 0x07, 0xce, 0x70, 0xf7, 0xec, 0xf5, 0xd5, 0xcd, 0x51,
	0xe3, 0xe7, 0xcd, 0xd1, 0x13, 0x7b, 0x34, 0x2a, 0x99, 0x07, 0x4c, 0x84, 0x9c, 0xea, 0x59, 0xf0,
	0x1e, 0x52, 0x1a, 0xaf, 0x46, 0x10, 0x7f, 0xff, 0x76, 0x42, 0xca, 0x93, 0x1b, 0x41, 0x3c, 0xe9,
	0xda, 0xd6, 0xb8, 0x4e, 0xb9, 0x97, 0xa4, 0xc3, 0x29, 0xa6, 0x2c, 0x8f, 0x00, 0x51, 0xa0, 0x47,
	0xfe, 0x37, 0xdd, 0xb6, 0x99, 0xf3, 0xa2, 0xe2, 0x5e, 0x92, 0x03, 0xce, 0x72, 0xc6, 0x17, 0x3c,
	0x4a, 0x40, 0x0a, 0xc5, 0xb4, 0xd7, 0x36, 0xe1, 0xe3, 0x32, 0xfc, 0xf0, 0xdf, 0xf0, 0x45, 0xae,
	0xb7, 0x92, 0x17, 0xb9, 0x9e, 0xec, 0x97, 0x8d, 0x91, 0x4d, 0x9c, 0xbd, 0xbb, 0x5a, 0xfb, 0xce,
	0xf5, 0xda, 0x77, 0x7e, 0xad, 0x7d, 0xe7, 0xeb, 0xc6, 0x6f, 0x5c, 0x6f, 0xfc, 0xc6, 0x8f, 0x8d,
	0xdf, 0xf8, 0x78, 0x9a, 0x32, 0x3d, 0x5b, 0x4c, 0x83, 0x58, 0xf0, 0xb0, 0xb8, 0xdd, 0x27, 0x39,
	0xe8, 0xa5, 0xc0, 0xb9, 0x59, 0x84, 0x5f, 0x6e, 0xfd, 0x0b, 0x7a, 0x25, 0x41, 0x4d, 0x5b, 0xe6,
	0x22, 0xbf, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xa6, 0x97, 0x6c, 0x2c, 0x03, 0x00, 0x00,
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
		size := m.MinimumDeposit.Size()
		i -= size
		if _, err := m.MinimumDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.MarginError.Size()
		i -= size
		if _, err := m.MarginError.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	{
		size := m.RewardPercentage.Size()
		i -= size
		if _, err := m.RewardPercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.LimitProcessOrder != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.LimitProcessOrder))
		i--
		dAtA[i] = 0x40
	}
	if m.LeverageEnabled {
		i--
		if m.LeverageEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.RewardEnabled {
		i--
		if m.RewardEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.PerpetualEnabled {
		i--
		if m.PerpetualEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.SwapEnabled {
		i--
		if m.SwapEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.ProcessOrdersEnabled {
		i--
		if m.ProcessOrdersEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.StakeEnabled {
		i--
		if m.StakeEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.MarketOrderEnabled {
		i--
		if m.MarketOrderEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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
	if m.MarketOrderEnabled {
		n += 2
	}
	if m.StakeEnabled {
		n += 2
	}
	if m.ProcessOrdersEnabled {
		n += 2
	}
	if m.SwapEnabled {
		n += 2
	}
	if m.PerpetualEnabled {
		n += 2
	}
	if m.RewardEnabled {
		n += 2
	}
	if m.LeverageEnabled {
		n += 2
	}
	if m.LimitProcessOrder != 0 {
		n += 1 + sovParams(uint64(m.LimitProcessOrder))
	}
	l = m.RewardPercentage.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MarginError.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinimumDeposit.Size()
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketOrderEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.MarketOrderEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.StakeEnabled = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProcessOrdersEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.ProcessOrdersEnabled = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.SwapEnabled = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.PerpetualEnabled = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.RewardEnabled = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeverageEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.LeverageEnabled = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitProcessOrder", wireType)
			}
			m.LimitProcessOrder = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LimitProcessOrder |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPercentage", wireType)
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
			if err := m.RewardPercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarginError", wireType)
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
			if err := m.MarginError.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumDeposit", wireType)
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
			if err := m.MinimumDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
