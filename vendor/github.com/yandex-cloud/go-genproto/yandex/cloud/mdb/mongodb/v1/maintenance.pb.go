// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/mdb/mongodb/v1/maintenance.proto

package mongodb

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type WeeklyMaintenanceWindow_WeekDay int32

const (
	WeeklyMaintenanceWindow_WEEK_DAY_UNSPECIFIED WeeklyMaintenanceWindow_WeekDay = 0
	WeeklyMaintenanceWindow_MON                  WeeklyMaintenanceWindow_WeekDay = 1
	WeeklyMaintenanceWindow_TUE                  WeeklyMaintenanceWindow_WeekDay = 2
	WeeklyMaintenanceWindow_WED                  WeeklyMaintenanceWindow_WeekDay = 3
	WeeklyMaintenanceWindow_THU                  WeeklyMaintenanceWindow_WeekDay = 4
	WeeklyMaintenanceWindow_FRI                  WeeklyMaintenanceWindow_WeekDay = 5
	WeeklyMaintenanceWindow_SAT                  WeeklyMaintenanceWindow_WeekDay = 6
	WeeklyMaintenanceWindow_SUN                  WeeklyMaintenanceWindow_WeekDay = 7
)

// Enum value maps for WeeklyMaintenanceWindow_WeekDay.
var (
	WeeklyMaintenanceWindow_WeekDay_name = map[int32]string{
		0: "WEEK_DAY_UNSPECIFIED",
		1: "MON",
		2: "TUE",
		3: "WED",
		4: "THU",
		5: "FRI",
		6: "SAT",
		7: "SUN",
	}
	WeeklyMaintenanceWindow_WeekDay_value = map[string]int32{
		"WEEK_DAY_UNSPECIFIED": 0,
		"MON":                  1,
		"TUE":                  2,
		"WED":                  3,
		"THU":                  4,
		"FRI":                  5,
		"SAT":                  6,
		"SUN":                  7,
	}
)

func (x WeeklyMaintenanceWindow_WeekDay) Enum() *WeeklyMaintenanceWindow_WeekDay {
	p := new(WeeklyMaintenanceWindow_WeekDay)
	*p = x
	return p
}

func (x WeeklyMaintenanceWindow_WeekDay) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WeeklyMaintenanceWindow_WeekDay) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_enumTypes[0].Descriptor()
}

func (WeeklyMaintenanceWindow_WeekDay) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_enumTypes[0]
}

func (x WeeklyMaintenanceWindow_WeekDay) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WeeklyMaintenanceWindow_WeekDay.Descriptor instead.
func (WeeklyMaintenanceWindow_WeekDay) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescGZIP(), []int{2, 0}
}

// A maintenance window settings.
type MaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maintenance policy in effect.
	//
	// Types that are assignable to Policy:
	//	*MaintenanceWindow_Anytime
	//	*MaintenanceWindow_WeeklyMaintenanceWindow
	Policy isMaintenanceWindow_Policy `protobuf_oneof:"policy"`
}

func (x *MaintenanceWindow) Reset() {
	*x = MaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceWindow) ProtoMessage() {}

func (x *MaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceWindow.ProtoReflect.Descriptor instead.
func (*MaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescGZIP(), []int{0}
}

func (m *MaintenanceWindow) GetPolicy() isMaintenanceWindow_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (x *MaintenanceWindow) GetAnytime() *AnytimeMaintenanceWindow {
	if x, ok := x.GetPolicy().(*MaintenanceWindow_Anytime); ok {
		return x.Anytime
	}
	return nil
}

func (x *MaintenanceWindow) GetWeeklyMaintenanceWindow() *WeeklyMaintenanceWindow {
	if x, ok := x.GetPolicy().(*MaintenanceWindow_WeeklyMaintenanceWindow); ok {
		return x.WeeklyMaintenanceWindow
	}
	return nil
}

type isMaintenanceWindow_Policy interface {
	isMaintenanceWindow_Policy()
}

type MaintenanceWindow_Anytime struct {
	// Maintenance operation can be scheduled anytime.
	Anytime *AnytimeMaintenanceWindow `protobuf:"bytes,1,opt,name=anytime,proto3,oneof"`
}

type MaintenanceWindow_WeeklyMaintenanceWindow struct {
	// Maintenance operation can be scheduled on a weekly basis.
	WeeklyMaintenanceWindow *WeeklyMaintenanceWindow `protobuf:"bytes,2,opt,name=weekly_maintenance_window,json=weeklyMaintenanceWindow,proto3,oneof"`
}

func (*MaintenanceWindow_Anytime) isMaintenanceWindow_Policy() {}

func (*MaintenanceWindow_WeeklyMaintenanceWindow) isMaintenanceWindow_Policy() {}

type AnytimeMaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnytimeMaintenanceWindow) Reset() {
	*x = AnytimeMaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnytimeMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnytimeMaintenanceWindow) ProtoMessage() {}

func (x *AnytimeMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnytimeMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*AnytimeMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescGZIP(), []int{1}
}

// Weelky maintenance window settings.
type WeeklyMaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Day of the week (in `DDD` format).
	Day WeeklyMaintenanceWindow_WeekDay `protobuf:"varint,1,opt,name=day,proto3,enum=yandex.cloud.mdb.mongodb.v1.WeeklyMaintenanceWindow_WeekDay" json:"day,omitempty"`
	// Hour of the day in UTC (in `HH` format).
	Hour int64 `protobuf:"varint,2,opt,name=hour,proto3" json:"hour,omitempty"`
}

