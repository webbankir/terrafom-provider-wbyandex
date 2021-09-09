// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/mdb/kafka/v1/common.proto

package kafka

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

type CompressionType int32

const (
	CompressionType_COMPRESSION_TYPE_UNSPECIFIED CompressionType = 0
	// no codec (uncompressed).
	CompressionType_COMPRESSION_TYPE_UNCOMPRESSED CompressionType = 1
	// Zstandard codec.
	CompressionType_COMPRESSION_TYPE_ZSTD CompressionType = 2
	// LZ4 codec.
	CompressionType_COMPRESSION_TYPE_LZ4 CompressionType = 3
	// Snappy codec.
	CompressionType_COMPRESSION_TYPE_SNAPPY CompressionType = 4
	// GZip codec.
	CompressionType_COMPRESSION_TYPE_GZIP CompressionType = 5
	// the codec to use is set by a producer (can be any of `ZSTD`, `LZ4`, `GZIP` or `SNAPPY` codecs).
	CompressionType_COMPRESSION_TYPE_PRODUCER CompressionType = 6
)

// Enum value maps for CompressionType.
var (
	CompressionType_name = map[int32]string{
		0: "COMPRESSION_TYPE_UNSPECIFIED",
		1: "COMPRESSION_TYPE_UNCOMPRESSED",
		2: "COMPRESSION_TYPE_ZSTD",
		3: "COMPRESSION_TYPE_LZ4",
		4: "COMPRESSION_TYPE_SNAPPY",
		5: "COMPRESSION_TYPE_GZIP",
		6: "COMPRESSION_TYPE_PRODUCER",
	}
	CompressionType_value = map[string]int32{
		"COMPRESSION_TYPE_UNSPECIFIED":  0,
		"COMPRESSION_TYPE_UNCOMPRESSED": 1,
		"COMPRESSION_TYPE_ZSTD":         2,
		"COMPRESSION_TYPE_LZ4":          3,
		"COMPRESSION_TYPE_SNAPPY":       4,
		"COMPRESSION_TYPE_GZIP":         5,
		"COMPRESSION_TYPE_PRODUCER":     6,
	}
)

func (x CompressionType) Enum() *CompressionType {
	p := new(CompressionType)
	*p = x
	return p
}

func (x CompressionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompressionType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_kafka_v1_common_proto_enumTypes[0].Descriptor()
}

func (CompressionType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_kafka_v1_common_proto_enumTypes[0]
}

func (x CompressionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompressionType.Descriptor instead.
func (CompressionType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescGZIP(), []int{0}
}

var File_yandex_cloud_mdb_kafka_v1_common_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_kafka_v1_common_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x2e, 0x76, 0x31, 0x2a, 0xe2, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x4f, 0x4d, 0x50, 0x52,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d,
	0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15,
	0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x5a, 0x53, 0x54, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4f, 0x4d, 0x50, 0x52,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x5a, 0x34, 0x10,
	0x03, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x50, 0x59, 0x10, 0x04, 0x12, 0x19,
	0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x47, 0x5a, 0x49, 0x50, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d,
	0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x44, 0x55, 0x43, 0x45, 0x52, 0x10, 0x06, 0x42, 0x64, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62,
	0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f,
	0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescData = file_yandex_cloud_mdb_kafka_v1_common_proto_rawDesc
)

func file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_kafka_v1_common_proto_rawDescData
}

var file_yandex_cloud_mdb_kafka_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_kafka_v1_common_proto_goTypes = []interface{}{
	(CompressionType)(0), // 0: yandex.cloud.mdb.kafka.v1.CompressionType
}
var file_yandex_cloud_mdb_kafka_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_kafka_v1_common_proto_init() }
func file_yandex_cloud_mdb_kafka_v1_common_proto_init() {
	if File_yandex_cloud_mdb_kafka_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_mdb_kafka_v1_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_kafka_v1_common_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_kafka_v1_common_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_kafka_v1_common_proto_enumTypes,
	}.Build()
	File_yandex_cloud_mdb_kafka_v1_common_proto = out.File
	file_yandex_cloud_mdb_kafka_v1_common_proto_rawDesc = nil
	file_yandex_cloud_mdb_kafka_v1_common_proto_goTypes = nil
	file_yandex_cloud_mdb_kafka_v1_common_proto_depIdxs = nil
}
