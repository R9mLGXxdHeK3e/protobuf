// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.12.3
// source: menu_api.proto

package menuv1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type MenusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *MenusRequest) Reset() {
	*x = MenusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenusRequest) ProtoMessage() {}

func (x *MenusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenusRequest.ProtoReflect.Descriptor instead.
func (*MenusRequest) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{0}
}

func (x *MenusRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type MenusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuList []*Menu `protobuf:"bytes,1,rep,name=menu_list,json=menuList,proto3" json:"menu_list,omitempty"`
}

func (x *MenusResponse) Reset() {
	*x = MenusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenusResponse) ProtoMessage() {}

func (x *MenusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenusResponse.ProtoReflect.Descriptor instead.
func (*MenusResponse) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{1}
}

func (x *MenusResponse) GetMenuList() []*Menu {
	if x != nil {
		return x.MenuList
	}
	return nil
}

type MenuDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId string `protobuf:"bytes,1,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
}

func (x *MenuDetailsRequest) Reset() {
	*x = MenuDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuDetailsRequest) ProtoMessage() {}

func (x *MenuDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuDetailsRequest.ProtoReflect.Descriptor instead.
func (*MenuDetailsRequest) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{2}
}

func (x *MenuDetailsRequest) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

type CategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId string `protobuf:"bytes,1,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
}

func (x *CategoriesRequest) Reset() {
	*x = CategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesRequest) ProtoMessage() {}

func (x *CategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesRequest.ProtoReflect.Descriptor instead.
func (*CategoriesRequest) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{3}
}

func (x *CategoriesRequest) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

type CategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryList []*Category `protobuf:"bytes,1,rep,name=category_list,json=categoryList,proto3" json:"category_list,omitempty"`
}

func (x *CategoriesResponse) Reset() {
	*x = CategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoriesResponse) ProtoMessage() {}

func (x *CategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoriesResponse.ProtoReflect.Descriptor instead.
func (*CategoriesResponse) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{4}
}

func (x *CategoriesResponse) GetCategoryList() []*Category {
	if x != nil {
		return x.CategoryList
	}
	return nil
}

// TODO: DELETE IT PROBABLY IN FUTURE
type CategoryDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *CategoryDetailsRequest) Reset() {
	*x = CategoryDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryDetailsRequest) ProtoMessage() {}

func (x *CategoryDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryDetailsRequest.ProtoReflect.Descriptor instead.
func (*CategoryDetailsRequest) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{5}
}

