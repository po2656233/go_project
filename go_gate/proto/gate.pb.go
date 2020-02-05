// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gate.proto

package gate

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 设置用户地址(ip)req
type MsgSetUserAddrReq struct {
	Addr                 string   `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgSetUserAddrReq) Reset()         { *m = MsgSetUserAddrReq{} }
func (m *MsgSetUserAddrReq) String() string { return proto.CompactTextString(m) }
func (*MsgSetUserAddrReq) ProtoMessage()    {}
func (*MsgSetUserAddrReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}

func (m *MsgSetUserAddrReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSetUserAddrReq.Unmarshal(m, b)
}
func (m *MsgSetUserAddrReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSetUserAddrReq.Marshal(b, m, deterministic)
}
func (m *MsgSetUserAddrReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetUserAddrReq.Merge(m, src)
}
func (m *MsgSetUserAddrReq) XXX_Size() int {
	return xxx_messageInfo_MsgSetUserAddrReq.Size(m)
}
func (m *MsgSetUserAddrReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetUserAddrReq.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetUserAddrReq proto.InternalMessageInfo

func (m *MsgSetUserAddrReq) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgSetUserAddrReq)(nil), "gate.MsgSetUserAddrReq")
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor_743bb58a714d8b7d) }

var fileDescriptor_743bb58a714d8b7d = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4f, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0xd4, 0xb9, 0x04, 0x7d, 0x8b,
	0xd3, 0x83, 0x53, 0x4b, 0x42, 0x8b, 0x53, 0x8b, 0x1c, 0x53, 0x52, 0x8a, 0x82, 0x52, 0x0b, 0x85,
	0x84, 0xb8, 0x58, 0x40, 0x4c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x3b, 0x89, 0x0d,
	0xac, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x14, 0x4e, 0x97, 0x70, 0x43, 0x00, 0x00, 0x00,
}