func (x *WeeklyMaintenanceWindow) Reset() {
	*x = WeeklyMaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeeklyMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklyMaintenanceWindow) ProtoMessage() {}

func (x *WeeklyMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklyMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*WeeklyMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescGZIP(), []int{2}
}

func (x *WeeklyMaintenanceWindow) GetDay() WeeklyMaintenanceWindow_WeekDay {
	if x != nil {
		return x.Day
	}
	return WeeklyMaintenanceWindow_WEEK_DAY_UNSPECIFIED
}

func (x *WeeklyMaintenanceWindow) GetHour() int64 {
	if x != nil {
		return x.Hour
	}
	return 0
}

// A planned maintenance operation.
type MaintenanceOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information about this maintenance operation.
	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	// Time until which this maintenance operation is delayed.
	DelayedUntil *timestamp.Timestamp `protobuf:"bytes,2,opt,name=delayed_until,json=delayedUntil,proto3" json:"delayed_until,omitempty"`
}

func (x *MaintenanceOperation) Reset() {
	*x = MaintenanceOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintenanceOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceOperation) ProtoMessage() {}

func (x *MaintenanceOperation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceOperation.ProtoReflect.Descriptor instead.
func (*MaintenanceOperation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescGZIP(), []int{3}
}

func (x *MaintenanceOperation) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *MaintenanceOperation) GetDelayedUntil() *timestamp.Timestamp {
	if x != nil {
		return x.DelayedUntil
	}
	return nil
}

var File_yandex_cloud_mdb_mongodb_v1_maintenance_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a,
	0x11, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x12, 0x51, 0x0a, 0x07, 0x61, 0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x48, 0x00, 0x52, 0x07, 0x61, 0x6e,
	0x79, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x72, 0x0a, 0x19, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f,
	0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67,
	0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x4d, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x48, 0x00,
	0x52, 0x17, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x42, 0x0e, 0x0a, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x6e, 0x79,
	0x74, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0xeb, 0x01, 0x0a, 0x17, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79,
	0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x12, 0x4e, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x12, 0x1c, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x08, 0xfa, 0xc7, 0x31, 0x04, 0x31, 0x2d, 0x32, 0x34, 0x52, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x22,
	0x62, 0x0a, 0x07, 0x57, 0x65, 0x65, 0x6b, 0x44, 0x61, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x45,
	0x45, 0x4b, 0x5f, 0x44, 0x41, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x07, 0x0a,
	0x03, 0x54, 0x55, 0x45, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x54, 0x48, 0x55, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x52, 0x49, 0x10,
	0x05, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x41, 0x54, 0x10, 0x06, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x55,
	0x4e, 0x10, 0x07, 0x22, 0x76, 0x0a, 0x14, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c,
	0x3d, 0x32, 0x35, 0x36, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0d, 0x64, 0x65,
	0x6c, 0x61, 0x79, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x65, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x42, 0x6a, 0x0a, 0x1f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x5a, 0x47,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescData = file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDesc
)

func file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDescData
}

var file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_goTypes = []interface{}{
	(WeeklyMaintenanceWindow_WeekDay)(0), // 0: yandex.cloud.mdb.mongodb.v1.WeeklyMaintenanceWindow.WeekDay
	(*MaintenanceWindow)(nil),            // 1: yandex.cloud.mdb.mongodb.v1.MaintenanceWindow
	(*AnytimeMaintenanceWindow)(nil),     // 2: yandex.cloud.mdb.mongodb.v1.AnytimeMaintenanceWindow
	(*WeeklyMaintenanceWindow)(nil),      // 3: yandex.cloud.mdb.mongodb.v1.WeeklyMaintenanceWindow
	(*MaintenanceOperation)(nil),         // 4: yandex.cloud.mdb.mongodb.v1.MaintenanceOperation
	(*timestamp.Timestamp)(nil),          // 5: google.protobuf.Timestamp
}
var file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.mdb.mongodb.v1.MaintenanceWindow.anytime:type_name -> yandex.cloud.mdb.mongodb.v1.AnytimeMaintenanceWindow
	3, // 1: yandex.cloud.mdb.mongodb.v1.MaintenanceWindow.weekly_maintenance_window:type_name -> yandex.cloud.mdb.mongodb.v1.WeeklyMaintenanceWindow
	0, // 2: yandex.cloud.mdb.mongodb.v1.WeeklyMaintenanceWindow.day:type_name -> yandex.cloud.mdb.mongodb.v1.WeeklyMaintenanceWindow.WeekDay
	5, // 3: yandex.cloud.mdb.mongodb.v1.MaintenanceOperation.delayed_until:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_init() }
func file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_init() {
	if File_yandex_cloud_mdb_mongodb_v1_maintenance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintenanceWindow); i {
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
		file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnytimeMaintenanceWindow); i {
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
		file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeeklyMaintenanceWindow); i {
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
		file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintenanceOperation); i {
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
	file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*MaintenanceWindow_Anytime)(nil),
		(*MaintenanceWindow_WeeklyMaintenanceWindow)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_mongodb_v1_maintenance_proto = out.File
	file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_rawDesc = nil
	file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_goTypes = nil
	file_yandex_cloud_mdb_mongodb_v1_maintenance_proto_depIdxs = nil
}
