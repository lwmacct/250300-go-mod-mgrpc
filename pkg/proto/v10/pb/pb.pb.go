// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: pkg/proto/v10/pb/pb.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type KvBool struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           bool                   `protobuf:"varint,2,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KvBool) Reset() {
	*x = KvBool{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KvBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvBool) ProtoMessage() {}

func (x *KvBool) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvBool.ProtoReflect.Descriptor instead.
func (*KvBool) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{0}
}

func (x *KvBool) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvBool) GetVal() bool {
	if x != nil {
		return x.Val
	}
	return false
}

type KvBytes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           []byte                 `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KvBytes) Reset() {
	*x = KvBytes{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KvBytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvBytes) ProtoMessage() {}

func (x *KvBytes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvBytes.ProtoReflect.Descriptor instead.
func (*KvBytes) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{1}
}

func (x *KvBytes) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvBytes) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type KvInt64 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           int64                  `protobuf:"varint,2,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KvInt64) Reset() {
	*x = KvInt64{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KvInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvInt64) ProtoMessage() {}

func (x *KvInt64) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvInt64.ProtoReflect.Descriptor instead.
func (*KvInt64) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{2}
}

func (x *KvInt64) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvInt64) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type KvFloat64 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           float64                `protobuf:"fixed64,2,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KvFloat64) Reset() {
	*x = KvFloat64{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KvFloat64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvFloat64) ProtoMessage() {}

func (x *KvFloat64) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvFloat64.ProtoReflect.Descriptor instead.
func (*KvFloat64) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{3}
}

func (x *KvFloat64) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvFloat64) GetVal() float64 {
	if x != nil {
		return x.Val
	}
	return 0
}

type KvString struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           string                 `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KvString) Reset() {
	*x = KvString{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KvString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvString) ProtoMessage() {}

func (x *KvString) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvString.ProtoReflect.Descriptor instead.
func (*KvString) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{4}
}

func (x *KvString) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvString) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type KvStringArray struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           []string               `protobuf:"bytes,2,rep,name=val,proto3" json:"val,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KvStringArray) Reset() {
	*x = KvStringArray{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KvStringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KvStringArray) ProtoMessage() {}

func (x *KvStringArray) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KvStringArray.ProtoReflect.Descriptor instead.
func (*KvStringArray) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{5}
}

func (x *KvStringArray) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KvStringArray) GetVal() []string {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapStringBool struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           map[string]bool        `protobuf:"bytes,2,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapStringBool) Reset() {
	*x = MapStringBool{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapStringBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapStringBool) ProtoMessage() {}

func (x *MapStringBool) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapStringBool.ProtoReflect.Descriptor instead.
func (*MapStringBool) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{6}
}

func (x *MapStringBool) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapStringBool) GetVal() map[string]bool {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapStringBytes struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           map[string][]byte      `protobuf:"bytes,2,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapStringBytes) Reset() {
	*x = MapStringBytes{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapStringBytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapStringBytes) ProtoMessage() {}

func (x *MapStringBytes) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapStringBytes.ProtoReflect.Descriptor instead.
func (*MapStringBytes) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{7}
}

func (x *MapStringBytes) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapStringBytes) GetVal() map[string][]byte {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapStringInt64 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           map[string]int64       `protobuf:"bytes,2,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapStringInt64) Reset() {
	*x = MapStringInt64{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapStringInt64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapStringInt64) ProtoMessage() {}

func (x *MapStringInt64) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapStringInt64.ProtoReflect.Descriptor instead.
func (*MapStringInt64) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{8}
}

func (x *MapStringInt64) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapStringInt64) GetVal() map[string]int64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapStringFloat64 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           map[string]float64     `protobuf:"bytes,2,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapStringFloat64) Reset() {
	*x = MapStringFloat64{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapStringFloat64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapStringFloat64) ProtoMessage() {}

func (x *MapStringFloat64) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapStringFloat64.ProtoReflect.Descriptor instead.
func (*MapStringFloat64) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{9}
}

func (x *MapStringFloat64) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapStringFloat64) GetVal() map[string]float64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapStringString struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val           map[string]string      `protobuf:"bytes,2,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapStringString) Reset() {
	*x = MapStringString{}
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapStringString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapStringString) ProtoMessage() {}

