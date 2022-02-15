// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: test.proto

package internal

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1 string `protobuf:"bytes,1,opt,name=field1,proto3" json:"field1,omitempty"`
}

func (x *TestMessage1) Reset() {
	*x = TestMessage1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage1) ProtoMessage() {}

func (x *TestMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage1.ProtoReflect.Descriptor instead.
func (*TestMessage1) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage1) GetField1() string {
	if x != nil {
		return x.Field1
	}
	return ""
}

type TestMessage2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field1    string      `protobuf:"bytes,1,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2    int32       `protobuf:"varint,2,opt,name=field2,proto3" json:"field2,omitempty"`
	Field3Sub *SubMessage `protobuf:"bytes,3,opt,name=field3_sub,json=field3Sub,proto3" json:"field3_sub,omitempty"`
}

func (x *TestMessage2) Reset() {
	*x = TestMessage2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage2) ProtoMessage() {}

func (x *TestMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage2.ProtoReflect.Descriptor instead.
func (*TestMessage2) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestMessage2) GetField1() string {
	if x != nil {
		return x.Field1
	}
	return ""
}

func (x *TestMessage2) GetField2() int32 {
	if x != nil {
		return x.Field2
	}
	return 0
}

func (x *TestMessage2) GetField3Sub() *SubMessage {
	if x != nil {
		return x.Field3Sub
	}
	return nil
}

type SubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubField  string  `protobuf:"bytes,1,opt,name=sub_field,json=subField,proto3" json:"sub_field,omitempty"`
	SubRepeat []int32 `protobuf:"varint,2,rep,packed,name=sub_repeat,json=subRepeat,proto3" json:"sub_repeat,omitempty"`
}

func (x *SubMessage) Reset() {
	*x = SubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubMessage) ProtoMessage() {}

func (x *SubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubMessage.ProtoReflect.Descriptor instead.
func (*SubMessage) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{2}
}

func (x *SubMessage) GetSubField() string {
	if x != nil {
		return x.SubField
	}
	return ""
}

func (x *SubMessage) GetSubRepeat() []int32 {
	if x != nil {
		return x.SubRepeat
	}
	return nil
}

type TestMessage3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weird_FieldName_1_ string        `protobuf:"bytes,1,opt,name=weird_FieldName_1_,json=weirdFieldName1,proto3" json:"weird_FieldName_1_,omitempty"`
	ABool2             bool          `protobuf:"varint,2,opt,name=a_bool2,json=aBool2,proto3" json:"a_bool2,omitempty"`
	AInt_3             int32         `protobuf:"varint,3,opt,name=a_int_3,json=aInt3,proto3" json:"a_int_3,omitempty"`
	ARepeat            []string      `protobuf:"bytes,4,rep,name=a_repeat,json=aRepeat,proto3" json:"a_repeat,omitempty"`
	ASubmsgRepeat      []*SubMessage `protobuf:"bytes,5,rep,name=a_submsg_repeat,json=aSubmsgRepeat,proto3" json:"a_submsg_repeat,omitempty"`
}

func (x *TestMessage3) Reset() {
	*x = TestMessage3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage3) ProtoMessage() {}

func (x *TestMessage3) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage3.ProtoReflect.Descriptor instead.
func (*TestMessage3) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{3}
}

func (x *TestMessage3) GetWeird_FieldName_1_() string {
	if x != nil {
		return x.Weird_FieldName_1_
	}
	return ""
}

func (x *TestMessage3) GetABool2() bool {
	if x != nil {
		return x.ABool2
	}
	return false
}

func (x *TestMessage3) GetAInt_3() int32 {
	if x != nil {
		return x.AInt_3
	}
	return 0
}

func (x *TestMessage3) GetARepeat() []string {
	if x != nil {
		return x.ARepeat
	}
	return nil
}

func (x *TestMessage3) GetASubmsgRepeat() []*SubMessage {
	if x != nil {
		return x.ASubmsgRepeat
	}
	return nil
}

type TestMessage4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ABool       bool            `protobuf:"varint,1,opt,name=a_bool,json=aBool,proto3" json:"a_bool,omitempty"`
	AInt32      int32           `protobuf:"varint,2,opt,name=a_int32,json=aInt32,proto3" json:"a_int32,omitempty"`
	ASint32     int32           `protobuf:"zigzag32,3,opt,name=a_sint32,json=aSint32,proto3" json:"a_sint32,omitempty"`
	ASfixed32   int32           `protobuf:"fixed32,4,opt,name=a_sfixed32,json=aSfixed32,proto3" json:"a_sfixed32,omitempty"`
	AUint32     uint32          `protobuf:"varint,5,opt,name=a_uint32,json=aUint32,proto3" json:"a_uint32,omitempty"`
	AFixed32    uint32          `protobuf:"fixed32,6,opt,name=a_fixed32,json=aFixed32,proto3" json:"a_fixed32,omitempty"`
	AInt64      int64           `protobuf:"varint,7,opt,name=a_int64,json=aInt64,proto3" json:"a_int64,omitempty"`
	ASint64     int64           `protobuf:"zigzag64,8,opt,name=a_sint64,json=aSint64,proto3" json:"a_sint64,omitempty"`
	ASfixed64   int64           `protobuf:"fixed64,9,opt,name=a_sfixed64,json=aSfixed64,proto3" json:"a_sfixed64,omitempty"`
	AUint64     uint64          `protobuf:"varint,10,opt,name=a_uint64,json=aUint64,proto3" json:"a_uint64,omitempty"`
	AFixed64    uint64          `protobuf:"fixed64,11,opt,name=a_fixed64,json=aFixed64,proto3" json:"a_fixed64,omitempty"`
	AFloat      float32         `protobuf:"fixed32,12,opt,name=a_float,json=aFloat,proto3" json:"a_float,omitempty"`
	ADouble     float64         `protobuf:"fixed64,13,opt,name=a_double,json=aDouble,proto3" json:"a_double,omitempty"`
	AString     string          `protobuf:"bytes,14,opt,name=a_string,json=aString,proto3" json:"a_string,omitempty"`
	ABytes      []byte          `protobuf:"bytes,15,opt,name=a_bytes,json=aBytes,proto3" json:"a_bytes,omitempty"`
	AStringList []string        `protobuf:"bytes,16,rep,name=a_string_list,json=aStringList,proto3" json:"a_string_list,omitempty"`
	AMessage    *TestMessage1   `protobuf:"bytes,17,opt,name=a_message,json=aMessage,proto3" json:"a_message,omitempty"`
	AMap        map[string]bool `protobuf:"bytes,18,rep,name=a_map,json=aMap,proto3" json:"a_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *TestMessage4) Reset() {
	*x = TestMessage4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage4) ProtoMessage() {}

