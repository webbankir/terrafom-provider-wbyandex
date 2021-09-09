// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/ydb/v1/location_service.proto

package ydb

import (
	context "context"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the location to return.
	LocationId string `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *GetLocationRequest) Reset() {
	*x = GetLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationRequest) ProtoMessage() {}

func (x *GetLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationRequest.ProtoReflect.Descriptor instead.
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_location_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetLocationRequest) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

type ListLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than `page_size`, the service returns a `next_page_token` that can be used
	// to get the next page of results in subsequent ListLocations requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set `page_token` to the `next_page_token` returned by a previous ListLocations
	// request to get the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListLocationsRequest) Reset() {
	*x = ListLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsRequest) ProtoMessage() {}

func (x *ListLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsRequest.ProtoReflect.Descriptor instead.
func (*ListLocationsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_location_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListLocationsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListLocationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Requested list of locations.
	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
	// This token allows you to get the next page of results for ListLocations requests,
	// if the number of results is larger than `page_size` specified in the request.
	// To get the next page, specify the value of `next_page_token` as a value for
	// the `page_token` parameter in the next ListLocations request. Subsequent ListLocations
	// requests will have their own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListLocationsResponse) Reset() {
	*x = ListLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsResponse) ProtoMessage() {}

func (x *ListLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsResponse.ProtoReflect.Descriptor instead.
func (*ListLocationsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_location_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListLocationsResponse) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *ListLocationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_ydb_v1_location_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ydb_v1_location_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79,
	0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79, 0x64,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8,
	0xc7, 0x31, 0x01, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x69, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06,
	0x30, 0x2d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7c, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x83, 0x02, 0x0a, 0x0f, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x79, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x12, 0x78, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x79, 0x64,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x56,
	0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x79, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x79, 0x64, 0x62, 0x2f,
	0x76, 0x31, 0x3b, 0x79, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ydb_v1_location_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ydb_v1_location_service_proto_rawDescData = file_yandex_cloud_ydb_v1_location_service_proto_rawDesc
)

func file_yandex_cloud_ydb_v1_location_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ydb_v1_location_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ydb_v1_location_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ydb_v1_location_service_proto_rawDescData)
	})
	return file_yandex_cloud_ydb_v1_location_service_proto_rawDescData
}

var file_yandex_cloud_ydb_v1_location_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_ydb_v1_location_service_proto_goTypes = []interface{}{
	(*GetLocationRequest)(nil),    // 0: yandex.cloud.ydb.v1.GetLocationRequest
	(*ListLocationsRequest)(nil),  // 1: yandex.cloud.ydb.v1.ListLocationsRequest
	(*ListLocationsResponse)(nil), // 2: yandex.cloud.ydb.v1.ListLocationsResponse
	(*Location)(nil),              // 3: yandex.cloud.ydb.v1.Location
}
var file_yandex_cloud_ydb_v1_location_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.ydb.v1.ListLocationsResponse.locations:type_name -> yandex.cloud.ydb.v1.Location
	0, // 1: yandex.cloud.ydb.v1.LocationService.Get:input_type -> yandex.cloud.ydb.v1.GetLocationRequest
	1, // 2: yandex.cloud.ydb.v1.LocationService.List:input_type -> yandex.cloud.ydb.v1.ListLocationsRequest
	3, // 3: yandex.cloud.ydb.v1.LocationService.Get:output_type -> yandex.cloud.ydb.v1.Location
	2, // 4: yandex.cloud.ydb.v1.LocationService.List:output_type -> yandex.cloud.ydb.v1.ListLocationsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ydb_v1_location_service_proto_init() }
func file_yandex_cloud_ydb_v1_location_service_proto_init() {
	if File_yandex_cloud_ydb_v1_location_service_proto != nil {
		return
	}
	file_yandex_cloud_ydb_v1_location_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationRequest); i {
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
		file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocationsRequest); i {
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
		file_yandex_cloud_ydb_v1_location_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocationsResponse); i {
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
			RawDescriptor: file_yandex_cloud_ydb_v1_location_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ydb_v1_location_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ydb_v1_location_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ydb_v1_location_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ydb_v1_location_service_proto = out.File
	file_yandex_cloud_ydb_v1_location_service_proto_rawDesc = nil
	file_yandex_cloud_ydb_v1_location_service_proto_goTypes = nil
	file_yandex_cloud_ydb_v1_location_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocationServiceClient interface {
	// Returns the specified location.
	Get(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*Location, error)
	// Returns the list of available locations.
	List(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) Get(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*Location, error) {
	out := new(Location)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.LocationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) List(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	out := new(ListLocationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.LocationService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationServiceServer is the server API for LocationService service.
type LocationServiceServer interface {
	// Returns the specified location.
	Get(context.Context, *GetLocationRequest) (*Location, error)
	// Returns the list of available locations.
	List(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
}

// UnimplementedLocationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLocationServiceServer struct {
}

func (*UnimplementedLocationServiceServer) Get(context.Context, *GetLocationRequest) (*Location, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedLocationServiceServer) List(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterLocationServiceServer(s *grpc.Server, srv LocationServiceServer) {
	s.RegisterService(&_LocationService_serviceDesc, srv)
}

func _LocationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.LocationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).Get(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.LocationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).List(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ydb.v1.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _LocationService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LocationService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ydb/v1/location_service.proto",
}
