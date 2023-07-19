// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: matchmaking.proto

package simple_mm

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

type MatchmakingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Mmr      int64  `protobuf:"varint,2,opt,name=mmr,proto3" json:"mmr,omitempty"`
}

func (x *MatchmakingRequest) Reset() {
	*x = MatchmakingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matchmaking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchmakingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchmakingRequest) ProtoMessage() {}

func (x *MatchmakingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_matchmaking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchmakingRequest.ProtoReflect.Descriptor instead.
func (*MatchmakingRequest) Descriptor() ([]byte, []int) {
	return file_matchmaking_proto_rawDescGZIP(), []int{0}
}

func (x *MatchmakingRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MatchmakingRequest) GetMmr() int64 {
	if x != nil {
		return x.Mmr
	}
	return 0
}

type MatchmakingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usernames []string `protobuf:"bytes,1,rep,name=usernames,proto3" json:"usernames,omitempty"`
	Mmr       []int64  `protobuf:"varint,2,rep,packed,name=mmr,proto3" json:"mmr,omitempty"`
}

func (x *MatchmakingResponse) Reset() {
	*x = MatchmakingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_matchmaking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchmakingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchmakingResponse) ProtoMessage() {}

func (x *MatchmakingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_matchmaking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchmakingResponse.ProtoReflect.Descriptor instead.
func (*MatchmakingResponse) Descriptor() ([]byte, []int) {
	return file_matchmaking_proto_rawDescGZIP(), []int{1}
}

func (x *MatchmakingResponse) GetUsernames() []string {
	if x != nil {
		return x.Usernames
	}
	return nil
}

func (x *MatchmakingResponse) GetMmr() []int64 {
	if x != nil {
		return x.Mmr
	}
	return nil
}

var File_matchmaking_proto protoreflect.FileDescriptor

var file_matchmaking_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x12, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x6d, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x6d, 0x72, 0x22, 0x45, 0x0a,
	0x13, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x6d, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x03, 0x6d, 0x6d, 0x72, 0x32, 0xa3, 0x01, 0x0a, 0x12, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x46, 0x61, 0x6b, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x2d, 0x6d, 0x6d, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_matchmaking_proto_rawDescOnce sync.Once
	file_matchmaking_proto_rawDescData = file_matchmaking_proto_rawDesc
)

func file_matchmaking_proto_rawDescGZIP() []byte {
	file_matchmaking_proto_rawDescOnce.Do(func() {
		file_matchmaking_proto_rawDescData = protoimpl.X.CompressGZIP(file_matchmaking_proto_rawDescData)
	})
	return file_matchmaking_proto_rawDescData
}

var file_matchmaking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_matchmaking_proto_goTypes = []interface{}{
	(*MatchmakingRequest)(nil),  // 0: user.MatchmakingRequest
	(*MatchmakingResponse)(nil), // 1: user.MatchmakingResponse
}
var file_matchmaking_proto_depIdxs = []int32{
	0, // 0: user.MatchmakingService.AddMatchmaking:input_type -> user.MatchmakingRequest
	0, // 1: user.MatchmakingService.AddFakeUser:input_type -> user.MatchmakingRequest
	1, // 2: user.MatchmakingService.AddMatchmaking:output_type -> user.MatchmakingResponse
	1, // 3: user.MatchmakingService.AddFakeUser:output_type -> user.MatchmakingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_matchmaking_proto_init() }
func file_matchmaking_proto_init() {
	if File_matchmaking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_matchmaking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchmakingRequest); i {
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
		file_matchmaking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchmakingResponse); i {
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
			RawDescriptor: file_matchmaking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_matchmaking_proto_goTypes,
		DependencyIndexes: file_matchmaking_proto_depIdxs,
		MessageInfos:      file_matchmaking_proto_msgTypes,
	}.Build()
	File_matchmaking_proto = out.File
	file_matchmaking_proto_rawDesc = nil
	file_matchmaking_proto_goTypes = nil
	file_matchmaking_proto_depIdxs = nil
}