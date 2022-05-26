// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: pb.proto

// protoc --go_out='./math' --go_opt=paths=source_relative --go-grpc_out='./math' --go-grpc_opt=paths=source_relative pb.proto

package math

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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vals []int64 `protobuf:"varint,1,rep,packed,name=vals,proto3" json:"vals,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetVals() []int64 {
	if x != nil {
		return x.Vals
	}
	return nil
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum int64 `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetSum() int64 {
	if x != nil {
		return x.Sum
	}
	return 0
}

type RandomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max int64 `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *RandomRequest) Reset() {
	*x = RandomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomRequest) ProtoMessage() {}

func (x *RandomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomRequest.ProtoReflect.Descriptor instead.
func (*RandomRequest) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{2}
}

func (x *RandomRequest) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

type RandomResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *RandomResponse) Reset() {
	*x = RandomResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RandomResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RandomResponse) ProtoMessage() {}

func (x *RandomResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RandomResponse.ProtoReflect.Descriptor instead.
func (*RandomResponse) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{3}
}

func (x *RandomResponse) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

var File_pb_proto protoreflect.FileDescriptor

var file_pb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x65, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x22, 0x20, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76,
	0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x04, 0x76, 0x61, 0x6c, 0x73, 0x22,
	0x1f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x75, 0x6d,
	0x22, 0x21, 0x0a, 0x0d, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x6d, 0x61, 0x78, 0x22, 0x22, 0x0a, 0x0e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x32, 0x9f, 0x01, 0x0a, 0x04, 0x4d, 0x61, 0x74, 0x68,
	0x12, 0x44, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x12, 0x20, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x13, 0x5a, 0x11, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_rawDescOnce sync.Once
	file_pb_proto_rawDescData = file_pb_proto_rawDesc
)

func file_pb_proto_rawDescGZIP() []byte {
	file_pb_proto_rawDescOnce.Do(func() {
		file_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_rawDescData)
	})
	return file_pb_proto_rawDescData
}

var file_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_proto_goTypes = []interface{}{
	(*AddRequest)(nil),     // 0: reverse_grpc.math.AddRequest
	(*AddResponse)(nil),    // 1: reverse_grpc.math.AddResponse
	(*RandomRequest)(nil),  // 2: reverse_grpc.math.RandomRequest
	(*RandomResponse)(nil), // 3: reverse_grpc.math.RandomResponse
}
var file_pb_proto_depIdxs = []int32{
	0, // 0: reverse_grpc.math.Math.Add:input_type -> reverse_grpc.math.AddRequest
	2, // 1: reverse_grpc.math.Math.Random:input_type -> reverse_grpc.math.RandomRequest
	1, // 2: reverse_grpc.math.Math.Add:output_type -> reverse_grpc.math.AddResponse
	3, // 3: reverse_grpc.math.Math.Random:output_type -> reverse_grpc.math.RandomResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_proto_init() }
func file_pb_proto_init() {
	if File_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomRequest); i {
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
		file_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RandomResponse); i {
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
			RawDescriptor: file_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_proto_goTypes,
		DependencyIndexes: file_pb_proto_depIdxs,
		MessageInfos:      file_pb_proto_msgTypes,
	}.Build()
	File_pb_proto = out.File
	file_pb_proto_rawDesc = nil
	file_pb_proto_goTypes = nil
	file_pb_proto_depIdxs = nil
}
