// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: hlds.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_hlds_proto protoreflect.FileDescriptor

var file_hlds_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x72, 0x75,
	0x2e, 0x74, 0x6f, 0x6e, 0x36, 0x31, 0x38, 0x2e, 0x68, 0x6c, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x68, 0x6c, 0x64, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x81, 0x01, 0x0a,
	0x17, 0x48, 0x61, 0x6c, 0x66, 0x4c, 0x69, 0x66, 0x65, 0x44, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x66, 0x0a, 0x12, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x52, 0x63, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x24,
	0x2e, 0x72, 0x75, 0x2e, 0x74, 0x6f, 0x6e, 0x36, 0x31, 0x38, 0x2e, 0x68, 0x6c, 0x64, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x63, 0x6f, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x2a, 0x2e, 0x72, 0x75, 0x2e, 0x74, 0x6f, 0x6e, 0x36, 0x31, 0x38,
	0x2e, 0x68, 0x6c, 0x64, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x63, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x42, 0x4b, 0x0a, 0x11, 0x72, 0x75, 0x2e, 0x74, 0x6f, 0x6e, 0x36, 0x31, 0x38, 0x2e, 0x68, 0x6c,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x48, 0x61, 0x6c, 0x66, 0x4c, 0x69, 0x66, 0x65, 0x44,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x01,
	0x5a, 0x1b, 0x72, 0x75, 0x2e, 0x74, 0x6f, 0x6e, 0x36, 0x31, 0x38, 0x2e, 0x68, 0x6c, 0x64, 0x73,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_hlds_proto_goTypes = []interface{}{
	(*RconCommand)(nil),       // 0: ru.ton618.hlds.model.v1.RconCommand
	(*RconCommandResult)(nil), // 1: ru.ton618.hlds.model.v1.RconCommandResult
}
var file_hlds_proto_depIdxs = []int32{
	0, // 0: ru.ton618.hlds.service.v1.HalfLifeDedicatedServer.ExecuteRconCommand:input_type -> ru.ton618.hlds.model.v1.RconCommand
	1, // 1: ru.ton618.hlds.service.v1.HalfLifeDedicatedServer.ExecuteRconCommand:output_type -> ru.ton618.hlds.model.v1.RconCommandResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hlds_proto_init() }
func file_hlds_proto_init() {
	if File_hlds_proto != nil {
		return
	}
	file_hlds_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_hlds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hlds_proto_goTypes,
		DependencyIndexes: file_hlds_proto_depIdxs,
	}.Build()
	File_hlds_proto = out.File
	file_hlds_proto_rawDesc = nil
	file_hlds_proto_goTypes = nil
	file_hlds_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HalfLifeDedicatedServerClient is the client API for HalfLifeDedicatedServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HalfLifeDedicatedServerClient interface {
	ExecuteRconCommand(ctx context.Context, in *RconCommand, opts ...grpc.CallOption) (*RconCommandResult, error)
}

type halfLifeDedicatedServerClient struct {
	cc grpc.ClientConnInterface
}

func NewHalfLifeDedicatedServerClient(cc grpc.ClientConnInterface) HalfLifeDedicatedServerClient {
	return &halfLifeDedicatedServerClient{cc}
}

func (c *halfLifeDedicatedServerClient) ExecuteRconCommand(ctx context.Context, in *RconCommand, opts ...grpc.CallOption) (*RconCommandResult, error) {
	out := new(RconCommandResult)
	err := c.cc.Invoke(ctx, "/ru.ton618.hlds.service.v1.HalfLifeDedicatedServer/ExecuteRconCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HalfLifeDedicatedServerServer is the server API for HalfLifeDedicatedServer service.
type HalfLifeDedicatedServerServer interface {
	ExecuteRconCommand(context.Context, *RconCommand) (*RconCommandResult, error)
}

// UnimplementedHalfLifeDedicatedServerServer can be embedded to have forward compatible implementations.
type UnimplementedHalfLifeDedicatedServerServer struct {
}

func (*UnimplementedHalfLifeDedicatedServerServer) ExecuteRconCommand(context.Context, *RconCommand) (*RconCommandResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteRconCommand not implemented")
}

func RegisterHalfLifeDedicatedServerServer(s *grpc.Server, srv HalfLifeDedicatedServerServer) {
	s.RegisterService(&_HalfLifeDedicatedServer_serviceDesc, srv)
}

func _HalfLifeDedicatedServer_ExecuteRconCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RconCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HalfLifeDedicatedServerServer).ExecuteRconCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ru.ton618.hlds.service.v1.HalfLifeDedicatedServer/ExecuteRconCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HalfLifeDedicatedServerServer).ExecuteRconCommand(ctx, req.(*RconCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _HalfLifeDedicatedServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ru.ton618.hlds.service.v1.HalfLifeDedicatedServer",
	HandlerType: (*HalfLifeDedicatedServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteRconCommand",
			Handler:    _HalfLifeDedicatedServer_ExecuteRconCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hlds.proto",
}
