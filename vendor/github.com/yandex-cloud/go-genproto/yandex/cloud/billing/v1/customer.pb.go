// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/billing/v1/customer.proto

package billing

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

// A Customer resource.
type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the customer.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the [yandex.cloud.billing.v1.BillingAccount] assigned to the customer.
	BillingAccountId string `protobuf:"bytes,2,opt,name=billing_account_id,json=billingAccountId,proto3" json:"billing_account_id,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetBillingAccountId() string {
	if x != nil {
		return x.BillingAccountId
	}
	return ""
}

// Person of the customer. Contains legal information.
type CustomerPerson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the person.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Long name of the person.
	Longname string `protobuf:"bytes,2,opt,name=longname,proto3" json:"longname,omitempty"`
	// Phone of the person.
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	// Email of the person.
	Email string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// Post code of the person.
	PostCode string `protobuf:"bytes,5,opt,name=post_code,json=postCode,proto3" json:"post_code,omitempty"`
	// Post address of the person.
	PostAddress string `protobuf:"bytes,6,opt,name=post_address,json=postAddress,proto3" json:"post_address,omitempty"`
	// Legal address of the person.
	LegalAddress string `protobuf:"bytes,7,opt,name=legal_address,json=legalAddress,proto3" json:"legal_address,omitempty"`
	// Tax identification number of the person.
	Tin string `protobuf:"bytes,8,opt,name=tin,proto3" json:"tin,omitempty"`
}

func (x *CustomerPerson) Reset() {
	*x = CustomerPerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerPerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerPerson) ProtoMessage() {}

func (x *CustomerPerson) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerPerson.ProtoReflect.Descriptor instead.
func (*CustomerPerson) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerPerson) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerPerson) GetLongname() string {
	if x != nil {
		return x.Longname
	}
	return ""
}

func (x *CustomerPerson) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CustomerPerson) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CustomerPerson) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *CustomerPerson) GetPostAddress() string {
	if x != nil {
		return x.PostAddress
	}
	return ""
}

func (x *CustomerPerson) GetLegalAddress() string {
	if x != nil {
		return x.LegalAddress
	}
	return ""
}

func (x *CustomerPerson) GetTin() string {
	if x != nil {
		return x.Tin
	}
	return ""
}

var File_yandex_cloud_billing_v1_customer_proto protoreflect.FileDescriptor

var file_yandex_cloud_billing_v1_customer_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x22, 0x48, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xe3, 0x01, 0x0a, 0x0e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x6e, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x6f, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x65,
	0x67, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69,
	0x6e, 0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_billing_v1_customer_proto_rawDescOnce sync.Once
	file_yandex_cloud_billing_v1_customer_proto_rawDescData = file_yandex_cloud_billing_v1_customer_proto_rawDesc
)

func file_yandex_cloud_billing_v1_customer_proto_rawDescGZIP() []byte {
	file_yandex_cloud_billing_v1_customer_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_billing_v1_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_billing_v1_customer_proto_rawDescData)
	})
	return file_yandex_cloud_billing_v1_customer_proto_rawDescData
}

var file_yandex_cloud_billing_v1_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_billing_v1_customer_proto_goTypes = []interface{}{
	(*Customer)(nil),       // 0: yandex.cloud.billing.v1.Customer
	(*CustomerPerson)(nil), // 1: yandex.cloud.billing.v1.CustomerPerson
}
var file_yandex_cloud_billing_v1_customer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_billing_v1_customer_proto_init() }
func file_yandex_cloud_billing_v1_customer_proto_init() {
	if File_yandex_cloud_billing_v1_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_billing_v1_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_yandex_cloud_billing_v1_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerPerson); i {
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
			RawDescriptor: file_yandex_cloud_billing_v1_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_billing_v1_customer_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_billing_v1_customer_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_billing_v1_customer_proto_msgTypes,
	}.Build()
	File_yandex_cloud_billing_v1_customer_proto = out.File
	file_yandex_cloud_billing_v1_customer_proto_rawDesc = nil
	file_yandex_cloud_billing_v1_customer_proto_goTypes = nil
	file_yandex_cloud_billing_v1_customer_proto_depIdxs = nil
}
