// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.19.6
// source: game/uncollected-garbage.proto

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

type UncollectedGarbage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *math.Point `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Garbage  *Garbage    `protobuf:"bytes,2,opt,name=garbage,proto3" json:"garbage,omitempty"`
}

func (x *UncollectedGarbage) Reset() {
	*x = UncollectedGarbage{}
	mi := &file_game_uncollected_garbage_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UncollectedGarbage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UncollectedGarbage) ProtoMessage() {}

func (x *UncollectedGarbage) ProtoReflect() protoreflect.Message {
	mi := &file_game_uncollected_garbage_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UncollectedGarbage.ProtoReflect.Descriptor instead.
func (*UncollectedGarbage) Descriptor() ([]byte, []int) {
	return file_game_uncollected_garbage_proto_rawDescGZIP(), []int{0}
}

func (x *UncollectedGarbage) GetLocation() *math.Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *UncollectedGarbage) GetGarbage() *Garbage {
	if x != nil {
		return x.Garbage
	}
	return nil
}

var File_game_uncollected_garbage_proto protoreflect.FileDescriptor

var file_game_uncollected_garbage_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x75, 0x6e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x2d, 0x67, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x1a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x67, 0x61, 0x72,
	0x62, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6d, 0x61, 0x74, 0x68,
	0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x12,
	0x55, 0x6e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x47, 0x61, 0x72, 0x62, 0x61,
	0x67, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x07, 0x67,
	0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x52, 0x07, 0x67, 0x61, 0x72,
	0x62, 0x61, 0x67, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x69, 0x6b, 0x6f, 0x6e, 0x65, 0x4b, 0x69, 0x73, 0x75, 0x2f, 0x72,
	0x65, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x2d, 0x69, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x67, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_uncollected_garbage_proto_rawDescOnce sync.Once
	file_game_uncollected_garbage_proto_rawDescData = file_game_uncollected_garbage_proto_rawDesc
)

func file_game_uncollected_garbage_proto_rawDescGZIP() []byte {
	file_game_uncollected_garbage_proto_rawDescOnce.Do(func() {
		file_game_uncollected_garbage_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_uncollected_garbage_proto_rawDescData)
	})
	return file_game_uncollected_garbage_proto_rawDescData
}

var file_game_uncollected_garbage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_game_uncollected_garbage_proto_goTypes = []any{
	(*UncollectedGarbage)(nil), // 0: game.UncollectedGarbage
	(*math.Point)(nil),         // 1: math.Point
	(*Garbage)(nil),            // 2: game.Garbage
}
var file_game_uncollected_garbage_proto_depIdxs = []int32{
	1, // 0: game.UncollectedGarbage.location:type_name -> math.Point
	2, // 1: game.UncollectedGarbage.garbage:type_name -> game.Garbage
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_game_uncollected_garbage_proto_init() }
func file_game_uncollected_garbage_proto_init() {
	if File_game_uncollected_garbage_proto != nil {
		return
	}
	file_game_garbage_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_uncollected_garbage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_uncollected_garbage_proto_goTypes,
		DependencyIndexes: file_game_uncollected_garbage_proto_depIdxs,
		MessageInfos:      file_game_uncollected_garbage_proto_msgTypes,
	}.Build()
	File_game_uncollected_garbage_proto = out.File
	file_game_uncollected_garbage_proto_rawDesc = nil
	file_game_uncollected_garbage_proto_goTypes = nil
	file_game_uncollected_garbage_proto_depIdxs = nil
}