func (x *CategoryDetailsRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type ProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId string `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
}

func (x *ProductsRequest) Reset() {
	*x = ProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsRequest) ProtoMessage() {}

func (x *ProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsRequest.ProtoReflect.Descriptor instead.
func (*ProductsRequest) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{6}
}

func (x *ProductsRequest) GetCategoryId() string {
	if x != nil {
		return x.CategoryId
	}
	return ""
}

type ProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductList []*Product `protobuf:"bytes,1,rep,name=product_list,json=productList,proto3" json:"product_list,omitempty"`
}

func (x *ProductsResponse) Reset() {
	*x = ProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductsResponse) ProtoMessage() {}

func (x *ProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductsResponse.ProtoReflect.Descriptor instead.
func (*ProductsResponse) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{7}
}

func (x *ProductsResponse) GetProductList() []*Product {
	if x != nil {
		return x.ProductList
	}
	return nil
}

type ProductDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *ProductDetailsRequest) Reset() {
	*x = ProductDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menu_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailsRequest) ProtoMessage() {}

func (x *ProductDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menu_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetailsRequest.ProtoReflect.Descriptor instead.
func (*ProductDetailsRequest) Descriptor() ([]byte, []int) {
	return file_menu_api_proto_rawDescGZIP(), []int{8}
}

func (x *ProductDetailsRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

var File_menu_api_proto protoreflect.FileDescriptor

var file_menu_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x1a, 0x12, 0x6d, 0x65, 0x6e, 0x75, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a,
	0x0c, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x6e, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x6d, 0x65, 0x6e,
	0x75, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x08, 0x6d, 0x65, 0x6e,
	0x75, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x12, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d,
	0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x6e, 0x75, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e,
	0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75,
	0x49, 0x64, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x39, 0x0a, 0x16, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22,
	0x47, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x32, 0xa4, 0x03, 0x0a, 0x10, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x05, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x15,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b,
	0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65,
	0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0f, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x6e,
	0x75, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x6e, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x32, 0x6d, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x65, 0x6e, 0x75, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menu_api_proto_rawDescOnce sync.Once
	file_menu_api_proto_rawDescData = file_menu_api_proto_rawDesc
)

func file_menu_api_proto_rawDescGZIP() []byte {
	file_menu_api_proto_rawDescOnce.Do(func() {
		file_menu_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_menu_api_proto_rawDescData)
	})
	return file_menu_api_proto_rawDescData
}

var file_menu_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_menu_api_proto_goTypes = []interface{}{
	(*MenusRequest)(nil),           // 0: menu.v1.MenusRequest
	(*MenusResponse)(nil),          // 1: menu.v1.MenusResponse
	(*MenuDetailsRequest)(nil),     // 2: menu.v1.MenuDetailsRequest
	(*CategoriesRequest)(nil),      // 3: menu.v1.CategoriesRequest
	(*CategoriesResponse)(nil),     // 4: menu.v1.CategoriesResponse
	(*CategoryDetailsRequest)(nil), // 5: menu.v1.CategoryDetailsRequest
	(*ProductsRequest)(nil),        // 6: menu.v1.ProductsRequest
	(*ProductsResponse)(nil),       // 7: menu.v1.ProductsResponse
	(*ProductDetailsRequest)(nil),  // 8: menu.v1.ProductDetailsRequest
	(*Menu)(nil),                   // 9: menu.v1.Menu
	(*Category)(nil),               // 10: menu.v1.Category
	(*Product)(nil),                // 11: menu.v1.Product
}
var file_menu_api_proto_depIdxs = []int32{
	9,  // 0: menu.v1.MenusResponse.menu_list:type_name -> menu.v1.Menu
	10, // 1: menu.v1.CategoriesResponse.category_list:type_name -> menu.v1.Category
	11, // 2: menu.v1.ProductsResponse.product_list:type_name -> menu.v1.Product
	0,  // 3: menu.v1.MenuDataProducer.Menus:input_type -> menu.v1.MenusRequest
	2,  // 4: menu.v1.MenuDataProducer.MenuDetails:input_type -> menu.v1.MenuDetailsRequest
	3,  // 5: menu.v1.MenuDataProducer.Categories:input_type -> menu.v1.CategoriesRequest
	5,  // 6: menu.v1.MenuDataProducer.CategoryDetails:input_type -> menu.v1.CategoryDetailsRequest
	6,  // 7: menu.v1.MenuDataProducer.Products:input_type -> menu.v1.ProductsRequest
	8,  // 8: menu.v1.MenuDataProducer.ProductDetails:input_type -> menu.v1.ProductDetailsRequest
	1,  // 9: menu.v1.MenuDataProducer.Menus:output_type -> menu.v1.MenusResponse
	9,  // 10: menu.v1.MenuDataProducer.MenuDetails:output_type -> menu.v1.Menu
	4,  // 11: menu.v1.MenuDataProducer.Categories:output_type -> menu.v1.CategoriesResponse
	10, // 12: menu.v1.MenuDataProducer.CategoryDetails:output_type -> menu.v1.Category
	7,  // 13: menu.v1.MenuDataProducer.Products:output_type -> menu.v1.ProductsResponse
	11, // 14: menu.v1.MenuDataProducer.ProductDetails:output_type -> menu.v1.Product
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_menu_api_proto_init() }
func file_menu_api_proto_init() {
	if File_menu_api_proto != nil {
		return
	}
	file_menu_v1_menu_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_menu_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenusRequest); i {
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
		file_menu_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenusResponse); i {
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
		file_menu_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuDetailsRequest); i {
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
		file_menu_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesRequest); i {
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
		file_menu_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoriesResponse); i {
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
		file_menu_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryDetailsRequest); i {
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
		file_menu_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsRequest); i {
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
		file_menu_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductsResponse); i {
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
		file_menu_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDetailsRequest); i {
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
			RawDescriptor: file_menu_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menu_api_proto_goTypes,
		DependencyIndexes: file_menu_api_proto_depIdxs,
		MessageInfos:      file_menu_api_proto_msgTypes,
	}.Build()
	File_menu_api_proto = out.File
	file_menu_api_proto_rawDesc = nil
	file_menu_api_proto_goTypes = nil
	file_menu_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MenuDataProducerClient is the client API for MenuDataProducer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MenuDataProducerClient interface {
	// Menu methods
	Menus(ctx context.Context, in *MenusRequest, opts ...grpc.CallOption) (*MenusResponse, error)
	MenuDetails(ctx context.Context, in *MenuDetailsRequest, opts ...grpc.CallOption) (*Menu, error)
	// Category methods
	Categories(ctx context.Context, in *CategoriesRequest, opts ...grpc.CallOption) (*CategoriesResponse, error)
	CategoryDetails(ctx context.Context, in *CategoryDetailsRequest, opts ...grpc.CallOption) (*Category, error)
	// Product methods
	Products(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error)
	ProductDetails(ctx context.Context, in *ProductDetailsRequest, opts ...grpc.CallOption) (*Product, error)
}

type menuDataProducerClient struct {
	cc grpc.ClientConnInterface
}

func NewMenuDataProducerClient(cc grpc.ClientConnInterface) MenuDataProducerClient {
	return &menuDataProducerClient{cc}
}

func (c *menuDataProducerClient) Menus(ctx context.Context, in *MenusRequest, opts ...grpc.CallOption) (*MenusResponse, error) {
	out := new(MenusResponse)
	err := c.cc.Invoke(ctx, "/menu.v1.MenuDataProducer/Menus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuDataProducerClient) MenuDetails(ctx context.Context, in *MenuDetailsRequest, opts ...grpc.CallOption) (*Menu, error) {
	out := new(Menu)
	err := c.cc.Invoke(ctx, "/menu.v1.MenuDataProducer/MenuDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuDataProducerClient) Categories(ctx context.Context, in *CategoriesRequest, opts ...grpc.CallOption) (*CategoriesResponse, error) {
	out := new(CategoriesResponse)
	err := c.cc.Invoke(ctx, "/menu.v1.MenuDataProducer/Categories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuDataProducerClient) CategoryDetails(ctx context.Context, in *CategoryDetailsRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/menu.v1.MenuDataProducer/CategoryDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuDataProducerClient) Products(ctx context.Context, in *ProductsRequest, opts ...grpc.CallOption) (*ProductsResponse, error) {
	out := new(ProductsResponse)
	err := c.cc.Invoke(ctx, "/menu.v1.MenuDataProducer/Products", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuDataProducerClient) ProductDetails(ctx context.Context, in *ProductDetailsRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/menu.v1.MenuDataProducer/ProductDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuDataProducerServer is the server API for MenuDataProducer service.
type MenuDataProducerServer interface {
	// Menu methods
	Menus(context.Context, *MenusRequest) (*MenusResponse, error)
	MenuDetails(context.Context, *MenuDetailsRequest) (*Menu, error)
	// Category methods
	Categories(context.Context, *CategoriesRequest) (*CategoriesResponse, error)
	CategoryDetails(context.Context, *CategoryDetailsRequest) (*Category, error)
	// Product methods
	Products(context.Context, *ProductsRequest) (*ProductsResponse, error)
	ProductDetails(context.Context, *ProductDetailsRequest) (*Product, error)
}

// UnimplementedMenuDataProducerServer can be embedded to have forward compatible implementations.
type UnimplementedMenuDataProducerServer struct {
}

func (*UnimplementedMenuDataProducerServer) Menus(context.Context, *MenusRequest) (*MenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Menus not implemented")
}
func (*UnimplementedMenuDataProducerServer) MenuDetails(context.Context, *MenuDetailsRequest) (*Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuDetails not implemented")
}
func (*UnimplementedMenuDataProducerServer) Categories(context.Context, *CategoriesRequest) (*CategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Categories not implemented")
}
func (*UnimplementedMenuDataProducerServer) CategoryDetails(context.Context, *CategoryDetailsRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryDetails not implemented")
}
func (*UnimplementedMenuDataProducerServer) Products(context.Context, *ProductsRequest) (*ProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Products not implemented")
}
func (*UnimplementedMenuDataProducerServer) ProductDetails(context.Context, *ProductDetailsRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProductDetails not implemented")
}

func RegisterMenuDataProducerServer(s *grpc.Server, srv MenuDataProducerServer) {
	s.RegisterService(&_MenuDataProducer_serviceDesc, srv)
}

func _MenuDataProducer_Menus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuDataProducerServer).Menus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.v1.MenuDataProducer/Menus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuDataProducerServer).Menus(ctx, req.(*MenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuDataProducer_MenuDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuDataProducerServer).MenuDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.v1.MenuDataProducer/MenuDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuDataProducerServer).MenuDetails(ctx, req.(*MenuDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuDataProducer_Categories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuDataProducerServer).Categories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.v1.MenuDataProducer/Categories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuDataProducerServer).Categories(ctx, req.(*CategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuDataProducer_CategoryDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuDataProducerServer).CategoryDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.v1.MenuDataProducer/CategoryDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuDataProducerServer).CategoryDetails(ctx, req.(*CategoryDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuDataProducer_Products_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuDataProducerServer).Products(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.v1.MenuDataProducer/Products",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuDataProducerServer).Products(ctx, req.(*ProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuDataProducer_ProductDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuDataProducerServer).ProductDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/menu.v1.MenuDataProducer/ProductDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuDataProducerServer).ProductDetails(ctx, req.(*ProductDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MenuDataProducer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "menu.v1.MenuDataProducer",
	HandlerType: (*MenuDataProducerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Menus",
			Handler:    _MenuDataProducer_Menus_Handler,
		},
		{
			MethodName: "MenuDetails",
			Handler:    _MenuDataProducer_MenuDetails_Handler,
		},
		{
			MethodName: "Categories",
			Handler:    _MenuDataProducer_Categories_Handler,
		},
		{
			MethodName: "CategoryDetails",
			Handler:    _MenuDataProducer_CategoryDetails_Handler,
		},
		{
			MethodName: "Products",
			Handler:    _MenuDataProducer_Products_Handler,
		},
		{
			MethodName: "ProductDetails",
			Handler:    _MenuDataProducer_ProductDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "menu_api.proto",
}