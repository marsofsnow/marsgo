//*
// Copyright (C) 2014 Open Whisper Systems
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: PubSubMessage.proto

package textsecure

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

type PubSubMessage_Type int32

const (
	PubSubMessage_UNKNOWN   PubSubMessage_Type = 0
	PubSubMessage_QUERY_DB  PubSubMessage_Type = 1
	PubSubMessage_DELIVER   PubSubMessage_Type = 2 //递送
	PubSubMessage_KEEPALIVE PubSubMessage_Type = 3
	PubSubMessage_CLOSE     PubSubMessage_Type = 4
	PubSubMessage_CONNECTED PubSubMessage_Type = 5
)

// Enum value maps for PubSubMessage_Type.
var (
	PubSubMessage_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "QUERY_DB",
		2: "DELIVER",
		3: "KEEPALIVE",
		4: "CLOSE",
		5: "CONNECTED",
	}
	PubSubMessage_Type_value = map[string]int32{
		"UNKNOWN":   0,
		"QUERY_DB":  1,
		"DELIVER":   2,
		"KEEPALIVE": 3,
		"CLOSE":     4,
		"CONNECTED": 5,
	}
)

func (x PubSubMessage_Type) Enum() *PubSubMessage_Type {
	p := new(PubSubMessage_Type)
	*p = x
	return p
}

func (x PubSubMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PubSubMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_PubSubMessage_proto_enumTypes[0].Descriptor()
}

func (PubSubMessage_Type) Type() protoreflect.EnumType {
	return &file_PubSubMessage_proto_enumTypes[0]
}

func (x PubSubMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PubSubMessage_Type.Descriptor instead.
func (PubSubMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_PubSubMessage_proto_rawDescGZIP(), []int{0, 0}
}

// 订阅消息
type PubSubMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    PubSubMessage_Type `protobuf:"varint,1,opt,name=type,proto3,enum=textsecure.PubSubMessage_Type" json:"type,omitempty"`
	Content []byte             `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PubSubMessage) Reset() {
	*x = PubSubMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PubSubMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubSubMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubSubMessage) ProtoMessage() {}

func (x *PubSubMessage) ProtoReflect() protoreflect.Message {
	mi := &file_PubSubMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubSubMessage.ProtoReflect.Descriptor instead.
func (*PubSubMessage) Descriptor() ([]byte, []int) {
	return file_PubSubMessage_proto_rawDescGZIP(), []int{0}
}

func (x *PubSubMessage) GetType() PubSubMessage_Type {
	if x != nil {
		return x.Type
	}
	return PubSubMessage_UNKNOWN
}

func (x *PubSubMessage) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_PubSubMessage_proto protoreflect.FileDescriptor

var file_PubSubMessage_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x65, 0x22, 0xb6, 0x01, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x2e, 0x50,
	0x75, 0x62, 0x53, 0x75, 0x62, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x57, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f,
	0x44, 0x42, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x45, 0x45, 0x50, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x05, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_PubSubMessage_proto_rawDescOnce sync.Once
	file_PubSubMessage_proto_rawDescData = file_PubSubMessage_proto_rawDesc
)

func file_PubSubMessage_proto_rawDescGZIP() []byte {
	file_PubSubMessage_proto_rawDescOnce.Do(func() {
		file_PubSubMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_PubSubMessage_proto_rawDescData)
	})
	return file_PubSubMessage_proto_rawDescData
}

var file_PubSubMessage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PubSubMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PubSubMessage_proto_goTypes = []interface{}{
	(PubSubMessage_Type)(0), // 0: textsecure.PubSubMessage.Type
	(*PubSubMessage)(nil),   // 1: textsecure.PubSubMessage
}
var file_PubSubMessage_proto_depIdxs = []int32{
	0, // 0: textsecure.PubSubMessage.type:type_name -> textsecure.PubSubMessage.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_PubSubMessage_proto_init() }
func file_PubSubMessage_proto_init() {
	if File_PubSubMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PubSubMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubSubMessage); i {
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
			RawDescriptor: file_PubSubMessage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PubSubMessage_proto_goTypes,
		DependencyIndexes: file_PubSubMessage_proto_depIdxs,
		EnumInfos:         file_PubSubMessage_proto_enumTypes,
		MessageInfos:      file_PubSubMessage_proto_msgTypes,
	}.Build()
	File_PubSubMessage_proto = out.File
	file_PubSubMessage_proto_rawDesc = nil
	file_PubSubMessage_proto_goTypes = nil
	file_PubSubMessage_proto_depIdxs = nil
}