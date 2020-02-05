// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

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

//注册
type Register struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	SecurityCode         string   `protobuf:"bytes,3,opt,name=SecurityCode,proto3" json:"SecurityCode,omitempty"`
	MachineCode          string   `protobuf:"bytes,4,opt,name=MachineCode,proto3" json:"MachineCode,omitempty"`
	InvitationCode       string   `protobuf:"bytes,5,opt,name=InvitationCode,proto3" json:"InvitationCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Register) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Register) GetSecurityCode() string {
	if m != nil {
		return m.SecurityCode
	}
	return ""
}

func (m *Register) GetMachineCode() string {
	if m != nil {
		return m.MachineCode
	}
	return ""
}

func (m *Register) GetInvitationCode() string {
	if m != nil {
		return m.InvitationCode
	}
	return ""
}

//注册结果
type RegisterResult struct {
	State                uint32   `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
	Hints                string   `protobuf:"bytes,2,opt,name=Hints,proto3" json:"Hints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResult) Reset()         { *m = RegisterResult{} }
func (m *RegisterResult) String() string { return proto.CompactTextString(m) }
func (*RegisterResult) ProtoMessage()    {}
func (*RegisterResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *RegisterResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResult.Unmarshal(m, b)
}
func (m *RegisterResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResult.Marshal(b, m, deterministic)
}
func (m *RegisterResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResult.Merge(m, src)
}
func (m *RegisterResult) XXX_Size() int {
	return xxx_messageInfo_RegisterResult.Size(m)
}
func (m *RegisterResult) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResult.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResult proto.InternalMessageInfo

func (m *RegisterResult) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *RegisterResult) GetHints() string {
	if m != nil {
		return m.Hints
	}
	return ""
}

//登录
type Login struct {
	Account              string   `protobuf:"bytes,1,opt,name=Account,proto3" json:"Account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	SecurityCode         string   `protobuf:"bytes,3,opt,name=SecurityCode,proto3" json:"SecurityCode,omitempty"`
	MachineCode          string   `protobuf:"bytes,4,opt,name=MachineCode,proto3" json:"MachineCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
}

func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (m *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(m, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Login) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Login) GetSecurityCode() string {
	if m != nil {
		return m.SecurityCode
	}
	return ""
}

func (m *Login) GetMachineCode() string {
	if m != nil {
		return m.MachineCode
	}
	return ""
}

//结果反馈
type ResResult struct {
	State                uint32   `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
	Hints                string   `protobuf:"bytes,2,opt,name=Hints,proto3" json:"Hints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResResult) Reset()         { *m = ResResult{} }
func (m *ResResult) String() string { return proto.CompactTextString(m) }
func (*ResResult) ProtoMessage()    {}
func (*ResResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{3}
}

func (m *ResResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResResult.Unmarshal(m, b)
}
func (m *ResResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResResult.Marshal(b, m, deterministic)
}
func (m *ResResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResResult.Merge(m, src)
}
func (m *ResResult) XXX_Size() int {
	return xxx_messageInfo_ResResult.Size(m)
}
func (m *ResResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ResResult.DiscardUnknown(m)
}

var xxx_messageInfo_ResResult proto.InternalMessageInfo

func (m *ResResult) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *ResResult) GetHints() string {
	if m != nil {
		return m.Hints
	}
	return ""
}

/////////////list/////////////////
//任务信息
type TaskItem struct {
	TaskID               uint32   `protobuf:"varint,1,opt,name=TaskID,proto3" json:"TaskID,omitempty"`
	Twice                uint32   `protobuf:"varint,2,opt,name=Twice,proto3" json:"Twice,omitempty"`
	Hints                string   `protobuf:"bytes,3,opt,name=Hints,proto3" json:"Hints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskItem) Reset()         { *m = TaskItem{} }
func (m *TaskItem) String() string { return proto.CompactTextString(m) }
func (*TaskItem) ProtoMessage()    {}
func (*TaskItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{4}
}

