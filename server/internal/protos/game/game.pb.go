// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.6
// source: game/game.proto

package game

import (
	math "github.com/DaikoneKisu/recycle-it/server/internal/protos/math"
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

type Game struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Settings               *GameSettings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	Stage                  *GameStage    `protobuf:"bytes,3,opt,name=stage,proto3" json:"stage,omitempty"`
	TimeRemainingInSeconds int32         `protobuf:"varint,4,opt,name=timeRemainingInSeconds,proto3" json:"timeRemainingInSeconds,omitempty"`
}

func (x *Game) Reset() {
	*x = Game{}
	mi := &file_game_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Game) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Game) ProtoMessage() {}

func (x *Game) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Game.ProtoReflect.Descriptor instead.
func (*Game) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{0}
}

func (x *Game) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Game) GetSettings() *GameSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Game) GetStage() *GameStage {
	if x != nil {
		return x.Stage
	}
	return nil
}

func (x *Game) GetTimeRemainingInSeconds() int32 {
	if x != nil {
		return x.TimeRemainingInSeconds
	}
	return 0
}

type GameSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequiredPlayerAmount  int32 `protobuf:"varint,1,opt,name=requiredPlayerAmount,proto3" json:"requiredPlayerAmount,omitempty"`
	GameDurationInSeconds int32 `protobuf:"varint,2,opt,name=gameDurationInSeconds,proto3" json:"gameDurationInSeconds,omitempty"`
}

func (x *GameSettings) Reset() {
	*x = GameSettings{}
	mi := &file_game_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameSettings) ProtoMessage() {}

func (x *GameSettings) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameSettings.ProtoReflect.Descriptor instead.
func (*GameSettings) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameSettings) GetRequiredPlayerAmount() int32 {
	if x != nil {
		return x.RequiredPlayerAmount
	}
	return 0
}

func (x *GameSettings) GetGameDurationInSeconds() int32 {
	if x != nil {
		return x.GameDurationInSeconds
	}
	return 0
}

type GameStage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GarbageCollectors          []*GarbageCollector `protobuf:"bytes,1,rep,name=garbageCollectors,proto3" json:"garbageCollectors,omitempty"`
	UncollectedGarbage         Garbage             `protobuf:"varint,2,opt,name=uncollectedGarbage,proto3,enum=game.Garbage" json:"uncollectedGarbage,omitempty"`
	UncollectedGarbageLocation *math.Point2D       `protobuf:"bytes,3,opt,name=uncollectedGarbageLocation,proto3" json:"uncollectedGarbageLocation,omitempty"`
}

func (x *GameStage) Reset() {
	*x = GameStage{}
	mi := &file_game_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameStage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStage) ProtoMessage() {}

func (x *GameStage) ProtoReflect() protoreflect.Message {
	mi := &file_game_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStage.ProtoReflect.Descriptor instead.
func (*GameStage) Descriptor() ([]byte, []int) {
	return file_game_game_proto_rawDescGZIP(), []int{2}
}

func (x *GameStage) GetGarbageCollectors() []*GarbageCollector {
	if x != nil {
		return x.GarbageCollectors
	}
	return nil
}

func (x *GameStage) GetUncollectedGarbage() Garbage {
	if x != nil {
		return x.UncollectedGarbage
	}
	return Garbage_GARBAGE_UNKNOWN
}

func (x *GameStage) GetUncollectedGarbageLocation() *math.Point2D {
	if x != nil {
		return x.UncollectedGarbageLocation
	}
	return nil
}

var File_game_game_proto protoreflect.FileDescriptor

var file_game_game_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61,
	0x72, 0x62, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x61, 0x6d,
	0x65, 0x2f, 0x67, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x2d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6d, 0x61, 0x74, 0x68, 0x2f,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2d, 0x32, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5,
	0x01, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x08, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x12, 0x36,
	0x0a, 0x16, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x78, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x67, 0x61,
	0x6d, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x67, 0x61, 0x6d, 0x65, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73,
	0x22, 0xdf, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x67, 0x65, 0x12, 0x44,
	0x0a, 0x11, 0x67, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x52, 0x11, 0x67, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x3d, 0x0a, 0x12, 0x75, 0x6e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x52,
	0x12, 0x75, 0x6e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x47, 0x61, 0x72, 0x62,
	0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x1a, 0x75, 0x6e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x32, 0x44, 0x52, 0x1a, 0x75, 0x6e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x69, 0x6b, 0x6f, 0x6e, 0x65, 0x4b, 0x69, 0x73, 0x75, 0x2f, 0x72, 0x65, 0x63,
	0x79, 0x63, 0x6c, 0x65, 0x2d, 0x69, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_game_proto_rawDescOnce sync.Once
	file_game_game_proto_rawDescData = file_game_game_proto_rawDesc
)

func file_game_game_proto_rawDescGZIP() []byte {
	file_game_game_proto_rawDescOnce.Do(func() {
		file_game_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_game_proto_rawDescData)
	})
	return file_game_game_proto_rawDescData
}

var file_game_game_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_game_game_proto_goTypes = []any{
	(*Game)(nil),             // 0: game.Game
	(*GameSettings)(nil),     // 1: game.GameSettings
	(*GameStage)(nil),        // 2: game.GameStage
	(*GarbageCollector)(nil), // 3: game.GarbageCollector
	(Garbage)(0),             // 4: game.Garbage
	(*math.Point2D)(nil),     // 5: math.Point2D
}
var file_game_game_proto_depIdxs = []int32{
	1, // 0: game.Game.settings:type_name -> game.GameSettings
	2, // 1: game.Game.stage:type_name -> game.GameStage
	3, // 2: game.GameStage.garbageCollectors:type_name -> game.GarbageCollector
	4, // 3: game.GameStage.uncollectedGarbage:type_name -> game.Garbage
	5, // 4: game.GameStage.uncollectedGarbageLocation:type_name -> math.Point2D
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_game_game_proto_init() }
func file_game_game_proto_init() {
	if File_game_game_proto != nil {
		return
	}
	file_game_garbage_proto_init()
	file_game_garbage_collector_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_game_proto_goTypes,
		DependencyIndexes: file_game_game_proto_depIdxs,
		MessageInfos:      file_game_game_proto_msgTypes,
	}.Build()
	File_game_game_proto = out.File
	file_game_game_proto_rawDesc = nil
	file_game_game_proto_goTypes = nil
	file_game_game_proto_depIdxs = nil
}