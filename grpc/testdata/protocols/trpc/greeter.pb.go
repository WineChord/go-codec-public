//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 THL A29 Limited, a Tencent company.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: greeter.proto

package greeter

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "trpc.group/trpc-go/trpc-codec/grpc/testdata/protocols/common"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_greeter_proto protoreflect.FileDescriptor

var file_greeter_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4a,
	0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x05, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e,
	0x74, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x6f, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x63,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_greeter_proto_goTypes = []interface{}{
	(*common.HelloReq)(nil), // 0: trpc.app.common.HelloReq
	(*common.HelloRsp)(nil), // 1: trpc.app.common.HelloRsp
}
var file_greeter_proto_depIdxs = []int32{
	0, // 0: trpc.app.server.Greeter.Hello:input_type -> trpc.app.common.HelloReq
	1, // 1: trpc.app.server.Greeter.Hello:output_type -> trpc.app.common.HelloRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greeter_proto_init() }
func file_greeter_proto_init() {
	if File_greeter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_greeter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greeter_proto_goTypes,
		DependencyIndexes: file_greeter_proto_depIdxs,
	}.Build()
	File_greeter_proto = out.File
	file_greeter_proto_rawDesc = nil
	file_greeter_proto_goTypes = nil
	file_greeter_proto_depIdxs = nil
}
