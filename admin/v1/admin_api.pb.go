// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.12.3
// source: admin_api.proto

package adminv1

import (
	context "context"
	menuv1 "github.com/2menus/protobuf/menuv1"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	FullName string `protobuf:"bytes,3,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Country  string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *CreateMenuRequest) Reset() {
	*x = CreateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMenuRequest) ProtoMessage() {}

func (x *CreateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMenuRequest.ProtoReflect.Descriptor instead.
func (*CreateMenuRequest) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMenuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateMenuRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMenuRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateMenuRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type UpdateMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Country  string `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *UpdateMenuRequest) Reset() {
	*x = UpdateMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMenuRequest) ProtoMessage() {}

func (x *UpdateMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMenuRequest.ProtoReflect.Descriptor instead.
func (*UpdateMenuRequest) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateMenuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMenuRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *UpdateMenuRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type DeleteMenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMenuRequest) Reset() {
	*x = DeleteMenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMenuRequest) ProtoMessage() {}

func (x *DeleteMenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMenuRequest.ProtoReflect.Descriptor instead.
func (*DeleteMenuRequest) Descriptor() ([]byte, []int) {
	return file_admin_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMenuRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_admin_api_proto protoreflect.FileDescriptor

var file_admin_api_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x6d, 0x65, 0x6e,
	0x75, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x5a, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xc4, 0x01,
	0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x3a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e,
	0x75, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e,
	0x75, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1b, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x32, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_admin_api_proto_rawDescOnce sync.Once
	file_admin_api_proto_rawDescData = file_admin_api_proto_rawDesc
)

func file_admin_api_proto_rawDescGZIP() []byte {
	file_admin_api_proto_rawDescOnce.Do(func() {
		file_admin_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_api_proto_rawDescData)
	})
	return file_admin_api_proto_rawDescData
}

var file_admin_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_admin_api_proto_goTypes = []interface{}{
	(*CreateMenuRequest)(nil), // 0: admin.v1.CreateMenuRequest
	(*UpdateMenuRequest)(nil), // 1: admin.v1.UpdateMenuRequest
	(*DeleteMenuRequest)(nil), // 2: admin.v1.DeleteMenuRequest
	(*menuv1.Menu)(nil),       // 3: menu.v1.Menu
	(*empty.Empty)(nil),       // 4: google.protobuf.Empty
}
var file_admin_api_proto_depIdxs = []int32{
	0, // 0: admin.v1.Admin.CreateMenu:input_type -> admin.v1.CreateMenuRequest
	1, // 1: admin.v1.Admin.UpdateMenu:input_type -> admin.v1.UpdateMenuRequest
	2, // 2: admin.v1.Admin.DeleteMenu:input_type -> admin.v1.DeleteMenuRequest
	3, // 3: admin.v1.Admin.CreateMenu:output_type -> menu.v1.Menu
	3, // 4: admin.v1.Admin.UpdateMenu:output_type -> menu.v1.Menu
	4, // 5: admin.v1.Admin.DeleteMenu:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_api_proto_init() }
func file_admin_api_proto_init() {
	if File_admin_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMenuRequest); i {
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
		file_admin_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMenuRequest); i {
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
		file_admin_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMenuRequest); i {
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
			RawDescriptor: file_admin_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_api_proto_goTypes,
		DependencyIndexes: file_admin_api_proto_depIdxs,
		MessageInfos:      file_admin_api_proto_msgTypes,
	}.Build()
	File_admin_api_proto = out.File
	file_admin_api_proto_rawDesc = nil
	file_admin_api_proto_goTypes = nil
	file_admin_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*menuv1.Menu, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*menuv1.Menu, error)
	DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type adminClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClient(cc grpc.ClientConnInterface) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*menuv1.Menu, error) {
	out := new(menuv1.Menu)
	err := c.cc.Invoke(ctx, "/admin.v1.Admin/CreateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*menuv1.Menu, error) {
	out := new(menuv1.Menu)
	err := c.cc.Invoke(ctx, "/admin.v1.Admin/UpdateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/admin.v1.Admin/DeleteMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	CreateMenu(context.Context, *CreateMenuRequest) (*menuv1.Menu, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*menuv1.Menu, error)
	DeleteMenu(context.Context, *DeleteMenuRequest) (*empty.Empty, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) CreateMenu(context.Context, *CreateMenuRequest) (*menuv1.Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (*UnimplementedAdminServer) UpdateMenu(context.Context, *UpdateMenuRequest) (*menuv1.Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (*UnimplementedAdminServer) DeleteMenu(context.Context, *DeleteMenuRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenu not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.v1.Admin/CreateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).CreateMenu(ctx, req.(*CreateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.v1.Admin/UpdateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_DeleteMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).DeleteMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin.v1.Admin/DeleteMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).DeleteMenu(ctx, req.(*DeleteMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "admin.v1.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMenu",
			Handler:    _Admin_CreateMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _Admin_UpdateMenu_Handler,
		},
		{
			MethodName: "DeleteMenu",
			Handler:    _Admin_DeleteMenu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin_api.proto",
}