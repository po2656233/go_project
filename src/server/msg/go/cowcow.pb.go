// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cowcow.proto

package _go

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

//【百人类游戏】
//入场 (场景)
type GameCowcowEnter struct {
	UserInfo             *PlayerInfo `protobuf:"bytes,1,opt,name=UserInfo,proto3" json:"UserInfo,omitempty"`
	TimeStamp            int64       `protobuf:"varint,2,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
	Chips                []int32     `protobuf:"varint,3,rep,packed,name=Chips,proto3" json:"Chips,omitempty"`
	AwardAreas           [][]byte    `protobuf:"bytes,4,rep,name=AwardAreas,proto3" json:"AwardAreas,omitempty"`
	FreeTime             uint32      `protobuf:"varint,5,opt,name=FreeTime,proto3" json:"FreeTime,omitempty"`
	BetTime              uint32      `protobuf:"varint,6,opt,name=BetTime,proto3" json:"BetTime,omitempty"`
	OpenTime             uint32      `protobuf:"varint,7,opt,name=OpenTime,proto3" json:"OpenTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GameCowcowEnter) Reset()         { *m = GameCowcowEnter{} }
func (m *GameCowcowEnter) String() string { return proto.CompactTextString(m) }
func (*GameCowcowEnter) ProtoMessage()    {}
func (*GameCowcowEnter) Descriptor() ([]byte, []int) {
	return fileDescriptor_626b438dd4edf163, []int{0}
}

func (m *GameCowcowEnter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameCowcowEnter.Unmarshal(m, b)
}
func (m *GameCowcowEnter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameCowcowEnter.Marshal(b, m, deterministic)
}
func (m *GameCowcowEnter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameCowcowEnter.Merge(m, src)
}
func (m *GameCowcowEnter) XXX_Size() int {
	return xxx_messageInfo_GameCowcowEnter.Size(m)
}
func (m *GameCowcowEnter) XXX_DiscardUnknown() {
	xxx_messageInfo_GameCowcowEnter.DiscardUnknown(m)
}

var xxx_messageInfo_GameCowcowEnter proto.InternalMessageInfo

func (m *GameCowcowEnter) GetUserInfo() *PlayerInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *GameCowcowEnter) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func (m *GameCowcowEnter) GetChips() []int32 {
	if m != nil {
		return m.Chips
	}
	return nil
}

func (m *GameCowcowEnter) GetAwardAreas() [][]byte {
	if m != nil {
		return m.AwardAreas
	}
	return nil
}

func (m *GameCowcowEnter) GetFreeTime() uint32 {
	if m != nil {
		return m.FreeTime
	}
	return 0
}

func (m *GameCowcowEnter) GetBetTime() uint32 {
	if m != nil {
		return m.BetTime
	}
	return 0
}

func (m *GameCowcowEnter) GetOpenTime() uint32 {
	if m != nil {
		return m.OpenTime
	}
	return 0
}

