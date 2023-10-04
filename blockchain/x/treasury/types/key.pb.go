// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fusionchain/treasury/key.proto

package types

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

// KeyRequestStatus indicates the status of a key request.
// A request starts as "pending", waiting to be picked up. Then it can move to
// either "approved" or "rejected", depending on the decision of the MPC nodes.
type KeyRequestStatus int32

const (
	// The request is missing the status field.
	KeyRequestStatus_KEY_REQUEST_STATUS_UNSPECIFIED KeyRequestStatus = 0
	// The request is waiting to be fulfilled. This is the initial state of a
	// request.
	KeyRequestStatus_KEY_REQUEST_STATUS_PENDING KeyRequestStatus = 1
	// The request was fulfilled. This is a final state for a request.
	KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED KeyRequestStatus = 2
	// The request was rejected. This is a final state for a request.
	KeyRequestStatus_KEY_REQUEST_STATUS_REJECTED KeyRequestStatus = 3
)

var KeyRequestStatus_name = map[int32]string{
	0: "KEY_REQUEST_STATUS_UNSPECIFIED",
	1: "KEY_REQUEST_STATUS_PENDING",
	2: "KEY_REQUEST_STATUS_FULFILLED",
	3: "KEY_REQUEST_STATUS_REJECTED",
}

var KeyRequestStatus_value = map[string]int32{
	"KEY_REQUEST_STATUS_UNSPECIFIED": 0,
	"KEY_REQUEST_STATUS_PENDING":     1,
	"KEY_REQUEST_STATUS_FULFILLED":   2,
	"KEY_REQUEST_STATUS_REJECTED":    3,
}

func (x KeyRequestStatus) String() string {
	return proto.EnumName(KeyRequestStatus_name, int32(x))
}

func (KeyRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4c8664de1491f4e, []int{0}
}

// KeyType indicates what crypto scheme will be used by this key (e.g.
// ECDSA). Its public key will be one of the specified type.
type KeyType int32

const (
	// The key type is missing.
	KeyType_KEY_TYPE_UNSPECIFIED KeyType = 0
	// The key is an ECDSA secp256k1 key.
	KeyType_KEY_TYPE_ECDSA_SECP256K1 KeyType = 1
	// The key is an EdDSA Ed25519 key.
	KeyType_KEY_TYPE_EDDSA_ED25519 KeyType = 2
)

var KeyType_name = map[int32]string{
	0: "KEY_TYPE_UNSPECIFIED",
	1: "KEY_TYPE_ECDSA_SECP256K1",
	2: "KEY_TYPE_EDDSA_ED25519",
}

var KeyType_value = map[string]int32{
	"KEY_TYPE_UNSPECIFIED":     0,
	"KEY_TYPE_ECDSA_SECP256K1": 1,
	"KEY_TYPE_EDDSA_ED25519":   2,
}

func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}

func (KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4c8664de1491f4e, []int{1}
}

