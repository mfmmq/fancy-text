// Recommendations storage and management

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: pkg/proto/text.proto

package text

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

type FancyText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *FancyText) Reset() {
	*x = FancyText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_text_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FancyText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FancyText) ProtoMessage() {}

func (x *FancyText) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_text_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FancyText.ProtoReflect.Descriptor instead.
func (*FancyText) Descriptor() ([]byte, []int) {
	return file_pkg_proto_text_proto_rawDescGZIP(), []int{0}
}

func (x *FancyText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_pkg_proto_text_proto protoreflect.FileDescriptor

var file_pkg_proto_text_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x1f, 0x0a, 0x09,
	0x46, 0x61, 0x6e, 0x63, 0x79, 0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x32, 0x48, 0x0a,
	0x12, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x32, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x63, 0x79, 0x54,
	0x65, 0x78, 0x74, 0x12, 0x0f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x46, 0x61, 0x6e, 0x63, 0x79,
	0x54, 0x65, 0x78, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x46, 0x61, 0x6e, 0x63,
	0x79, 0x54, 0x65, 0x78, 0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_text_proto_rawDescOnce sync.Once
	file_pkg_proto_text_proto_rawDescData = file_pkg_proto_text_proto_rawDesc
)

func file_pkg_proto_text_proto_rawDescGZIP() []byte {
	file_pkg_proto_text_proto_rawDescOnce.Do(func() {
		file_pkg_proto_text_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_text_proto_rawDescData)
	})
	return file_pkg_proto_text_proto_rawDescData
}

var file_pkg_proto_text_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_text_proto_goTypes = []interface{}{
	(*FancyText)(nil), // 0: text.FancyText
}
var file_pkg_proto_text_proto_depIdxs = []int32{
	0, // 0: text.FancyTextGenerator.GetFancyText:input_type -> text.FancyText
	0, // 1: text.FancyTextGenerator.GetFancyText:output_type -> text.FancyText
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_text_proto_init() }
func file_pkg_proto_text_proto_init() {
	if File_pkg_proto_text_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_text_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FancyText); i {
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
			RawDescriptor: file_pkg_proto_text_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_text_proto_goTypes,
		DependencyIndexes: file_pkg_proto_text_proto_depIdxs,
		MessageInfos:      file_pkg_proto_text_proto_msgTypes,
	}.Build()
	File_pkg_proto_text_proto = out.File
	file_pkg_proto_text_proto_rawDesc = nil
	file_pkg_proto_text_proto_goTypes = nil
	file_pkg_proto_text_proto_depIdxs = nil
}