//游戏消息
//抢庄
type GameCowcowHost struct {
	UserID               uint64   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	IsWant               bool     `protobuf:"varint,2,opt,name=IsWant,proto3" json:"IsWant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameCowcowHost) Reset()         { *m = GameCowcowHost{} }
func (m *GameCowcowHost) String() string { return proto.CompactTextString(m) }
func (*GameCowcowHost) ProtoMessage()    {}
func (*GameCowcowHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_626b438dd4edf163, []int{1}
}

func (m *GameCowcowHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameCowcowHost.Unmarshal(m, b)
}
func (m *GameCowcowHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameCowcowHost.Marshal(b, m, deterministic)
}
func (m *GameCowcowHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameCowcowHost.Merge(m, src)
}
func (m *GameCowcowHost) XXX_Size() int {
	return xxx_messageInfo_GameCowcowHost.Size(m)
}
func (m *GameCowcowHost) XXX_DiscardUnknown() {
	xxx_messageInfo_GameCowcowHost.DiscardUnknown(m)
}

var xxx_messageInfo_GameCowcowHost proto.InternalMessageInfo

func (m *GameCowcowHost) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GameCowcowHost) GetIsWant() bool {
	if m != nil {
		return m.IsWant
	}
	return false
}

//超级抢庄
type GameCowcowSuperHost struct {
	UserID               uint64   `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	IsWant               bool     `protobuf:"varint,2,opt,name=IsWant,proto3" json:"IsWant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameCowcowSuperHost) Reset()         { *m = GameCowcowSuperHost{} }
func (m *GameCowcowSuperHost) String() string { return proto.CompactTextString(m) }
func (*GameCowcowSuperHost) ProtoMessage()    {}
func (*GameCowcowSuperHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_626b438dd4edf163, []int{2}
}

func (m *GameCowcowSuperHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameCowcowSuperHost.Unmarshal(m, b)
}
func (m *GameCowcowSuperHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameCowcowSuperHost.Marshal(b, m, deterministic)
}
func (m *GameCowcowSuperHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameCowcowSuperHost.Merge(m, src)
}
func (m *GameCowcowSuperHost) XXX_Size() int {
	return xxx_messageInfo_GameCowcowSuperHost.Size(m)
}
func (m *GameCowcowSuperHost) XXX_DiscardUnknown() {
	xxx_messageInfo_GameCowcowSuperHost.DiscardUnknown(m)
}

var xxx_messageInfo_GameCowcowSuperHost proto.InternalMessageInfo

func (m *GameCowcowSuperHost) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GameCowcowSuperHost) GetIsWant() bool {
	if m != nil {
		return m.IsWant
	}
	return false
}

//游戏中
type GameCowcowPlaying struct {
	BetArea              int32    `protobuf:"varint,1,opt,name=BetArea,proto3" json:"BetArea,omitempty"`
	BetScore             int64    `protobuf:"varint,2,opt,name=BetScore,proto3" json:"BetScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameCowcowPlaying) Reset()         { *m = GameCowcowPlaying{} }
func (m *GameCowcowPlaying) String() string { return proto.CompactTextString(m) }
func (*GameCowcowPlaying) ProtoMessage()    {}
func (*GameCowcowPlaying) Descriptor() ([]byte, []int) {
	return fileDescriptor_626b438dd4edf163, []int{3}
}

func (m *GameCowcowPlaying) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameCowcowPlaying.Unmarshal(m, b)
}
func (m *GameCowcowPlaying) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameCowcowPlaying.Marshal(b, m, deterministic)
}
func (m *GameCowcowPlaying) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameCowcowPlaying.Merge(m, src)
}
func (m *GameCowcowPlaying) XXX_Size() int {
	return xxx_messageInfo_GameCowcowPlaying.Size(m)
}
func (m *GameCowcowPlaying) XXX_DiscardUnknown() {
	xxx_messageInfo_GameCowcowPlaying.DiscardUnknown(m)
}

var xxx_messageInfo_GameCowcowPlaying proto.InternalMessageInfo

func (m *GameCowcowPlaying) GetBetArea() int32 {
	if m != nil {
		return m.BetArea
	}
	return 0
}

func (m *GameCowcowPlaying) GetBetScore() int64 {
	if m != nil {
		return m.BetScore
	}
	return 0
}

//下注结果
type GameCowcowBetResult struct {
	State                int32    `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
	Hints                string   `protobuf:"bytes,2,opt,name=Hints,proto3" json:"Hints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameCowcowBetResult) Reset()         { *m = GameCowcowBetResult{} }
func (m *GameCowcowBetResult) String() string { return proto.CompactTextString(m) }
func (*GameCowcowBetResult) ProtoMessage()    {}
func (*GameCowcowBetResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_626b438dd4edf163, []int{4}
}

func (m *GameCowcowBetResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameCowcowBetResult.Unmarshal(m, b)
}
func (m *GameCowcowBetResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameCowcowBetResult.Marshal(b, m, deterministic)
}
func (m *GameCowcowBetResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameCowcowBetResult.Merge(m, src)
}
func (m *GameCowcowBetResult) XXX_Size() int {
	return xxx_messageInfo_GameCowcowBetResult.Size(m)
}
func (m *GameCowcowBetResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GameCowcowBetResult.DiscardUnknown(m)
}

var xxx_messageInfo_GameCowcowBetResult proto.InternalMessageInfo

func (m *GameCowcowBetResult) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *GameCowcowBetResult) GetHints() string {
	if m != nil {
		return m.Hints
	}
	return ""
}