func (m *TaskItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskItem.Unmarshal(m, b)
}
func (m *TaskItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskItem.Marshal(b, m, deterministic)
}
func (m *TaskItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskItem.Merge(m, src)
}
func (m *TaskItem) XXX_Size() int {
	return xxx_messageInfo_TaskItem.Size(m)
}
func (m *TaskItem) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskItem.DiscardUnknown(m)
}

var xxx_messageInfo_TaskItem proto.InternalMessageInfo

func (m *TaskItem) GetTaskID() uint32 {
	if m != nil {
		return m.TaskID
	}
	return 0
}

func (m *TaskItem) GetTwice() uint32 {
	if m != nil {
		return m.Twice
	}
	return 0
}

func (m *TaskItem) GetHints() string {
	if m != nil {
		return m.Hints
	}
	return ""
}

//任务列表
type TaskList struct {
	Task                 []*TaskItem `protobuf:"bytes,1,rep,name=Task,proto3" json:"Task,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{5}
}

func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (m *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(m, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

func (m *TaskList) GetTask() []*TaskItem {
	if m != nil {
		return m.Task
	}
	return nil
}

//游戏列表
type GameList struct {
	Items                []*GameItem `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GameList) Reset()         { *m = GameList{} }
func (m *GameList) String() string { return proto.CompactTextString(m) }
func (*GameList) ProtoMessage()    {}
func (*GameList) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{6}
}

func (m *GameList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameList.Unmarshal(m, b)
}
func (m *GameList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameList.Marshal(b, m, deterministic)
}
func (m *GameList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameList.Merge(m, src)
}
func (m *GameList) XXX_Size() int {
	return xxx_messageInfo_GameList.Size(m)
}
func (m *GameList) XXX_DiscardUnknown() {
	xxx_messageInfo_GameList.DiscardUnknown(m)
}

var xxx_messageInfo_GameList proto.InternalMessageInfo

func (m *GameList) GetItems() []*GameItem {
	if m != nil {
		return m.Items
	}
	return nil
}

/////////////info//////////////////
//个人信息
type UserInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Accounts             string   `protobuf:"bytes,2,opt,name=Accounts,proto3" json:"Accounts,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	FaceID               uint32   `protobuf:"varint,4,opt,name=FaceID,proto3" json:"FaceID,omitempty"`
	Gender               uint32   `protobuf:"varint,5,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Age                  uint32   `protobuf:"varint,6,opt,name=Age,proto3" json:"Age,omitempty"`
	Level                uint32   `protobuf:"varint,7,opt,name=Level,proto3" json:"Level,omitempty"`
	Gold                 int64    `protobuf:"varint,8,opt,name=Gold,proto3" json:"Gold,omitempty"`
	PassPortID           string   `protobuf:"bytes,9,opt,name=PassPortID,proto3" json:"PassPortID,omitempty"`
	Compellation         string   `protobuf:"bytes,10,opt,name=Compellation,proto3" json:"Compellation,omitempty"`
	AgentID              uint32   `protobuf:"varint,11,opt,name=AgentID,proto3" json:"AgentID,omitempty"`
	SpreaderGameID       uint32   `protobuf:"varint,12,opt,name=SpreaderGameID,proto3" json:"SpreaderGameID,omitempty"`
	ClientAddr           uint32   `protobuf:"varint,13,opt,name=ClientAddr,proto3" json:"ClientAddr,omitempty"`
	MachineCode          string   `protobuf:"bytes,14,opt,name=MachineCode,proto3" json:"MachineCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{7}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetAccounts() string {
	if m != nil {
		return m.Accounts
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserInfo) GetFaceID() uint32 {
	if m != nil {
		return m.FaceID
	}
	return 0
}

func (m *UserInfo) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UserInfo) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *UserInfo) GetGold() int64 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *UserInfo) GetPassPortID() string {
	if m != nil {
		return m.PassPortID
	}
	return ""
}

func (m *UserInfo) GetCompellation() string {
	if m != nil {
		return m.Compellation
	}
	return ""
}

func (m *UserInfo) GetAgentID() uint32 {
	if m != nil {
		return m.AgentID
	}
	return 0
}

func (m *UserInfo) GetSpreaderGameID() uint32 {
	if m != nil {
		return m.SpreaderGameID
	}
	return 0
}

func (m *UserInfo) GetClientAddr() uint32 {
	if m != nil {
		return m.ClientAddr
	}
	return 0
}

func (m *UserInfo) GetMachineCode() string {
	if m != nil {
		return m.MachineCode
	}
	return ""
}

//房间信息
type RoomInfo struct {
	RoomNum              uint32    `protobuf:"varint,1,opt,name=RoomNum,proto3" json:"RoomNum,omitempty"`
	RoomKey              string    `protobuf:"bytes,2,opt,name=RoomKey,proto3" json:"RoomKey,omitempty"`
	RoomName             string    `protobuf:"bytes,3,opt,name=RoomName,proto3" json:"RoomName,omitempty"`
	Games                *GameList `protobuf:"bytes,4,opt,name=Games,proto3" json:"Games,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{8}
}