func (x *MapStringString) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_v10_pb_pb_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapStringString.ProtoReflect.Descriptor instead.
func (*MapStringString) Descriptor() ([]byte, []int) {
	return file_pkg_proto_v10_pb_pb_proto_rawDescGZIP(), []int{10}
}

func (x *MapStringString) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *MapStringString) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_pkg_proto_v10_pb_pb_proto protoreflect.FileDescriptor

const file_pkg_proto_v10_pb_pb_proto_rawDesc = "" +
	"\n" +
	"\x19pkg/proto/v10/pb/pb.proto\x12\tmgrpc.v10\x1a\x19google/protobuf/any.proto\",\n" +
	"\x06KvBool\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x10\n" +
	"\x03val\x18\x02 \x01(\bR\x03val\"-\n" +
	"\aKvBytes\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x10\n" +
	"\x03val\x18\x02 \x01(\fR\x03val\"-\n" +
	"\aKvInt64\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x10\n" +
	"\x03val\x18\x02 \x01(\x03R\x03val\"/\n" +
	"\tKvFloat64\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x10\n" +
	"\x03val\x18\x02 \x01(\x01R\x03val\".\n" +
	"\bKvString\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x10\n" +
	"\x03val\x18\x02 \x01(\tR\x03val\"3\n" +
	"\rKvStringArray\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x10\n" +
	"\x03val\x18\x02 \x03(\tR\x03val\"\x8e\x01\n" +
	"\rMapStringBool\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x123\n" +
	"\x03val\x18\x02 \x03(\v2!.mgrpc.v10.MapStringBool.ValEntryR\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\bR\x05value:\x028\x01\"\x90\x01\n" +
	"\x0eMapStringBytes\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x124\n" +
	"\x03val\x18\x02 \x03(\v2\".mgrpc.v10.MapStringBytes.ValEntryR\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\fR\x05value:\x028\x01\"\x90\x01\n" +
	"\x0eMapStringInt64\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x124\n" +
	"\x03val\x18\x02 \x03(\v2\".mgrpc.v10.MapStringInt64.ValEntryR\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value:\x028\x01\"\x94\x01\n" +
	"\x10MapStringFloat64\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x126\n" +
	"\x03val\x18\x02 \x03(\v2$.mgrpc.v10.MapStringFloat64.ValEntryR\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x01R\x05value:\x028\x01\"\x92\x01\n" +
	"\x0fMapStringString\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x125\n" +
	"\x03val\x18\x02 \x03(\v2#.mgrpc.v10.MapStringString.ValEntryR\x03val\x1a6\n" +
	"\bValEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x012\x95\x12\n" +
	"\x03Rpc\x121\n" +
	"\x03Any\x12\x14.google.protobuf.Any\x1a\x14.google.protobuf.Any\x127\n" +
	"\fInt64AskBool\x12\x12.mgrpc.v10.KvInt64\x1a\x11.mgrpc.v10.KvBool\"\x00\x129\n" +
	"\rInt64AskBytes\x12\x12.mgrpc.v10.KvInt64\x1a\x12.mgrpc.v10.KvBytes\"\x00\x129\n" +
	"\rInt64AskInt64\x12\x12.mgrpc.v10.KvInt64\x1a\x12.mgrpc.v10.KvInt64\"\x00\x12=\n" +
	"\x0fInt64AskFloat64\x12\x12.mgrpc.v10.KvInt64\x1a\x14.mgrpc.v10.KvFloat64\"\x00\x12;\n" +
	"\x0eInt64AskString\x12\x12.mgrpc.v10.KvInt64\x1a\x13.mgrpc.v10.KvString\"\x00\x12E\n" +
	"\x13Int64AskStringArray\x12\x12.mgrpc.v10.KvInt64\x1a\x18.mgrpc.v10.KvStringArray\"\x00\x12G\n" +
	"\x15Int64AskMapStringBool\x12\x12.mgrpc.v10.KvInt64\x1a\x18.mgrpc.v10.MapStringBool\"\x00\x12I\n" +
	"\x16Int64AskMapStringInt64\x12\x12.mgrpc.v10.KvInt64\x1a\x19.mgrpc.v10.MapStringInt64\"\x00\x12M\n" +
	"\x18Int64AskMapStringFloat64\x12\x12.mgrpc.v10.KvInt64\x1a\x1b.mgrpc.v10.MapStringFloat64\"\x00\x12K\n" +
	"\x17Int64AskMapStringString\x12\x12.mgrpc.v10.KvInt64\x1a\x1a.mgrpc.v10.MapStringString\"\x00\x12I\n" +
	"\x16Int64AskMapStringBytes\x12\x12.mgrpc.v10.KvInt64\x1a\x19.mgrpc.v10.MapStringBytes\"\x00\x129\n" +
	"\rStringAskBool\x12\x13.mgrpc.v10.KvString\x1a\x11.mgrpc.v10.KvBool\"\x00\x12;\n" +
	"\x0eStringAskBytes\x12\x13.mgrpc.v10.KvString\x1a\x12.mgrpc.v10.KvBytes\"\x00\x12;\n" +
	"\x0eStringAskInt64\x12\x13.mgrpc.v10.KvString\x1a\x12.mgrpc.v10.KvInt64\"\x00\x12?\n" +
	"\x10StringAskFloat64\x12\x13.mgrpc.v10.KvString\x1a\x14.mgrpc.v10.KvFloat64\"\x00\x12=\n" +
	"\x0fStringAskString\x12\x13.mgrpc.v10.KvString\x1a\x13.mgrpc.v10.KvString\"\x00\x12G\n" +
	"\x14StringAskStringArray\x12\x13.mgrpc.v10.KvString\x1a\x18.mgrpc.v10.KvStringArray\"\x00\x12I\n" +
	"\x16StringAskMapStringBool\x12\x13.mgrpc.v10.KvString\x1a\x18.mgrpc.v10.MapStringBool\"\x00\x12K\n" +
	"\x17StringAskMapStringInt64\x12\x13.mgrpc.v10.KvString\x1a\x19.mgrpc.v10.MapStringInt64\"\x00\x12O\n" +
	"\x19StringAskMapStringFloat64\x12\x13.mgrpc.v10.KvString\x1a\x1b.mgrpc.v10.MapStringFloat64\"\x00\x12M\n" +
	"\x18StringAskMapStringString\x12\x13.mgrpc.v10.KvString\x1a\x1a.mgrpc.v10.MapStringString\"\x00\x12K\n" +
	"\x17StringAskMapStringBytes\x12\x13.mgrpc.v10.KvString\x1a\x19.mgrpc.v10.MapStringBytes\"\x00\x127\n" +
	"\fBytesAskBool\x12\x12.mgrpc.v10.KvBytes\x1a\x11.mgrpc.v10.KvBool\"\x00\x129\n" +
	"\rBytesAskBytes\x12\x12.mgrpc.v10.KvBytes\x1a\x12.mgrpc.v10.KvBytes\"\x00\x129\n" +
	"\rBytesAskInt64\x12\x12.mgrpc.v10.KvBytes\x1a\x12.mgrpc.v10.KvInt64\"\x00\x12=\n" +
	"\x0fBytesAskFloat64\x12\x12.mgrpc.v10.KvBytes\x1a\x14.mgrpc.v10.KvFloat64\"\x00\x12;\n" +
	"\x0eBytesAskString\x12\x12.mgrpc.v10.KvBytes\x1a\x13.mgrpc.v10.KvString\"\x00\x12E\n" +
	"\x13BytesAskStringArray\x12\x12.mgrpc.v10.KvBytes\x1a\x18.mgrpc.v10.KvStringArray\"\x00\x12G\n" +
	"\x15BytesAskMapStringBool\x12\x12.mgrpc.v10.KvBytes\x1a\x18.mgrpc.v10.MapStringBool\"\x00\x12I\n" +
	"\x16BytesAskMapStringBytes\x12\x12.mgrpc.v10.KvBytes\x1a\x19.mgrpc.v10.MapStringBytes\"\x00\x12I\n" +
	"\x16BytesAskMapStringInt64\x12\x12.mgrpc.v10.KvBytes\x1a\x19.mgrpc.v10.MapStringInt64\"\x00\x12M\n" +
	"\x18BytesAskMapStringFloat64\x12\x12.mgrpc.v10.KvBytes\x1a\x1b.mgrpc.v10.MapStringFloat64\"\x00\x12K\n" +
	"\x17BytesAskMapStringString\x12\x12.mgrpc.v10.KvBytes\x1a\x1a.mgrpc.v10.MapStringString\"\x00B\x06Z\x04.;pbb\x06proto3"