//结束
type GameCowcowOver struct {
	AwardArea            []byte   `protobuf:"bytes,1,opt,name=AwardArea,proto3" json:"AwardArea,omitempty"`
	PlayerCard           []byte   `protobuf:"bytes,2,opt,name=PlayerCard,proto3" json:"PlayerCard,omitempty"`
	BankerCard           []byte   `protobuf:"bytes,3,opt,name=BankerCard,proto3" json:"BankerCard,omitempty"`
	Acquire              int64    `protobuf:"varint,4,opt,name=Acquire,proto3" json:"Acquire,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameCowcowOver) Reset()         { *m = GameCowcowOver{} }
func (m *GameCowcowOver) String() string { return proto.CompactTextString(m) }
func (*GameCowcowOver) ProtoMessage()    {}
func (*GameCowcowOver) Descriptor() ([]byte, []int) {
	return fileDescriptor_626b438dd4edf163, []int{5}
}

func (m *GameCowcowOver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameCowcowOver.Unmarshal(m, b)
}
func (m *GameCowcowOver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameCowcowOver.Marshal(b, m, deterministic)
}
func (m *GameCowcowOver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameCowcowOver.Merge(m, src)
}
func (m *GameCowcowOver) XXX_Size() int {
	return xxx_messageInfo_GameCowcowOver.Size(m)
}
func (m *GameCowcowOver) XXX_DiscardUnknown() {
	xxx_messageInfo_GameCowcowOver.DiscardUnknown(m)
}

var xxx_messageInfo_GameCowcowOver proto.InternalMessageInfo

func (m *GameCowcowOver) GetAwardArea() []byte {
	if m != nil {
		return m.AwardArea
	}
	return nil
}

func (m *GameCowcowOver) GetPlayerCard() []byte {
	if m != nil {
		return m.PlayerCard
	}
	return nil
}

func (m *GameCowcowOver) GetBankerCard() []byte {
	if m != nil {
		return m.BankerCard
	}
	return nil
}

func (m *GameCowcowOver) GetAcquire() int64 {
	if m != nil {
		return m.Acquire
	}
	return 0
}

func init() {
	proto.RegisterType((*GameCowcowEnter)(nil), "go.GameCowcowEnter")
	proto.RegisterType((*GameCowcowHost)(nil), "go.GameCowcowHost")
	proto.RegisterType((*GameCowcowSuperHost)(nil), "go.GameCowcowSuperHost")
	proto.RegisterType((*GameCowcowPlaying)(nil), "go.GameCowcowPlaying")
	proto.RegisterType((*GameCowcowBetResult)(nil), "go.GameCowcowBetResult")
	proto.RegisterType((*GameCowcowOver)(nil), "go.GameCowcowOver")
}

func init() { proto.RegisterFile("cowcow.proto", fileDescriptor_626b438dd4edf163) }

var fileDescriptor_626b438dd4edf163 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x4e, 0xe3, 0x30,
	0x14, 0x85, 0x95, 0xa6, 0xe9, 0x8f, 0x27, 0xd3, 0xd1, 0x78, 0x46, 0x28, 0xaa, 0x10, 0x8a, 0xb2,
	0x8a, 0x58, 0x74, 0x01, 0x2f, 0x40, 0x52, 0x0a, 0xed, 0xaa, 0xc8, 0x01, 0xb1, 0x36, 0xe9, 0x25,
	0x44, 0x34, 0x71, 0xb0, 0x5d, 0x2a, 0xde, 0x80, 0xf7, 0xe4, 0x45, 0x90, 0xed, 0xfc, 0xad, 0x59,
	0x7e, 0xe7, 0xc4, 0x47, 0x39, 0xf7, 0x5e, 0xe4, 0xa6, 0xec, 0x98, 0xb2, 0xe3, 0xa2, 0xe2, 0x4c,
	0x32, 0x3c, 0xc8, 0xd8, 0x7c, 0x96, 0xd1, 0x02, 0x52, 0x56, 0x14, 0x46, 0x0b, 0xbe, 0x2c, 0xf4,
	0xe7, 0x96, 0x16, 0xb0, 0xd4, 0x1f, 0xae, 0x4a, 0x09, 0x1c, 0x9f, 0xa3, 0xc9, 0x83, 0x00, 0xbe,
	0x29, 0x9f, 0x99, 0x67, 0xf9, 0x56, 0xf8, 0xeb, 0x62, 0xb6, 0xc8, 0xd8, 0xe2, 0x6e, 0x4f, 0x3f,
	0x8c, 0x4a, 0x5a, 0x1f, 0x9f, 0xa2, 0xe9, 0x7d, 0x5e, 0x40, 0x22, 0x69, 0x51, 0x79, 0x03, 0xdf,
	0x0a, 0x6d, 0xd2, 0x09, 0xf8, 0x3f, 0x72, 0x96, 0x2f, 0x79, 0x25, 0x3c, 0xdb, 0xb7, 0x43, 0x87,
	0x18, 0xc0, 0x67, 0x08, 0x45, 0x47, 0xca, 0x77, 0x11, 0x07, 0x2a, 0xbc, 0xa1, 0x6f, 0x87, 0x2e,
	0xe9, 0x29, 0x78, 0x8e, 0x26, 0x37, 0x1c, 0x40, 0xc5, 0x78, 0x8e, 0x6f, 0x85, 0xbf, 0x49, 0xcb,
	0xd8, 0x43, 0xe3, 0x18, 0xa4, 0xb6, 0x46, 0xda, 0x6a, 0x50, 0xbd, 0xda, 0x56, 0x50, 0x6a, 0x6b,
	0x6c, 0x5e, 0x35, 0x1c, 0x5c, 0xa1, 0x59, 0x57, 0x72, 0xcd, 0x84, 0xc4, 0x27, 0x68, 0xa4, 0x3b,
	0x5c, 0xeb, 0x86, 0x43, 0x52, 0x93, 0xd2, 0x37, 0xe2, 0x91, 0x96, 0x52, 0x97, 0x99, 0x90, 0x9a,
	0x82, 0x15, 0xfa, 0xd7, 0x25, 0x24, 0x87, 0x0a, 0xf8, 0x8f, 0x62, 0x36, 0xe8, 0x6f, 0x17, 0xa3,
	0x06, 0x9a, 0x97, 0x59, 0xdd, 0x49, 0x75, 0xd7, 0x29, 0x0e, 0x69, 0x50, 0x75, 0x8a, 0x41, 0x26,
	0x29, 0xe3, 0x50, 0x0f, 0xb7, 0xe5, 0x20, 0xea, 0xff, 0x51, 0x0c, 0x92, 0x80, 0x38, 0xec, 0xa5,
	0x1a, 0x79, 0x22, 0xa9, 0x84, 0x3a, 0xca, 0x80, 0x52, 0xd7, 0x79, 0x29, 0x85, 0x4e, 0x99, 0x12,
	0x03, 0xc1, 0xa7, 0xd5, 0x9f, 0xcb, 0xf6, 0x1d, 0xb8, 0xda, 0x67, 0xbb, 0x09, 0x1d, 0xe1, 0x92,
	0x4e, 0x50, 0x9b, 0x33, 0x57, 0xb0, 0xa4, 0x7c, 0xa7, 0xb3, 0x5c, 0xd2, 0x53, 0x94, 0x1f, 0xd3,
	0xf2, 0xb5, 0xf6, 0x6d, 0xe3, 0x77, 0x8a, 0x6a, 0x1a, 0xa5, 0x6f, 0x87, 0x9c, 0x83, 0x37, 0xd4,
	0x75, 0x1a, 0x7c, 0x1a, 0xe9, 0x73, 0xbc, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x40, 0xf0, 0x2d,
	0x5d, 0xb2, 0x02, 0x00, 0x00,
}