func (x *TestMessage4) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage4.ProtoReflect.Descriptor instead.
func (*TestMessage4) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{4}
}

func (x *TestMessage4) GetABool() bool {
	if x != nil {
		return x.ABool
	}
	return false
}

func (x *TestMessage4) GetAInt32() int32 {
	if x != nil {
		return x.AInt32
	}
	return 0
}

func (x *TestMessage4) GetASint32() int32 {
	if x != nil {
		return x.ASint32
	}
	return 0
}

func (x *TestMessage4) GetASfixed32() int32 {
	if x != nil {
		return x.ASfixed32
	}
	return 0
}

func (x *TestMessage4) GetAUint32() uint32 {
	if x != nil {
		return x.AUint32
	}
	return 0
}

func (x *TestMessage4) GetAFixed32() uint32 {
	if x != nil {
		return x.AFixed32
	}
	return 0
}

func (x *TestMessage4) GetAInt64() int64 {
	if x != nil {
		return x.AInt64
	}
	return 0
}

func (x *TestMessage4) GetASint64() int64 {
	if x != nil {
		return x.ASint64
	}
	return 0
}

func (x *TestMessage4) GetASfixed64() int64 {
	if x != nil {
		return x.ASfixed64
	}
	return 0
}

func (x *TestMessage4) GetAUint64() uint64 {
	if x != nil {
		return x.AUint64
	}
	return 0
}

