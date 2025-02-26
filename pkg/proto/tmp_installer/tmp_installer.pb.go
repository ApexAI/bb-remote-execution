// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: pkg/proto/tmp_installer/tmp_installer.proto

package tmp_installer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type InstallTemporaryDirectoryRequest struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	TemporaryDirectory string                 `protobuf:"bytes,1,opt,name=temporary_directory,json=temporaryDirectory,proto3" json:"temporary_directory,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *InstallTemporaryDirectoryRequest) Reset() {
	*x = InstallTemporaryDirectoryRequest{}
	mi := &file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstallTemporaryDirectoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstallTemporaryDirectoryRequest) ProtoMessage() {}

func (x *InstallTemporaryDirectoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstallTemporaryDirectoryRequest.ProtoReflect.Descriptor instead.
func (*InstallTemporaryDirectoryRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescGZIP(), []int{0}
}

func (x *InstallTemporaryDirectoryRequest) GetTemporaryDirectory() string {
	if x != nil {
		return x.TemporaryDirectory
	}
	return ""
}

var File_pkg_proto_tmp_installer_tmp_installer_proto protoreflect.FileDescriptor

var file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6d, 0x70, 0x5f,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x20, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x54, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x32, 0xcf, 0x01, 0x0a, 0x1b, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x6e, 0x0a, 0x19, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x72, 0x79, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x39, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62,
	0x61, 0x72, 0x6e, 0x2e, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x72, 0x79, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2d, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescOnce sync.Once
	file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescData []byte
)

func file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescGZIP() []byte {
	file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescOnce.Do(func() {
		file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc), len(file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc)))
	})
	return file_pkg_proto_tmp_installer_tmp_installer_proto_rawDescData
}

var file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_tmp_installer_tmp_installer_proto_goTypes = []any{
	(*InstallTemporaryDirectoryRequest)(nil), // 0: buildbarn.tmp_installer.InstallTemporaryDirectoryRequest
	(*emptypb.Empty)(nil),                    // 1: google.protobuf.Empty
}
var file_pkg_proto_tmp_installer_tmp_installer_proto_depIdxs = []int32{
	1, // 0: buildbarn.tmp_installer.TemporaryDirectoryInstaller.CheckReadiness:input_type -> google.protobuf.Empty
	0, // 1: buildbarn.tmp_installer.TemporaryDirectoryInstaller.InstallTemporaryDirectory:input_type -> buildbarn.tmp_installer.InstallTemporaryDirectoryRequest
	1, // 2: buildbarn.tmp_installer.TemporaryDirectoryInstaller.CheckReadiness:output_type -> google.protobuf.Empty
	1, // 3: buildbarn.tmp_installer.TemporaryDirectoryInstaller.InstallTemporaryDirectory:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_tmp_installer_tmp_installer_proto_init() }
func file_pkg_proto_tmp_installer_tmp_installer_proto_init() {
	if File_pkg_proto_tmp_installer_tmp_installer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc), len(file_pkg_proto_tmp_installer_tmp_installer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_tmp_installer_tmp_installer_proto_goTypes,
		DependencyIndexes: file_pkg_proto_tmp_installer_tmp_installer_proto_depIdxs,
		MessageInfos:      file_pkg_proto_tmp_installer_tmp_installer_proto_msgTypes,
	}.Build()
	File_pkg_proto_tmp_installer_tmp_installer_proto = out.File
	file_pkg_proto_tmp_installer_tmp_installer_proto_goTypes = nil
	file_pkg_proto_tmp_installer_tmp_installer_proto_depIdxs = nil
}
