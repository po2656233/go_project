// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.7.0
// source: brcowcow.proto

package _go

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// kindID 2003
//入场
type BrcowcowSceneResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeStamp  int64           `protobuf:"varint,1,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`      //时间戳
	Chips      []int32         `protobuf:"varint,2,rep,packed,name=Chips,proto3" json:"Chips,omitempty"`       //筹码
	AwardAreas [][]byte        `protobuf:"bytes,3,rep,name=AwardAreas,proto3" json:"AwardAreas,omitempty"`     //开奖记录(路单)
	AreaBets   []int64         `protobuf:"varint,4,rep,packed,name=AreaBets,proto3" json:"AreaBets,omitempty"` //各下注区当前总下注额
	MyBets     []int64         `protobuf:"varint,5,rep,packed,name=MyBets,proto3" json:"MyBets,omitempty"`     //我在各下注区的总下注额
	Inning     string          `protobuf:"bytes,6,opt,name=Inning,proto3" json:"Inning,omitempty"`             // 牌局号
	AllPlayers *PlayerListInfo `protobuf:"bytes,7,opt,name=AllPlayers,proto3" json:"AllPlayers,omitempty"`     //玩家列表
	HostID     uint64          `protobuf:"varint,8,opt,name=HostID,proto3" json:"HostID,omitempty"`
}

func (x *BrcowcowSceneResp) Reset() {
	*x = BrcowcowSceneResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowSceneResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowSceneResp) ProtoMessage() {}

func (x *BrcowcowSceneResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowSceneResp.ProtoReflect.Descriptor instead.
func (*BrcowcowSceneResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{0}
}

func (x *BrcowcowSceneResp) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

func (x *BrcowcowSceneResp) GetChips() []int32 {
	if x != nil {
		return x.Chips
	}
	return nil
}

func (x *BrcowcowSceneResp) GetAwardAreas() [][]byte {
	if x != nil {
		return x.AwardAreas
	}
	return nil
}

func (x *BrcowcowSceneResp) GetAreaBets() []int64 {
	if x != nil {
		return x.AreaBets
	}
	return nil
}

func (x *BrcowcowSceneResp) GetMyBets() []int64 {
	if x != nil {
		return x.MyBets
	}
	return nil
}

func (x *BrcowcowSceneResp) GetInning() string {
	if x != nil {
		return x.Inning
	}
	return ""
}

func (x *BrcowcowSceneResp) GetAllPlayers() *PlayerListInfo {
	if x != nil {
		return x.AllPlayers
	}
	return nil
}

func (x *BrcowcowSceneResp) GetHostID() uint64 {
	if x != nil {
		return x.HostID
	}
	return 0
}

//状态
// 服务端推送
//(空闲 - 喊庄)
type BrcowcowStateFreeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times *TimeInfo `protobuf:"bytes,1,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (x *BrcowcowStateFreeResp) Reset() {
	*x = BrcowcowStateFreeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowStateFreeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowStateFreeResp) ProtoMessage() {}

func (x *BrcowcowStateFreeResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowStateFreeResp.ProtoReflect.Descriptor instead.
func (*BrcowcowStateFreeResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{1}
}

func (x *BrcowcowStateFreeResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

//(开始 - 定庄)
type BrcowcowStateStartResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times  *TimeInfo `protobuf:"bytes,1,opt,name=Times,proto3" json:"Times,omitempty"`
	HostID uint64    `protobuf:"varint,2,opt,name=HostID,proto3" json:"HostID,omitempty"`
	Inning string    `protobuf:"bytes,3,opt,name=Inning,proto3" json:"Inning,omitempty"` // 牌局号
}

func (x *BrcowcowStateStartResp) Reset() {
	*x = BrcowcowStateStartResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowStateStartResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowStateStartResp) ProtoMessage() {}

func (x *BrcowcowStateStartResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowStateStartResp.ProtoReflect.Descriptor instead.
func (*BrcowcowStateStartResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{2}
}

func (x *BrcowcowStateStartResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

func (x *BrcowcowStateStartResp) GetHostID() uint64 {
	if x != nil {
		return x.HostID
	}
	return 0
}

func (x *BrcowcowStateStartResp) GetInning() string {
	if x != nil {
		return x.Inning
	}
	return ""
}

//(游戏中 - 下注)
type BrcowcowStatePlayingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times *TimeInfo `protobuf:"bytes,1,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (x *BrcowcowStatePlayingResp) Reset() {
	*x = BrcowcowStatePlayingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowStatePlayingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowStatePlayingResp) ProtoMessage() {}

func (x *BrcowcowStatePlayingResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowStatePlayingResp.ProtoReflect.Descriptor instead.
func (*BrcowcowStatePlayingResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{3}
}

func (x *BrcowcowStatePlayingResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

//(开奖)
type BrcowcowStateOpenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times    *TimeInfo         `protobuf:"bytes,1,opt,name=Times,proto3" json:"Times,omitempty"`
	OpenInfo *BrcowcowOpenResp `protobuf:"bytes,2,opt,name=OpenInfo,proto3" json:"OpenInfo,omitempty"`
}

func (x *BrcowcowStateOpenResp) Reset() {
	*x = BrcowcowStateOpenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowStateOpenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowStateOpenResp) ProtoMessage() {}

func (x *BrcowcowStateOpenResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowStateOpenResp.ProtoReflect.Descriptor instead.
func (*BrcowcowStateOpenResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{4}
}

func (x *BrcowcowStateOpenResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

func (x *BrcowcowStateOpenResp) GetOpenInfo() *BrcowcowOpenResp {
	if x != nil {
		return x.OpenInfo
	}
	return nil
}

//(结束)
type BrcowcowStateOverResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Times *TimeInfo `protobuf:"bytes,1,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (x *BrcowcowStateOverResp) Reset() {
	*x = BrcowcowStateOverResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowStateOverResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowStateOverResp) ProtoMessage() {}

func (x *BrcowcowStateOverResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowStateOverResp.ProtoReflect.Descriptor instead.
func (*BrcowcowStateOverResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{5}
}

func (x *BrcowcowStateOverResp) GetTimes() *TimeInfo {
	if x != nil {
		return x.Times
	}
	return nil
}

//下注
type BrcowcowBetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BetArea  int32 `protobuf:"varint,1,opt,name=BetArea,proto3" json:"BetArea,omitempty"`   //下注区域
	BetScore int64 `protobuf:"varint,2,opt,name=BetScore,proto3" json:"BetScore,omitempty"` //下注金额
}

func (x *BrcowcowBetReq) Reset() {
	*x = BrcowcowBetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowBetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowBetReq) ProtoMessage() {}

func (x *BrcowcowBetReq) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowBetReq.ProtoReflect.Descriptor instead.
func (*BrcowcowBetReq) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{6}
}

func (x *BrcowcowBetReq) GetBetArea() int32 {
	if x != nil {
		return x.BetArea
	}
	return 0
}

func (x *BrcowcowBetReq) GetBetScore() int64 {
	if x != nil {
		return x.BetScore
	}
	return 0
}

//下注结果：广播
type BrcowcowBetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	BetArea  int32  `protobuf:"varint,2,opt,name=BetArea,proto3" json:"BetArea,omitempty"`   //下注区域
	BetScore int64  `protobuf:"varint,3,opt,name=BetScore,proto3" json:"BetScore,omitempty"` //下注金额
}

func (x *BrcowcowBetResp) Reset() {
	*x = BrcowcowBetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowBetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowBetResp) ProtoMessage() {}

func (x *BrcowcowBetResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowBetResp.ProtoReflect.Descriptor instead.
func (*BrcowcowBetResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{7}
}

func (x *BrcowcowBetResp) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BrcowcowBetResp) GetBetArea() int32 {
	if x != nil {
		return x.BetArea
	}
	return 0
}

func (x *BrcowcowBetResp) GetBetScore() int64 {
	if x != nil {
		return x.BetScore
	}
	return 0
}

//开牌
type BrcowcowOpenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AwardArea  []byte    `protobuf:"bytes,1,opt,name=AwardArea,proto3" json:"AwardArea,omitempty"`   //各区域胜负：1胜，0负
	BankerCard *CardInfo `protobuf:"bytes,2,opt,name=BankerCard,proto3" json:"BankerCard,omitempty"` //天
	TianCard   *CardInfo `protobuf:"bytes,3,opt,name=TianCard,proto3" json:"TianCard,omitempty"`     //天
	XuanCard   *CardInfo `protobuf:"bytes,4,opt,name=XuanCard,proto3" json:"XuanCard,omitempty"`     //玄
	DiCard     *CardInfo `protobuf:"bytes,5,opt,name=DiCard,proto3" json:"DiCard,omitempty"`         //地
	HuangCard  *CardInfo `protobuf:"bytes,6,opt,name=HuangCard,proto3" json:"HuangCard,omitempty"`   //黄
}

func (x *BrcowcowOpenResp) Reset() {
	*x = BrcowcowOpenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowOpenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowOpenResp) ProtoMessage() {}

func (x *BrcowcowOpenResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowOpenResp.ProtoReflect.Descriptor instead.
func (*BrcowcowOpenResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{8}
}

func (x *BrcowcowOpenResp) GetAwardArea() []byte {
	if x != nil {
		return x.AwardArea
	}
	return nil
}

func (x *BrcowcowOpenResp) GetBankerCard() *CardInfo {
	if x != nil {
		return x.BankerCard
	}
	return nil
}

func (x *BrcowcowOpenResp) GetTianCard() *CardInfo {
	if x != nil {
		return x.TianCard
	}
	return nil
}

func (x *BrcowcowOpenResp) GetXuanCard() *CardInfo {
	if x != nil {
		return x.XuanCard
	}
	return nil
}

func (x *BrcowcowOpenResp) GetDiCard() *CardInfo {
	if x != nil {
		return x.DiCard
	}
	return nil
}

func (x *BrcowcowOpenResp) GetHuangCard() *CardInfo {
	if x != nil {
		return x.HuangCard
	}
	return nil
}

//结算
type BrcowcowOverResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MyAcquire       int64   `protobuf:"varint,1,opt,name=MyAcquire,proto3" json:"MyAcquire,omitempty"`                    //个人所得
	TotalSettlement []int64 `protobuf:"varint,2,rep,packed,name=TotalSettlement,proto3" json:"TotalSettlement,omitempty"` //统计：庄家+各区域输赢钱数额结算
}

func (x *BrcowcowOverResp) Reset() {
	*x = BrcowcowOverResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowOverResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowOverResp) ProtoMessage() {}

func (x *BrcowcowOverResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowOverResp.ProtoReflect.Descriptor instead.
func (*BrcowcowOverResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{9}
}

func (x *BrcowcowOverResp) GetMyAcquire() int64 {
	if x != nil {
		return x.MyAcquire
	}
	return 0
}

func (x *BrcowcowOverResp) GetTotalSettlement() []int64 {
	if x != nil {
		return x.TotalSettlement
	}
	return nil
}

//----------------------------------------------------------------------------------
//抢庄
type BrcowcowHostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsWant bool `protobuf:"varint,1,opt,name=IsWant,proto3" json:"IsWant,omitempty"` //true上庄 false取消上庄
}

func (x *BrcowcowHostReq) Reset() {
	*x = BrcowcowHostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowHostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowHostReq) ProtoMessage() {}

func (x *BrcowcowHostReq) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowHostReq.ProtoReflect.Descriptor instead.
func (*BrcowcowHostReq) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{10}
}

func (x *BrcowcowHostReq) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

type BrcowcowHostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	IsWant bool   `protobuf:"varint,2,opt,name=IsWant,proto3" json:"IsWant,omitempty"` //true上庄 false取消上庄
}

func (x *BrcowcowHostResp) Reset() {
	*x = BrcowcowHostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowHostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowHostResp) ProtoMessage() {}

func (x *BrcowcowHostResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowHostResp.ProtoReflect.Descriptor instead.
func (*BrcowcowHostResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{11}
}

func (x *BrcowcowHostResp) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *BrcowcowHostResp) GetIsWant() bool {
	if x != nil {
		return x.IsWant
	}
	return false
}

//待上庄列表
type BrcowcowHostListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BrcowcowHostListReq) Reset() {
	*x = BrcowcowHostListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowHostListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowHostListReq) ProtoMessage() {}

func (x *BrcowcowHostListReq) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowHostListReq.ProtoReflect.Descriptor instead.
func (*BrcowcowHostListReq) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{12}
}

type BrcowcowHostListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurHost  *PlayerInfo `protobuf:"bytes,1,opt,name=CurHost,proto3" json:"CurHost,omitempty"`           //当前庄家
	Waitlist []uint64    `protobuf:"varint,2,rep,packed,name=Waitlist,proto3" json:"Waitlist,omitempty"` //待上庄列表
}

func (x *BrcowcowHostListResp) Reset() {
	*x = BrcowcowHostListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brcowcow_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrcowcowHostListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrcowcowHostListResp) ProtoMessage() {}

func (x *BrcowcowHostListResp) ProtoReflect() protoreflect.Message {
	mi := &file_brcowcow_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrcowcowHostListResp.ProtoReflect.Descriptor instead.
func (*BrcowcowHostListResp) Descriptor() ([]byte, []int) {
	return file_brcowcow_proto_rawDescGZIP(), []int{13}
}

func (x *BrcowcowHostListResp) GetCurHost() *PlayerInfo {
	if x != nil {
		return x.CurHost
	}
	return nil
}

func (x *BrcowcowHostListResp) GetWaitlist() []uint64 {
	if x != nil {
		return x.Waitlist
	}
	return nil
}

var File_brcowcow_proto protoreflect.FileDescriptor

var file_brcowcow_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x67, 0x6f, 0x1a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x01, 0x0a, 0x11, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f,
	0x77, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x68, 0x69, 0x70,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x43, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x41, 0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0a, 0x41, 0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x41, 0x72, 0x65, 0x61, 0x42, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03,
	0x52, 0x08, 0x41, 0x72, 0x65, 0x61, 0x42, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x79,
	0x42, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x4d, 0x79, 0x42, 0x65,
	0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x32, 0x0a, 0x0a, 0x41, 0x6c,
	0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0a, 0x41, 0x6c, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x15, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63,
	0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x72, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x22, 0x0a, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x67, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x16, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a,
	0x05, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67,
	0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x49, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x22, 0x3e, 0x0a, 0x18, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a,
	0x05, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67,
	0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x22, 0x6d, 0x0a, 0x15, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x30,
	0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x2e, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x4f, 0x70,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x3b, 0x0a, 0x15, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x4f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x22, 0x0a, 0x05, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x46, 0x0a,
	0x0e, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x42, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x42, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x42, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x65, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x42, 0x65, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x5f, 0x0a, 0x0f, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f,
	0x77, 0x42, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x42, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x42, 0x65, 0x74, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x65,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x42, 0x65,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x84, 0x02, 0x0a, 0x10, 0x42, 0x72, 0x63, 0x6f, 0x77,
	0x63, 0x6f, 0x77, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x41, 0x77, 0x61, 0x72, 0x64, 0x41, 0x72, 0x65, 0x61, 0x12, 0x2c, 0x0a, 0x0a, 0x42, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x67, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x42, 0x61, 0x6e,
	0x6b, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x54, 0x69, 0x61, 0x6e, 0x43,
	0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x54, 0x69, 0x61, 0x6e, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x28, 0x0a, 0x08, 0x58, 0x75, 0x61, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x58, 0x75, 0x61, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x44,
	0x69, 0x43, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f,
	0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x44, 0x69, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x2a, 0x0a, 0x09, 0x48, 0x75, 0x61, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x48, 0x75, 0x61, 0x6e, 0x67, 0x43, 0x61, 0x72, 0x64, 0x22, 0x5a, 0x0a,
	0x10, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x4f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x79, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x4d, 0x79, 0x41, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x42, 0x72, 0x63,
	0x6f, 0x77, 0x63, 0x6f, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73,
	0x57, 0x61, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x10, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77,
	0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x49, 0x73, 0x57, 0x61, 0x6e, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x42, 0x72, 0x63, 0x6f,
	0x77, 0x63, 0x6f, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22,
	0x5c, 0x0a, 0x14, 0x42, 0x72, 0x63, 0x6f, 0x77, 0x63, 0x6f, 0x77, 0x48, 0x6f, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x07, 0x43, 0x75, 0x72, 0x48, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x6f, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x43, 0x75, 0x72, 0x48, 0x6f, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x57, 0x61, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x08, 0x57, 0x61, 0x69, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_brcowcow_proto_rawDescOnce sync.Once
	file_brcowcow_proto_rawDescData = file_brcowcow_proto_rawDesc
)

func file_brcowcow_proto_rawDescGZIP() []byte {
	file_brcowcow_proto_rawDescOnce.Do(func() {
		file_brcowcow_proto_rawDescData = protoimpl.X.CompressGZIP(file_brcowcow_proto_rawDescData)
	})
	return file_brcowcow_proto_rawDescData
}

var file_brcowcow_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_brcowcow_proto_goTypes = []interface{}{
	(*BrcowcowSceneResp)(nil),        // 0: go.BrcowcowSceneResp
	(*BrcowcowStateFreeResp)(nil),    // 1: go.BrcowcowStateFreeResp
	(*BrcowcowStateStartResp)(nil),   // 2: go.BrcowcowStateStartResp
	(*BrcowcowStatePlayingResp)(nil), // 3: go.BrcowcowStatePlayingResp
	(*BrcowcowStateOpenResp)(nil),    // 4: go.BrcowcowStateOpenResp
	(*BrcowcowStateOverResp)(nil),    // 5: go.BrcowcowStateOverResp
	(*BrcowcowBetReq)(nil),           // 6: go.BrcowcowBetReq
	(*BrcowcowBetResp)(nil),          // 7: go.BrcowcowBetResp
	(*BrcowcowOpenResp)(nil),         // 8: go.BrcowcowOpenResp
	(*BrcowcowOverResp)(nil),         // 9: go.BrcowcowOverResp
	(*BrcowcowHostReq)(nil),          // 10: go.BrcowcowHostReq
	(*BrcowcowHostResp)(nil),         // 11: go.BrcowcowHostResp
	(*BrcowcowHostListReq)(nil),      // 12: go.BrcowcowHostListReq
	(*BrcowcowHostListResp)(nil),     // 13: go.BrcowcowHostListResp
	(*PlayerListInfo)(nil),           // 14: go.PlayerListInfo
	(*TimeInfo)(nil),                 // 15: go.TimeInfo
	(*CardInfo)(nil),                 // 16: go.CardInfo
	(*PlayerInfo)(nil),               // 17: go.PlayerInfo
}
var file_brcowcow_proto_depIdxs = []int32{
	14, // 0: go.BrcowcowSceneResp.AllPlayers:type_name -> go.PlayerListInfo
	15, // 1: go.BrcowcowStateFreeResp.Times:type_name -> go.TimeInfo
	15, // 2: go.BrcowcowStateStartResp.Times:type_name -> go.TimeInfo
	15, // 3: go.BrcowcowStatePlayingResp.Times:type_name -> go.TimeInfo
	15, // 4: go.BrcowcowStateOpenResp.Times:type_name -> go.TimeInfo
	8,  // 5: go.BrcowcowStateOpenResp.OpenInfo:type_name -> go.BrcowcowOpenResp
	15, // 6: go.BrcowcowStateOverResp.Times:type_name -> go.TimeInfo
	16, // 7: go.BrcowcowOpenResp.BankerCard:type_name -> go.CardInfo
	16, // 8: go.BrcowcowOpenResp.TianCard:type_name -> go.CardInfo
	16, // 9: go.BrcowcowOpenResp.XuanCard:type_name -> go.CardInfo
	16, // 10: go.BrcowcowOpenResp.DiCard:type_name -> go.CardInfo
	16, // 11: go.BrcowcowOpenResp.HuangCard:type_name -> go.CardInfo
	17, // 12: go.BrcowcowHostListResp.CurHost:type_name -> go.PlayerInfo
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_brcowcow_proto_init() }
func file_brcowcow_proto_init() {
	if File_brcowcow_proto != nil {
		return
	}
	file_gamecomm_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_brcowcow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowSceneResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowStateFreeResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowStateStartResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowStatePlayingResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowStateOpenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowStateOverResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowBetReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowBetResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowOpenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowOverResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowHostReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowHostResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowHostListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_brcowcow_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrcowcowHostListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_brcowcow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_brcowcow_proto_goTypes,
		DependencyIndexes: file_brcowcow_proto_depIdxs,
		MessageInfos:      file_brcowcow_proto_msgTypes,
	}.Build()
	File_brcowcow_proto = out.File
	file_brcowcow_proto_rawDesc = nil
	file_brcowcow_proto_goTypes = nil
	file_brcowcow_proto_depIdxs = nil
}