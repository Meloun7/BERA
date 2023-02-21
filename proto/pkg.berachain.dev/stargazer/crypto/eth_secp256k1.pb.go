// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargazer/crypto/v1/eth_secp256k1.proto

package crypto

import (
	fmt "fmt"
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

// `EthSecp256k1PubKey` defines a type alias for an `ecdsa.PublicKey` that implements
// CometBFT's `PubKey` interface. It represents the 33-byte compressed public
// key format.
type EthSecp256K1PubKey struct {
	// `key` is the public key in byte form.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *EthSecp256K1PubKey) Reset()         { *m = EthSecp256K1PubKey{} }
func (m *EthSecp256K1PubKey) String() string { return proto.CompactTextString(m) }
func (*EthSecp256K1PubKey) ProtoMessage()    {}
func (*EthSecp256K1PubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7e1c79185bb538d, []int{0}
}
func (m *EthSecp256K1PubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthSecp256K1PubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthSecp256K1PubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthSecp256K1PubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthSecp256K1PubKey.Merge(m, src)
}
func (m *EthSecp256K1PubKey) XXX_Size() int {
	return m.Size()
}
func (m *EthSecp256K1PubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EthSecp256K1PubKey.DiscardUnknown(m)
}

var xxx_messageInfo_EthSecp256K1PubKey proto.InternalMessageInfo

func (m *EthSecp256K1PubKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

// `EthSecp256k1PrivKey` defines a type alias for a n`ecdsa.PrivateKey` that implements
// CometBFT's `PrivateKey` interface.
type EthSecp256K1PrivKey struct {
	// `key` is the private key in byte form.
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *EthSecp256K1PrivKey) Reset()         { *m = EthSecp256K1PrivKey{} }
func (m *EthSecp256K1PrivKey) String() string { return proto.CompactTextString(m) }
func (*EthSecp256K1PrivKey) ProtoMessage()    {}
func (*EthSecp256K1PrivKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7e1c79185bb538d, []int{1}
}
func (m *EthSecp256K1PrivKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthSecp256K1PrivKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthSecp256K1PrivKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthSecp256K1PrivKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthSecp256K1PrivKey.Merge(m, src)
}
func (m *EthSecp256K1PrivKey) XXX_Size() int {
	return m.Size()
}
func (m *EthSecp256K1PrivKey) XXX_DiscardUnknown() {
	xxx_messageInfo_EthSecp256K1PrivKey.DiscardUnknown(m)
}

var xxx_messageInfo_EthSecp256K1PrivKey proto.InternalMessageInfo

func (m *EthSecp256K1PrivKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*EthSecp256K1PubKey)(nil), "stargazer.crypto.v1.EthSecp256k1PubKey")
	proto.RegisterType((*EthSecp256K1PrivKey)(nil), "stargazer.crypto.v1.EthSecp256k1PrivKey")
}

func init() {
	proto.RegisterFile("stargazer/crypto/v1/eth_secp256k1.proto", fileDescriptor_c7e1c79185bb538d)
}

var fileDescriptor_c7e1c79185bb538d = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x2e, 0x49, 0x2c,
	0x4a, 0x4f, 0xac, 0x4a, 0x2d, 0xd2, 0x4f, 0x2e, 0xaa, 0x2c, 0x28, 0xc9, 0xd7, 0x2f, 0x33, 0xd4,
	0x4f, 0x2d, 0xc9, 0x88, 0x2f, 0x4e, 0x4d, 0x2e, 0x30, 0x32, 0x35, 0xcb, 0x36, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x86, 0x2b, 0xd4, 0x83, 0x28, 0xd4, 0x2b, 0x33, 0x54, 0x52, 0xe3,
	0x12, 0x72, 0x2d, 0xc9, 0x08, 0x86, 0x29, 0x0d, 0x28, 0x4d, 0xf2, 0x4e, 0xad, 0x14, 0x12, 0xe0,
	0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0x31, 0x95, 0xd4, 0xb9,
	0x84, 0x51, 0xd4, 0x15, 0x65, 0x96, 0x61, 0x55, 0xe8, 0x64, 0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3,
	0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x4a, 0x05, 0xd9, 0xe9, 0x7a, 0x49, 0xa9, 0x45, 0x89, 0xc9, 0x19,
	0x89, 0x99, 0x79, 0x7a, 0x29, 0xa9, 0x65, 0xfa, 0xe8, 0x4e, 0x4f, 0x62, 0x03, 0x3b, 0xd5, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xa1, 0x26, 0x45, 0xd5, 0x00, 0x00, 0x00,
}

func (m *EthSecp256K1PubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthSecp256K1PubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthSecp256K1PubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintEthSecp256K1(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EthSecp256K1PrivKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthSecp256K1PrivKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthSecp256K1PrivKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintEthSecp256K1(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEthSecp256K1(dAtA []byte, offset int, v uint64) int {
	offset -= sovEthSecp256K1(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthSecp256K1PubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEthSecp256K1(uint64(l))
	}
	return n
}

func (m *EthSecp256K1PrivKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovEthSecp256K1(uint64(l))
	}
	return n
}

func sovEthSecp256K1(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEthSecp256K1(x uint64) (n int) {
	return sovEthSecp256K1(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthSecp256K1PubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthSecp256K1
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
			return fmt.Errorf("proto: EthSecp256k1PubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthSecp256k1PubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthSecp256K1
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
				return ErrInvalidLengthEthSecp256K1
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEthSecp256K1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthSecp256K1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthSecp256K1
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
func (m *EthSecp256K1PrivKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthSecp256K1
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
			return fmt.Errorf("proto: EthSecp256k1PrivKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthSecp256k1PrivKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthSecp256K1
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
				return ErrInvalidLengthEthSecp256K1
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEthSecp256K1
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthSecp256K1(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthSecp256K1
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
func skipEthSecp256K1(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthSecp256K1
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
					return 0, ErrIntOverflowEthSecp256K1
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
					return 0, ErrIntOverflowEthSecp256K1
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
				return 0, ErrInvalidLengthEthSecp256K1
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEthSecp256K1
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEthSecp256K1
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEthSecp256K1        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthSecp256K1          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEthSecp256K1 = fmt.Errorf("proto: unexpected end of group")
)
