// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: lobby.proto

package lobby

import (
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

type JoinRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *JoinRoomRequest) Reset() {
	*x = JoinRoomRequest{}
	mi := &file_lobby_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomRequest) ProtoMessage() {}

func (x *JoinRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomRequest.ProtoReflect.Descriptor instead.
func (*JoinRoomRequest) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRoomRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type JoinRoomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtherPlayers []*Player `protobuf:"bytes,2,rep,name=other_players,json=otherPlayers,proto3" json:"other_players,omitempty"`
	Me           *Player   `protobuf:"bytes,3,opt,name=me,proto3" json:"me,omitempty"`
	MaxPlayers   int32     `protobuf:"varint,4,opt,name=max_players,json=maxPlayers,proto3" json:"max_players,omitempty"`
}

func (x *JoinRoomResponse) Reset() {
	*x = JoinRoomResponse{}
	mi := &file_lobby_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRoomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRoomResponse) ProtoMessage() {}

func (x *JoinRoomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRoomResponse.ProtoReflect.Descriptor instead.
func (*JoinRoomResponse) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{1}
}

func (x *JoinRoomResponse) GetOtherPlayers() []*Player {
	if x != nil {
		return x.OtherPlayers
	}
	return nil
}

func (x *JoinRoomResponse) GetMe() *Player {
	if x != nil {
		return x.Me
	}
	return nil
}

func (x *JoinRoomResponse) GetMaxPlayers() int32 {
	if x != nil {
		return x.MaxPlayers
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_lobby_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_lobby_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_lobby_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

var File_lobby_proto protoreflect.FileDescriptor

var file_lobby_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a,
	0x0f, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7a, 0x0a,
	0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x0d, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x0c, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12,
	0x17, 0x0a, 0x02, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x02, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x25, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x32, 0x3f, 0x0a, 0x0c, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2f, 0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x10, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x75, 0x69, 0x73, 0x67, 0x36, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f,
	0x62, 0x62, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lobby_proto_rawDescOnce sync.Once
	file_lobby_proto_rawDescData = file_lobby_proto_rawDesc
)

func file_lobby_proto_rawDescGZIP() []byte {
	file_lobby_proto_rawDescOnce.Do(func() {
		file_lobby_proto_rawDescData = protoimpl.X.CompressGZIP(file_lobby_proto_rawDescData)
	})
	return file_lobby_proto_rawDescData
}

var file_lobby_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lobby_proto_goTypes = []any{
	(*JoinRoomRequest)(nil),  // 0: JoinRoomRequest
	(*JoinRoomResponse)(nil), // 1: JoinRoomResponse
	(*Player)(nil),           // 2: Player
}
var file_lobby_proto_depIdxs = []int32{
	2, // 0: JoinRoomResponse.other_players:type_name -> Player
	2, // 1: JoinRoomResponse.me:type_name -> Player
	0, // 2: LobbyService.JoinRoom:input_type -> JoinRoomRequest
	1, // 3: LobbyService.JoinRoom:output_type -> JoinRoomResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lobby_proto_init() }
func file_lobby_proto_init() {
	if File_lobby_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lobby_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lobby_proto_goTypes,
		DependencyIndexes: file_lobby_proto_depIdxs,
		MessageInfos:      file_lobby_proto_msgTypes,
	}.Build()
	File_lobby_proto = out.File
	file_lobby_proto_rawDesc = nil
	file_lobby_proto_goTypes = nil
	file_lobby_proto_depIdxs = nil
}