func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInfo.Unmarshal(m, b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
}
func (m *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(m, src)
}
func (m *RoomInfo) XXX_Size() int {
	return xxx_messageInfo_RoomInfo.Size(m)
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetRoomNum() uint32 {
	if m != nil {
		return m.RoomNum
	}
	return 0
}

func (m *RoomInfo) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *RoomInfo) GetRoomName() string {
	if m != nil {
		return m.RoomName
	}
	return ""
}

func (m *RoomInfo) GetGames() *GameList {
	if m != nil {
		return m.Games
	}
	return nil
}

//游戏信息
type GameBaseInfo struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	KindID               uint32   `protobuf:"varint,2,opt,name=KindID,proto3" json:"KindID,omitempty"`
	Level                uint32   `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	EnterScore           uint32   `protobuf:"varint,5,opt,name=EnterScore,proto3" json:"EnterScore,omitempty"`
	LessScore            uint32   `protobuf:"varint,6,opt,name=LessScore,proto3" json:"LessScore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameBaseInfo) Reset()         { *m = GameBaseInfo{} }
func (m *GameBaseInfo) String() string { return proto.CompactTextString(m) }
func (*GameBaseInfo) ProtoMessage()    {}
func (*GameBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{9}
}

func (m *GameBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameBaseInfo.Unmarshal(m, b)
}
func (m *GameBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameBaseInfo.Marshal(b, m, deterministic)
}
func (m *GameBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameBaseInfo.Merge(m, src)
}
func (m *GameBaseInfo) XXX_Size() int {
	return xxx_messageInfo_GameBaseInfo.Size(m)
}
func (m *GameBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_GameBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_GameBaseInfo proto.InternalMessageInfo

func (m *GameBaseInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *GameBaseInfo) GetKindID() uint32 {
	if m != nil {
		return m.KindID
	}
	return 0
}

func (m *GameBaseInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *GameBaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GameBaseInfo) GetEnterScore() uint32 {
	if m != nil {
		return m.EnterScore
	}
	return 0
}

func (m *GameBaseInfo) GetLessScore() uint32 {
	if m != nil {
		return m.LessScore
	}
	return 0
}

//子游戏
type GameItem struct {
	ID                   uint32        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Info                 *GameBaseInfo `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GameItem) Reset()         { *m = GameItem{} }
func (m *GameItem) String() string { return proto.CompactTextString(m) }
func (*GameItem) ProtoMessage()    {}
func (*GameItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{10}
}

func (m *GameItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameItem.Unmarshal(m, b)
}
func (m *GameItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameItem.Marshal(b, m, deterministic)
}
func (m *GameItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameItem.Merge(m, src)
}
func (m *GameItem) XXX_Size() int {
	return xxx_messageInfo_GameItem.Size(m)
}
func (m *GameItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GameItem.DiscardUnknown(m)
}

var xxx_messageInfo_GameItem proto.InternalMessageInfo

func (m *GameItem) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *GameItem) GetInfo() *GameBaseInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

//主页信息
type MasterInfo struct {
	UserInfo             *UserInfo   `protobuf:"bytes,1,opt,name=UserInfo,proto3" json:"UserInfo,omitempty"`
	RoomsInfo            []*RoomInfo `protobuf:"bytes,2,rep,name=RoomsInfo,proto3" json:"RoomsInfo,omitempty"`
	Tasks                *TaskList   `protobuf:"bytes,3,opt,name=Tasks,proto3" json:"Tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MasterInfo) Reset()         { *m = MasterInfo{} }
func (m *MasterInfo) String() string { return proto.CompactTextString(m) }
func (*MasterInfo) ProtoMessage()    {}
func (*MasterInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{11}
}

func (m *MasterInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MasterInfo.Unmarshal(m, b)
}
func (m *MasterInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MasterInfo.Marshal(b, m, deterministic)
}
func (m *MasterInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterInfo.Merge(m, src)
}
func (m *MasterInfo) XXX_Size() int {
	return xxx_messageInfo_MasterInfo.Size(m)
}
func (m *MasterInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MasterInfo proto.InternalMessageInfo

func (m *MasterInfo) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *MasterInfo) GetRoomsInfo() []*RoomInfo {
	if m != nil {
		return m.RoomsInfo
	}
	return nil
}

func (m *MasterInfo) GetTasks() *TaskList {
	if m != nil {
		return m.Tasks
	}
	return nil
}

//////////////客户端发起(上行)/////////////////////////////
//进入房间
type ReqEnterRoom struct {
	RoomNum              uint32   `protobuf:"varint,1,opt,name=RoomNum,proto3" json:"RoomNum,omitempty"`
	RoomKey              string   `protobuf:"bytes,2,opt,name=RoomKey,proto3" json:"RoomKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEnterRoom) Reset()         { *m = ReqEnterRoom{} }
func (m *ReqEnterRoom) String() string { return proto.CompactTextString(m) }
func (*ReqEnterRoom) ProtoMessage()    {}
func (*ReqEnterRoom) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{12}
}

func (m *ReqEnterRoom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEnterRoom.Unmarshal(m, b)
}
func (m *ReqEnterRoom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEnterRoom.Marshal(b, m, deterministic)
}
func (m *ReqEnterRoom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEnterRoom.Merge(m, src)
}
func (m *ReqEnterRoom) XXX_Size() int {
	return xxx_messageInfo_ReqEnterRoom.Size(m)
}
func (m *ReqEnterRoom) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEnterRoom.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEnterRoom proto.InternalMessageInfo

func (m *ReqEnterRoom) GetRoomNum() uint32 {
	if m != nil {
		return m.RoomNum
	}
	return 0
}

func (m *ReqEnterRoom) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

//进入游戏
type ReqEnterGame struct {
	GameID               uint32   `protobuf:"varint,1,opt,name=GameID,proto3" json:"GameID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEnterGame) Reset()         { *m = ReqEnterGame{} }
func (m *ReqEnterGame) String() string { return proto.CompactTextString(m) }
func (*ReqEnterGame) ProtoMessage()    {}
func (*ReqEnterGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{13}
}

func (m *ReqEnterGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEnterGame.Unmarshal(m, b)
}
func (m *ReqEnterGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEnterGame.Marshal(b, m, deterministic)
}
func (m *ReqEnterGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEnterGame.Merge(m, src)
}
func (m *ReqEnterGame) XXX_Size() int {
	return xxx_messageInfo_ReqEnterGame.Size(m)
}
func (m *ReqEnterGame) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEnterGame.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEnterGame proto.InternalMessageInfo

func (m *ReqEnterGame) GetGameID() uint32 {
	if m != nil {
		return m.GameID
	}
	return 0
}

//退出游戏
type ReqExitGame struct {
	GameID               uint32   `protobuf:"varint,1,opt,name=GameID,proto3" json:"GameID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqExitGame) Reset()         { *m = ReqExitGame{} }
func (m *ReqExitGame) String() string { return proto.CompactTextString(m) }
func (*ReqExitGame) ProtoMessage()    {}
func (*ReqExitGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{14}
}

func (m *ReqExitGame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqExitGame.Unmarshal(m, b)
}
func (m *ReqExitGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqExitGame.Marshal(b, m, deterministic)
}
func (m *ReqExitGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqExitGame.Merge(m, src)
}
func (m *ReqExitGame) XXX_Size() int {
	return xxx_messageInfo_ReqExitGame.Size(m)
}
func (m *ReqExitGame) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqExitGame.DiscardUnknown(m)
}

var xxx_messageInfo_ReqExitGame proto.InternalMessageInfo

func (m *ReqExitGame) GetGameID() uint32 {
	if m != nil {
		return m.GameID
	}
	return 0
}

func init() {
	proto.RegisterType((*Register)(nil), "go.Register")
	proto.RegisterType((*RegisterResult)(nil), "go.RegisterResult")
	proto.RegisterType((*Login)(nil), "go.Login")
	proto.RegisterType((*ResResult)(nil), "go.ResResult")
	proto.RegisterType((*TaskItem)(nil), "go.TaskItem")
	proto.RegisterType((*TaskList)(nil), "go.TaskList")
	proto.RegisterType((*GameList)(nil), "go.GameList")
	proto.RegisterType((*UserInfo)(nil), "go.UserInfo")
	proto.RegisterType((*RoomInfo)(nil), "go.RoomInfo")
	proto.RegisterType((*GameBaseInfo)(nil), "go.GameBaseInfo")
	proto.RegisterType((*GameItem)(nil), "go.GameItem")
	proto.RegisterType((*MasterInfo)(nil), "go.MasterInfo")
	proto.RegisterType((*ReqEnterRoom)(nil), "go.ReqEnterRoom")
	proto.RegisterType((*ReqEnterGame)(nil), "go.ReqEnterGame")
	proto.RegisterType((*ReqExitGame)(nil), "go.ReqExitGame")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xe3, 0x24, 0x4d, 0x26, 0x3f, 0xaa, 0x56, 0xa8, 0xb2, 0x10, 0xaa, 0x22, 0x0b, 0xaa,
	0x08, 0xa1, 0x1c, 0xca, 0x81, 0x0b, 0x07, 0xda, 0x06, 0x4a, 0xd4, 0xb4, 0xaa, 0x36, 0xe5, 0x01,
	0x4c, 0x3c, 0x04, 0x8b, 0xc4, 0x1b, 0xbc, 0x9b, 0x96, 0x5e, 0xb8, 0x20, 0x71, 0xe0, 0x31, 0x38,
	0xf1, 0x98, 0x68, 0x66, 0x77, 0x13, 0x37, 0x08, 0x0e, 0x95, 0xb8, 0xcd, 0xf7, 0xcd, 0xec, 0xee,
	0xfc, 0x7c, 0x63, 0x43, 0x6b, 0xae, 0x66, 0x59, 0x3e, 0x58, 0x16, 0xca, 0x28, 0x51, 0x99, 0xa9,
	0xf8, 0x57, 0x00, 0x0d, 0x89, 0xb3, 0x4c, 0x1b, 0x2c, 0x84, 0x80, 0xea, 0x45, 0xb2, 0xc0, 0x28,
	0xe8, 0x05, 0xfd, 0xa6, 0x64, 0x5b, 0x3c, 0x84, 0xc6, 0x65, 0xa2, 0xf5, 0x8d, 0x2a, 0xd2, 0xa8,
	0xc2, 0xfc, 0x1a, 0x8b, 0x18, 0xda, 0x13, 0x9c, 0xae, 0x8a, 0xcc, 0xdc, 0x9e, 0xa8, 0x14, 0xa3,
	0x90, 0xfd, 0x77, 0x38, 0xd1, 0x83, 0xd6, 0x79, 0x32, 0xfd, 0x98, 0xe5, 0xc8, 0x21, 0x55, 0x0e,
	0x29, 0x53, 0xe2, 0x00, 0xba, 0xa3, 0xfc, 0x3a, 0x33, 0x89, 0xc9, 0x54, 0xce, 0x41, 0x35, 0x0e,
	0xda, 0x62, 0xe3, 0x97, 0xd0, 0xf5, 0x99, 0x4a, 0xd4, 0xab, 0xb9, 0x11, 0x0f, 0xa0, 0x36, 0x31,
	0x89, 0xb1, 0x09, 0x77, 0xa4, 0x05, 0xc4, 0xbe, 0xcd, 0x72, 0xa3, 0x5d, 0xba, 0x16, 0xc4, 0xdf,
	0x02, 0xa8, 0x8d, 0xa9, 0x78, 0x11, 0xc1, 0xce, 0xd1, 0x74, 0xaa, 0x56, 0xb9, 0x71, 0x85, 0x7a,
	0xf8, 0xff, 0x6b, 0x8d, 0x5f, 0x40, 0x53, 0xa2, 0xbe, 0x47, 0xfa, 0x17, 0xd0, 0xb8, 0x4a, 0xf4,
	0xa7, 0x91, 0xc1, 0x85, 0xd8, 0x83, 0x3a, 0xdb, 0x43, 0x77, 0xd0, 0x21, 0x3a, 0x79, 0x75, 0x93,
	0x4d, 0x91, 0x4f, 0x76, 0xa4, 0x05, 0x9b, 0xfb, 0xc2, 0xf2, 0x7d, 0xcf, 0xec, 0x7d, 0xe3, 0x4c,
	0x1b, 0xd1, 0x83, 0x2a, 0xd9, 0x51, 0xd0, 0x0b, 0xfb, 0xad, 0xc3, 0xf6, 0x60, 0xa6, 0x06, 0xfe,
	0x2d, 0xc9, 0x9e, 0x78, 0x00, 0x8d, 0xd3, 0x64, 0x81, 0x1c, 0x1d, 0x43, 0x8d, 0x3c, 0xba, 0x1c,
	0x4e, 0x4e, 0x0e, 0xb7, 0xae, 0xf8, 0x47, 0x08, 0x8d, 0x77, 0x1a, 0x8b, 0x51, 0xfe, 0x41, 0xfd,
	0x4d, 0x55, 0xae, 0xe9, 0xbe, 0xce, 0x35, 0xbe, 0x33, 0x85, 0x70, 0x6b, 0x0a, 0x7b, 0x50, 0x7f,
	0x93, 0x4c, 0x71, 0x34, 0xe4, 0xe6, 0x76, 0xa4, 0x43, 0xc4, 0x9f, 0x62, 0x9e, 0x62, 0xc1, 0xda,
	0xe9, 0x48, 0x87, 0xc4, 0x2e, 0x84, 0x47, 0x33, 0x8c, 0xea, 0x4c, 0x92, 0x49, 0xed, 0x18, 0xe3,
	0x35, 0xce, 0xa3, 0x1d, 0xdb, 0x24, 0x06, 0x94, 0xe3, 0xa9, 0x9a, 0xa7, 0x51, 0xa3, 0x17, 0xf4,
	0x43, 0xc9, 0xb6, 0xd8, 0x07, 0xa0, 0x77, 0x2f, 0x55, 0x61, 0x46, 0xc3, 0xa8, 0xc9, 0x99, 0x94,
	0x18, 0x52, 0xc4, 0x89, 0x5a, 0x2c, 0x71, 0x3e, 0x67, 0x8d, 0x46, 0x60, 0x15, 0x51, 0xe6, 0x58,
	0x6b, 0x33, 0xcc, 0xe9, 0x82, 0x16, 0xbf, 0xe7, 0x21, 0xa9, 0x7e, 0xb2, 0x2c, 0x30, 0x49, 0xb1,
	0xe0, 0xee, 0x0d, 0xa3, 0x36, 0x07, 0x6c, 0xb1, 0x94, 0xc5, 0xc9, 0x3c, 0xc3, 0xdc, 0x1c, 0xa5,
	0x69, 0x11, 0x75, 0x38, 0xa6, 0xc4, 0x6c, 0x6b, 0xae, 0xfb, 0xa7, 0xe6, 0xbe, 0x42, 0x43, 0x2a,
	0xb5, 0xe0, 0x59, 0x44, 0xb0, 0x43, 0xf6, 0xc5, 0x6a, 0xe1, 0xb4, 0xe3, 0xa1, 0xf7, 0x9c, 0xe1,
	0xad, 0x1b, 0x88, 0x87, 0x34, 0x0f, 0x0e, 0xa2, 0x19, 0xba, 0x79, 0x78, 0x4c, 0x62, 0xa0, 0x3c,
	0x35, 0x8f, 0xa3, 0x24, 0x06, 0x52, 0x8a, 0xb4, 0xae, 0xf8, 0x67, 0x00, 0x6d, 0xb2, 0x8e, 0x13,
	0x8d, 0x5e, 0x10, 0x57, 0xb7, 0x4b, 0x2f, 0x7b, 0xb6, 0x69, 0x80, 0x67, 0x59, 0x9e, 0x8e, 0x86,
	0x4e, 0xbc, 0x0e, 0x6d, 0xc6, 0x15, 0x6e, 0x8d, 0x8b, 0xd3, 0xa9, 0x96, 0x24, 0xb5, 0x0f, 0xf0,
	0x3a, 0x37, 0x58, 0x4c, 0xa6, 0xaa, 0x40, 0x27, 0x83, 0x12, 0x23, 0x1e, 0x41, 0x73, 0x8c, 0x5a,
	0x5b, 0xb7, 0x15, 0xc4, 0x86, 0x88, 0x5f, 0x59, 0x85, 0xf3, 0x7e, 0x75, 0xa1, 0xb2, 0xde, 0xad,
	0xca, 0x68, 0x28, 0x1e, 0x43, 0x95, 0xf2, 0xe6, 0xcc, 0x5a, 0x87, 0xbb, 0xbe, 0x46, 0x5f, 0x8f,
	0x64, 0x6f, 0xfc, 0x3d, 0x00, 0x38, 0x4f, 0xe8, 0xeb, 0xc4, 0x45, 0xf6, 0x37, 0x1b, 0xc0, 0x57,
	0xb9, 0xe6, 0x78, 0x4e, 0x6e, 0xf6, 0xe3, 0x29, 0x34, 0xa9, 0x9f, 0xda, 0xbd, 0xb1, 0x5e, 0x2a,
	0x3f, 0x34, 0xb9, 0x71, 0x53, 0xbf, 0x69, 0x21, 0xed, 0x32, 0x97, 0x76, 0xd5, 0xf6, 0x9b, 0x5d,
	0xf1, 0x31, 0xb4, 0x25, 0x7e, 0xe6, 0xca, 0xe9, 0xe0, 0x7d, 0x66, 0x1e, 0x1f, 0x6c, 0xee, 0xa0,
	0x52, 0x79, 0xbf, 0xac, 0x4a, 0xdd, 0x27, 0xc7, 0xa2, 0xf8, 0x09, 0xb4, 0x28, 0xee, 0x4b, 0x66,
	0xfe, 0x15, 0xf6, 0xbe, 0xce, 0x3f, 0x9c, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x96,
	0x83, 0xae, 0x7f, 0x06, 0x00, 0x00,
}