type KeyRequest struct {
	Id            uint64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator       string           `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	WorkspaceAddr string           `protobuf:"bytes,3,opt,name=workspace_addr,json=workspaceAddr,proto3" json:"workspace_addr,omitempty"`
	KeyringId     uint64           `protobuf:"varint,4,opt,name=keyring_id,json=keyringId,proto3" json:"keyring_id,omitempty"`
	KeyType       KeyType          `protobuf:"varint,5,opt,name=key_type,json=keyType,proto3,enum=fusionchain.treasury.KeyType" json:"key_type,omitempty"`
	Status        KeyRequestStatus `protobuf:"varint,6,opt,name=status,proto3,enum=fusionchain.treasury.KeyRequestStatus" json:"status,omitempty"`
	RejectReason  string           `protobuf:"bytes,7,opt,name=reject_reason,json=rejectReason,proto3" json:"reject_reason,omitempty"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4c8664de1491f4e, []int{0}
}
func (m *KeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequest.Merge(m, src)
}
func (m *KeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *KeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequest proto.InternalMessageInfo

func (m *KeyRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KeyRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *KeyRequest) GetWorkspaceAddr() string {
	if m != nil {
		return m.WorkspaceAddr
	}
	return ""
}

func (m *KeyRequest) GetKeyringId() uint64 {
	if m != nil {
		return m.KeyringId
	}
	return 0
}

func (m *KeyRequest) GetKeyType() KeyType {
	if m != nil {
		return m.KeyType
	}
	return KeyType_KEY_TYPE_UNSPECIFIED
}

func (m *KeyRequest) GetStatus() KeyRequestStatus {
	if m != nil {
		return m.Status
	}
	return KeyRequestStatus_KEY_REQUEST_STATUS_UNSPECIFIED
}

func (m *KeyRequest) GetRejectReason() string {
	if m != nil {
		return m.RejectReason
	}
	return ""
}

type Key struct {
	Id            uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	WorkspaceAddr string  `protobuf:"bytes,2,opt,name=workspace_addr,json=workspaceAddr,proto3" json:"workspace_addr,omitempty"`
	KeyringId     uint64  `protobuf:"varint,3,opt,name=keyring_id,json=keyringId,proto3" json:"keyring_id,omitempty"`
	Type          KeyType `protobuf:"varint,4,opt,name=type,proto3,enum=fusionchain.treasury.KeyType" json:"type,omitempty"`
	PublicKey     []byte  `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4c8664de1491f4e, []int{1}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Key) GetWorkspaceAddr() string {
	if m != nil {
		return m.WorkspaceAddr
	}
	return ""
}

func (m *Key) GetKeyringId() uint64 {
	if m != nil {
		return m.KeyringId
	}
	return 0
}

func (m *Key) GetType() KeyType {
	if m != nil {
		return m.Type
	}
	return KeyType_KEY_TYPE_UNSPECIFIED
}

func (m *Key) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func init() {
	proto.RegisterEnum("fusionchain.treasury.KeyRequestStatus", KeyRequestStatus_name, KeyRequestStatus_value)
	proto.RegisterEnum("fusionchain.treasury.KeyType", KeyType_name, KeyType_value)
	proto.RegisterType((*KeyRequest)(nil), "fusionchain.treasury.KeyRequest")
	proto.RegisterType((*Key)(nil), "fusionchain.treasury.Key")
}

func init() { proto.RegisterFile("fusionchain/treasury/key.proto", fileDescriptor_c4c8664de1491f4e) }

var fileDescriptor_c4c8664de1491f4e = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0x69, 0x6d, 0xed, 0x61, 0xb7, 0x84, 0x61, 0x91, 0x61, 0xdd, 0x8d, 0xa5, 0xa2,
	0x94, 0x05, 0x5b, 0x5a, 0xa9, 0xe8, 0x8d, 0x50, 0x9b, 0xe9, 0x52, 0x53, 0x4a, 0x4d, 0xd2, 0x8b,
	0x15, 0x24, 0xa4, 0xc9, 0xb8, 0x1b, 0xa3, 0x4d, 0x77, 0x92, 0xa0, 0xf3, 0x16, 0xde, 0x8a, 0xcf,
	0xe0, 0x7b, 0x78, 0xb9, 0x97, 0x5e, 0x4a, 0xfb, 0x22, 0x92, 0x69, 0xf7, 0x0f, 0xdd, 0x8a, 0xde,
	0xcd, 0x7c, 0xdf, 0x77, 0x0e, 0xe7, 0x77, 0xe0, 0x80, 0xf6, 0x3e, 0x8d, 0x83, 0x68, 0xe6, 0x9d,
	0xb9, 0xc1, 0xac, 0x99, 0x70, 0xe6, 0xc6, 0x29, 0x17, 0xcd, 0x90, 0x89, 0xc6, 0x9c, 0x47, 0x49,
	0x84, 0xf7, 0x6e, 0xf8, 0x8d, 0x4b, 0xbf, 0xf6, 0x5d, 0x01, 0x30, 0x98, 0x30, 0xd9, 0x79, 0xca,
	0xe2, 0x04, 0x57, 0x40, 0x09, 0x7c, 0x82, 0xaa, 0xa8, 0x5e, 0x30, 0x95, 0xc0, 0xc7, 0x04, 0x4a,
	0x1e, 0x67, 0x6e, 0x12, 0x71, 0xa2, 0x54, 0x51, 0xbd, 0x6c, 0x5e, 0x7e, 0xf1, 0x23, 0xa8, 0x7c,
	0x8e, 0x78, 0x18, 0xcf, 0x5d, 0x8f, 0x39, 0xae, 0xef, 0x73, 0x92, 0x97, 0x81, 0xdd, 0x2b, 0xb5,
	0xeb, 0xfb, 0x1c, 0x1f, 0x02, 0x84, 0x4c, 0xf0, 0x60, 0x76, 0xea, 0x04, 0x3e, 0x29, 0xc8, 0xc6,
	0xe5, 0xb5, 0x32, 0xf0, 0xf1, 0x73, 0xb8, 0x1b, 0x32, 0xe1, 0x24, 0x62, 0xce, 0xc8, 0x9d, 0x2a,
	0xaa, 0x57, 0xda, 0x87, 0x8d, 0x6d, 0x73, 0x36, 0x0c, 0x26, 0x6c, 0x31, 0x67, 0x66, 0x29, 0x5c,
	0x3d, 0xf0, 0x4b, 0x28, 0xc6, 0x89, 0x9b, 0xa4, 0x31, 0x29, 0xca, 0xba, 0xc7, 0x7f, 0xad, 0x5b,
	0xb3, 0x59, 0x32, 0x6d, 0xae, 0xab, 0xf0, 0x43, 0xd8, 0xe5, 0xec, 0x03, 0xf3, 0x12, 0x27, 0x8b,
	0x46, 0x33, 0x52, 0x92, 0xe3, 0xef, 0xac, 0x44, 0x53, 0x6a, 0xb5, 0x1f, 0x08, 0xf2, 0x06, 0x13,
	0xb7, 0xd6, 0x72, 0x1b, 0x5e, 0xf9, 0x37, 0x7c, 0x7e, 0x13, 0xbe, 0x05, 0x05, 0x09, 0x5e, 0xf8,
	0x1f, 0x70, 0x19, 0xcd, 0x3a, 0xce, 0xd3, 0xe9, 0xc7, 0xc0, 0x73, 0x42, 0x26, 0xe4, 0xc6, 0x76,
	0xcc, 0xf2, 0x4a, 0x31, 0x98, 0x38, 0xfa, 0x86, 0x40, 0xdd, 0x24, 0xc6, 0x35, 0xd0, 0x0c, 0x7a,
	0xe2, 0x98, 0xf4, 0xcd, 0x84, 0x5a, 0xb6, 0x63, 0xd9, 0x5d, 0x7b, 0x62, 0x39, 0x93, 0x91, 0x35,
	0xa6, 0xbd, 0x41, 0x7f, 0x40, 0x75, 0x35, 0x87, 0x35, 0xd8, 0xdf, 0x92, 0x19, 0xd3, 0x91, 0x3e,
	0x18, 0x1d, 0xab, 0x08, 0x57, 0xe1, 0x60, 0x8b, 0xdf, 0x9f, 0x0c, 0xfb, 0x83, 0xe1, 0x90, 0xea,
	0xaa, 0x82, 0x1f, 0xc0, 0xfd, 0x2d, 0x09, 0x93, 0xbe, 0xa6, 0x3d, 0x9b, 0xea, 0x6a, 0xfe, 0xe8,
	0x1d, 0x94, 0xd6, 0x2c, 0x98, 0xc0, 0x5e, 0x96, 0xb5, 0x4f, 0xc6, 0x74, 0x63, 0x8e, 0x03, 0x20,
	0x57, 0x0e, 0xed, 0xe9, 0x56, 0xd7, 0xb1, 0x68, 0x6f, 0xdc, 0xee, 0x3c, 0x33, 0x5a, 0x2a, 0xc2,
	0xfb, 0x70, 0xef, 0xda, 0xd5, 0x33, 0x97, 0xea, 0xed, 0x4e, 0xa7, 0xf5, 0x42, 0x55, 0x5e, 0x1d,
	0xff, 0x5c, 0x68, 0xe8, 0x62, 0xa1, 0xa1, 0xdf, 0x0b, 0x0d, 0x7d, 0x5d, 0x6a, 0xb9, 0x8b, 0xa5,
	0x96, 0xfb, 0xb5, 0xd4, 0x72, 0x6f, 0x9f, 0x9c, 0x06, 0xc9, 0x59, 0x3a, 0x6d, 0x78, 0xd1, 0xa7,
	0xe6, 0x39, 0x67, 0x7e, 0xd4, 0xbc, 0x79, 0x29, 0x5f, 0xae, 0x6f, 0x25, 0xdb, 0x70, 0x3c, 0x2d,
	0xca, 0x73, 0x79, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xfe, 0x69, 0xa5, 0x50, 0x03, 0x00,
	0x00,
}

func (m *KeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RejectReason) > 0 {
		i -= len(m.RejectReason)
		copy(dAtA[i:], m.RejectReason)
		i = encodeVarintKey(dAtA, i, uint64(len(m.RejectReason)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if m.KeyType != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.KeyType))
		i--
		dAtA[i] = 0x28
	}
	if m.KeyringId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.KeyringId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.WorkspaceAddr) > 0 {
		i -= len(m.WorkspaceAddr)
		copy(dAtA[i:], m.WorkspaceAddr)
		i = encodeVarintKey(dAtA, i, uint64(len(m.WorkspaceAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintKey(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintKey(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Type != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if m.KeyringId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.KeyringId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.WorkspaceAddr) > 0 {
		i -= len(m.WorkspaceAddr)
		copy(dAtA[i:], m.WorkspaceAddr)
		i = encodeVarintKey(dAtA, i, uint64(len(m.WorkspaceAddr)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovKey(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	l = len(m.WorkspaceAddr)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	if m.KeyringId != 0 {
		n += 1 + sovKey(uint64(m.KeyringId))
	}
	if m.KeyType != 0 {
		n += 1 + sovKey(uint64(m.KeyType))
	}
	if m.Status != 0 {
		n += 1 + sovKey(uint64(m.Status))
	}
	l = len(m.RejectReason)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	return n
}

func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovKey(uint64(m.Id))
	}
	l = len(m.WorkspaceAddr)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	if m.KeyringId != 0 {
		n += 1 + sovKey(uint64(m.KeyringId))
	}
	if m.Type != 0 {
		n += 1 + sovKey(uint64(m.Type))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	return n
}

func sovKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKey(x uint64) (n int) {
	return sovKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKey
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
			return fmt.Errorf("proto: KeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkspaceAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkspaceAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyringId", wireType)
			}
			m.KeyringId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyringId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			m.KeyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyType |= KeyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= KeyRequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RejectReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKey
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
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKey
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkspaceAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkspaceAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyringId", wireType)
			}
			m.KeyringId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyringId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= KeyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKey
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
func skipKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
				return 0, ErrInvalidLengthKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKey = fmt.Errorf("proto: unexpected end of group")
)