var (
	file_pkg_proto_v10_pb_pb_proto_rawDescOnce sync.Once
	file_pkg_proto_v10_pb_pb_proto_rawDescData []byte
)

func file_pkg_proto_v10_pb_pb_proto_rawDescGZIP() []byte {
	file_pkg_proto_v10_pb_pb_proto_rawDescOnce.Do(func() {
		file_pkg_proto_v10_pb_pb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_v10_pb_pb_proto_rawDesc), len(file_pkg_proto_v10_pb_pb_proto_rawDesc)))
	})
	return file_pkg_proto_v10_pb_pb_proto_rawDescData
}

var file_pkg_proto_v10_pb_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_pkg_proto_v10_pb_pb_proto_goTypes = []any{
	(*KvBool)(nil),           // 0: mgrpc.v10.KvBool
	(*KvBytes)(nil),          // 1: mgrpc.v10.KvBytes
	(*KvInt64)(nil),          // 2: mgrpc.v10.KvInt64
	(*KvFloat64)(nil),        // 3: mgrpc.v10.KvFloat64
	(*KvString)(nil),         // 4: mgrpc.v10.KvString
	(*KvStringArray)(nil),    // 5: mgrpc.v10.KvStringArray
	(*MapStringBool)(nil),    // 6: mgrpc.v10.MapStringBool
	(*MapStringBytes)(nil),   // 7: mgrpc.v10.MapStringBytes
	(*MapStringInt64)(nil),   // 8: mgrpc.v10.MapStringInt64
	(*MapStringFloat64)(nil), // 9: mgrpc.v10.MapStringFloat64
	(*MapStringString)(nil),  // 10: mgrpc.v10.MapStringString
	nil,                      // 11: mgrpc.v10.MapStringBool.ValEntry
	nil,                      // 12: mgrpc.v10.MapStringBytes.ValEntry
	nil,                      // 13: mgrpc.v10.MapStringInt64.ValEntry
	nil,                      // 14: mgrpc.v10.MapStringFloat64.ValEntry
	nil,                      // 15: mgrpc.v10.MapStringString.ValEntry
	(*anypb.Any)(nil),        // 16: google.protobuf.Any
}
var file_pkg_proto_v10_pb_pb_proto_depIdxs = []int32{
	11, // 0: mgrpc.v10.MapStringBool.val:type_name -> mgrpc.v10.MapStringBool.ValEntry
	12, // 1: mgrpc.v10.MapStringBytes.val:type_name -> mgrpc.v10.MapStringBytes.ValEntry
	13, // 2: mgrpc.v10.MapStringInt64.val:type_name -> mgrpc.v10.MapStringInt64.ValEntry
	14, // 3: mgrpc.v10.MapStringFloat64.val:type_name -> mgrpc.v10.MapStringFloat64.ValEntry
	15, // 4: mgrpc.v10.MapStringString.val:type_name -> mgrpc.v10.MapStringString.ValEntry
	16, // 5: mgrpc.v10.Rpc.Any:input_type -> google.protobuf.Any
	2,  // 6: mgrpc.v10.Rpc.Int64AskBool:input_type -> mgrpc.v10.KvInt64
	2,  // 7: mgrpc.v10.Rpc.Int64AskBytes:input_type -> mgrpc.v10.KvInt64
	2,  // 8: mgrpc.v10.Rpc.Int64AskInt64:input_type -> mgrpc.v10.KvInt64
	2,  // 9: mgrpc.v10.Rpc.Int64AskFloat64:input_type -> mgrpc.v10.KvInt64
	2,  // 10: mgrpc.v10.Rpc.Int64AskString:input_type -> mgrpc.v10.KvInt64
	2,  // 11: mgrpc.v10.Rpc.Int64AskStringArray:input_type -> mgrpc.v10.KvInt64
	2,  // 12: mgrpc.v10.Rpc.Int64AskMapStringBool:input_type -> mgrpc.v10.KvInt64
	2,  // 13: mgrpc.v10.Rpc.Int64AskMapStringInt64:input_type -> mgrpc.v10.KvInt64
	2,  // 14: mgrpc.v10.Rpc.Int64AskMapStringFloat64:input_type -> mgrpc.v10.KvInt64
	2,  // 15: mgrpc.v10.Rpc.Int64AskMapStringString:input_type -> mgrpc.v10.KvInt64
	2,  // 16: mgrpc.v10.Rpc.Int64AskMapStringBytes:input_type -> mgrpc.v10.KvInt64
	4,  // 17: mgrpc.v10.Rpc.StringAskBool:input_type -> mgrpc.v10.KvString
	4,  // 18: mgrpc.v10.Rpc.StringAskBytes:input_type -> mgrpc.v10.KvString
	4,  // 19: mgrpc.v10.Rpc.StringAskInt64:input_type -> mgrpc.v10.KvString
	4,  // 20: mgrpc.v10.Rpc.StringAskFloat64:input_type -> mgrpc.v10.KvString
	4,  // 21: mgrpc.v10.Rpc.StringAskString:input_type -> mgrpc.v10.KvString
	4,  // 22: mgrpc.v10.Rpc.StringAskStringArray:input_type -> mgrpc.v10.KvString
	4,  // 23: mgrpc.v10.Rpc.StringAskMapStringBool:input_type -> mgrpc.v10.KvString
	4,  // 24: mgrpc.v10.Rpc.StringAskMapStringInt64:input_type -> mgrpc.v10.KvString
	4,  // 25: mgrpc.v10.Rpc.StringAskMapStringFloat64:input_type -> mgrpc.v10.KvString
	4,  // 26: mgrpc.v10.Rpc.StringAskMapStringString:input_type -> mgrpc.v10.KvString
	4,  // 27: mgrpc.v10.Rpc.StringAskMapStringBytes:input_type -> mgrpc.v10.KvString
	1,  // 28: mgrpc.v10.Rpc.BytesAskBool:input_type -> mgrpc.v10.KvBytes
	1,  // 29: mgrpc.v10.Rpc.BytesAskBytes:input_type -> mgrpc.v10.KvBytes
	1,  // 30: mgrpc.v10.Rpc.BytesAskInt64:input_type -> mgrpc.v10.KvBytes
	1,  // 31: mgrpc.v10.Rpc.BytesAskFloat64:input_type -> mgrpc.v10.KvBytes
	1,  // 32: mgrpc.v10.Rpc.BytesAskString:input_type -> mgrpc.v10.KvBytes
	1,  // 33: mgrpc.v10.Rpc.BytesAskStringArray:input_type -> mgrpc.v10.KvBytes
	1,  // 34: mgrpc.v10.Rpc.BytesAskMapStringBool:input_type -> mgrpc.v10.KvBytes
	1,  // 35: mgrpc.v10.Rpc.BytesAskMapStringBytes:input_type -> mgrpc.v10.KvBytes
	1,  // 36: mgrpc.v10.Rpc.BytesAskMapStringInt64:input_type -> mgrpc.v10.KvBytes
	1,  // 37: mgrpc.v10.Rpc.BytesAskMapStringFloat64:input_type -> mgrpc.v10.KvBytes
	1,  // 38: mgrpc.v10.Rpc.BytesAskMapStringString:input_type -> mgrpc.v10.KvBytes
	16, // 39: mgrpc.v10.Rpc.Any:output_type -> google.protobuf.Any
	0,  // 40: mgrpc.v10.Rpc.Int64AskBool:output_type -> mgrpc.v10.KvBool
	1,  // 41: mgrpc.v10.Rpc.Int64AskBytes:output_type -> mgrpc.v10.KvBytes
	2,  // 42: mgrpc.v10.Rpc.Int64AskInt64:output_type -> mgrpc.v10.KvInt64
	3,  // 43: mgrpc.v10.Rpc.Int64AskFloat64:output_type -> mgrpc.v10.KvFloat64
	4,  // 44: mgrpc.v10.Rpc.Int64AskString:output_type -> mgrpc.v10.KvString
	5,  // 45: mgrpc.v10.Rpc.Int64AskStringArray:output_type -> mgrpc.v10.KvStringArray
	6,  // 46: mgrpc.v10.Rpc.Int64AskMapStringBool:output_type -> mgrpc.v10.MapStringBool
	8,  // 47: mgrpc.v10.Rpc.Int64AskMapStringInt64:output_type -> mgrpc.v10.MapStringInt64
	9,  // 48: mgrpc.v10.Rpc.Int64AskMapStringFloat64:output_type -> mgrpc.v10.MapStringFloat64
	10, // 49: mgrpc.v10.Rpc.Int64AskMapStringString:output_type -> mgrpc.v10.MapStringString
	7,  // 50: mgrpc.v10.Rpc.Int64AskMapStringBytes:output_type -> mgrpc.v10.MapStringBytes
	0,  // 51: mgrpc.v10.Rpc.StringAskBool:output_type -> mgrpc.v10.KvBool
	1,  // 52: mgrpc.v10.Rpc.StringAskBytes:output_type -> mgrpc.v10.KvBytes
	2,  // 53: mgrpc.v10.Rpc.StringAskInt64:output_type -> mgrpc.v10.KvInt64
	3,  // 54: mgrpc.v10.Rpc.StringAskFloat64:output_type -> mgrpc.v10.KvFloat64
	4,  // 55: mgrpc.v10.Rpc.StringAskString:output_type -> mgrpc.v10.KvString
	5,  // 56: mgrpc.v10.Rpc.StringAskStringArray:output_type -> mgrpc.v10.KvStringArray
	6,  // 57: mgrpc.v10.Rpc.StringAskMapStringBool:output_type -> mgrpc.v10.MapStringBool
	8,  // 58: mgrpc.v10.Rpc.StringAskMapStringInt64:output_type -> mgrpc.v10.MapStringInt64
	9,  // 59: mgrpc.v10.Rpc.StringAskMapStringFloat64:output_type -> mgrpc.v10.MapStringFloat64
	10, // 60: mgrpc.v10.Rpc.StringAskMapStringString:output_type -> mgrpc.v10.MapStringString
	7,  // 61: mgrpc.v10.Rpc.StringAskMapStringBytes:output_type -> mgrpc.v10.MapStringBytes
	0,  // 62: mgrpc.v10.Rpc.BytesAskBool:output_type -> mgrpc.v10.KvBool
	1,  // 63: mgrpc.v10.Rpc.BytesAskBytes:output_type -> mgrpc.v10.KvBytes
	2,  // 64: mgrpc.v10.Rpc.BytesAskInt64:output_type -> mgrpc.v10.KvInt64
	3,  // 65: mgrpc.v10.Rpc.BytesAskFloat64:output_type -> mgrpc.v10.KvFloat64
	4,  // 66: mgrpc.v10.Rpc.BytesAskString:output_type -> mgrpc.v10.KvString
	5,  // 67: mgrpc.v10.Rpc.BytesAskStringArray:output_type -> mgrpc.v10.KvStringArray
	6,  // 68: mgrpc.v10.Rpc.BytesAskMapStringBool:output_type -> mgrpc.v10.MapStringBool
	7,  // 69: mgrpc.v10.Rpc.BytesAskMapStringBytes:output_type -> mgrpc.v10.MapStringBytes
	8,  // 70: mgrpc.v10.Rpc.BytesAskMapStringInt64:output_type -> mgrpc.v10.MapStringInt64
	9,  // 71: mgrpc.v10.Rpc.BytesAskMapStringFloat64:output_type -> mgrpc.v10.MapStringFloat64
	10, // 72: mgrpc.v10.Rpc.BytesAskMapStringString:output_type -> mgrpc.v10.MapStringString
	39, // [39:73] is the sub-list for method output_type
	5,  // [5:39] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_proto_v10_pb_pb_proto_init() }
func file_pkg_proto_v10_pb_pb_proto_init() {
	if File_pkg_proto_v10_pb_pb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_v10_pb_pb_proto_rawDesc), len(file_pkg_proto_v10_pb_pb_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_v10_pb_pb_proto_goTypes,
		DependencyIndexes: file_pkg_proto_v10_pb_pb_proto_depIdxs,
		MessageInfos:      file_pkg_proto_v10_pb_pb_proto_msgTypes,
	}.Build()
	File_pkg_proto_v10_pb_pb_proto = out.File
	file_pkg_proto_v10_pb_pb_proto_goTypes = nil
	file_pkg_proto_v10_pb_pb_proto_depIdxs = nil
}