func (x *TestMessage4) GetAFixed64() uint64 {
	if x != nil {
		return x.AFixed64
	}
	return 0
}

func (x *TestMessage4) GetAFloat() float32 {
	if x != nil {
		return x.AFloat
	}
	return 0
}

func (x *TestMessage4) GetADouble() float64 {
	if x != nil {
		return x.ADouble
	}
	return 0
}

func (x *TestMessage4) GetAString() string {
	if x != nil {
		return x.AString
	}
	return ""
}

func (x *TestMessage4) GetABytes() []byte {
	if x != nil {
		return x.ABytes
	}
	return nil
}

func (x *TestMessage4) GetAStringList() []string {
	if x != nil {
		return x.AStringList
	}
	return nil
}

func (x *TestMessage4) GetAMessage() *TestMessage1 {
	if x != nil {
		return x.AMessage
	}
	return nil
}

func (x *TestMessage4) GetAMap() map[string]bool {
	if x != nil {
		return x.AMap
	}
	return nil
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0c,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x31, 0x22, 0x6a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x32, 0x12, 0x2a, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x5f, 0x73,
	0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x75, 0x62, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x33, 0x53, 0x75, 0x62,
	0x22, 0x48, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x75, 0x62, 0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x75, 0x62, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x12, 0x2b, 0x0a, 0x12, 0x77,
	0x65, 0x69, 0x72, 0x64, 0x5f, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x31,
	0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x65, 0x69, 0x72, 0x64, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x31, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x5f, 0x62, 0x6f,
	0x6f, 0x6c, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x42, 0x6f, 0x6f, 0x6c,
	0x32, 0x12, 0x16, 0x0a, 0x07, 0x61, 0x5f, 0x69, 0x6e, 0x74, 0x5f, 0x33, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x61, 0x49, 0x6e, 0x74, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x52, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x12, 0x33, 0x0a, 0x0f, 0x61, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x73, 0x67,
	0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x61, 0x53, 0x75, 0x62,
	0x6d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x22, 0xda, 0x04, 0x0a, 0x0c, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x34, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x5f,
	0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x42, 0x6f, 0x6f,
	0x6c, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x11, 0x52, 0x07, 0x61, 0x53,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x09, 0x61, 0x53, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x33, 0x32, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x07, 0x52, 0x08, 0x61, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x12, 0x52, 0x07, 0x61, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x10, 0x52, 0x09, 0x61, 0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x61, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x5f,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x06, 0x52, 0x08, 0x61,
	0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x5f, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x07, 0x61, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x22, 0x0a, 0x0d, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x61, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x31, 0x52, 0x08, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2c, 0x0a, 0x05, 0x61, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x34, 0x2e, 0x41, 0x4d,
	0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x4d, 0x61, 0x70, 0x1a, 0x37, 0x0a,
	0x09, 0x41, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x66, 0x6f, 0x78, 0x79, 0x67, 0x6f,
	0x2e, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x72,
	0x75, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_test_proto_goTypes = []interface{}{
	(*TestMessage1)(nil), // 0: TestMessage1
	(*TestMessage2)(nil), // 1: TestMessage2
	(*SubMessage)(nil),   // 2: SubMessage
	(*TestMessage3)(nil), // 3: TestMessage3
	(*TestMessage4)(nil), // 4: TestMessage4
	nil,                  // 5: TestMessage4.AMapEntry
}
var file_test_proto_depIdxs = []int32{
	2, // 0: TestMessage2.field3_sub:type_name -> SubMessage
	2, // 1: TestMessage3.a_submsg_repeat:type_name -> SubMessage
	0, // 2: TestMessage4.a_message:type_name -> TestMessage1
	5, // 3: TestMessage4.a_map:type_name -> TestMessage4.AMapEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage1); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage2); i {
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
		file_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubMessage); i {
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
		file_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage3); i {
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
		file_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage4); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
