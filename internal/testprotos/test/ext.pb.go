// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/test/ext.proto

package test

import (
	protoreflect "github.com/xjz9600/protobuf/reflect/protoreflect"
	protoimpl "github.com/xjz9600/protobuf/runtime/protoimpl"
	reflect "reflect"
)

var file_internal_testprotos_test_ext_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*TestAllExtensions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         2000,
		Name:          "goproto.proto.test.foreign_int32_extension",
		Tag:           "varint,2000,opt,name=foreign_int32_extension",
		Filename:      "internal/testprotos/test/ext.proto",
	},
}

// Extension fields to TestAllExtensions.
var (
	// optional int32 foreign_int32_extension = 2000;
	E_ForeignInt32Extension = &file_internal_testprotos_test_ext_proto_extTypes[0]
)

var File_internal_testprotos_test_ext_proto protoreflect.FileDescriptor

var file_internal_testprotos_test_ext_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x5e, 0x0a,
	0x17, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd0, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74,
}

var file_internal_testprotos_test_ext_proto_goTypes = []interface{}{
	(*TestAllExtensions)(nil), // 0: goproto.proto.test.TestAllExtensions
}
var file_internal_testprotos_test_ext_proto_depIdxs = []int32{
	0, // 0: goproto.proto.test.foreign_int32_extension:extendee -> goproto.proto.test.TestAllExtensions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_test_ext_proto_init() }
func file_internal_testprotos_test_ext_proto_init() {
	if File_internal_testprotos_test_ext_proto != nil {
		return
	}
	file_internal_testprotos_test_test_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_test_ext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_test_ext_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_test_ext_proto_depIdxs,
		ExtensionInfos:    file_internal_testprotos_test_ext_proto_extTypes,
	}.Build()
	File_internal_testprotos_test_ext_proto = out.File
	file_internal_testprotos_test_ext_proto_rawDesc = nil
	file_internal_testprotos_test_ext_proto_goTypes = nil
	file_internal_testprotos_test_ext_proto_depIdxs = nil
}
